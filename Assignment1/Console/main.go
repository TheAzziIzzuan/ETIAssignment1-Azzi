package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//Passenger Structure
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

type Trip struct {
	TripID        int    `json:"tripid" gorm:"primaryKey"`
	PassengerID   int    `json:"passengerid"`
	PickupPoint   string `json:"pickup"`
	DropoffPoint  string `json:"dropoff"`
	DriverID      int    `json:"driverid"`
	Carlicensenum string `json:"Carlicensenum"`
	Status        string `json:"status" gorm:"default:Waiting for Driver"`
}

//Menu upon running main.go
func main() {
	for {
		fmt.Println("Welcome to getaride")
		fmt.Println("[1] Login Passenger")
		fmt.Println("[2] Create Passenger")
		fmt.Println("[3] Login Driver")
		fmt.Println("[4] Create Driver")

		//scan for user input
		var userInput string
		fmt.Scanln(&userInput)

		//if user input is 1 redirect to login
		if userInput == "1" {
			passengerhome()
		} else if userInput == "2" {
			createPassengeraccount()
		} else if userInput == "3" {
			driverhome()
		} else if userInput == "4" {
			createDriveraccount()
		} else {
			fmt.Println("Please Select a valid option")
		}
	}
}

//menu of passenger when logging in
func passengerhome() {
	//check if passenger email is valid
	fmt.Println("Enter your Passenger Email: ")
	var validemail string
	fmt.Scanln(&validemail)

	//check if email exist in databsae
	passenger := getPassengerbyemail(validemail)
	if (passenger != Passenger{}) {
		fmt.Println("Welcome", validemail)
	} else {
		fmt.Println("The Passenger email you entered is not registered")
		return
	}

	for {
		fmt.Println("[1] Update account")
		fmt.Println("[2] Display Trip")
		fmt.Println("[3] Create Trip")
		fmt.Println("[0] Exit")

		var userInput string
		fmt.Scanln(&userInput)
		
		//different options for passenger
		if userInput == "1" {
			updatePassengeraccount(validemail)
		} else if userInput == "2" {
			trips := getPassengerTripbyID(passenger.PassengerID)
			displaypassengertrips(trips)
		} else if userInput == "3" {
			createTrip(validemail)
		} else if userInput == "0" {
			main()
		}
	}
}

//driver home menu upon logging in
func driverhome() {
	//validate driver email address
	fmt.Println("Enter your Driver Email:")
	var validemail string
	fmt.Scanln(&validemail)

	//check if email exist in databsae
	driver := getDriverbyemail(validemail)
	if (driver != Driver{}) {
		fmt.Println("Welcome", validemail)
	} else {
		fmt.Println("The Driver email you entered is not registered")
		return
	}
	for {
		fmt.Println("[1] Update account")
		fmt.Println("[2] Display Trip")
		fmt.Println("[3] Start Trip")
		fmt.Println("[4] End Trip")
		fmt.Println("[0] Exit")

		var userInput string
		fmt.Scanln(&userInput)
		if userInput == "1" {
			updateDriveraccount(validemail)
		} else if userInput == "2" { //Display Driver
			trips := getDriverTripbyID(driver.DriverID)
			displaydrivertrips(trips)
		} else if userInput == "3" {
			startTripDriver(driver.DriverID)
		} else if userInput == "4" {
			endTripDriver(driver.DriverID)
		} else if userInput == "0" {
			main()
		} else {
			fmt.Println("Please enter valid input")
		}
	}
}

//get passenger email by calling passenger api using email address
func getPassengerbyemail(email string) Passenger {
	url := "http://localhost:5001/passenger/trip/" + email

	response, err := http.Get(url)

	var passenger Passenger
	if err == nil {
		if response.StatusCode == http.StatusCreated {
			json.NewDecoder(response.Body).Decode(&passenger)
			response.Body.Close()
			return passenger
		}
	} else {
		return passenger
	}

	return passenger
}

