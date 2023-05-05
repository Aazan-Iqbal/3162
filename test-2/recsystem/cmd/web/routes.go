package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//create multiplexer
	router := httprouter.New()
	// create file server
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer)) //exclude resource and go to static


	// For test 1 crud
	router.HandlerFunc(http.MethodGet, "/admin/manage-equipment", app.ManageEquipment)


	router.HandlerFunc(http.MethodGet, "/", app.Home)
	router.HandlerFunc(http.MethodGet, "/about", app.About)

	router.HandlerFunc(http.MethodGet, "/login", app.Login)
	router.HandlerFunc(http.MethodPost, "/login-auth", app.LoginSubmit)
	router.HandlerFunc(http.MethodGet, "/sign-in", app.SignIn)
	router.HandlerFunc(http.MethodPost, "/sign-in-auth", app.SignInSubmit)
	router.HandlerFunc(http.MethodGet, "/scan-qr-code", app.ScanQrCode)
	router.HandlerFunc(http.MethodPost, "/scan-qr-code-check", app.ScanQrCodeSubmit)

	return router
}
