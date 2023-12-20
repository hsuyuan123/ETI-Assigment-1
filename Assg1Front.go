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
	var user User
	var userID string
	fmt.Print("Enter the ID of the user to be updated: ")
	fmt.Scanf("%v", &userID)
	fmt.Print("Enter the new first name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	course.FirstName = strings.TrimSpace(input)
	fmt.Print("Enter the new last name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	course.LastName = strings.TrimSpace(input)
	fmt.Print("Enter the new mobile number: ")
	fmt.Scanf("%d", &(user.MobileNo))
	fmt.Print("Enter the new Email: ")
	fmt.Scanf("%d", &(user.Email))

	postBody, _ := json.Marshal(course)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5001/api/v1/drivers/"+driverID, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Driver", driverID, "updated successfully")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - driver", driverID, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func updateDriver() {
	var driver Driver
	var driverID string
	fmt.Print("Enter the ID of the driver to be updated: ")
	fmt.Scanf("%v", &driverID)
	fmt.Print("Enter the new first name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	course.FirstName = strings.TrimSpace(input)
	fmt.Print("Enter the new last name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	course.LastName = strings.TrimSpace(input)
	fmt.Print("Enter the new mobile number: ")
	fmt.Scanf("%d", &(driver.MobileNo))
	fmt.Print("Enter the new Email: ")
	fmt.Scanf("%d", &(driver.Email))
	fmt.Print("Enter the new Vehicle Number: ")
	fmt.Scanf("%d", &(driver.VehicleNo))
	fmt.Print("Enter the new License Number: ")
	fmt.Scanf("%d", &(driver.LicenseNo))

	postBody, _ := json.Marshal(course)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/drivers/"+driverID, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Driver", driverID, "updated successfully")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - driver", driverID, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func deleteDriver() {
	var course string
	fmt.Print("Enter the ID of the driver to be deleted: ")
	fmt.Scanf("%v", &driver)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodDelete, "http://localhost:5000/api/v1/driver/"+course, nil); err == nil {
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

//Update trips for driver

//Update trips for users