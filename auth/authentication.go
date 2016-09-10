package auth

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

/* Set up a global string for our secret */
var mySigningKey = []byte("secret")

type Token struct {
	Token string `json:"token"`
}

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/* Set token claims */
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(mySigningKey)
	jwtToken := Token{tokenString}
	b, _ := json.Marshal(jwtToken)
	w.Write(b)
}
