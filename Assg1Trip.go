package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Link to user and driver
func main() {
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db")

	if err != nil {
		panic(err.Error())
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/user", newUser).Methods("GET", "DELETE", "POST", "PATCH", "PUT")
	//User
	router.HandleFunc("/api/v1/user", User)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
	//Driver
	router.HandleFunc("/api/v1/driver", Driver)
	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
	//Trip
	router.HandleFunc("/api/v1/trip", Trip)
	fmt.Println("Listening at port 5002")
	log.Fatal(http.ListenAndServe(":5002", router))
}

type Trip struct {
	TripID          string `json:"TripID"`
	DriverId        string `json:"User ID"`
	PickupLoc       string `json:"PickupLoc"`
	AltPickupLoc    string `json:"AltPickupLoc"`
	StartTravelTime string `json:"StartTravelTime"`
	DestAddress     string `json:"DestAddress"`
	NoOfPsgr        string `json:"NoOfPsgr"`
	TripStatus      string `json:"TripStatus"`
	CustomerID      string `json:"CustomerID"`
}

func newTrip(id string, t Trip, u User, d Driver) {
	_, err := db.Exec("insert into trip values(?,?,?,?,?)", id, d.DriverId, t.PickupLoc, t.AltPickupLoc, t.StartTravelTime, t.DestAddress, t.NoOfPsgr, t.TripStatus, u.UserId)
	if err != nil {
		panic(err.Error())
	}
}

func updateTrip(id string, t Trip) {
	_, err := db.Exec("update course set name=?, intake=?, mingpa=?, maxgpa=? where id=?", t.PickupLoc, t.AltPickupLoc, t.StartTravelTime, t.DestAddress, t.NoOfPsgr, id)
	if err != nil {
		panic(err.Error())
	}
}

//Enroll by changing userid of trip to id of current logged in user
