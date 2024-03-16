package user

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Struct untuk representasi user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Fungsi untuk membuat koneksi ke database
func connectDB() (*sql.DB, error) {
	dbURL := "postgresql://postgres:IOIvlMqzIeTraabYeMaZfbIJtmYZbXyU@roundhouse.proxy.rlwy.net:26710/railway"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Fungsi untuk membuat user baru
func createUser(db *sql.DB, user User) error {
	_, err := db.Exec("INSERT INTO customers (username, email) VALUES ($1, $2)", user.Username, user.Email)
	if err != nil {
		return err
	}
	return nil
}

// Fungsi untuk membaca user berdasarkan ID
func getUserByID(db *sql.DB, userID int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, email FROM customers WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Fungsi untuk memperbarui informasi user
func updateUser(db *sql.DB, user User) error {
	_, err := db.Exec("UPDATE customers SET username = $1, email = $2 WHERE id = $3", user.Username, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// Fungsi untuk menghapus user berdasarkan ID
func deleteUser(db *sql.DB, userID int) error {
	_, err := db.Exec("DELETE FROM customers WHERE id = $1", userID)
	if err != nil {
		return err
	}
	return nil
}

func Crud() {
	// Membuat koneksi ke database
	db, err := connectDB()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Membuat sebuah user baru
	newUser := User{Username: "john_doe", Email: "john@example.com"}
	err = createUser(db, newUser)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println("User created successfully!")

	// Membaca user berdasarkan ID
	userID := 1
	user, err := getUserByID(db, userID)
	if err != nil {
		fmt.Println("Error getting user:", err)
		return
	}
	fmt.Println("User:", user)

	// Memperbarui informasi user
	user.Username = "updated_username"
	user.Email = "updated_email@example.com"
	err = updateUser(db, user)
	if err != nil {
		fmt.Println("Error updating user:", err)
		return
	}
	fmt.Println("User updated successfully!")

	// Menghapus user berdasarkan ID
	err = deleteUser(db, userID)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return
	}
	fmt.Println("User deleted successfully!")

}
