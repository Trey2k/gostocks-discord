package td

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/Trey2k/gostocks-discord/webapp"
	"github.com/pkg/browser"
)

var accessToken string
var refreshToken string
var clientKey string
var clientCode string

//CallbackAddress callback address
var CallbackAddress string

//AuthURL URL to auth
var AuthURL string

//Init run before auth
func Init() {
	clientKey = utils.Config.TD.ClientKey
	clientCode = clientKey + "@AMER.OAUTHAP"
	CallbackAddress = utils.Config.TD.CallbackAddress
	callbackURL := "https://" + CallbackAddress
	AuthURL = "https://auth.tdameritrade.com/auth?response_type=code&redirect_uri=" + callbackURL + "&client_id=" + clientCode
}

//Auth Authenticate
func Auth() {

	token, err := checkRefreshToken()
	utils.ErrCheck("Error checking saved token", err)
	if !token {
		browser.OpenURL(AuthURL)

		oauth := <-webapp.OauthChan //Holding call]]
		var response RequestTokensResponse

		err := requestTokens(oauth, clientCode, &response)
		utils.ErrCheck("Error requesting TD Tokens", err)

		accessToken = response.AccessToken
		refreshToken = response.RefreshToken
		err = saveRefreshToken(refreshToken)
		utils.ErrCheck("Error saveing refresh token", err)
	}

	fmt.Println("TD Ameritrade authenticated.")
}

func saveRefreshToken(refreshToken string) error {
	if _, err := os.Stat("./info"); os.IsNotExist(err) {
		os.Mkdir("./info", 0700)
	}
	file, err := os.Create("./info/refresh.token")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, refreshToken)
	if err != nil {
		return err
	}
	return file.Sync()
}

func checkRefreshToken() (bool, error) {
	if utils.FileExists("./info/refresh.token") {
		data, err := ioutil.ReadFile("./info/refresh.token")
		if err != nil {
			return false, err
		}
		token := string(data)
		var response RequestTokensResponse
		err = refreshTokens(token, clientCode, &response)
		if err != nil || response.AccessToken == "" {
			return false, nil
		}
		accessToken = response.AccessToken
		refreshToken = token
		return true, nil
	}
	return false, nil

}
