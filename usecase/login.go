package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gaurav/golang-jwt-project/helpers"
	"github.com/gaurav/golang-jwt-project/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// create jwt token and send it in response
	w.Header().Set("Content-Type", "application/json")
	// decoding
	var user *models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal("error decoding the request body")
		return
	}

	// check creds
	if user.UserName != "gaurav" || user.Password != "12233" {
		log.Fatal("invalid creds")
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		// create token
		token, err := helpers.CreateJwtToken(user.UserName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("token", token)
	}

}

func Protected(w http.ResponseWriter, r *http.Request) {
	// take token from header and verify it
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "token not provided")
		return
	}

	fmt.Print("token ", token)
	token = token[len("Bearer "):]

	if err := helpers.VerifyJwtToken(token); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "token invalid")
		return
	}
	fmt.Fprint(w, "welcome ! you are a verified user")
}
