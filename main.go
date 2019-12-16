package main

import (
	"firebase_auth/auth"
	"firebase_auth/users"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	router := httprouter.New()

	router.Handler(http.MethodPut, "/user/:phone", auth.IsAuthorised(users.UpdateUser()))

	router.Handler(http.MethodGet, "/users/:phone", auth.IsAuthorised(users.GetUserFromPhone()))

	router.Handler(http.MethodGet, "/weather", auth.IsAuthorised(getWeather()))
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getWeather() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		url := "https://api.darksky.net/forecast/9702fb99b96c28f5e83c814002d8f2cf/37.8267,-122.4233"

		req, _ := http.NewRequest("GET", url, nil)

		res, _ := http.DefaultClient.Do(req)

		res.Header.Add("Content-Type", "application/json")
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		_, _ = io.WriteString(w, string(body))
	})

}
