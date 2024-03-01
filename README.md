## Quick start

Requiements:

    docker
    docker-compose
    make
    go 1.18 or higher

## 1. Run server

`make dev-up`

`go run ./cmd/server/ -e .env`

## 2. Start client

For first time registration:

`go run ./cmd/client/ -a localhost:8080 signup -u username -p password`

Login:

`go run ./cmd/client/ -a localhost:8080 login -u username -p password`

## Description

Passman is a client-server application used for storing private data:

    Passwords
    Bank cards
    Notes
    Files

All data encryption/decryption operations is carried out on the server side.
Server encrypts data using symmetric encryption based on the AES algorithm 
(block size = 128, key size = 256, mode = CTR) with HMAC SHA-512 validation.
AES keys generates for each users vault during registration and stored on the
server in encrypted form using AEG-GCM and user password. When user logged in 
server decrypts user vault AES key using user password. 
Private data is not stored locally on client side.
Files are transferred in chunks of 4 KiB.
For user authorisation and authentication are used JWT tokens. Token is not 
stored on the client side and generates for each user signup/login.


## Run server

Server uses postgresql for storing all types of private data.
Use `make dev-up` to start the database with docker-compose.

Server gets all configuration parameters from environment variables.
It also can parse env vars from the file with -e flag: `go run ./cmd/server/ -e .env` 

## Using the client

Client needs only the server address "host:port", user login and password.
Run `go run ./cmd/client/ help` for help.

## Encrypted data format

                        bits
    +-----------------------------------------+
    | version | randIV | encryptedData | hmac |
    +-----------------------------------------+
        (8)     (128)      (varies)      (512)