package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/hashicorp/go-version"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var A_VERSION = "1.5.10"
var B_VERSION = "1.0.1"

func CheckForUpdates() {
	latest_script, err := http.Get("https://raw.githubusercontent.com/cass-dlcm/splatstats-uploader-go/main/splatstatsuploader.go")
	if err != nil {
		fmt.Println("Error retrieving the latest version.")
	}
	defer latest_script.Body.Close()
	body, _ := io.ReadAll(latest_script.Body)
	re := regexp.MustCompile("B_VERSION = \"([\\d.]*)\"")
	new_version := re.FindString(string(body))
	v1, err := version.NewVersion(B_VERSION)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	v2, err := version.NewVersion(new_version[13 : len(new_version)-1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if v1.LessThan(v2) {
		fmt.Println("New version availbile at https://github.com/cass-dlcm/splatstats-uploader-go.")
		fmt.Println("Please download the new version before continuing.")
		latest_script.Body.Close()
		os.Exit(0)
	}
}

func SetLanguage() {
	fmt.Println("Default locale is en-US. Press Enter to accept, or enter your own (see readme for list).")
	var locale string
	// Taking input from user
	fmt.Scanln(&locale)
	if locale == "" {
		viper.Set("user_lang", "en-US")
	} else {
		language_list := map[string]string{
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
		_, exists := language_list[locale]
		for !exists {
			fmt.Println("Invalid language code. Please try entering it again.")
			fmt.Scanln(&locale)
			_, exists = language_list[locale]
		}
		viper.Set("user_lang", locale)
	}
	viper.WriteConfig()
}

func SetApiToken() {
	var username string
	fmt.Println("SplatStats username: ")
	fmt.Scanln(&username)
	password, _ := term.ReadPassword(int(os.Stdin.Fd()))
	url := "http://localhost:8000/auth/api-token/"
	client := &http.Client{}
	auth_json, _ := json.Marshal(map[string]string{
		"username": username, "password": string(password),
	})
	auth_body := bytes.NewReader(auth_json)
	req, _ := http.NewRequest("POST", url, auth_body)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()
	viper.Set("api_key", newStr[10:len(newStr)-2])
	viper.WriteConfig()
}

func main() {

	CheckForUpdates()

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

	if !(viper.IsSet("api_key")) {
		SetApiToken()
	}

	if !(viper.IsSet("user_lang")) || viper.GetString("user_lang") == "" {
		SetLanguage()
	}
}
