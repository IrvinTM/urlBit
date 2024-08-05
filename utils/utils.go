package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	//TODO add error management
	json.NewEncoder(w).Encode(data)
}

func GenShort() string{
	var chars = []string {"a", "b", "c", "d", "e", "f", "g", "h", "i","j", "k", "l", "m", "n", "o", "p", "q", "r","s", "t", "v", "w", "x", "y", "z", "1", "2","3", "4", "5", "6", "7", "8", "9", "0"}
	random := ""
	for i := 0; i < 7; i++ {
		random += chars[rand.Intn(len(chars))]
	}
	return random
}
