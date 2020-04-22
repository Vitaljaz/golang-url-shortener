package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/recoilme/slowpoke"
	"github.com/speps/go-hashids"
)

var file = "db/data.db"

func BuildShortURL(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	longURL := strings.Join(params["url"], "")
	_, err := http.Get(longURL)

	if err != nil {
		fmt.Fprintf(w, "Bad url! example: https://www.google.com/")
	} else {
		hash := GetHash(longURL)
		fmt.Println("[Build]: " + longURL + " => " + GetShortURL(hash))
		SaveToDB(hash, longURL)
		fmt.Fprintf(w, "Short url: "+GetShortURL(hash))
	}
}

func Redirect(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	longURL, err := GetFromDB(params["hash"])
	fmt.Println("[Redirect]: " + GetShortURL(params["hash"]) + " => " + longURL)

	if err != nil {
		fmt.Fprintf(w, "Bad url!")
	} else {
		http.Redirect(w, req, longURL, 301)
	}
}

func MainPage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "To shorten the URL, please create a GET request in this form:")
	fmt.Fprintln(w, "TEMPLATE: http://localhost:9090/build?url=LONG_URL")
	fmt.Fprintln(w, "FOR EXAMPLE: http://localhost:9090/build?url=https://www.google.com/")
}

func GetHash(longURL string) string {
	hd := hashids.NewData()
	hd.Salt = longURL
	h, _ := hashids.NewWithData(hd)
	id, _ := h.Encode([]int{1, 2, 3})
	return id
}

func GetShortURL(hash string) string {
	return "http://localhost:9090/" + hash
}

func main() {
	defer slowpoke.CloseAll()

	router := mux.NewRouter()
	router.HandleFunc("/", MainPage).Methods("GET")
	router.HandleFunc("/build", BuildShortURL).Methods("GET")
	router.HandleFunc("/{hash}", Redirect).Methods("GET")
	http.ListenAndServe(":9090", router)
}
