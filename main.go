package main

import (
	"database/sql"
	"fmt"
	"sanbercode-final-project-Sahrul-Yoyo/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
)

var mySigningKey = []byte("secret")

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {

	//KONEKSI DATABASE RAILWAY.APP ------------------------------------------->
	dbURL := "postgresql://postgres:IOIvlMqzIeTraabYeMaZfbIJtmYZbXyU@roundhouse.proxy.rlwy.net:26710/railway"

	// Membuka koneksi ke database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Memeriksa apakah koneksi berhasil
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}

	fmt.Println("Connected to the database successfully!")

	// MEMBUAT TOKEN ------------------------------------------------------->
	token := jwt.New(jwt.SigningMethodHS256)

	// Set klaim
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "John Doe"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign token dengan kunci rahasia
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("error creating token:", err)
		return
	}

	fmt.Println("token:", tokenString)

	// Verifikasi token yang diterima
	receivedToken := tokenString //"token_yang_diterima_dari_klien"
	token, err = jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		// Verifikasi tipe algoritma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("claims:")
		fmt.Println("authorized:", claims["authorized"])
		fmt.Println("user:", claims["user"])
	} else {
		fmt.Println("invalid token:", err)
	}

	//CRUD OPERATION -------------------------------------------------------------------->
	user.Crud()
}
