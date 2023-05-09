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
	dynamicMiddleware := alice.New(app.sessionsManager.LoadAndSave)                              //This was edited

	// For test 1 crud
	router.Handler(http.MethodGet, "/admin/manage-equipment", dynamicMiddleware.ThenFunc(app.ManageEquipment))
	router.Handler(http.MethodPost, "/admin/add-equipment", dynamicMiddleware.ThenFunc(app.AddEquipment))
	router.Handler(http.MethodPost, "/admin/edit-equipment", dynamicMiddleware.ThenFunc(app.EditEquipment))
	router.Handler(http.MethodPost, "/admin/update-equipment", dynamicMiddleware.ThenFunc(app.UpdateEquipment))
	router.Handler(http.MethodPost, "/admin/delete-equipment", dynamicMiddleware.ThenFunc(app.DeleteEquipment))

	router.Handler(http.MethodGet, "/user/login", dynamicMiddleware.ThenFunc(app.userLogin))
	router.Handler(http.MethodPost, "/user/login", dynamicMiddleware.ThenFunc(app.userLoginSubmit))
	router.Handler(http.MethodGet, "/user/sign-up", dynamicMiddleware.ThenFunc(app.userSignup))
	router.Handler(http.MethodPost, "/user/sign-up", dynamicMiddleware.ThenFunc(app.userSignupSubmit))
	router.Handler(http.MethodPost, "/user/logout", dynamicMiddleware.ThenFunc(app.userLogoutSubmit))

	router.Handler(http.MethodGet, "/", dynamicMiddleware.ThenFunc(app.Home))
	router.Handler(http.MethodGet, "/about", dynamicMiddleware.ThenFunc(app.About))

	router.Handler(http.MethodGet, "/sign-in", dynamicMiddleware.ThenFunc(app.SignIn))
	router.Handler(http.MethodPost, "/sign-in-auth", dynamicMiddleware.ThenFunc(app.SignInSubmit))
	router.Handler(http.MethodGet, "/scan-qr-code", dynamicMiddleware.ThenFunc(app.ScanQrCode))
	router.Handler(http.MethodPost, "/scan-qr-code-check", dynamicMiddleware.ThenFunc(app.ScanQrCodeSubmit))

	//tidy up the middleware chain
	standardMiddleware := alice.New(app.RecoverPanicMiddleware, //new function
		app.logRequestMiddleware,
		securityHeadersMiddleware)

	return standardMiddleware.Then(router)
}
