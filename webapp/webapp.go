package webapp

import (
	"html/template"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

//OauthChan channel oauth is sent through
var OauthChan chan string = make(chan string)
var oauthURL string
var oathSent bool

//Start start webapp
func Start(callbackAddress string, authURL string) {
	oathSent = false
	oauthURL = authURL

	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler).Methods("GET")

	listen, err := net.Listen("tcp", callbackAddress)
	if err != nil {
		panic(err)
	}

	err = http.ServeTLS(listen, mux, "webapp/certs/localhost.crt", "webapp/certs/localhost.key")
	if err != nil {
		panic(err)
	}
}

type pageData struct {
	AppName string
	AuthURL string
	Authed  bool
}

var data = pageData{
	AppName: "GoStocks-Discord",
	AuthURL: oauthURL,
	Authed:  false,
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	templates := template.Must(template.ParseGlob("webapp/templates/*.html"))

	oauth := r.URL.Query().Get("code")
	if oauth != "" && oathSent == false {
		OauthChan <- oauth
		data.Authed = true
		oathSent = true
	}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
