package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	//TODO add error management
	json.NewEncoder(w).Encode(data)
}

func GenShort() string {
    const chars = "abcdefghijklmnopqrstuvwxyz1234567890"
    var builder strings.Builder
    for i := 0; i < 7; i++ {
        builder.WriteByte(chars[rand.Intn(len(chars))])
    }
    return builder.String()
}
