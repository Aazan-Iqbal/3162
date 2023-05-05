package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	//create multiplexer
	router := httprouter.New()
	// create file server
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer)) //exclude resource and go to static
	dynamicMiddleware := alice.New(app.sessionsManager.LoadAndSave, noSurf)                      //This was edited

	// For test 1 crud
	router.Handler(http.MethodGet, "/admin/manage-equipment", dynamicMiddleware.ThenFunc(app.ManageEquipment))

	// for test-2
	router.Handler(http.MethodGet, "/user/login", dynamicMiddleware.ThenFunc(app.Login))
	router.Handler(http.MethodPost, "/user/login-auth", dynamicMiddleware.ThenFunc(app.LoginSubmit))
	router.Handler(http.MethodPost, "/user/logout", dynamicMiddleware.ThenFunc(app.userLogoutSubmit))

	router.HandlerFunc(http.MethodGet, "/", app.Home)
	router.HandlerFunc(http.MethodGet, "/about", app.About)

	router.HandlerFunc(http.MethodGet, "/sign-in", app.SignIn)
	router.HandlerFunc(http.MethodPost, "/sign-in-auth", app.SignInSubmit)
	router.HandlerFunc(http.MethodGet, "/scan-qr-code", app.ScanQrCode)
	router.HandlerFunc(http.MethodPost, "/scan-qr-code-check", app.ScanQrCodeSubmit)

	//tidy up the middleware chain
	standardMiddleware := alice.New(app.RecoverPanicMiddleware, //new function
		app.logRequestMiddleware,
		securityHeadersMiddleware)

	return standardMiddleware.Then(router)
}
