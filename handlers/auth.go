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

type Route struct {
	Path   string
	Method string
}

var noAuthRoutes []Route = []Route{
	Route{Path: "api/v1/users", Method: "POST"},
	Route{Path: "api/v1/authenticate", Method: "POST"},
	Route{Path: "", Method: "OPTIONS"},
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
		if isNoAuth(r) {
			h.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		splitHeader := strings.Split(authHeader, "Bearer ")
		if len(splitHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := splitHeader[1]
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

func isNoAuth(r *http.Request) bool {
	for _, route := range noAuthRoutes {
		if strings.Contains(r.URL.Path, route.Path) && r.Method == route.Method {
			return true
		}
	}

	return false
}

func Logout(w http.ResponseWriter, r *http.Request) {
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	user := models.User{}
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
