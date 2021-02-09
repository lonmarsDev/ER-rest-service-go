# ER Golang API Service


### Installation

golang requires [golang.org](https://golang.org/) go1.11+ to run.

Install the dependencies and devDependencies and start the server.


## How to run the app in local
In the root project directory.
```bash
$ go run cmd/score_service/main.go
```

## How to run the app in local docker image
In the root project directory.
```bash
$ ./scripts/docker/run.sh
```

## How to use 
Use postman in localhost:8080/
method: POST


## Use in CURL
```bash
curl --location --request POST 'http://localhost:8080/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "scores": {
        "managers": [
            { "userId": 1, "score": 1 },
            { "userId": 2, "score": 5 }
        ],
        "team": [
            { "userId": 4, "score": 1 },
            { "userId": 5, "score": 5 },
            { "userId": 6, "score": 3 },
            { "userId": 7, "score": 3 }
        ],
        "others": [
            { "userId": 8, "score": 1 },
            { "userId": 9, "score": 5 }
            ]
        }
}'
```