//get driver email by calling driver api using email address
func getDriverbyemail(email string) Driver {
	url := "http://localhost:5010/driver/" + email

	response, err := http.Get(url)

	var driver Driver
	if err == nil {
		if response.StatusCode == http.StatusCreated {
			json.NewDecoder(response.Body).Decode(&driver)
			response.Body.Close()
			//driver with objects
			return driver
		}
	} else {
		//empty driver
		return driver
	}
	//empty driver
	return driver
}
//get driver trip by calling trip api using driver id
func getDriverTripbyID(driver_id int) []Trip {
	url := fmt.Sprintf("http://localhost:5020/trip/driverhistory/%d", driver_id)

	response, err := http.Get(url)

	var trips []Trip
	if err == nil {
		if response.StatusCode == http.StatusCreated {
			json.NewDecoder(response.Body).Decode(&trips)
			response.Body.Close()
			return trips
		}
	} else {
		//return empty trips
		return trips
	}

	//return empty trips
	return trips
}
//get passenger trip by calling trip api using passenger id
func getPassengerTripbyID(passenger_id int) []Trip {
	url := fmt.Sprintf("http://localhost:5020/trip/passengerhistory/%d", passenger_id)

	response, err := http.Get(url)

	var trips []Trip
	if err == nil {
		if response.StatusCode == http.StatusCreated {
			json.NewDecoder(response.Body).Decode(&trips)
			response.Body.Close()
			return trips
		}
	} else {
		//return empty trips
		return trips
	}

	//return empty trips
	return trips
}

//function to print and display driver trips
func displaydrivertrips(trips []Trip) {
	fmt.Println("DRIVER TRIPS")
	for _, trip := range trips {
		fmt.Println("TripID: ", trip.TripID)
		fmt.Println("PassengerID: ", trip.PassengerID)
		fmt.Println("Pickuppoint: ", trip.PickupPoint)
		fmt.Println("dropoffpoint: ", trip.DropoffPoint)
		fmt.Println("CarLicenceNum: ", trip.Carlicensenum)
		fmt.Println("Status: ", trip.Status)
		fmt.Println("")
	}
}
//function to print and display passenger trips
func displaypassengertrips(trips []Trip) {
	fmt.Println("PASSENGER HISTORY TRIPS")
	for _, trip := range trips {
		fmt.Println("TripID: ", trip.TripID)
		fmt.Println("PassengerID: ", trip.PassengerID)
		fmt.Println("Pickuppoint: ", trip.PickupPoint)
		fmt.Println("dropoffpoint: ", trip.DropoffPoint)
		fmt.Println("CarLicenceNum: ", trip.Carlicensenum)
		fmt.Println("Status: ", trip.Status)
		fmt.Println("")
	}
}

//CREATE PASSENGER ACCOUNT INPUT
func createPassengeraccount() {
	fmt.Println("Please enter your first name: ")
	var firstname string
	fmt.Scanln(&firstname)

	fmt.Println("Please enter your last name: ")
	var lastname string
	fmt.Scanln(&lastname)

	fmt.Println("Please enter your phone number: ")
	var phonenumber string
	fmt.Scanln(&phonenumber)

	fmt.Println("Please enter your Email Address: ")
	var emailaddress string
	fmt.Scanln(&emailaddress)

	account := createAccountforpassengerapi(firstname, lastname, phonenumber, emailaddress)
	if (account == Passenger{}) {
		fmt.Println("Account could not be created")
	} else {
		fmt.Println("Account has been created")
	}

}

//CREATE DRIVER ACCOUNT INPUT
func createDriveraccount() {
	fmt.Println("Please enter your first name: ")
	var firstname string
	fmt.Scanln(&firstname)

	fmt.Println("Please enter your last name: ")
	var lastname string
	fmt.Scanln(&lastname)

	fmt.Println("Please enter your phone number: ")
	var phonenumber string
	fmt.Scanln(&phonenumber)

	fmt.Println("Please enter your Email Address: ")
	var emailaddress string
	fmt.Scanln(&emailaddress)

	fmt.Println("Please enter your IC number: ")
	var icnumber string
	fmt.Scanln(&icnumber)

	fmt.Println("Please enter your Car license number: ")
	var carlicensenum string
	fmt.Scanln(&carlicensenum)

	account := createAccountfordriverapi(firstname, lastname, phonenumber, emailaddress, icnumber, carlicensenum)
	if (account == Driver{}) {
		fmt.Println("Account could not be created")
	} else {
		fmt.Println("Account has been created")
	}

}

