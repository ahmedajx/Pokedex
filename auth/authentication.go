package auth

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

/* Set up a global string for our secret */
var mySigningKey = []byte("secret")

type Token struct {
	Token string `json:"token"`
}

//http://www.alexedwards.net/blog/making-and-using-middleware
func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		//check if key Authorized exists in header and value has a valid token.
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
	})
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
