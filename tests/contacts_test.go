package tests

import (
	"encoding/json"
	"exercises/addressapi/services"
	"testing"
)

func TestCreateContact(t *testing.T) {
	var response services.Response
	contacts := services.Contact{}
	result := contacts.CreateContact()
	json.Unmarshal([]byte(result), &response)

	if response.Status != "success" {
		t.Fatal("Test Failed Woefully ", result)
	}
	t.Log("Creating User Passed Successfully")
}
