## PHONE NUMBER VALIDATION API
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#testing">Testing</a></li>
  </ol>
</details>


## About The Project
You are responsible for validating the users phone number and categorizing them based on the there
country code.

## Getting Started

### Prerequisites

Before this project can be run locally you will need to install [Go](https://golang.org/doc/install)

### Installation

To utilize the project, the following needs to be done:
1. Clone this repository
2. Install the dependencies, using the following command:
```
go mod tidy
```

## Usage

1. To run the project locally, use the following command:
```
make run
```
2. To retrieve the mobile numbers based on the country and code , use Postman to make a GET request to the following URL:
```
http://localhost/api/v1/user/mobile/:code/:state
```
3. To download all the frontend dependencies Run : npm install
## Testing
Tests can be run using the following command:
```
make test
```


TO PREPOPULATE YOU DATABASE RUN THE SQL on the SQLITE3

INSERT INTO users ( user_name, mobile) VALUES (“Mary Cockill", "(212) 934962214”);
INSERT INTO users ( user_name, mobile) VALUES (“John Cockill", "(212) 595925230”);
INSERT INTO users ( user_name, mobile) VALUES (“Peter Cockill", "(212)  789157465”);
INSERT INTO users ( user_name, mobile) VALUES (“Taiwan Cockill", "(212)   6611488”);
INSERT INTO users ( user_name, mobile) VALUES (“Emeka Cockill", "(251) 446871685”);
INSERT INTO users ( user_name, mobile) VALUES (“Sandra Cockill", "(251) 306229918”);
INSERT INTO users ( user_name, mobile) VALUES (“Henry Cockill", "(251) 946163891”);
INSERT INTO users ( user_name, mobile) VALUES (“Sean Cockill", "(251) 107459230”);
INSERT INTO users ( user_name, mobile) VALUES (“Joy Cockill", "(258) 854934161”);
INSERT INTO users ( user_name, mobile) VALUES (“Anthony Cockill", "(258) 284851159”);
INSERT INTO users ( user_name, mobile) VALUES (“Faith Cockill", "(258) 21389685”);
INSERT INTO users ( user_name, mobile) VALUES (“Michael Cockill", "(258) 20823916”);
INSERT INTO users ( user_name, mobile) VALUES (“Collin Cockill", "(256) 712531962”);
INSERT INTO users ( user_name, mobile) VALUES (“Bayo Cockill", "(256) 139238264”);
INSERT INTO users ( user_name, mobile) VALUES (“Mark Cockill", "(256) 152121547”);
INSERT INTO users ( user_name, mobile) VALUES (“Lauretta Cockill", "(256) 753477588”);
INSERT INTO users ( user_name, mobile) VALUES ("Mary Cockill", "(212) 934962214");
INSERT INTO users ( user_name, mobile) VALUES ("John Cockill", "(212) 595925230");
INSERT INTO users ( user_name, mobile) VALUES ("Peter Cockill", "(212)  789157465");
INSERT INTO users ( user_name, mobile) VALUES ("Taiwan Cockill", "(212)   6611488");
INSERT INTO users ( user_name, mobile) VALUES ("Emeka Cockill", "(251) 446871685");
INSERT INTO users ( user_name, mobile) VALUES ("Sandra Cockill", "(251) 306229918");
INSERT INTO users ( user_name, mobile) VALUES ("Henry Cockill", "(251) 946163891");
INSERT INTO users ( user_name, mobile) VALUES ("Sean Cockill", "(251) 107459230");
INSERT INTO users ( user_name, mobile) VALUES ("Joy Cockill", "(258) 854934161");
INSERT INTO users ( user_name, mobile) VALUES ("Anthony Cockill", "(258) 284851159");
INSERT INTO users ( user_name, mobile) VALUES ("Faith Cockill", "(258) 21389685");
INSERT INTO users ( user_name, mobile) VALUES ("Michael Cockill", "(258) 20823916");
INSERT INTO users ( user_name, mobile) VALUES ("Collin Cockill", "(256) 712531962");
INSERT INTO users ( user_name, mobile) VALUES ("Bayo Cockill", "(256) 139238264");
INSERT INTO users ( user_name, mobile) VALUES ("Mark Cockill", "(256) 152121547");
INSERT INTO users ( user_name, mobile) VALUES ("Lauretta Cockill", "(256) 753477588");
INSERT INTO users ( user_name, mobile) VALUES ("Faith Cockill", "(212) 21389685");
INSERT INTO users ( user_name, mobile) VALUES ("Michael Cockill", "(212) 595925230");
INSERT INTO users ( user_name, mobile) VALUES ("Collin Cockill", "(212) 661148867");
INSERT INTO users ( user_name, mobile) VALUES ("Bayo Cockill", "(212) 911325714");
INSERT INTO users ( user_name, mobile) VALUES ("Mark Cockill", "(212) 862213494");