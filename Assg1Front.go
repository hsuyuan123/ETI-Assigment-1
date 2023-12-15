package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
) 
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/user", newUser).Methods("POST")//More methods GET

	//router.HandleFunc("/api/v1/user", users)
	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
} 
func console(){
	fmt.Printf("1. Create New User")
	fmt.Printf("2. Login")
	input, _ := reader.ReadString('\n')
	if input == 1{
		newUser()
	} else if input == 2{
		fmt.Printf("Input Email")
		input2, _ := reader.ReadString('\n')
		results, err := db.Query("select * from User")
		if err != nil {
			panic(err.Error())
		}
		for results.Next() {
			var u User
			err = results.Scan(&u.ID, &u.FirstName, &u.LastName, &u.MobileNo, &u.Email, &u.LicenseNo, &u.VehicleNo, &u.AccStatus)
			if err != nil {
				panic(err.Error())
			}
		}
		for r in results {
			if r.Email == input2{
				fmt.Printf("Hello" + r.FirstName + r.LastName)

			}
		}
		//Logged in
		fmt.Printf("1. Update User")
		fmt.Printf("2. Delete User")
		fmt.Printf("3. Enroll in Trip")
		if r.LicenseNo != nil(
			fmt.Printf("4. Create Trip")
		)
		input3, _ := reader.ReadString('\n')
		if input3 == 1 {
			updateUser()
		} else if input3 == 2 {
			deleteUser()
		} else if input3 == 3{
			enrollTrip()
		} else if input3 == 4 {
			newTrip()
		}	
	} else {
		fmt.Printf("Input a valid number.")
	}
}