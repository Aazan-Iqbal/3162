package main

import (
	"log"
	"net/http"
	"strconv"
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

// ----------------------HANDLERS FOR MANAGING EQUIPMENT---------------------------
// handler for manage equipment and diplaying equipment list
func (app *application) ManageEquipment(w http.ResponseWriter, r *http.Request) {

	// list, err := app.equipment.Read()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// data := &templateData{
	// 	Equipment: list,
	// }
	RenderTemplate(w, "equipment-management.page.tmpl", nil)

}

func (app *application) AddEquipment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil { // check for errors in parsing form
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	name := r.PostForm.Get("name")
	image := r.PostForm.Get("image")
	equipment_type_id, err := strconv.ParseInt(r.PostForm.Get("equipment_type_id"), 10, 64)
	status, err := strconv.ParseBool(r.PostForm.Get("status"))
	availability, err := strconv.ParseBool(r.PostForm.Get("availability"))
	log.Printf("%s, %s, %v, %v, %v, \n", name, image, equipment_type_id, status, availability)

	err = app.equipment.Insert(name, []byte(image), equipment_type_id, status, availability)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/admin/manage-equipment", http.StatusSeeOther)
}
func (app *application) EditEquipment(w http.ResponseWriter, r *http.Request) {

}
func (app *application) UpdateEquipment(w http.ResponseWriter, r *http.Request) {

}

func (app *application) DeleteEquipment(w http.ResponseWriter, r *http.Request) {

}

//---------------------------END OF EQUIPMENT HANDLERS----------------------------

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
