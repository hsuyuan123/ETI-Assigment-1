package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Driver struct {
	ID        string `json:"ID"`
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	MobileNo  int    `json:"MobileNo"`
	Email     string `json:"Email"`
	LicenseNo string `json:"LicenseNo"`
	VehicleNo string `json:"VehicleNo"`
	AccStatus string `json:"AccStatus"`
}

func GetData(db *sql.DB) {
	results, err := db.Query("select * from Driver")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var d Driver
		err = results.Scan(&d.ID, &d.FirstName, &d.LastName, &d.MobileNo, &d.Email, &d.LicenseNo, &d.VehicleNo, &d.AccStatus)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(d.ID, d.FirstName, d.LastName, d.MobileNo, d.Email)
	}
}

func newUser(w http.ResponseWriter, r *http.Request) {
	var respond Driver
	body, _ := io.ReadAll(r.Body) //from http.Request
	err := json.Unmarshal(body, &respond)
	fmt.Printf("First Name: %s", respond.FirstName)
	fmt.Printf("Last Name: %s", respond.LastName)
	fmt.Printf("Mobile Number: %s", respond.MobileNo)
	fmt.Printf("Email: %s", respond.Email)
	fmt.Printf("License Number: %s", respond.LicenseNo)
	fmt.Printf("Vehicle Number: %s", respond.VehicleNo)
	fmt.Printf(err.Error())
}

func users(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")
	defer db.Close()

	results, _ := db.Query("select * from users")

	var users map[string]User = map[string]User{} //////ADDED
	for results.Next() {
		var u User
		_ = results.Scan(&u.ID, &u.FirstName, &u.LastName, &u.MobileNo, &u.Email, &u.LicenseNo, &u.VehicleNo, &u.AccStatus)

		fmt.Println(u)
		users[u.ID] = u //////ADDED
	}
	json.NewEncoder(w).Encode(users) //////ADDED
}

func updateUser(db *sql.DB) {
	var u User
	result, err := db.Exec("insert into user (firstname, lastname,  mobileno, email)     values(?, ?, ?, ?, ?, ?)", u.FirstName, u.LastName, u.MobileNo, u.Email, &u.LicenseNo, &u.VehicleNo, &u.AccStatus)
	if err != nil {
		panic(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		print(id)
		panic(err.Error())
	}
}

func deleteUser() {
	params := mux.Vars(r)

	data, ok := users[params["email"]]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No user found")
	} else {
		fmt.Println(data)
		delete(users, params["email"])
		fmt.Fprintf(w, "User deleted")
	}
}
