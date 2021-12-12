package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DataMigrations() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		fmt.Println(err.Error())
		panic("Database Unable to connect")
	}

	db.AutoMigrate(&Passenger{})
}
func main() { /*Connect to DB*/
	DataMigrations()

	router := mux.NewRouter()
	router.HandleFunc("/passenger", CreatePassenger).Methods("POST")
	router.HandleFunc("/passenger/trip/{email}", GetPassenger).Methods("GET")
	router.HandleFunc("/passenger/trip/{passenger_id}", GetPassengerTripbyID).Methods("GET")
	router.HandleFunc("/passenger/{email}", UpdatePassenger).Methods("PUT")
	router.HandleFunc("/passenger/{email}", DeletePassenger).Methods("DELETE")
	http.ListenAndServe(":5001", router)
	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}

var db *gorm.DB
var err error

const dsn = "root:root@tcp(127.0.0.1:3306)/assignment1?charset=utf8mb4&parseTime=True&loc=Local"

type Passenger struct {
	//auto increament ID
	PassengerID  int    `json:"passengerid" gorm:"primaryKey"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	MobileNo     string `json:"mobileno"`
	EmailAddress string `json:"emailaddress"`
}

type Trip struct {
	TripID        int    `json:"tripid" gorm:"primaryKey"`
	PassengerID   int    `json:"passengerid"`
	PickupPoint   string `json:"pickup"`
	DropoffPoint  string `json:"dropoff"`
	DriverID      int    `json:"driverid"`
	Carlicensenum string `json:"Carlicensenum"`
	Status        string `json:"status"`
}

func CreatePassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newpassenger Passenger
	var dbpassenger Passenger

	if err == nil {
		json.Unmarshal(reqBody, &newpassenger)

		if newpassenger.Firstname == "" {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "Please Enter your First Name")
			return
		} else if newpassenger.Lastname == "" {

			{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Please Enter your Last Name"))
				return
			}

		} else if newpassenger.MobileNo == "" {
			{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Please Enter your Mobile No"))
				return
			}
		} else if newpassenger.EmailAddress == "" {
			{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Please Enter your Email Address"))
				return
			}
		}
	}
	err := db.Where("email_address = ?", newpassenger.EmailAddress).First(&dbpassenger).Error
	if err == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "The email address has already existed")
		return
	}

	err1 := db.Where("mobile_no = ?", newpassenger.MobileNo).First(&dbpassenger).Error
	if err1 == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "the phone number already register")
		return
	}

	//if pass all validations come here
	db.Create(&newpassenger)
	json.NewEncoder(w).Encode(newpassenger)

}

func GetPassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var passenger Passenger
	err := db.Where("email_address = ?", params["email"]).First(&passenger).Error
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(passenger)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "not found")
	}
}

func DeletePassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var passenger Passenger
	err := db.Where("email_address = ?", params["email"]).First(&passenger).Error
	if err == nil {
		fmt.Fprintf(w, "Unable to Delete")

	} else {
		fmt.Fprintf(w, "Email you enter is not registered ")

	}
}

func UpdatePassenger(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	var passenger Passenger

	if err != nil {
		fmt.Printf("  The email you enter is not registered  ")
		return
	} else {
		json.NewDecoder(router.Body).Decode(&passenger)
		db.Model(&Passenger{}).Where("email_address=?", params["email"]).Updates(passenger)

		var newPassenger Passenger
		db.Where("email_address=?", passenger.EmailAddress).First(&newPassenger)
		json.NewEncoder(w).Encode(newPassenger)
	}
}

func getPassengertrip(passengerId string) ([]Trip, error) {
	url := "http://localhost:5020/trip/" + passengerId

	response, err := http.Get(url)
	var trips []Trip

	if err == nil {
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &trips)

		return trips, nil
	} else {
		return trips, err
	}
}

func GetPassengerTripbyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var passenger Passenger
	//check that passenger exist
	err := db.Where("passenger_id = ?", params["passenger_id"]).First(&passenger).Error
	if err == nil {
		trips, err2 := getPassengertrip(params["passenger_id"])
		if err2 != nil {
			fmt.Fprintf(w, "HTTP Error")
			return
		}
		if len(trips) <= 0 {
			fmt.Fprintf(w, "Passenger does not have any trips")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(trips)
	} else {
		fmt.Fprintf(w, "Passenger not registered")
	}

}
