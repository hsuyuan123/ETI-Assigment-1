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
type Trip struct {
	TripID        string `json:"TripID"`
	UserID string `json:"User ID"`
	PickupLoc  string `json:"PickupLoc"`
	AltPickupLoc  string    `json:"AltPickupLoc"`
	StartTravelTime     string `json:"StartTravelTime"`
	DestAddress string `json:"DestAddress"`
	NoOfPsgr string `json:"NoOfPsgr"`
	TripStatus string `json:"TripStatus"`
	CustomerID string `json:"CustomerID"`
}

func newTrip(w http.ResponseWriter, r *http.Request) {
	var respond Trip[]
	body, _ := io.ReadAll(r.Body) //from http.Request
	err := json.Unmarshal(body, &respond)
	//Current User ID
	fmt.Printf("Pickup Location: %s", respond.PickupLoc)
	fmt.Printf("Alternate Pickup Location: %s", respond.AltPickupLoc)
	fmt.Printf("Start Travel Time: %s", respond.StartTravelTime)
	fmt.Printf("Destination Address: %s", respond.DestAddress)
	fmt.Printf("Max Number of Passengers: %s", respond.NoOfPsgr)
	CustomerID = nil
	fmt.Printf(err.Error())
}

func enrollTrip(w http.ResponseWriter, r *http.Request){
	//Display all Trips
	//Allow user to choose one
	//Looks at user's current trip StartTravelTime to check for clashes
	//Saves current user id at the trip entry
}