//CALL PASSENGER API TO CREATE PASSENGER ACCOUNT
func createAccountforpassengerapi(firstname string, lastname string, phoneNumber string, emailaddress string) Passenger {
	url := "http://localhost:5001/passenger"

	newPassenger := Passenger{
		Firstname:    firstname,
		Lastname:     lastname,
		MobileNo:     phoneNumber,
		EmailAddress: emailaddress,
	}

	jsonNewPassenger, _ := json.Marshal(newPassenger)

	request, err := http.NewRequest(http.MethodPost,
		url, bytes.NewBuffer([]byte(jsonNewPassenger)))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	var passenger Passenger
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode != 200 {
			var errorMsg string
			json.NewDecoder(response.Body).Decode(&errorMsg)
			return passenger
		}

		data, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		json.Unmarshal(data, &passenger)
	}

	return passenger
}

//UPDATE PASSENGER INPUT
func updatePassengeraccount(validemail string) {
	fmt.Println("Please enter your first name: ")
	var firstname string
	fmt.Scanln(&firstname)

	fmt.Println("Please enter your last name: ")
	var lastname string
	fmt.Scanln(&lastname)

	fmt.Println("Please enter your phone number: ")
	var phonenumber string
	fmt.Scanln(&phonenumber)

	fmt.Println("Please enter your Email Address: ")
	var emailaddress string
	fmt.Scanln(&emailaddress)

	account := updateAccountforpassengerapi(validemail, firstname, lastname, phonenumber, emailaddress)
	if (account == Passenger{}) {
		fmt.Println("Account could not be update")
	} else {
		fmt.Println("Account has been updated")
		fmt.Println("Please relog using the new email")
		passengerhome()
	}

}

//UPDATE DRIVER ACCOUNT INPUT
func updateDriveraccount(validemail string) {
	fmt.Println("Please enter your first name: ")
	var firstname string
	fmt.Scanln(&firstname)

	fmt.Println("Please enter your last name: ")
	var lastname string
	fmt.Scanln(&lastname)

	fmt.Println("Please enter your phone number: ")
	var phonenumber string
	fmt.Scanln(&phonenumber)

	fmt.Println("Please enter your Email Address: ")
	var emailaddress string
	fmt.Scanln(&emailaddress)

	fmt.Println("Please enter your Car license number: ")
	var carlicensenum string
	fmt.Scanln(&carlicensenum)

	driveraccount := updateAccountfordriverapi(validemail, firstname, lastname, phonenumber, emailaddress, carlicensenum)
	if (driveraccount == Driver{}) {
		fmt.Println("Account could not be update")

	} else {
		fmt.Println("Account has been updated")
		fmt.Println("Please relog using the new email")
		driverhome()
	}

}

//CALL DRIVER API TO CREATE ACCOUNT
func createAccountfordriverapi(firstname string, lastname string, phoneNumber string, emailaddress string, icnumber string, carlicensenum string) Driver {
	url := "http://localhost:5010/driver"

	newDriver := Driver{
		Firstname:     firstname,
		Lastname:      lastname,
		MobileNo:      phoneNumber,
		EmailAddress:  emailaddress,
		IcNum:         icnumber,
		Carlicensenum: carlicensenum,
		Available:     true,
	}

	jsonNewPassenger, _ := json.Marshal(newDriver)

	request, err := http.NewRequest(http.MethodPost,
		url, bytes.NewBuffer([]byte(jsonNewPassenger)))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	var driver Driver
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode != 200 {
			var errorMsg string
			json.NewDecoder(response.Body).Decode(&errorMsg)
			return driver
		}

		data, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		json.Unmarshal(data, &driver)
	}

	return driver
}

