package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	UserId    string `json:"User ID"`
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	MobileNo  int    `json:"MobileNo"`
	Email     string `json:"Email"`
	LicenseNo string `json:"LicenseNo"`
	VehicleNo string `json:"VehicleNo"`
	AccStatus string `json:"AccStatus"`
}

var (
	db  *sql.DB
	err error
)

// Link to trip
func main() {
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db")

	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/user", newUser).Methods("GET", "DELETE", "POST", "PATCH", "PUT", "OPTIONS")
	//User
	router.HandleFunc("/api/v1/user", User)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
	//Trip
	router.HandleFunc("/api/v1/trip", Trip)
	fmt.Println("Listening at port 5002")
	log.Fatal(http.ListenAndServe(":5002", router))
}

func insertUser(id string, u User) {
	_, err := db.Exec("insert into user values(?,?,?,?,?)", id, u.FirstName, u.LastName, u.MobileNo, u.Email)
	if err != nil {
		panic(err.Error())
	}
}

func updateCourse(id string, c Course) {
	_, err := db.Exec("update course set name=?, intake=?, mingpa=?, maxgpa=? where id=?", c.Name, c.Intake, c.MinGPA, c.MaxGPA, id)
	if err != nil {
		panic(err.Error())
	}
}

func deleteUser(id string) (int64, error) {
	result, err := db.Exec("delete from course where id=?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func insertUser(id string, u User) {
	_, err := db.Exec("insert into course values(?,?,?,?,?)", id, u.FirstName, u.LastName, u.MobileNo, u.Email)
	if err != nil {
		panic(err.Error())
	}
}

func updateUser(id string, u User) {
	_, err := db.Exec("update user set name=?, intake=?, mingpa=?, maxgpa=? where id=?", u.FirstName, u.LastName, u.MobileNo, u.Email, id)
	if err != nil {
		panic(err.Error())
	}
}

func user(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data User

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := users[params["userid"]]; !ok {
					fmt.Println(data)
					users[params["userid"]] = data

					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Fprintf(w, "User ID exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data User

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := users[params["userid"]]; ok {
					fmt.Println(data)

					users[params["userid"]] = data
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "User ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PATCH" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data map[string]interface{}

			if err := json.Unmarshal(body, &data); err == nil {
				if orig, ok := user[params["userid"]]; ok {
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
						}
					}
					user[params["userid"]] = orig
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Course ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
		//delete user
	} else if val, ok := user[params["userid"]]; ok {
		if r.Method == "DELETE" {
			fmt.Fprintf(w, params["userid"]+" Deleted")
			delete(user, params["userid"])
		} else {
			json.NewEncoder(w).Encode(val)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid User ID")
	}
}

func allusers(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	found := false
	results := map[string]User{}

	if value := query.Get("q"); len(value) > 0 {
		for k, v := range users {
			if strings.Contains(strings.ToLower(v.Name), strings.ToLower(value)) {
				results[k] = v
				found = true
			}
		}

		if !found {
			fmt.Fprintf(w, "No user found")
		} else {
			json.NewEncoder(w).Encode(struct {
				Results map[string]User `json:"Search Results"`
			}{results})
		}
	} else if value = query.Get("gpa"); len(value) > 0 {
		found := false
		results := map[string]User{}
		gpa, _ := strconv.Atoi(value)

		for k, v := range user {
			if gpa <= v.MaxGPA {
				results[k] = v
				found = true
			}
		}

		if !found {
			fmt.Fprintf(w, "No course eligible")
		} else {
			json.NewEncoder(w).Encode(struct {
				Results map[string]Course `json:"Eligible course(s)"`
			}{results})
		}
	} else {
		coursesWrapper := struct {
			Courses map[string]Course `json:"Courses"`
		}{courses}
		json.NewEncoder(w).Encode(coursesWrapper)
		return
	}
}
