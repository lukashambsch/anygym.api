package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var mySigningKey = []byte("secret")

func SetToken(email string, password string) (string, error) {
	userMatches, err := datastore.GetUserList(fmt.Sprintf("WHERE email = '%s'", email))
	if err != nil {
		return "", err
	}

	if len(userMatches) == 0 {
		return "", fmt.Errorf("No user matches that email address")
	}
	user := userMatches[0]

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		return "", fmt.Errorf("Invalid Password")
	}

	expireToken := time.Now().Add(time.Hour * 1).Unix()

	claims := jwt.StandardClaims{
		ExpiresAt: expireToken,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString(mySigningKey)

	return signedToken, nil
}

func VerifyToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "authenticate") || r.Method == "OPTIONS" {
			h.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	user := models.UserLogin{}
	err := json.Unmarshal(body, &user)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, APIErrorMessage{Message: err.Error()})
	}

	signedToken, err := SetToken(user.Email, user.Password)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, APIErrorMessage{Message: err.Error()})
	}

	w.Write([]byte(signedToken))
}
