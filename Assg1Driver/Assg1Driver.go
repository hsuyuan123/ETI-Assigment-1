package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Driver struct {
	DriverId  string `json:"Driver ID"`
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	MobileNo  int    `json:"MobileNo"`
	Email     string `json:"Email"`
	LicenseNo string `json:"LicenseNo"`
	VehicleNo string `json:"VehicleNo"`
	AccStatus string `json:"AccStatus"`
}

func main() {
	db, err = sql.Open("mysql", "user1:password@tcp(127.0.0.1:3306)/my_db")

	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/drivers {driverid}", driver).Methods("GET", "DELETE", "POST", "PATCH", "PUT", "OPTIONS")
	//Driver
	router.HandleFunc("/api/v1/driver", Driver)
	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
	//Trip
	router.HandleFunc("/api/v1/trip", Trip)
	fmt.Println("Listening at port 5002")
	log.Fatal(http.ListenAndServe(":5002", router))
}

func GetData(db *sql.DB) {
	results, err := db.Query("select * from Driver")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var d Driver
		err = results.Scan(&d.DriverId, &d.FirstName, &d.LastName, &d.MobileNo, &d.Email, &d.LicenseNo, &d.VehicleNo, &d.AccStatus)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(d.DriverId, d.FirstName, d.LastName, d.MobileNo, d.Email)
	}
}

func drivers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := drivers[params["driverid"]]; !ok {
					fmt.Println(data)
					drivers[params["driverid"]] = data

					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Fprintf(w, "Driver ID exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := drivers[params["druverid"]]; ok {
					fmt.Println(data)

					drivers[params["driverid"]] = data
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Driver ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PATCH" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data map[string]interface{}

			if err := json.Unmarshal(body, &data); err == nil {
				if orig, ok := driver[params["driverid"]]; ok {
					fmt.Println(data)

					for k, v := range data {
						switch k {
						case "First Name":
							orig.FirstName = v.(string)
						case "Last Name":
							orig.LastName = v.(string)
						case "Mobile No":
							orig.MobileNo = int(v.(float64))
						case "Email":
							orig.Email = v.(string)
						case "VehicleNo":
							orig.Email = v.(string)
						case "LicenseNo":
							orig.Email = v.(string)
						}
					}
					driver[params["driverid"]] = orig
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Driver ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
		//delete user
	} else if val, ok := driver[params["driverid"]]; ok {
		if r.Method == "DELETE" {
			fmt.Fprintf(w, params["driverid"]+" Deleted")
			delete(user, params["driverid"])
		} else {
			json.NewEncoder(w).Encode(val)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Driver ID")
	}
}

func updateDriver(db *sql.DB) {
	var u User
	result, err := db.Exec("insert into driver (firstname, lastname,  mobileno, email, vehicleno, licenseno)     values(?, ?, ?, ?, ?, ?, ?, ?)", u.FirstName, u.LastName, u.MobileNo, u.Email, &u.LicenseNo, &u.VehicleNo, &u.AccStatus)
	if err != nil {
		panic(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		print(id)
		panic(err.Error())
	}
}

func delDriver(id string) (int64, error) {
	result, err := db.Exec("delete from driver where id=?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
