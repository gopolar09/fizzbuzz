# FizzBuzz
Simple application showing FizzBuzz


## Frontend Installation

To run this application, you need to have Node.js and NPM (or Yarn) installed on your computer.

```
cd frontend
npm install
```

Create a .env file in the root directory by coping .env.example

### Start the application:

In the project directory, you can run:

`npm run start`

Open [http://localhost:3000](http://localhost:3000) to view it in the browser.


## Backend Installation

To run this application, you need to have Go installed on your computer.

```
cd backend
go mod download
```

Create a .env file in the root directory by coping .env.example

### Start the application:

In the project directory, you can run:

`go run main.go`


## Usage

If the number is a multiple of 3, FE shows the message “Fizz”
If the number is a multiple of 5, FE shows the message “Buzz”
If the number is a multiple of 3 and 5, FE shows the message “FizzBuzz”. It doesn’t mean showing total 3 messages “Fizz”, “Buzz”, and “FizzBuzz”, but only “FizzBuzz”.

To use the API, send a POST request to /fizzbuzz with a JSON payload containing a count parameter:

The API will respond with a JSON object containing a message parameter with the FizzBuzz value for the given count

```
{
    "status": 200,
    "message": "FizzBuzz"
}
```



## License