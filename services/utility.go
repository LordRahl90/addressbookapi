package services

import (
	"encoding/json"
	"log"
)

//Response Struct
type Response struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Message     string `json:"message"`
}

//Data Interface
type Data interface{}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func returnError(err error) string {
	if err != nil {
		log.Println(err)
		response := Response{"error", "An error occurred", "We are looking into it right away!"}
		b, issues := json.Marshal(response)
		if issues != nil {
			log.Println("An issue occurred while reporting the error")
		}
		return string(b)
	}
	return ""
}

func (response *Response) returnResponse(status, description, message string) string {
	response.Status = status
	response.Description = description
	response.Message = message

	b, err := json.Marshal(response)
	checkError(err)
	return string(b)
}
