package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/cass-dlcm/splatstatsuploader/statink2splatstats"
	"net/http"
	"os"
	"time"

	"github.com/blang/semver"
	"github.com/cass-dlcm/splatstatsuploader/data"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var progVersion = "4.0.0"

func doSelfUpdate() {
	v := semver.MustParse(progVersion)

	latest, err := selfupdate.UpdateSelf(v, "cass-dlcm/splatstats-uploader-go")
	if err != nil {
		panic(err)
	}

	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		if _, err := fmt.Println("Current binary is the latest version", progVersion); err != nil {
			panic(err)
		}
	} else {
		if _, err := fmt.Println("Successfully updated to version", latest.Version); err != nil {
			panic(err)
		}
		if _, err := fmt.Println("Release note:\n", latest.ReleaseNotes); err != nil {
			panic(err)
		}
	}
}

func setLanguage() {
	if _, err := fmt.Println("Default locale is en-US. Press Enter to accept, or enter your own (see readme for list)."); err != nil {
		panic(err)
	}

	var locale string
	// Taking input from user
	if _, err := fmt.Scanln(&locale); err != nil {
		panic(err)
	}

	if locale == "" {
		viper.Set("user_lang", "en-US")
	} else {
		languageList := map[string]string{
			"en-US": "en-US",
			"es-MX": "es-MX",
			"fr-CA": "fr-CA",
			"ja-JP": "ja-JP",
			"en-GB": "en-GB",
			"es-ES": "es-ES",
			"fr-FR": "fr-FR",
			"de-DE": "de-DE",
			"it-IT": "it-IT",
			"nl-NL": "nl-NL",
			"ru-RU": "ru-RU",
		}
		_, exists := languageList[locale]
		for !exists {
			if _, err := fmt.Println("Invalid language code. Please try entering it again."); err != nil {
				panic(err)
			}

			if _, err := fmt.Scanln(&locale); err != nil {
				panic(err)
			}

			_, exists = languageList[locale]
		}
		viper.Set("user_lang", locale)
	}

	if err := viper.WriteConfig(); err != nil {
		panic(err)
	}
}

func setApiToken(client *http.Client) {
	var username string

	if _, err := fmt.Println("SplatStats username: "); err != nil {
		panic(err)
	}

	if _, err := fmt.Scanln(&username); err != nil {
		panic(err)
	}

	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
	}

	url := "https://splatstats.cass-dlcm.dev/auth/api-token/"

	authJson, err := json.Marshal(map[string]string{
		"username": username, "password": string(password),
	})
	if err != nil {
		fmt.Println(err)
	}

	authBody := bytes.NewReader(authJson)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, authBody)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		panic(err)
	}

	newStr := buf.String()
	viper.Set("api_key", newStr[10:len(newStr)-2])

	if err := viper.WriteConfig(); err != nil {
		panic(err)
	}
}

func getFlags() (int, bool, bool, bool, bool) {
	m := flag.Int("m", -1, "To monitor for new match results.")
	f := flag.Bool("f", false, "To upload battles/shifts from files.")
	s := flag.Bool("s", false, "To save battles/shifts to files.")
	statink := flag.Bool("statink", false, "To migrate from stat.ink to SplatStats.")
	salmon := flag.Bool("salmon", false, "To upload salmon run matches.")
	flag.Parse()

	if *f && *s {
		if _, err := fmt.Println("Cannot use -f and -s together. Exiting."); err != nil {
			panic(err)
		}

		os.Exit(1)
	}

	if *m != -1 && *statink {
		if _, err := fmt.Println("Cannot use -m and --statink together. Exiting."); err != nil {
			panic(err)
		}

		os.Exit(1)
	}

	if *f && *m != -1 {
		if _, err := fmt.Println("Cannot use -f and -m together. Exiting"); err != nil {
			panic(err)
		}

		os.Exit(1)
	}

	return *m, *f, *s, *salmon, *statink
}

func main() {
	doSelfUpdate()

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			if _, err := fmt.Println("No config file found. One will be created."); err != nil {
				panic(err)
			}
		} else {
			// Config file was found but another error was produced
			if _, err := fmt.Println("Error reading the config file."); err != nil {
				panic(err)
			}
			os.Exit(1)
		}
	}

	// Config file found and successfully parsed

	// Set undefined variables
	viper.SetDefault("api_key", "")
	viper.SetDefault("cookie", "")
	viper.SetDefault("session_token", "")
	viper.SetDefault("user_lang", "")
	viper.SetDefault("statink_api_key", "")

	client := &http.Client{}

	if !(viper.IsSet("api_key")) || viper.GetString("api_key") == "" {
		setApiToken(client)
	}

	if !(viper.IsSet("user_lang")) || viper.GetString("user_lang") == "" {
		setLanguage()
	}

	_, timezone := time.Now().Zone()
	timezone = -timezone / 60
	appHead := map[string]string{
		"Host":              "app.splatoon2.nintendo.net",
		"x-unique-id":       "32449507786579989235",
		"x-requested-with":  "XMLHttpRequest",
		"x-timezone-offset": fmt.Sprint(timezone),
		"User-Agent":        "Mozilla/5.0 (Linux; Android 7.1.2; Pixel Build/NJH47D; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/59.0.3071.125 Mobile Safari/537.36",
		"Accept":            "*/*",
		"Referer":           "https://app.splatoon2.nintendo.net/home",
		"Accept-Encoding":   "gzip deflate",
		"Accept-Language":   viper.GetString("user_lang"),
	}

	m, f, s, salmon, statink := getFlags()
	if m != -1 {
		data.Monitor(m, s, salmon, viper.GetString("api_key"), progVersion, appHead, client)
	} else if f {
		if statink {

		} else {
			data.File(salmon, viper.GetString("api_key"), client)
		}
	} else if statink {
		statink2splatstats.Migrate(s, salmon, viper.GetString("api_key"), client)
	} else {
		if salmon {
			data.GetSplatnetSalmon(s, viper.GetString("api_key"), progVersion, appHead, client)
		} else {
			data.GetSplatnetBattle(s, viper.GetString("api_key"), progVersion, appHead, client)
		}
	}
}
