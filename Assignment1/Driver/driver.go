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

	db.AutoMigrate(&Driver{})
}
func main() { /*Connect to DB*/
	DataMigrations()

	router := mux.NewRouter()
	router.HandleFunc("/driver", CreateDriver).Methods("POST")
	router.HandleFunc("/driver", GetDriver).Methods("GET")
	router.HandleFunc("/driver/{email}", GetDriverbyEmail).Methods("GET")
	router.HandleFunc("/driver/{email}", UpdateDriver).Methods("PUT")
	router.HandleFunc("/driver/{email}", DeleteDriver).Methods("DELETE")
	router.HandleFunc("/driver/starttrip/{driver_id}", startTrip).Methods("PUT")
	router.HandleFunc("/driver/endtrip/{driver_id}", endTrip).Methods("PUT")
	http.ListenAndServe(":5010", router)
	fmt.Println("Listening at port 5010")
	log.Fatal(http.ListenAndServe(":5010", router))
}

var db *gorm.DB
var err error

const dsn = "root:root@tcp(127.0.0.1:3306)/assignment1?charset=utf8mb4&parseTime=True&loc=Local"

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

type Trip struct {
	TripID        int    `json:"tripid" gorm:"primaryKey"`
	PassengerID   int    `json:"passengerid"`
	PickupPoint   string `json:"pickup"`
	DropoffPoint  string `json:"dropoff"`
	DriverID      int    `json:"driverid"`
	Carlicensenum string `json:"Carlicensenum"`
	Status        string `json:"status"`
}

func CreateDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newdriver Driver
	var dbdriver Driver

	if err == nil {
		json.Unmarshal(reqBody, &newdriver)

		if newdriver.Firstname == "" {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "Please Enter your First Name")
			return
		} else if newdriver.Lastname == "" {

			{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Please Enter your Last Name"))
				return
			}

		} else if newdriver.MobileNo == "" {
			{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Please Enter your Mobile No"))
				return
			}
		} else if newdriver.EmailAddress == "" {
			{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Please Enter your Email Address"))
				return
			}
		} else if newdriver.IcNum == "" {
			{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Please Enter your Identification Number"))
				return
			}
		} else if newdriver.Carlicensenum == "" {
			{
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte(
					"Please Enter your License Number"))
				return
			}
		}
	}

	err := db.Where("email_address = ?", newdriver.EmailAddress).First(&dbdriver).Error
	if err == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "The email address has already existed")
		return
	}

	err1 := db.Where("mobile_no = ?", newdriver.MobileNo).First(&dbdriver).Error
	if err1 == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "the phone number has already existed")
		return
	}

	err2 := db.Where("ic_num = ?", newdriver.IcNum).First(&dbdriver).Error
	if err2 == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "The Indentification Number has already existed")
		return
	}

	err3 := db.Where("carlicensenum= ?", newdriver.Carlicensenum).First(&dbdriver).Error
	if err3 == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "The License Number has already existed")
		return
	}

	//if pass all validations come here
	db.Create(&newdriver)
	json.NewEncoder(w).Encode(newdriver)

}

func GetDriver(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var driver Driver
		err := db.Where("available = ?", true).First(&driver).Error
		fmt.Println(driver)
		if err == nil {
			fmt.Println("Driver Found")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(driver)

		} else {
			fmt.Fprint(w, "No Available Driver Found")
		}
	}
}

func GetDriverbyEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var driver Driver
	err := db.Where("email_address = ?", params["email"]).First(&driver).Error
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(driver)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "not found")
	}
}

func DeleteDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var driver Driver
	err := db.Where("email_address = ?", params["email"]).First(&driver).Error
	if err == nil {
		fmt.Fprintf(w, "Unable to Delete")

	} else {
		fmt.Fprintf(w, "Email you enter is not registered ")

	}
}

func UpdateDriver(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	var driver Driver

	if err != nil {
		fmt.Printf("  The email you enter is not registered  ")
		return
	} else {
		json.NewDecoder(router.Body).Decode(&driver)
		db.Model(&Driver{}).Where("email_address=?", params["email"]).Updates(driver)

		var newDriver Driver
		err := db.Where("email_address=?", driver.EmailAddress).First(&newDriver).Error
		fmt.Println(err)
		json.NewEncoder(w).Encode(newDriver)
	}
}

func updatetripstatus(driverID string) error {
	url := "http://localhost:5020/trip/status"

	request, err := http.NewRequest(http.MethodPut,
		url+"/"+driverID, nil)

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
		//return nil
	}
	return nil
}

func updatetripstatusend(driverID string) error {
	url := "http://localhost:5020/trip/completed"

	request, err := http.NewRequest(http.MethodPut,
		url+"/"+driverID, nil)

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
		//return nil
	}
	return nil
}

//this fucntion become put
func startTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var driver Driver
	// to check if the email
	err := db.Where("driver_id = ?", params["driver_id"]).First(&driver).Error

	if err == nil {

		err2 := updatetripstatus(params["driver_id"])
		if err2 == nil {
			db.Model(&driver).Update("available", false)
			json.NewDecoder(r.Body).Decode(&driver)
			json.NewEncoder(w).Encode(driver)

		} else {
			fmt.Fprintf(w, "unable to update the trips status")

		}

		fmt.Fprintf(w, "Updated")

	} else {
		fmt.Fprintf(w, "Driver is not registered")

	}

}

func endTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var driver Driver
	// to check if the email
	err := db.Where("driver_id = ?", params["driver_id"]).First(&driver).Error

	if err == nil {

		err2 := updatetripstatusend(params["driver_id"])
		if err2 == nil {
			db.Model(&driver).Update("available", true)
			json.NewDecoder(r.Body).Decode(&driver)
			json.NewEncoder(w).Encode(driver)

		} else {
			fmt.Fprintf(w, "unable to update the trips status")

		}

		fmt.Fprintf(w, "Updated")

	} else {
		fmt.Fprintf(w, "Driver is not registered")

	}

}
