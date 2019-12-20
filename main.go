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
		middleware.GoKitIsAuthorised(transport.GetIsAdminEndpoint(svc)),
		transport.IsAdminRequestDecoder,
		httptransport.EncodeJSONResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext))

	changePermissionHandler := httptransport.NewServer(
		middleware.GoKitIsAuthorised(transport.GetChangePermissionEndpoint(svc)),
		transport.ChangePermissionRequestDecoder,
		httptransport.EncodeJSONResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext))

	helloHandler := httptransport.NewServer(
		middleware.GoKitIsAuthorised(svc.HelloWorld()),
		transport.GetHelloWorldDecoder,
		httptransport.EncodeJSONResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext))


	router := httprouter.New()
	router.Handler(http.MethodGet, "/admin", isAdminHandler)
	router.Handler(http.MethodGet, "/hello", helloHandler)
	router.Handler(http.MethodGet, "/update", changePermissionHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
