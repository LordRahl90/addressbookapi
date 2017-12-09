package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Contact Struct
type Contact struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Firstname  string        `json:"firstname"`
	Middlename string        `json:"middlename"`
	Surname    string        `json:"surname"`
	Phone      string        `json:"phone"`
	Email      string        `json:"email"`
}

//Manager Interface The data manager interface
type Manager interface {
	CleanData()
	VerifyData()
	SanitizeData()
}

//CreateContact function
func (contact Contact) CreateContact() string {

	var response Response
	// session, err := mgo.Dial("127.0.0.1")
	session, err := mgo.Dial("addressapi_database_1")
	resp := returnError(err)
	if strings.Compare(resp, "") != 0 {
		return resp
	}

	session.SetMode(mgo.Monotonic, true)
	//lets set index
	collection := session.DB(os.Getenv("DB_NAME")).C("contacts")
	index := mgo.Index{
		Key:        []string{"phone", "email"},
		DropDups:   true,
		Unique:     true,
		Sparse:     true,
		Background: true,
	}
	collection.EnsureIndex(index)

	//lets make sure this record doesnt exist already
	recordCount, err := collection.Find(bson.M{"$or": []bson.M{bson.M{"email": contact.Email}, bson.M{"phone": contact.Phone}}}).Count()
	resp = returnError(err)
	if strings.Compare(resp, "") != 0 {
		return resp
	}

	if recordCount <= 0 {
		err := collection.Insert(contact)
		resp = returnError(err)
		if strings.Compare(resp, "") != 0 {
			return resp
		}

		log.Println("New Record Created Successfully")
	} else {
		b := response.returnResponse("error", "Record Exists", "Duplicate Record, as this record exists already!")
		return b
	}

	b := response.returnResponse("success", "Successful", "Record Created Successfully!!")
	return b
}

//FindContact function
func (contact Contact) FindContact() {
	fmt.Println(contact)
}

//Contacts Multiple contacts
type Contacts struct {
	Contacts []Contact
}

//AllContacts function
func AllContacts() string {
	fmt.Println("Loading all contacts")
	contacts := []Contact{}
	var response Response

	session, err := mgo.Dial("addressapi_database_1")
	resp := returnError(err)
	if strings.Compare(resp, "") != 0 {
		return resp
	}

	session.SetMode(mgo.Monotonic, true)
	collections := session.DB(os.Getenv("DB_NAME")).C("contacts")

	err = collections.Find(nil).All(&contacts)
	resp = returnError(err)
	if strings.Compare(resp, "") != 0 {
		return resp
	}

	if len(contacts) <= 0 {
		b := response.returnResponse("success", "Successful", "Record Created Successfully!!")
		return b
	}

	b, err := json.Marshal(contacts)
	return string(b)
}
