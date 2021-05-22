package iksm

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/frankenbeanies/uuid4"
	"github.com/spf13/viper"
)

func EnterCookie() string {
	// Prompts the user to enter their iksm_session cookie
	var new_cookie string
	fmt.Println("Go to the page below to find instructions to obtain your iksm_session cookie:\nhttps://github.com/frozenpandaman/splatnet2statink/wiki/mitmproxy-instructions\nEnter it here: ")
	fmt.Scanln(&new_cookie)
	for len(new_cookie) != 40 {
		fmt.Println("Cookie is invalid. Please enter it again.\nCookie: ")
		fmt.Scanln(&new_cookie)
	}
	return new_cookie
}

func GetSessionToken(session_token_code string, auth_code_verifier string, client *http.Client) interface{} {
	app_head := map[string]string{
		"User-Agent":      "OnlineLounge/1.11.0 NASDKAPI Android",
		"Accept-Language": "en-US",
		"Accept":          "application/json",
		"Content-Type":    "application/x-www-form-urlencoded",
		"Content-Length":  "540",
		"Host":            "accounts.nintendo.com",
		"Connection":      "Keep-Alive",
		"Accept-Encoding": "gzip",
	}
	body := map[string]string{
		"client_id":                   "71b963c1b7b6d119",
		"session_token_code":          session_token_code,
		"session_token_code_verifier": strings.ReplaceAll(auth_code_verifier, "=", ""),
	}
	req_data := url.Values{}
	for key, element := range body {
		req_data.Set(key, element)
	}
	body_marshalled := strings.NewReader(req_data.Encode())
	url := "https://accounts.nintendo.com/connect/1.0.0/api/session_token"
	req, err := http.NewRequest("POST", url, body_marshalled)
	if err != nil {
		fmt.Println(err)
	}
	for key, element := range app_head {
		req.Header.Add(key, element)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	type SessionTokenData struct {
		Code         string `json:"code"`
		SessionToken string `json:"session_token"`
	}
	var data SessionTokenData
	json.NewDecoder(resp.Body).Decode(&data)
	fmt.Println(data)
	return data.SessionToken
}


func GetHashFromS2sApi(id_token string, timestamp int, version string, client *http.Client) string {
	api_app_head := map[string]string{"Content-Type": "application/x-www-form-urlencoded", "User-Agent": "splatstatsuploader/" + version}
	api_body := map[string]string{"naIdToken": id_token, "timestamp": fmt.Sprint(timestamp)}
	req_data := url.Values{}
	for key, element := range api_body {
		req_data.Set(key, element)
	}
	body_marshalled := strings.NewReader(req_data.Encode())
	fmt.Println(req_data.Encode())
	req, err := http.NewRequest("POST", "https://elifessler.com/s2s/api/gen2", body_marshalled)
	if err != nil {
		fmt.Println(err)
	}
	for key, element := range api_app_head {
		req.Header.Add(key, element)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	type S2sApiHash struct {
		Hash string `json:"hash"`
	}
	var api_response S2sApiHash
	json.NewDecoder(resp.Body).Decode(&api_response)
	return api_response.Hash
}

type FlapgApiData struct {
	Result struct {
		F  string `json:"f"`
		P1 string `json:"p1"`
		P2 int    `json:"p2"`
		P3 string `json:"p3"`
	} `json:"result"`
}

func CallFlapgApi(id_token string, guid string, timestamp int, f_type string, version string, client *http.Client) FlapgApiData {
	api_app_head := map[string]string{
		"x-token": id_token,
		"x-time":  fmt.Sprint(timestamp),
		"x-guid":  guid,
		"x-hash":  GetHashFromS2sApi(id_token, timestamp, version, client),
		"x-ver":   "3",
		"x-iid":   f_type,
	}
	req, err := http.NewRequest("GET", "https://flapg.com/ika2/api/login?public", nil)
	if err != nil {
		fmt.Println(err)
	}
	for key, element := range api_app_head {
		req.Header.Add(key, element)
	}
	// resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	var data FlapgApiData
	// json.NewDecoder(resp.Body).Decode(&data)
	return data
}

func GetCookie(version string, client *http.Client) (string, string) {
	timestamp := int(time.Now().Unix())
	guid := uuid4.New().String()
	user_lang := viper.GetString("user_lang")
	app_head := map[string]string{
		"Host":            "accounts.nintendo.com",
		"Accept-Encoding": "gzip deflate",
		"Content-Type":    "application/json; charset=utf-8",
		"Accept-Language": user_lang,
		"Content-Length":  "439",
		"Accept":          "application/json",
		"Connection":      "Keep-Alive",
		"User-Agent":      "OnlineLounge/1.11.0 NASDKAPI Android",
	}
	body := map[string]string{
		"client_id":     "71b963c1b7b6d119", // Splatoon 2 service
		"session_token": viper.GetString("session_token"),
		"grant_type":    "urn:ietf:params:oauth:grant-type:jwt-bearer-session-token",
	}
	body_marshalled, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
	}
	url := "https://accounts.nintendo.com/connect/1.0.0/api/token"
	req, err := http.NewRequest("POST", url, bytes.NewReader(body_marshalled))
	if err != nil {
		fmt.Println(err)
	}
	for key, element := range app_head {
		req.Header.Add(key, element)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	type IdResponse struct {
		AccessToken string   `json:"access_token"`
		ExpiresIn   int      `json:"expires_in"`
		IDToken     string   `json:"id_token"`
		Scope       []string `json:"scope"`
		TokenType   string   `json:"token_type"`
	}
	var id_response IdResponse
	json.NewDecoder(resp.Body).Decode(&id_response)
	fmt.Println(id_response)
	app_head = map[string]string{
		"User-Agent":      "OnlineLounge/1.11.0 NASDKAPI Android",
		"Accept-Language": user_lang,
		"Accept":          "application/json",
		"Authorization":   "Bearer " + id_response.AccessToken,
		"Host":            "api.accounts.nintendo.com",
		"Connection":      "Keep-Alive",
		"Accept-Encoding": "gzip deflate",
	}
	url = "https://api.accounts.nintendo.com/2.0.0/users/me"
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	for key, element := range app_head {
		req.Header.Add(key, element)
	}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	type UserInfo struct {
		Analyticsoptedin          bool `json:"analyticsOptedIn"`
		Analyticsoptedinupdatedat int  `json:"analyticsOptedInUpdatedAt"`
		Analyticspermissions      struct {
			Internalanalysis struct {
				Permitted bool `json:"permitted"`
				Updatedat int  `json:"updatedAt"`
			} `json:"internalAnalysis"`
			Targetmarketing struct {
				Permitted bool `json:"permitted"`
				Updatedat int  `json:"updatedAt"`
			} `json:"targetMarketing"`
		} `json:"analyticsPermissions"`
		Birthday      string `json:"birthday"`
		Candidatemiis []struct {
			Updatedat     int    `json:"updatedAt"`
			Favoritecolor string `json:"favoriteColor"`
			Type          string `json:"type"`
			Clientid      string `json:"clientId"`
			Storedata     struct {
				Num3 string `json:"3"`
			} `json:"storeData"`
			ID               string `json:"id"`
			Imageuritemplate string `json:"imageUriTemplate"`
			Imageorigin      string `json:"imageOrigin"`
			Etag             string `json:"etag"`
		} `json:"candidateMiis"`
		Clientfriendsoptedin          bool   `json:"clientFriendsOptedIn"`
		Clientfriendsoptedinupdatedat int    `json:"clientFriendsOptedInUpdatedAt"`
		Country                       string `json:"country"`
		Createdat                     int    `json:"createdAt"`
		Eachemailoptedin              struct {
			Deals struct {
				Optedin   bool `json:"optedIn"`
				Updatedat int  `json:"updatedAt"`
			} `json:"deals"`
			Survey struct {
				Optedin   bool `json:"optedIn"`
				Updatedat int  `json:"updatedAt"`
			} `json:"survey"`
		} `json:"eachEmailOptedIn"`
		Emailoptedin          bool   `json:"emailOptedIn"`
		Emailoptedinupdatedat int    `json:"emailOptedInUpdatedAt"`
		Emailverified         bool   `json:"emailVerified"`
		Gender                string `json:"gender"`
		ID                    string `json:"id"`
		Ischild               bool   `json:"isChild"`
		Language              string `json:"language"`
		Mii                   struct {
			Clientid string `json:"clientId"`
			Coredata struct {
				Num4 string `json:"4"`
			} `json:"coreData"`
			Etag             string `json:"etag"`
			Favoritecolor    string `json:"favoriteColor"`
			ID               string `json:"id"`
			Imageorigin      string `json:"imageOrigin"`
			Imageuritemplate string `json:"imageUriTemplate"`
			Storedata        struct {
				Num3 string `json:"3"`
			} `json:"storeData"`
			Type      string `json:"type"`
			Updatedat int    `json:"updatedAt"`
		} `json:"mii"`
		Nickname   string      `json:"nickname"`
		Region     interface{} `json:"region"`
		Screenname string      `json:"screenName"`
		Timezone   struct {
			ID               string `json:"id"`
			Name             string `json:"name"`
			Utcoffset        string `json:"utcOffset"`
			Utcoffsetseconds int    `json:"utcOffsetSeconds"`
		} `json:"timezone"`
		Updatedat int `json:"updatedAt"`
	}
	var user_info UserInfo
	json.NewDecoder(resp.Body).Decode(&user_info)
	nickname := user_info.Nickname
	app_head = map[string]string{
		"Host":             "api-lp1.znc.srv.nintendo.net",
		"Accept-Language":  user_lang,
		"User-Agent":       "com.nintendo.znca/1.11.0 (Android/7.1.2)",
		"Accept":           "application/json",
		"X-ProductVersion": "1.11.0",
		"Content-Type":     "application/json; charset=utf-8",
		"Connection":       "Keep-Alive",
		"Authorization":    "Bearer",
		"X-Platform":       "Android",
		"Accept-Encoding":  "gzip",
	}
	id_token := id_response.AccessToken
	flapg_nso := CallFlapgApi(id_token, guid, timestamp, "nso", version, client).Result
	parameter := map[string]interface{}{
		"f":          flapg_nso.F,
		"naIdToken":  flapg_nso.P1,
		"timestamp":  flapg_nso.P2,
		"requestId":  flapg_nso.P3,
		"naCountry":  user_info.Country,
		"naBirthday": user_info.Birthday,
		"language":   user_info.Language,
	}
	new_body := make(map[string]interface{})
	new_body["parameter"] = parameter
	url = "https://api-lp1.znc.srv.nintendo.net/v1/Account/Login"
	new_body_json, err := json.Marshal(new_body)
	if err != nil {
		fmt.Println(err)
	}
	req, err = http.NewRequest("POST", url, bytes.NewReader(new_body_json))
	if err != nil {
		fmt.Println(err)
	}
	for key, element := range app_head {
		req.Header.Add(key, element)
	}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	type SplatoonToken struct {
		Correlationid struct {
		} `json:"correlationId"`
		Result struct {
			Firebasecredential struct {
				Accesstoken interface{} `json:"accessToken"`
				Expiresin   int         `json:"expiresIn"`
			} `json:"firebaseCredential"`
			User struct {
				ID         int64  `json:"id"`
				Imageuri   string `json:"imageUri"`
				Membership struct {
					Active bool `json:"active"`
				} `json:"membership"`
				Name      string `json:"name"`
				Supportid string `json:"supportId"`
			} `json:"user"`
			Webapiservercredential struct {
				Accesstoken string `json:"accessToken"`
				Expiresin   int    `json:"expiresIn"`
			} `json:"webApiServerCredential"`
		} `json:"result"`
		Status struct {
		} `json:"status"`
	}
	var splatoon_token SplatoonToken
	json.NewDecoder(resp.Body).Decode(&splatoon_token)
	fmt.Println(splatoon_token)
	id_token = splatoon_token.Result.Webapiservercredential.Accesstoken
	flapg_app := CallFlapgApi(id_token, guid, timestamp, "app", version, client).Result
	app_head = map[string]string{
		"Host":             "api-lp1.znc.srv.nintendo.net",
		"User-Agent":       "com.nintendo.znca/1.11.0 (Android/7.1.2)",
		"Accept":           "application/json",
		"X-ProductVersion": "1.11.0",
		"Content-Type":     "application/json; charset=utf-8",
		"Connection":       "Keep-Alive",
		"Authorization":    "Bearer " + id_token,
		//"Content-Length":   "37",
		"X-Platform":      "Android",
		"Accept-Encoding": "gzip deflate",
		"Accept-Language": user_lang,
	}
	new_body_2 := map[string]map[string]interface{}{}
	parameter = map[string]interface{}{
		"id":                5741031244955648,
		"f":                 flapg_app.F,
		"registrationToken": flapg_app.P1,
		"timestamp":         flapg_app.P2,
		"requestId":         flapg_app.P3,
	}
	new_body_2["paramter"] = parameter
	url = "https://api-lp1.znc.srv.nintendo.net/v2/Game/GetWebServiceToken"
	body_json, err := json.Marshal(new_body_2)
	if err != nil {
		fmt.Println(err)
	}
	req, err = http.NewRequest("POST", url, bytes.NewReader(body_json))
	if err != nil {
		fmt.Println(err)
	}
	for key, element := range app_head {
		req.Header.Add(key, element)
	}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	var splatoon_access_token map[string]map[string]string
	json.NewDecoder(resp.Body).Decode(&splatoon_access_token)
	app_head = map[string]string{
		"Host":                    "app.splatoon2.nintendo.net",
		"X-IsAppAnalyticsOptedIn": "false",
		"Accept":                  "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Accept-Encoding":         "gzip deflate",
		"X-GameWebToken":          splatoon_access_token["result"]["accessToken"],
		"Accept-Language":         user_lang,
		"X-IsAnalyticsOptedIn":    "false",
		"Connection":              "keep-alive",
		"DNT":                     "0",
		"User-Agent":              "Mozilla/5.0 (Linux; Android 7.1.2; Pixel Build/NJH47D; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/59.0.3071.125 Mobile Safari/537.36",
		"X-Requested-With":        "com.nintendo.znca",
	}
	url = "https://app.splatoon2.nintendo.net/?lang=" + user_lang
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	for key, element := range app_head {
		req.Header.Add(key, element)
	}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "iksm_session" {
			return nickname, cookie.Value
		}
	}
	return nickname, ""
}

func GenNewCookie(reason string, version string, client *http.Client) {
    // Attempts to generate a new cookie in case the provided one is invalid.

    manual := false

    if reason == "blank" {
        fmt.Println("Blank cookie.")
	} else if reason == "auth" {  // authentication error
        fmt.Println("The stored cookie has expired.")
	} else {  // server error or player hasn't battled before
        fmt.Println("Cannot access SplatNet 2 without having played at least one battle online.")
        os.Exit(1)
	}
    if viper.GetString("session_token") == "" {
        fmt.Println("session_token is blank. Please log in to your Nintendo Account to obtain your session_token.")
        new_token := LogIn(version)
        if new_token == nil {
			fmt.Println("There was a problem logging you in. Please try again later.")
		} else {
            if *new_token == "skip" {	// user has opted to manually enter cookie
                manual = true
                fmt.Println("\nYou have opted against automatic cookie generation and must manually input your iksm_session cookie.\n")
			} else {
                fmt.Println("\nWrote session_token to config.txt.")
			}
            viper.Set("session_token", *new_token)
			viper.WriteConfig()
		}
	} else if viper.Get("session_token") == "skip" {
        manual = true
        fmt.Println("\nYou have opted against automatic cookie generation and must manually input your iksm_session cookie. You may clear this setting by removing \"skip\" from the session_token field in config.txt.\n")
	}

	var new_cookie string
	var acc_name string
    if manual {
        new_cookie = EnterCookie()
	} else {
        fmt.Println("Attempting to generate new cookie...")
        acc_name, new_cookie = GetCookie(version, client)
	}
    viper.Set("cookie", new_cookie)
    viper.WriteConfig()
    if manual {
		fmt.Println("Wrote iksm_session cookie to config.yaml.")
    } else {
		fmt.Println("Wrote iksm_session cookie for " + acc_name + " to config.yaml.")
	}
}

func LogIn(version string) *string {
	auth_state_unencoded := make([]byte, 36)
	_, err := rand.Read(auth_state_unencoded)
	if err != nil {
		fmt.Println(err)
	}
	auth_state := base64.RawURLEncoding.EncodeToString(auth_state_unencoded)
	auth_code_verifier_unencoded := make([]byte, 32)
	_, err = rand.Read(auth_code_verifier_unencoded)
	if err != nil {
		fmt.Println(err)
	}
	auth_code_verifier := base64.RawURLEncoding.EncodeToString(auth_code_verifier_unencoded)
	auth_code_hash := sha256.Sum256([]byte(strings.ReplaceAll(auth_code_verifier, "=", "")))
	auth_code_challenge := base64.RawURLEncoding.EncodeToString(auth_code_hash[:])
	app_head := map[string]string{
		"Host":                      "accounts.nintendo.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Linux; Android 7.1.2; Pixel Build/NJH47D; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/59.0.3071.125 Mobile Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8n",
		"DNT":                       "1",
		"Accept-Encoding":           "gzip,deflate,br",
	}
	body := map[string]string{
		"state":                               auth_state,
		"redirect_uri":                        "npf71b963c1b7b6d119://auth",
		"client_id":                           "71b963c1b7b6d119",
		"scope":                               "openid user user.birthday user.mii user.screenName",
		"response_type":                       "session_token_code",
		"session_token_code_challenge":        strings.ReplaceAll(auth_code_challenge, "=", ""),
		"session_token_code_challenge_method": "S256",
		"theme":                               "login_form",
	}
	data := url.Values{}
	for key, element := range body {
		data.Set(key, element)
	}
	url := "https://accounts.nintendo.com/connect/1.0.0/authorize"
	req, err := http.NewRequest("GET", url, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	q := req.URL.Query()
	for key, element := range body {
		q.Add(key, element)
	}
	req.URL.RawQuery = q.Encode()
	for key, element := range app_head {
		req.Header.Add(key, element)
	}
	client := &http.Client{}
	post_login := req.URL.String()
	fmt.Println("Navigate to this URL in your browser:")
	fmt.Println(post_login)
	fmt.Println("Log in, right click the \"Select this account\" button, copy the link address, and paste it below:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	use_account_url := scanner.Text()
	re := regexp.MustCompile("de=(.*)&")
	session_token_code := re.FindAllStringSubmatch(use_account_url, -1)
	session_token := GetSessionToken(session_token_code[0][1], auth_code_verifier, client).(string)
	return &session_token
}