//CALL UPDATE PASSENGER API
func updateAccountforpassengerapi(validemail string, firstname string, lastname string, phoneNumber string, emailaddress string) Passenger {
	url := "http://localhost:5001/passenger/" + validemail

	newPassenger := Passenger{
		Firstname:    firstname,
		Lastname:     lastname,
		MobileNo:     phoneNumber,
		EmailAddress: emailaddress,
	}

	jsonNewPassenger, _ := json.Marshal(newPassenger)

	request, err := http.NewRequest(http.MethodPut,
		url, bytes.NewBuffer([]byte(jsonNewPassenger)))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	var passenger Passenger
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode != 200 {
			var errorMsg string
			json.NewDecoder(response.Body).Decode(&errorMsg)
			return passenger
		}

		data, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		json.Unmarshal(data, &passenger)
	}

	return passenger
}

//CALL UPDATE DRIVER API
func updateAccountfordriverapi(validemail string, firstname string, lastname string, phoneNumber string, emailaddress string, carlicensenum string) Driver {
	url := "http://localhost:5010/driver/" + validemail

	newDriver := Driver{
		Firstname:     firstname,
		Lastname:      lastname,
		MobileNo:      phoneNumber,
		EmailAddress:  emailaddress,
		IcNum:         "",
		Carlicensenum: carlicensenum,
		Available:     true,
	}

	jsonNewDriver, _ := json.Marshal(newDriver)

	request, err := http.NewRequest(http.MethodPut,
		url, bytes.NewBuffer([]byte(jsonNewDriver)))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	var driver Driver
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode != 200 {
			var errorMsg string
			json.NewDecoder(response.Body).Decode(&errorMsg)
			return driver
		}

		data, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		json.Unmarshal(data, &driver)
	}

	return driver
}

//execute driver trip by calling driver api with put request
func startTripDriver(driverid int) Trip {
	url := fmt.Sprintf("http://localhost:5010/driver/starttrip/%d", driverid)

	request, err := http.NewRequest(http.MethodPut,
		url, nil)

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	var trip Trip
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode != 200 {
			var errorMsg string
			json.NewDecoder(response.Body).Decode(&errorMsg)
			return trip
		}

		data, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		json.Unmarshal(data, &trip)
	}

	return trip
}

//execute end driver trip by calling driver api with put request
func endTripDriver(driverid int) Trip {
	url := fmt.Sprintf("http://localhost:5010/driver/endtrip/%d", driverid)

	request, err := http.NewRequest(http.MethodPut,
		url, nil)

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	var trip Trip
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode != 200 {
			var errorMsg string
			json.NewDecoder(response.Body).Decode(&errorMsg)
			return trip
		}

		data, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		json.Unmarshal(data, &trip)
	}

	return trip
}

//function that ask user for input
func createTrip(validemail string) {
	fmt.Println("Please your PickUpPoint: ")
	var pickuppoint string
	fmt.Scanln(&pickuppoint)

	fmt.Println("Please your DropOffPoint: ")
	var dropoffpoint string
	fmt.Scanln(&dropoffpoint)
	account := createtripapi(validemail, pickuppoint, dropoffpoint)
	if (account == Trip{}) {
		fmt.Println("Trip could not be created")
	} else {
		fmt.Println("Trip has been created")
	}

}

//create trip using the function that is mentioned above
func createtripapi(validemail string, pickup string, dropoff string) Trip {
	url := "http://localhost:5020/trip/" + validemail

	newTrip := Trip{
		PickupPoint:  pickup,
		DropoffPoint: dropoff,
	}

	jsonNewTrip, _ := json.Marshal(newTrip)

	request, err := http.NewRequest(http.MethodPost,
		url, bytes.NewBuffer([]byte(jsonNewTrip)))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	var trip Trip
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode != http.StatusCreated {
			var errorMsg string
			json.NewDecoder(response.Body).Decode(&errorMsg)
			return trip
		}

		data, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		json.Unmarshal(data, &trip)
	}

	return trip
}
