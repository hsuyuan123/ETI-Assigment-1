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
	router.HandleFunc("/api/v1/user", newUser).Methods("GET", "DELETE", "POST", "PATCH", "PUT", "OPTIONS")
	//User
	router.HandleFunc("/api/v1/user", users)
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

func main() {
	outer:
		for {
			fmt.Println(strings.Repeat("=", 10))
			fmt.Println("CarPool Management Console\n",
				"1. List all Users\n",
				"2. Create new user\n",
				"3. Update user\n",
				"4. Delete user\n",
				"5.List all Drivers\n",
				"6. Update Driver\n",
				"7. Delete Driver\n",
				"8. Update Trips\n",
				"9. Delete Trips\n",
				"10. Quit")
			fmt.Print("Enter an option: ")
	
			var choice int
			fmt.Scanf("%d", &choice)
	
			switch choice {
			case 1:
				listAllUser()
			case 2:
				createUser()
			case 3:
				updateUser()
			case 4:
				deleteUser()
			case 5:
				listAllDriver()
			case 6:
				createDriver()
			case 7:
				updateDriver()
			case 8:
				listAllTrip()
			case 9:
				updateTrip()
			case 10:
				deleteTrip()
			case 11:
				break outer
			default:
				fmt.Println("### Invalid Input ###")
			}
		}
	}

/*func console(){
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
*/
	func deleteUser() {
		var course string
		fmt.Print("Enter the ID of the user to be deleted: ")
		fmt.Scanf("%v", &user)
	
		client := &http.Client{}
		if req, err := http.NewRequest(http.MethodDelete, "http://localhost:5000/api/v1/courses/"+course, nil); err == nil {
			if res, err := client.Do(req); err == nil {
				if res.StatusCode == 200 {
					fmt.Println("Course", course, "deleted successfully")
				} else if res.StatusCode == 404 {
					fmt.Println("Error - course", course, "does not exist")
				}
			} else {
				fmt.Println(2, err)
			}
		} else {
			fmt.Println(3, err)
		}
	}
}

func updateUser() {
	var course Course
	var courseID string
	fmt.Print("Enter the ID of the course to be updated: ")
	fmt.Scanf("%v", &courseID)
	fmt.Print("Enter the new name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	course.Name = strings.TrimSpace(input)
	fmt.Print("Enter the new planned intake number: ")
	fmt.Scanf("%d", &(course.Intake))
	fmt.Print("Enter the new minimum GPA: ")
	fmt.Scanf("%d", &(course.MinGPA))
	fmt.Print("Enter the new maximum GPA: ")
	fmt.Scanf("%d", &(course.MaxGPA))

	postBody, _ := json.Marshal(course)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/courses/"+courseID, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Course", courseID, "updated successfully")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - course", courseID, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}