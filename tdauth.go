package main

import (
	"html/template"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/browser"
)

var oauthChan chan string = make(chan string)
var accessToken string
var refreshToken string
var clientKey string
var clientCode string
var callbackAddress string
var callbackURL string
var authURL string

func tdauth() {
	clientKey = config.TD.ClientKey
	clientCode = clientKey + "@AMER.OAUTHAP"
	callbackAddress = config.TD.CallbackAddress
	callbackURL = "https://" + callbackAddress
	authURL = "https://auth.tdameritrade.com/auth?response_type=code&redirect_uri=" + callbackURL + "&client_id=" + clientCode

	browser.OpenURL(authURL)

	go startServer(callbackAddress)

	oauth := <-oauthChan //Holding call]]
	var response RequestTokensResponse

	requestTokens(oauth, clientCode, &response)

	accessToken = response.AccessToken
	refreshToken = response.RefreshToken

}

func startServer(callbackAddress string) {

	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler).Methods("GET")

	listen, err := net.Listen("tcp", callbackAddress)
	errCheck("Error creating web server", err)

	err = http.ServeTLS(listen, mux, "certs/localhost.crt", "certs/localhost.key")
	errCheck("Error serveing web server", err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	type pageData struct {
		AppName string
		AuthURL string
		Authed  bool
	}
	data := pageData{
		AppName: "GoStocks-Discord",
		AuthURL: authURL,
		Authed:  true,
	}
	templates := template.Must(template.ParseGlob("templates/*.html"))

	oauth := r.URL.Query().Get("code")
	if oauth != "" || accessToken != "" {
		oauthChan <- oauth
		data.Authed = false
	}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
