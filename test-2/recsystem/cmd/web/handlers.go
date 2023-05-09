package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/Aazan-Iqbal/3161/test-2/recsystem/internal/models"
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
	//remove the entry from the session manager
	flash := app.sessionsManager.PopString(r.Context(), "flash")
	data := &templateData{
		Flash:     flash,
		CSRFTOKEN: nosurf.Token(r), //added for authentication
	}
	RenderTemplate(w, "signup.page.tmpl", data)
}

// run when user submits sign up info
func (app *application) userSignupSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fname := r.PostForm.Get("fname")
	lname := r.PostForm.Get("lname")
	address := r.PostForm.Get("address")
	phone_number := r.PostForm.Get("phone_number")

	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	log.Println(password)

	newUser := models.User{
		Email:        email,
		First_name:   fname,
		Last_name:    lname,
		Dob:          "",
		Address:      address,
		Phone_number: phone_number,
		Roles_id:     1,
		Password:     password,
		CreatedAt:    "",
	}
	//write the data to the table
	err := app.users.Insert(newUser)
	log.Println(err)
	if err != nil {

		if errors.Is(err, models.ErrDuplicateEmail) {
			RenderTemplate(w, "signup.page.tmpl", nil)
		}
	}
	app.sessionsManager.Put(r.Context(), "flash", "Signup was successful")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

// create handler for login
func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {

	flash := app.sessionsManager.PopString(r.Context(), "flash")
	//remove the entry from the session manager
	data := &templateData{
		Flash:     flash,
		CSRFTOKEN: nosurf.Token(r), //added for authentication
	}
	RenderTemplate(w, "login.page.tmpl", data)

}

// create handler for submitting login information
func (app *application) userLoginSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	//write the data to the table
	id, err := app.users.Authenticate(email, password)
	log.Println(err)
	if err != nil {

		if errors.Is(err, models.ErrInvalidCredentials) {
			RenderTemplate(w, "login.page.tmpl", nil)
		}
		return
	}
	//add the user to the session cookie
	err = app.sessionsManager.RenewToken(r.Context())
	if err != nil {
		return
	}
	// add an authenticate entry
	app.sessionsManager.Put(r.Context(), "authenticatedUserID", id)
	http.Redirect(w, r, "/admin/manage-equipment", http.StatusSeeOther)
}

func (app *application) userLogoutSubmit(w http.ResponseWriter, r *http.Request) {
	//remove the entry from the session manager
	err := app.sessionsManager.RenewToken(r.Context())
	if err != nil {
		return
	}
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
