package main

import (
	"net/http"

	"github.com/justinas/nosurf"
	//"strconv"
)

// create handler for home
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return

	}
	RenderTemplate(w, "home.page.tmpl", nil)

}

// handler for manage equipment
func (app *application) ManageEquipment(w http.ResponseWriter, r *http.Request) {

	flash := app.sessionsManager.PopString(r.Context(), "flash")

	data := &templateData{
		Flash:     flash,
		CSRFTOKEN: nosurf.Token(r), //added for authentication
	}
	RenderTemplate(w, "equipment-management.page.tmpl", data)

}

// --------------------sign up, log in, and out functionality----------------------
// for user sign up
func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {

}

// run when user submits sign up info
func (app *application) userSignupSubmit(w http.ResponseWriter, r *http.Request) {

}

// create handler for login
func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {

}

// create handler for submitting login information
func (app *application) userLoginSubmit(w http.ResponseWriter, r *http.Request) {

}

func (app *application) userLogoutSubmit(w http.ResponseWriter, r *http.Request) {

}

// ---------------------------------------------------------------------------
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {

}

func (app *application) SignInSubmit(w http.ResponseWriter, r *http.Request) {

}

// create handler for ScanQrCode
func (app *application) ScanQrCode(w http.ResponseWriter, r *http.Request) {

}

// create handler for ScanQrCodeSubmit
func (app *application) ScanQrCodeSubmit(w http.ResponseWriter, r *http.Request) {

}

// create handler for about
func (app *application) About(w http.ResponseWriter, r *http.Request) {
}
