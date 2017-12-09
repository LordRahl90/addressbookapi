package handlers

import (
	"fmt"
	"net/http"

	"github.com/LordRahl90/addressapi/services"
)

//Index this function lists all the availaible contacts
func Index(w http.ResponseWriter, r *http.Request) {
	response := services.AllContacts()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, response)
}

//Create function
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//we render the remplate at this point
		fmt.Fprintf(w, "Sorry... you cant access this with GET")
	} else if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		firstname := r.FormValue("firstname")
		middlename := r.FormValue("middlename")
		surname := r.FormValue("surname")
		phone := r.FormValue("phone")
		email := r.FormValue("email")

		contact := services.Contact{
			Firstname:  firstname,
			Middlename: middlename,
			Surname:    surname,
			Phone:      phone,
			Email:      email,
		}

		result := contact.CreateContact()
		fmt.Fprintf(w, result)

	} else {
		fmt.Fprintf(w, "Unidentified request")
	}

}
