<div id="top"></div>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#Assignment">Assignment Objective</a>
      <ul>
        <li><a href="#AssignmentRequirements">Assignment Requirements</a></li>
        <li><a href="#AssignmentObjectives">Assignment Objectives</a></li>
        <li><a href="#DesignConsiderationsforthemicroservices">Design Considerations for the microservices</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Welcome to Getaride a simple ride hailing program made with GOLANG and GORM, involves CLI to use the application.
<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

* [GOLANG](https://go.dev/)
* [GORM](https://gorm.io/index.html)
* [MYSQL](https://www.mysql.com/)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- Assignment Objective-->
## AssignmentRequirements

* Minimum 2 microservices using GOLANG
* Persistent storage of information using database

## #AssignmentObjectives
* To be able to develop REST api
* Able to communicate between the api's

## DesignConsiderationsforthemicroservices
* The microservices have been created in such a way that they are uniquely individual.
* An example would be the passenger microservice, which was created solely for passengers and database communication.
* The same can be said for the rest of the microservices, such as Trip, which can only communicate with Trip database, and Driver, which can only communicate with Driver database.

*Using GORM, an Object Relational Mapping Library.
The GORM is a good Golang ORM package that attempts to be developer-friendly. It's an object-relational mapping (ORM) library for dealing with relational databases. The database/sql package is used to build this gorm library. an example would be instead of using query when excuting a SQL line, instead using GORM it simplifies the execution and creation of the database. Another example of using GORM is simplifying the database creation, if the table or database does not exist, GORM will automatically create and insert the table in the database along with the data that is posted.




<img src="images/Architecture Diagram.jpg" alt="Logo" width="1080" height="720">

For the Getaride application, there are 3 different microservices used and command line to execute the console application,
The rest API communicates with the used of HTTP GET POST PUT methods, such as creating the passenger account, it will issue a POST request and from there the information that is inputted will be send to the database for storing, The logic and data handling may then be handled within each Microservice, all while adhering to the loosely coupled philosophy that Microservices is known for.

The Getaride application consist of 

Passenger Microservices
* Create Passenger (POST)
* Update Passenger Details (PUT)
* Create Trip by calling the Trip Microservice (POST)
* Get all trips by calling the Trip Microservice (GET)

Driver Microservices
* Create Driver (POST)
* Update Driver Details (PUT)
* Start Trip (PUT) by calling the Trip Microservice (POST)
* Stop Trip (PUT) by calling the Trip Microservice (POST)
* Get all trips by driver id calling the Trip Microservice (GET)

Trip
* Creates TRIP (Passenger Creates Trip) - uses POST method to create trip, GET method to call passenger and driver microservices for the informations such as passenger ID and driver availibilty
* Start Trip (PUT) - when driver starts trip this function will be called by API
* End Trip (PUT) - when driver ends trip this function will be called by API


### Prerequisites

GOLANG and MYSQL must be installed in order for the program to work

1. SQL information
  ```sh
  Username : Root
  Password : Root
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/theazziizzuan/ETIAssignment1-Azzi.git
   ```
2. Install libraries
   ```sh
    go get -u github.com/go-sql-driver/mysql
    go get -u github.com/gorilla/mux
    go get -u github.com/gorilla/handlers
    go get -u gorm.io/gorm
   ```
3. Execute database script located in /database/InitDB.sql
    
    
<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1. Start the Microservices
   ```sh
   cd /Assignment1/Passenger
   go run Passenger.go
   ```
    ```sh
   cd /Assignment1/Driver
   go run Driver.go
   ```
    ```sh
   cd /Assignment1/Trip
   go run Trip.go 
   ```
   
2. Start the console app
   ```sh
    cd /Assignment1/Console
    go run main.go
   ```
<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap
```sh
- [1] Create Database
- [2] Create Microservices REST api using GOLANG
- [3] Console Application that call all the REST api
```


<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTACT -->
## Contact
School Email
```sh
S10189579@connect.np.edu.sg
```

<p align="right">(<a href="#top">back to top</a>)</p>




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/github_username/repo_name.svg?style=for-the-badge
[contributors-url]: https://github.com/github_username/repo_name/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/github_username/repo_name.svg?style=for-the-badge
[forks-url]: https://github.com/github_username/repo_name/network/members
[stars-shield]: https://img.shields.io/github/stars/github_username/repo_name.svg?style=for-the-badge
[stars-url]: https://github.com/github_username/repo_name/stargazers
[issues-shield]: https://img.shields.io/github/issues/github_username/repo_name.svg?style=for-the-badge
[issues-url]: https://github.com/github_username/repo_name/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/github_username/repo_name/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
