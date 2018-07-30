package Commands

import (
	"github.com/gin-gonic/gin"
	"log"
	"encoding/json"
)


type Response struct {
	Valid	bool `json:"valid"`
	Message	string	`json:"message"`
}

func ReturnErr(w gin.ResponseWriter, err error)  {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422)

	error :=map[string]string{
		"error": err.Error(),
	}
	log.Println("[ERROR] " + err.Error())

	if err := json.NewEncoder(w).Encode(error); err != nil {
		log.Println("[ERROR] " + err.Error())
	}
	panic(err)
}

func MakeResponse(responseWriter gin.ResponseWriter, response Response, status int){

	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(status)
	if err := json.NewEncoder(responseWriter).Encode(response); err != nil {
		log.Println(err)
		panic(err)
	}
}
