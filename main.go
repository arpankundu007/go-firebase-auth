package main

import (
	"go-firebase-auth/auth"
	"go-firebase-auth/users"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main(){
	router := httprouter.New()

	router.Handler(http.MethodPut, "/user/:phone", auth.IsAuthorised(users.UpdateUser()))

	router.Handler(http.MethodGet, "/users/:phone", auth.IsAuthorised(users.GetUserFromPhone()))

	router.Handler(http.MethodGet, "/weather", auth.IsAuthorised(getWeather()))

	router.Handler(http.MethodGet, "/promote/:phone", auth.IsAuthorised(users.PromoteUser()))

	router.Handler(http.MethodGet, "/demote/:phone", auth.IsAuthorised(users.DemoteUser()))

	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}

func getWeather() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		url := "https://samples.openweathermap.org/data/2.5/forecast?q=M%C3%BCnchen,DE&appid=b6907d289e10d714a6e88b30761fae22"

		req, _ := http.NewRequest("GET", url, nil)

		res, _ := http.DefaultClient.Do(req)

		res.Header.Add("Content-Type", "application/json")
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		_, _ = io.WriteString(w, string(body))
	})

}
