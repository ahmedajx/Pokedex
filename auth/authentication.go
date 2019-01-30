package auth

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
	"strings"
)

var mySigningKey = []byte("Ash_Ketchum")

type Token struct {
	Token string `json:"token"`
}

//http://www.alexedwards.net/blog/making-and-using-middleware
func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if (len(authorization) == 0) {
			w.WriteHeader(http.StatusUnauthorized)
			return;
		}

		tokenString := strings.Fields(authorization)[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		})
		
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return;
		}
		
		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized);
		}
		
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
