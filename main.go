package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"log"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	key := os.Getenv("SESSION_KEY")
	maxAge, _ := strconv.Atoi(os.Getenv("SESSION_MAX_AGE"))
	isProd, _ := strconv.ParseBool(os.Getenv("SESSION_SECURE"))

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), os.Getenv("GOOGLE_OAUTH_REDIRECT"), "email", "profile"),
	)

	p := mux.NewRouter()
	p.HandleFunc("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.ParseFiles("resources/success.html")
		t.Execute(res, user)
	})

	p.HandleFunc("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	p.HandleFunc("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})

	p.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("resources/index.html")
		t.Execute(res, false)
	})
	log.Println("listening on 3000")
	log.Fatal(http.ListenAndServe(":3000", p))
}
