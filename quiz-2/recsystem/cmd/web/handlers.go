package main

import (
	"net/http"
	//"strconv"
)

// include --about
// include --home

// create handler for home
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return

	} // w.Write([]byte("Welcome to my home page."))
	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for about
func (app *application) About(w http.ResponseWriter, r *http.Request) {
	// RenderTemplate(w, "about.page.tmpl", nil)
	// day := time.Now().Weekday()
	// w.Write([]byte(fmt.Sprintf("Welcome to my  about page, have a nice %s", day)))
	w.Write([]byte("Hello\n"))
}

// create handler for login
func (app *application) Login(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for LoginSubmit
func (app *application) LoginSubmit(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for SignIn
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for SignInSubmit
func (app *application) SignInSubmit(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for ScanQrCode
func (app *application) ScanQrCode(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for ScanQrCodeSubmit
func (app *application) ScanQrCodeSubmit(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for SignIn
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for SignIn
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for SignIn
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for SignIn
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for SignIn
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for SignIn
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for SignIn
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}
