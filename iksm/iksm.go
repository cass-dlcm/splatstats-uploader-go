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

// Prompts the user to enter their iksm_session cookie
func enterCookie() string {
	var newCookie string
	fmt.Println("Go to the page below to find instructions to obtain your iksm_session cookie:\nhttps://github.com/frozenpandaman/splatnet2statink/wiki/mitmproxy-instructions\nEnter it here: ")
	fmt.Scanln(&newCookie)
	for len(newCookie) != 40 {
		fmt.Println("Cookie is invalid. Please enter it again.\nCookie: ")
		fmt.Scanln(&newCookie)
	}
	return newCookie
}

// Helper function for logIn().
func getSessionToken(sessionTokenCode string, authCodeVerifier string, client *http.Client) interface{} {
	appHead := map[string]string{
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
		"session_token_code":          sessionTokenCode,
		"session_token_code_verifier": strings.ReplaceAll(authCodeVerifier, "=", ""),
	}
	reqData := url.Values{}
	for key, element := range body {
		reqData.Set(key, element)
	}
	bodyMarshalled := strings.NewReader(reqData.Encode())
	url := "https://accounts.nintendo.com/connect/1.0.0/api/session_token"
	req, err := http.NewRequest("POST", url, bodyMarshalled)
	if err != nil {
		panic(err)
	}
	for key, element := range appHead {
		req.Header.Add(key, element)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	type SessionTokenData struct {
		Code         string `json:"code"`
		SessionToken string `json:"session_token"`
	}
	var data SessionTokenData
	json.NewDecoder(resp.Body).Decode(&data)
	return data.SessionToken
}

// Passes an idToken and timestamp to the s2s API and fetches the resultant hash from the response.
func getHashFromS2sApi(idToken string, timestamp int, version string, client *http.Client) string {
	apiAppHead := map[string]string{"Content-Type": "application/x-www-form-urlencoded", "User-Agent": "splatstatsuploader/" + version}
	apiBody := map[string]string{"naIdToken": idToken, "timestamp": fmt.Sprint(timestamp)}
	reqData := url.Values{}
	for key, element := range apiBody {
		reqData.Set(key, element)
	}
	bodyMarshalled := strings.NewReader(reqData.Encode())
	req, err := http.NewRequest("POST", "https://elifessler.com/s2s/api/gen2", bodyMarshalled)
	if err != nil {
		panic(err)
	}
	for key, element := range apiAppHead {
		req.Header.Add(key, element)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	type S2sApiHash struct {
		Hash string `json:"hash"`
	}
	var apiResponse S2sApiHash
	json.NewDecoder(resp.Body).Decode(&apiResponse)
	return apiResponse.Hash
}

type flapgApiData struct {
	Result struct {
		F  string `json:"f"`
		P1 string `json:"p1"`
		P2 int    `json:"p2"`
		P3 string `json:"p3"`
	} `json:"result"`
}

// Passes in headers to the flapg API (Android emulator) and fetches the response.
func callFlapgApi(idToken string, guid string, timestamp int, fType string, version string, client *http.Client) flapgApiData {
	apiAppHead := map[string]string{
		"x-token": idToken,
		"x-time":  fmt.Sprint(timestamp),
		"x-guid":  guid,
		"x-hash":  getHashFromS2sApi(idToken, timestamp, version, client),
		"x-ver":   "3",
		"x-iid":   fType,
	}
	req, err := http.NewRequest("GET", "https://flapg.com/ika2/api/login?public", nil)
	if err != nil {
		panic(err)
	}
	for key, element := range apiAppHead {
		req.Header.Add(key, element)
	}
	// resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	var data flapgApiData
	// json.NewDecoder(resp.Body).Decode(&data)
	return data
}

type idResponseS struct {
	AccessToken string   `json:"access_token"`
	ExpiresIn   int      `json:"expires_in"`
	IDToken     string   `json:"id_token"`
	Scope       []string `json:"scope"`
	TokenType   string   `json:"token_type"`
}

func getIdResponse(userLang string, sessionToken string, client *http.Client) idResponseS {
	appHead := map[string]string{
		"Host":            "accounts.nintendo.com",
		"Accept-Encoding": "gzip deflate",
		"Content-Type":    "application/json; charset=utf-8",
		"Accept-Language": userLang,
		"Content-Length":  "439",
		"Accept":          "application/json",
		"Connection":      "Keep-Alive",
		"User-Agent":      "OnlineLounge/1.11.0 NASDKAPI Android",
	}
	body := map[string]string{
		"client_id":     "71b963c1b7b6d119", // Splatoon 2 service
		"session_token": sessionToken,
		"grant_type":    "urn:ietf:params:oauth:grant-type:jwt-bearer-session-token",
	}
	bodyMarshalled, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	url := "https://accounts.nintendo.com/connect/1.0.0/api/token"
	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyMarshalled))
	if err != nil {
		panic(err)
	}
	for key, element := range appHead {
		req.Header.Add(key, element)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var idResponse idResponseS
	json.NewDecoder(resp.Body).Decode(&idResponse)
	return idResponse
}

type userInfoS struct {
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

func getUserInfo(userLang string, idResponse idResponseS, client *http.Client) userInfoS {
	appHead := map[string]string{
		"User-Agent":      "OnlineLounge/1.11.0 NASDKAPI Android",
		"Accept-Language": userLang,
		"Accept":          "application/json",
		"Authorization":   "Bearer " + idResponse.AccessToken,
		"Host":            "api.accounts.nintendo.com",
		"Connection":      "Keep-Alive",
		"Accept-Encoding": "gzip deflate",
	}
	url := "https://api.accounts.nintendo.com/2.0.0/users/me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	for key, element := range appHead {
		req.Header.Add(key, element)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	var userInfo userInfoS
	json.NewDecoder(resp.Body).Decode(&userInfo)
	return userInfo
}

type splatoonTokenS struct {
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

func getSplatoonToken(userLang string, idResponse idResponseS, userInfo userInfoS, guid string, timestamp int, version string, client *http.Client) splatoonTokenS {
	appHead := map[string]string{
		"Host":             "api-lp1.znc.srv.nintendo.net",
		"Accept-Language":  userLang,
		"User-Agent":       "com.nintendo.znca/1.11.0 (Android/7.1.2)",
		"Accept":           "application/json",
		"X-ProductVersion": "1.11.0",
		"Content-Type":     "application/json; charset=utf-8",
		"Connection":       "Keep-Alive",
		"Authorization":    "Bearer",
		"X-Platform":       "Android",
		"Accept-Encoding":  "gzip",
	}
	idToken := idResponse.AccessToken
	flapgNso := callFlapgApi(idToken, guid, timestamp, "nso", version, client).Result
	parameter := map[string]interface{}{
		"f":          flapgNso.F,
		"naIdToken":  flapgNso.P1,
		"timestamp":  flapgNso.P2,
		"requestId":  flapgNso.P3,
		"naCountry":  userInfo.Country,
		"naBirthday": userInfo.Birthday,
		"language":   userInfo.Language,
	}
	newBody := make(map[string]interface{})
	newBody["parameter"] = parameter
	url := "https://api-lp1.znc.srv.nintendo.net/v1/Account/Login"
	newBodyJson, err := json.Marshal(newBody)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(newBodyJson))
	if err != nil {
		panic(err)
	}
	for key, element := range appHead {
		req.Header.Add(key, element)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	var splatoonToken splatoonTokenS
	json.NewDecoder(resp.Body).Decode(&splatoonToken)
	return splatoonToken
}

type splatoonAccessTokenS struct {
	Correlationid map[string]interface{} `json:"correlationId"`
	Result        struct {
		Accesstoken string `json:"accessToken"`
		Expiresin   int    `json:"expiresIn"`
	} `json:"result"`
	Status map[string]interface{} `json:"status"`
}

func getSplatoonAccessToken(userLang string, splatoonToken splatoonTokenS, guid string, timestamp int, version string, client *http.Client) splatoonAccessTokenS {
	idToken := splatoonToken.Result.Webapiservercredential.Accesstoken
	flapgApp := callFlapgApi(idToken, guid, timestamp, "app", version, client).Result
	appHead := map[string]string{
		"Host":             "api-lp1.znc.srv.nintendo.net",
		"User-Agent":       "com.nintendo.znca/1.11.0 (Android/7.1.2)",
		"Accept":           "application/json",
		"X-ProductVersion": "1.11.0",
		"Content-Type":     "application/json; charset=utf-8",
		"Connection":       "Keep-Alive",
		"Authorization":    "Bearer " + idToken,
		//"Content-Length":   "37",
		"X-Platform":      "Android",
		"Accept-Encoding": "gzip deflate",
		"Accept-Language": userLang,
	}
	newBody2 := map[string]map[string]interface{}{}
	parameter := map[string]interface{}{
		"id":                5741031244955648,
		"f":                 flapgApp.F,
		"registrationToken": flapgApp.P1,
		"timestamp":         flapgApp.P2,
		"requestId":         flapgApp.P3,
	}
	newBody2["parameter"] = parameter
	url := "https://api-lp1.znc.srv.nintendo.net/v2/Game/GetWebServiceToken"
	bodyJson, err := json.Marshal(newBody2)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyJson))
	if err != nil {
		panic(err)
	}
	for key, element := range appHead {
		req.Header.Add(key, element)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	var splatoonAccessToken splatoonAccessTokenS
	json.NewDecoder(resp.Body).Decode(&splatoonAccessToken)
	return splatoonAccessToken
}

// Returns a new cookie.
func getCookie(version string, client *http.Client) (string, string) {
	timestamp := int(time.Now().Unix())
	guid := uuid4.New().String()
	userLang := viper.GetString("user_lang")
	idResponse := getIdResponse(userLang, viper.GetString("session_token"), client)
	userInfo := getUserInfo(userLang, idResponse, client)
	nickname := userInfo.Nickname
	splatoonToken := getSplatoonToken(userLang, idResponse, userInfo, guid, timestamp, version, client)
	splatoonAccessToken := getSplatoonAccessToken(userLang, splatoonToken, guid, timestamp, version, client)
	appHead := map[string]string{
		"Host":                    "app.splatoon2.nintendo.net",
		"X-IsAppAnalyticsOptedIn": "false",
		"Accept":                  "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Accept-Encoding":         "gzip deflate",
		"X-GameWebToken":          splatoonAccessToken.Result.Accesstoken,
		"Accept-Language":         userLang,
		"X-IsAnalyticsOptedIn":    "false",
		"Connection":              "keep-alive",
		"DNT":                     "0",
		"User-Agent":              "Mozilla/5.0 (Linux; Android 7.1.2; Pixel Build/NJH47D; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/59.0.3071.125 Mobile Safari/537.36",
		"X-Requested-With":        "com.nintendo.znca",
	}
	url := "https://app.splatoon2.nintendo.net/?lang=" + userLang
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	for key, element := range appHead {
		req.Header.Add(key, element)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "iksm_session" {
			return nickname, cookie.Value
		}
	}
	return nickname, ""
}

// GenNewCookie attempts to generate a new cookie in case the provided one is invalid.
func GenNewCookie(reason string, version string, client *http.Client) {
	manual := false

	if reason == "blank" {
		fmt.Println("Blank cookie.")
	} else if reason == "auth" { // authentication error
		fmt.Println("The stored cookie has expired.")
	} else { // server error or player hasn't battled before
		fmt.Println("Cannot access SplatNet 2 without having played at least one battle online.")
		os.Exit(1)
	}
	if viper.GetString("session_token") == "" {
		fmt.Println("session_token is blank. Please log in to your Nintendo Account to obtain your session_token.")
		newToken := logIn(version)
		if newToken == nil {
			fmt.Println("There was a problem logging you in. Please try again later.")
		} else {
			if *newToken == "skip" { // user has opted to manually enter cookie
				manual = true
				fmt.Println("\nYou have opted against automatic cookie generation and must manually input your iksm_session cookie.\n")
			} else {
				fmt.Println("\nWrote session_token to config.txt.")
			}
			viper.Set("session_token", *newToken)
			viper.WriteConfig()
		}
	} else if viper.Get("session_token") == "skip" {
		manual = true
		fmt.Println("\nYou have opted against automatic cookie generation and must manually input your iksm_session cookie. You may clear this setting by removing \"skip\" from the session_token field in config.txt.\n")
	}

	var newCookie string
	var accName string
	if manual {
		newCookie = enterCookie()
	} else {
		fmt.Println("Attempting to generate new cookie...")
		accName, newCookie = getCookie(version, client)
	}
	viper.Set("cookie", newCookie)
	viper.WriteConfig()
	if manual {
		fmt.Println("Wrote iksm_session cookie to config.yaml.")
	} else {
		fmt.Println("Wrote iksm_session cookie for " + accName + " to config.yaml.")
	}
}

// Logs in to a Nintendo Account and returns a session_token.
func logIn(version string) *string {
	authStateUnencoded := make([]byte, 36)
	_, err := rand.Read(authStateUnencoded)
	if err != nil {
		panic(err)
	}
	authState := base64.RawURLEncoding.EncodeToString(authStateUnencoded)
	authCodeVerifierUnencoded := make([]byte, 32)
	_, err = rand.Read(authCodeVerifierUnencoded)
	if err != nil {
		panic(err)
	}
	authCodeVerifier := base64.RawURLEncoding.EncodeToString(authCodeVerifierUnencoded)
	authCodeHash := sha256.Sum256([]byte(strings.ReplaceAll(authCodeVerifier, "=", "")))
	authCodeChallenge := base64.RawURLEncoding.EncodeToString(authCodeHash[:])
	appHead := map[string]string{
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
		"state":                               authState,
		"redirect_uri":                        "npf71b963c1b7b6d119://auth",
		"client_id":                           "71b963c1b7b6d119",
		"scope":                               "openid user user.birthday user.mii user.screenName",
		"response_type":                       "session_token_code",
		"session_token_code_challenge":        strings.ReplaceAll(authCodeChallenge, "=", ""),
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
		panic(err)
	}
	q := req.URL.Query()
	for key, element := range body {
		q.Add(key, element)
	}
	req.URL.RawQuery = q.Encode()
	for key, element := range appHead {
		req.Header.Add(key, element)
	}
	client := &http.Client{}
	postLogin := req.URL.String()
	fmt.Println("Navigate to this URL in your browser:")
	fmt.Println(postLogin)
	fmt.Println("Log in, right click the \"Select this account\" button, copy the link address, and paste it below:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	useAccountUrl := scanner.Text()
	re := regexp.MustCompile("de=(.*)&")
	sessionTokenCode := re.FindAllStringSubmatch(useAccountUrl, -1)
	sessionToken := getSessionToken(sessionTokenCode[0][1], authCodeVerifier, client).(string)
	return &sessionToken
}
