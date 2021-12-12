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

	db.AutoMigrate(&Trip{})
}
func main() { /*Connect to DB*/
	DataMigrations()

	router := mux.NewRouter()
	router.HandleFunc("/trip/{email}", CreateTrip).Methods("POST")
	router.HandleFunc("/trip/passengerhistory/{passenger_id}", GetTrip).Methods("GET")
	router.HandleFunc("/trip/driverhistory/{driver_id}", GetDriverTrip).Methods("GET")
	router.HandleFunc("/trip/status/{driver_id}", startTripDriver).Methods("PUT")
	router.HandleFunc("/trip/completed/{driver_id}", endTripDriver).Methods("PUT")
	http.ListenAndServe(":5020", router)
	fmt.Println("Listening at port 5020")
	log.Fatal(http.ListenAndServe(":5020", router))
}

var db *gorm.DB
var err error

const dsn = "root:root@tcp(127.0.0.1:3306)/assignment1?charset=utf8mb4&parseTime=True&loc=Local"

type Trip struct {
	TripID        int    `json:"tripid" gorm:"primaryKey"`
	PassengerID   int    `json:"passengerid"`
	PickupPoint   string `json:"pickup"`
	DropoffPoint  string `json:"dropoff"`
	DriverID      int    `json:"driverid"`
	Carlicensenum string `json:"Carlicensenum"`
	Status        string `json:"status" gorm:"default:Waiting for Driver"`
}

type Passenger struct {
	//auto increament ID
	PassengerID  int    `json:"passengerid" gorm:"primaryKey"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	MobileNo     string `json:"mobileno"`
	EmailAddress string `json:"emailaddress"`
}

type Driver struct {
	//auto increament ID
	DriverID      int    `json:"driverid" gorm:"primaryKey"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	MobileNo      string `json:"mobileno"`
	EmailAddress  string `json:"emailaddress"`
	IcNum         string `json:"icnum"`
	Carlicensenum string `json:"carlicensenum"`
	Available     bool   `json:"available" gorm:"type:bool" gorm:"default:true"`
}

func GetAvailableDriver() (Driver, error) {
	url := "http://localhost:5010/driver"

	response, err := http.Get(url)

	var avaiDriver Driver
	if err == nil {
		if response.StatusCode == http.StatusAccepted {
			data, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()
			json.Unmarshal(data, &avaiDriver)
		}
	} else {
		return avaiDriver, err
	}

	return avaiDriver, nil
}

func CreateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	passenger, err := getPassengerbyemail(params["email"])

	reqBody, err := ioutil.ReadAll(r.Body)
	var newTrip Trip
	var driver Driver
	driver, err = GetAvailableDriver()
	if err != nil {
		fmt.Fprintf(w, "Driver not found")
	}
	fmt.Println(driver)
	if err == nil {
		if driver.EmailAddress == "" {
			fmt.Printf("driver not found")
		} else {
			// convert JSON to object
			json.Unmarshal(reqBody, &newTrip)
			fmt.Printf("Driver Accepted the trip")
			newTrip.PassengerID = passenger.PassengerID
			newTrip.DriverID = driver.DriverID
			newTrip.Carlicensenum = driver.Carlicensenum
			//cnewTrip.Status = "Waiting for Driver"
			db.Create(&newTrip)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newTrip)
		}

	} else {
		fmt.Printf("Error")
	}
}

func getPassengerbyemail(passengeremail string) (Passenger, error) {
	url := "http://localhost:5001/passenger/trip/" + passengeremail

	response, err := http.Get(url)
	var passenger Passenger

	if err == nil {
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &passenger)

		return passenger, nil
	} else {
		return passenger, err
	}
}

func GetTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var trips []Trip

	db.Where("passenger_id = ?", params["passenger_id"]).Find(&trips)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(trips)
	fmt.Println("Trips: ", trips)
}

func GetDriverTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var trips []Trip
	db.Where("driver_id = ?", params["driver_id"]).Find(&trips)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(trips)
	fmt.Println("Trips: ", trips)
}

func startTripDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var trips Trip
	err := db.Where("driver_id = ?", params["driver_id"]).Find(&trips)
	fmt.Println(params["driver_id"])
	if err == nil {
		fmt.Fprintf(w, "Driver does not have any trip")

	} else {
		db.Model(&Trip{}).Where("driver_id = ?", params["driver_id"]).Update("status", "In Progress")
		json.NewEncoder(w).Encode(trips)
		fmt.Println(params["driver_id"])
	}
}

func endTripDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var trips Trip
	err := db.Where("driver_id = ?", params["driver_id"]).Find(&trips)
	fmt.Println(params["driver_id"])
	if err == nil {
		fmt.Fprintf(w, "Driver does not have any trip")

	} else {
		db.Model(&Trip{}).Where("driver_id = ?", params["driver_id"]).Update("status", "Completed")
		json.NewEncoder(w).Encode(trips)
		fmt.Println(params["driver_id"])
	}
}
