/* package main

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

var db *sql.DB
var err error

type Passenger struct {
	PassengerID  int
	Firstname    string
	Lastname     string
	MobileNo     string
	EmailAddress string
	InTrip       bool
}

var passengers map[string]Passenger

func insertPassenger(db *sql.DB, newPassenger Passenger) {
	ID := newPassenger.PassengerID
	FN := newPassenger.Firstname
	LN := newPassenger.Lastname
	MN := newPassenger.MobileNo
	EA := newPassenger.EmailAddress
	query := fmt.Sprintf("INSERT INTO passenger (PassengerID, FirstName, LastName, MobileNo, EmailAddress) VALUES ('%d', '%s', '%s', '%s', '%s')",
		ID, FN, LN, MN, EA) //Create Query to insert passenger into db

	_, err := db.Query(query) //Run Query
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("insert success")

}

func GetPassenger(w http.ResponseWriter, router *http.Request) {
	fmt.Fprint(w, "hello")

}

func CreatePassenger(w http.ResponseWriter, router *http.Request) {

	params := mux.Vars(router)
	var newPassenger Passenger
	reqBody, err := ioutil.ReadAll(router.Body)
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/assignment1")
	json.Unmarshal(reqBody, &newPassenger)

	fmt.Println("Passenger " + newPassenger.EmailAddress)

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
		fmt.Println("Databse Close")

	}

	if err == nil {
		// convert JSON to object
		json.Unmarshal(reqBody, &newPassenger)

		if newPassenger.Firstname == "" {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			w.Write([]byte(
				"Please enter your First name"))
			return
		}

		if newPassenger.Lastname == "" {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			w.Write([]byte(
				"Please enter your Last name"))
			return
		}

		if newPassenger.MobileNo == "" {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			w.Write([]byte(
				"Please enter your mobile phone"))
			return
		}
		if newPassenger.EmailAddress == "" {
			w.WriteHeader(
				http.StatusUnprocessableEntity)
			w.Write([]byte(
				"Please enter your email address"))
			return
		}

		if _, ok := passengers[params["EmailAddress"]]; !ok {
			passengers[params["EmailAddress"]] = newPassenger
			w.WriteHeader(http.StatusCreated)
			json.Unmarshal(reqBody, &newPassenger)
			w.Write([]byte("Passenger account Created"))
		} else {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(
				"Duplicate email address"))
		}
	} else {
		w.WriteHeader(
			http.StatusUnprocessableEntity)
		w.Write([]byte("error processing"))
	}
	insertPassenger(db, newPassenger)

}

func UpdatePassenger(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Passenger Update")

}

func DeletePassenger(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Passenger Deleted")

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/passenger", GetPassenger).Methods("GET")
	router.HandleFunc("/passenger", CreatePassenger).Methods("POST")

	router.HandleFunc("/passenger ", DeletePassenger).Methods("DELETE")
	http.ListenAndServe(":5001", router)

	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/assignment")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	// Use mysql as driverName and a valid DSN as dataSourceName:

	fmt.Println("Database opened")

}
*/