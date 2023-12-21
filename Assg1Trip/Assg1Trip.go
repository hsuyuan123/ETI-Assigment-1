package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Link to user and driver
func main() {
	db, err = sql.Open("mysql", "user1:password@tcp(127.0.0.1:3306)/my_db")

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
	UserId      string `json:"UserId"`
}



func delUser(userid string) (int64, error) {
	result, err := db.Exec("delete from user where userid=?", userid)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func trip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Trip
			fmt.Println(string(body))
			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist(params["tripid"]); !ok {
					fmt.Println(data)
					//courses[params["courseid"]] = data
					insertTrip(params["tripid"], data)

					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Fprintf(w, "Trip ID exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Trip

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist(params["tripid"]); ok {
					fmt.Println(data)
					//courses[params["courseid"]] = data
					updateTrip(params["tripid"], data)
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Trip ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PATCH" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data map[string]interface{}

			if err := json.Unmarshal(body, &data); err == nil {
				if orig, ok := isExist(params["tripid"]); ok {
					fmt.Println(data)

					for k, v := range data {
						switch k {
						case "PickupLoc":
							orig.Name = v.(string)
						case "AltPickupLoc":
							orig.Intake = v.(float64)
						case "StartTravelTime":
							orig.MinGPA = time(v.(float64))
						case "DestAddress":
							orig.DestAddress = v.(float64)
						case "NoOfPsgr":
							orig.NoOfPsgr = int(v.(float64))
						}
						}
						
					}
					updateTrip(params["tripid"], orig)
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Trip ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if val, ok := isExist(params["tripid"]); ok {
		if r.Method == "DELETE" {
			fmt.Fprintf(w, params["tripid"]+" Deleted")
			delCourse(params["tripid"])
		} else {
			json.NewEncoder(w).Encode(val)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Trip ID")
	}
}
func alltrips(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	found := false
	results := map[string]Trip{}

	if value := query.Get("q"); len(value) > 0 {
		for k, v := range trips {
			if strings.Contains(strings.ToLower(v.tripid), strings.ToLower(value)) {
				results[k] = v
				found = true
			}
		}

		if !found {
			fmt.Fprintf(w, "No trips found")
		} else {
			json.NewEncoder(w).Encode(struct {
				Results map[string]User `json:"Search Results"`
			}{results})
		}
	} else if value = query.Get("starttraveltime"); len(value) > 0 {
		found := false
		results := map[string]User{}
		starttraveltime, _ := strconv.Atoi(value)

		for k, v := range trip {
			//Check for 30 mins
			if time.Now.Sub(time.starttraveltime) => 30 {
				results[k] = v
				found = true
			}
		}

		if !found {
			fmt.Fprintf(w, "No trip eligible to enroll")
		} else {
			json.NewEncoder(w).Encode(struct {
				Results map[string]Course `json:"Eligible trip(s)"`
			}{results})
		}
	} else {
		tripsWrapper := struct {
			Trips map[string]Trip `json:"Trips"`
		}{trips}
		json.NewEncoder(w).Encode(tripsWrapper)
		return
	}

	func getTrips() map[string]Trip {
		results, err := db.Query("select * from Trip")
		if err != nil {
			panic(err.Error())
		}
	
		var trips map[string]Trip = map[string]Trip{}
	
		for results.Next() {
			var t Trips
			var id string
			err = results.Scan(&id, &d.DriverId, &t.PickupLoc, &t.AltPickupLoc, &t.StartTravelTime, &t.DestAddress, &t.NoOfPsgr, &u.UserId)
			if err != nil {
				panic(err.Error())
			}
	
			trips[id] = t
		}
	
		return trips
	}

	func newTrip(id string, t Trip, u User, d Driver) {
		_, err := db.Exec("insert into trip values(?,?,?,?,?)", id, d.DriverId, t.PickupLoc, t.AltPickupLoc, t.StartTravelTime, t.DestAddress, t.NoOfPsgr, t.TripStatus, u.UserId)
		if err != nil {
			panic(err.Error())
		}
	}
	
	func updateTrip(id string, t Trip) {
		_, err := db.Exec("update trip set pickuploc=?, intake=?, altpickuploc=?, starttraveltime=?, destaddress = ?, noofpsgr = ?, tripstatus = ? where tripid=?", t.PickupLoc, t.AltPickupLoc, t.StartTravelTime, t.DestAddress, t.NoOfPsgr, t.TripStatus t.TripID)
		if err != nil {
			panic(err.Error())
		}
	}
	
	func isExist(tripid string) (Trip, bool) {
		var t Trip
	
		result := db.QueryRow("select * from trip where tripid=?", tripid)
		err := result.Scan(&tripid, &d.DriverId, &t.PickupLoc, &t.AltPickupLoc, &t.StartTravelTime, &t.DestAddress, &t.NoOfPsgr, &t.TripStatus, &u.UserId)
		if err == sql.ErrNoRows {
			return u, false
		}
	
		return u, true
	}
}