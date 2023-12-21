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

type User struct {
	UserId    string `json:"User ID"`
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	MobileNo  int    `json:"MobileNo"`
	Email     string `json:"Email"`
	AccStatus string `json:"AccStatus"`
}

var (
	db  *sql.DB
	err error
)

// Link to trip
func main() {
	db, err = sql.Open("mysql", "user1:password@tcp(127.0.0.1:3306)/my_db")

	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/users/{userid}", user).Methods("GET", "DELETE", "POST", "PATCH", "PUT", "OPTIONS")
	//User
	router.HandleFunc("/api/v1/users", User)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
	//Trip
	router.HandleFunc("/api/v1/trip", Trip)
	fmt.Println("Listening at port 5002")
	log.Fatal(http.ListenAndServe(":5002", router))
}

func users(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data User

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist[params["userid"]]; !ok {
					fmt.Println(data)
					insertUser[params["userid"]] = data

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
				if _, ok := isExist[params["userid"]]; ok {
					fmt.Println(data)

					updateUser[params["userid"]] = data
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
				if orig, ok := isExist[params["userid"]]; ok {
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
					updateUser[params["userid"]] = orig
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "User ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
		//delete user
	} else if val, ok := isExist[params["userid"]]; ok {
		if r.Method == "DELETE" {
			fmt.Fprintf(w, params["userid"]+" Deleted")
			delete(users, params["userid"])
		} else {
			json.NewEncoder(w).Encode(val)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid User ID")
	}

}

func getUsers() map[string]User {
	results, err := db.Query("select * from User")
	if err != nil {
		panic(err.Error())
	}

	var users map[string]User = map[string]User{}

	for results.Next() {
		var u User
		var id string
		err = results.Scan(&id, &u.FirstName, &u.LastName, &u.MobileNo, &u.Email)
		if err != nil {
			panic(err.Error())
		}

		users[id] = u
	}

	return users
}

func insertUser(userid string, u User) {
	u.AccStatus = "active"
	_, err := db.Exec("insert into user values(?,?,?,?,?)", userid, u.FirstName, u.LastName, u.MobileNo, u.Email, u.AccStatus)
	if err != nil {
		panic(err.Error())
	}
}
func updateUser(userid string, u User) {
	_, err := db.Exec("update user set firstname=?, lastname=?, mobileno=?, email=? where userid=?", u.FirstName, u.LastName, u.MobileNo, u.Email, userid)
	if err != nil {
		panic(err.Error())
	}
}

func isExist(userid string) (User, bool) {
	var u User

	result := db.QueryRow("select * from user where userid=?", userid)
	err := result.Scan(&userid, &u.FirstName, &u.LastName, &u.MobileNo, &u.Email)
	if err == sql.ErrNoRows {
		return u, false
	}

	return u, true
}

func delUser(userid string) (int64, error) {
	result, err := db.Exec("delete from user where userid=?", userid)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
