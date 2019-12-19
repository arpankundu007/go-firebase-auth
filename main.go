package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
	"go-firebase-auth/middleware"
	"go-firebase-auth/transport"
	"go-firebase-auth/usecase"
	"log"
	"net/http"
)


func main(){

	svc := usecase.AuthServiceInstance{}

	isAdminHandler := httptransport.NewServer(
		middleware.GoKitIsAuthorised()(transport.GetIsAdminEndpoint(svc)),
		transport.IsAdminRequestDecoder,
		httptransport.EncodeJSONResponse)

	changePermissionHandler := httptransport.NewServer(
		transport.GetChangePermissionEndpoint(svc),
		transport.ChangePermissionRequestDecoder,
		httptransport.EncodeJSONResponse)

	weatherHandler := httptransport.NewServer(
		middleware.GoKitIsAuthorised()(svc.HelloWorld()),
		transport.GetWeatherDecoder,
		httptransport.EncodeJSONResponse)


	router := httprouter.New()
	router.Handler(http.MethodGet, "/admin", isAdminHandler)
	router.Handler(http.MethodGet, "/weather", weatherHandler)
	router.Handler(http.MethodGet, "/update/:phone/:admin", middleware.IsAuthorized(changePermissionHandler))

	log.Fatal(http.ListenAndServe(":8080", router))
}
