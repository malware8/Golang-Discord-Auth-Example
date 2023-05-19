package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	dsc "github.com/realTristan/disgoauth"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	fmt.Println(clientID, clientSecret)
	// client
	ds := dsc.Init(&dsc.Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  "http://localhost:8000/callback",
		Scopes:       []string{dsc.ScopeIdentify, dsc.ScopeEmail},
	})

	// default handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ds.RedirectHandler(w, r, "")
	})

	// callback handler
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {

		code := r.URL.Query()["code"][0]

		token, _ := ds.GetOnlyAccessToken(code)

		result, _ := dsc.GetUserData(token)

		fmt.Println(result)
	})

	http.ListenAndServe(":8000", nil)
}
