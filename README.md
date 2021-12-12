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
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
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
## Assignment Requirements

* Minimum 2 microservices using GOLANG
* Persistent storage of information using database

## Assignment Objectives
* To be able to develop REST api
* Able to communicate between the api's

## Design Considerations for the microservices
* The microservices have been created in such a way that they are uniquely individual.
* An example would be the passenger microservice, which was created solely for passengers and database communication.
* The same can be said for the rest of the microservices, such as Trip, which can only communicate with Trip database, and Driver, which can only communicate with Driver database.

*Using GORM, an Object Relational Mapping Library.
The GORM is a good Golang ORM package that attempts to be developer-friendly. It's an object-relational mapping (ORM) library for dealing with relational databases. The database/sql package is used to build this gorm library. an example would be instead of using query when excuting a SQL line, instead using GORM it simplifies the execution and creation of the database. Another example of using GORM is simplifying the database creation, if the table or database does not exist, GORM will automatically create and insert the table in the database along with the data that is posted.

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
* Creates TRIP (Passenger Creates Trip) - uses POST method to create trip and GET method to call passenger and driver microservices
* 
* 

Trip Microservices


### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* npm
  ```sh
  npm install npm@latest -g
  ```

### Installation

1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo
   ```sh
   git clone https://github.com/github_username/repo_name.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
4. Enter your API in `config.js`
   ```js
   const API_KEY = 'ENTER YOUR API';
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [] Feature 1
- [] Feature 2
- [] Feature 3
    - [] Nested Feature

See the [open issues](https://github.com/github_username/repo_name/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Your Name - [@twitter_handle](https://twitter.com/twitter_handle) - email@email_client.com

Project Link: [https://github.com/github_username/repo_name](https://github.com/github_username/repo_name)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [y]()
* []()
* []()

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
