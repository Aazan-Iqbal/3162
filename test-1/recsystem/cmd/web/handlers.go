package main

import (
	"net/http"
	//"strconv"
)

// handler for manage equipment
func (app *application) ManageEquipment(w http.ResponseWriter, r *http.Request) {

	RenderTemplate(w, "./ui/static/html/equipment-management.page.tmpl")

}

// create handler for home
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return

	} // w.Write([]byte("Welcome to my home page."))
	RenderTemplate(w, "./ui/static/html/home.page.tmpl")

}

// create handler for about
func (app *application) About(w http.ResponseWriter, r *http.Request) {
	// RenderTemplate(w, "about.page.tmpl")
	// day := time.Now().Weekday()
	// w.Write([]byte(fmt.Sprintf("Welcome to my  about page, have a nice %s", day)))
	w.Write([]byte("Hello\n"))
}

// create handler for login
func (app *application) Login(w http.ResponseWriter, r *http.Request) {

	RenderTemplate(w, "./ui/static/html/home.page.tmpl")

}

// create handler for LoginSubmit
func (app *application) LoginSubmit(w http.ResponseWriter, r *http.Request) {

	RenderTemplate(w, "./ui/static/html/home.page.tmpl")

}

// create handler for SignIn
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {

	RenderTemplate(w, "./ui/static/html/home.page.tmpl")

}

// create handler for SignInSubmit
func (app *application) SignInSubmit(w http.ResponseWriter, r *http.Request) {

	RenderTemplate(w, "./ui/static/html/home.page.tmpl")

}

// create handler for ScanQrCode
func (app *application) ScanQrCode(w http.ResponseWriter, r *http.Request) {

	RenderTemplate(w, "./ui/static/html/home.page.tmpl")

}

// create handler for ScanQrCodeSubmit
func (app *application) ScanQrCodeSubmit(w http.ResponseWriter, r *http.Request) {

	RenderTemplate(w, "./ui/static/html/home.page.tmpl")

}
