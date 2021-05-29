package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"cass-dlcm.dev/splatstatsuploader/helpers"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var progVersion = "1.5.1"

func doSelfUpdate() {
	v := semver.MustParse(progVersion)
	latest, err := selfupdate.UpdateSelf(v, "cass-dlcm/splatstats-uploader-go")
	if err != nil {
		fmt.Println("Binary update failed:", err)
		return
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		fmt.Println("Current binary is the latest version", progVersion)
	} else {
		fmt.Println("Successfully updated to version", latest.Version)
		fmt.Println("Release note:\n", latest.ReleaseNotes)
	}
}

func setLanguage() {
	fmt.Println("Default locale is en-US. Press Enter to accept, or enter your own (see readme for list).")
	var locale string
	// Taking input from user
	fmt.Scanln(&locale)
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
			fmt.Println("Invalid language code. Please try entering it again.")
			fmt.Scanln(&locale)
			_, exists = languageList[locale]
		}
		viper.Set("user_lang", locale)
	}
	viper.WriteConfig()
}

func setApiToken(client *http.Client) {
	var username string
	fmt.Println("SplatStats username: ")
	fmt.Scanln(&username)
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
	}
	url := "http://localhost:8000/auth/api-token/"
	authJson, err := json.Marshal(map[string]string{
		"username": username, "password": string(password),
	})
	if err != nil {
		fmt.Println(err)
	}
	authBody := bytes.NewReader(authJson)
	req, err := http.NewRequest("POST", url, authBody)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()
	viper.Set("api_key", newStr[10:len(newStr)-2])
	viper.WriteConfig()
}

func getFlags() (int, bool, bool, bool) {
	m := flag.Int("m", -1, "To monitor for new match results.")
	f := flag.Bool("f", false, "To upload battles/shifts from files.")
	s := flag.Bool("s", false, "To save battles/shifts to files.")
	salmon := flag.Bool("salmon", false, "To upload salmon run matches.")
	flag.Parse()
	if *f && *s {
		fmt.Println("Cannot use -f and -s together. Exiting.")
		os.Exit(1)
	}
	if *f && *m != -1 {
		fmt.Println("Cannot use -f and -m together. Exiting")
		os.Exit(1)
	}
	return *m, *f, *s, *salmon
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
			fmt.Println("No config file found. One will be created.")
		} else {
			// Config file was found but another error was produced
			fmt.Println("Error reading the config file.")
			os.Exit(1)
		}
	}

	// Config file found and successfully parsed

	// Set undefined variables
	viper.SetDefault("api_key", "")
	viper.SetDefault("cookie", "")
	viper.SetDefault("session_token", "")
	viper.SetDefault("user_lang", "")

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

	m, f, s, salmon := getFlags()
	if m != -1 {
		helpers.Monitor(m, s, salmon, viper.GetString("api_key"), progVersion, appHead, client)
	} else if f {
		helpers.File(salmon, viper.GetString("api_key"), progVersion, client)
	} else {
		if salmon {
			helpers.GetSplatnetSalmon(s, viper.GetString("api_key"), progVersion, appHead, client)
		} else {
			helpers.GetSplatnetBattle(s, viper.GetString("api_key"), progVersion, appHead, client)
		}
	}
}
