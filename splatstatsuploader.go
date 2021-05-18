package splatstatsuploader

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

var A_VERSION = "1.5.10"
var B_VERSION = "1.0.0"

func CheckForUpdates() {
	latest_script, err := http.get(
		"https://raw.githubusercontent.com/cass-dlcm/splatstats-uploader/main/golang_version/src/splatstatsuploader.go"
	)
	if err {
		fmt.Println("Error retrieving the latest version.")
	}
	defer latest_script.Body.Close()
	body, err := io.ReadAll(resp.Body)
	re := regexp.MustCompile("B_VERSION = \"([\d.]*)\"")
	new_version := re.FindString(body)
	fmt.Println(new_version)
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
	}
	viper.Set("user_lang", locale)
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
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("api_key", "")
	viper.SetDefault("cookie", "")
	viper.SetDefault("session_token", "")
	viper.SetDefault("user_lang", "")

	if !(viper.IsSet("user_lang")) {
		SetLanguage()
	}
}
