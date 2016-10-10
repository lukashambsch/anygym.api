package handlers

import (
    "net/http"
    "strings"
    "time"

    "github.com/dgrijalva/jwt-go"
)

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

var mySigningKey = []byte("secret")

func SetToken(w http.ResponseWriter, r *http.Request) {
    expireToken := time.Now().Add(time.Hour * 1).Unix()
    expireCookie := time.Now().Add(time.Hour * 1)

    claims := jwt.StandardClaims{
        ExpiresAt: expireToken,
        Issuer: "test",
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    signedToken, _ := token.SignedString(mySigningKey)

    cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
    http.SetCookie(w, &cookie)

    w.Write([]byte(signedToken))
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
    return
}
