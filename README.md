# Posty

![posty cover](/assets/cover.png)

> Modular service to add posts written in golang

## Tasks

- [x] Create an User
- [x] Get a user using id
- [x] Create a Post
- [x] Get a post using id
- [x] List all posts of a user
- [x] Authentication without using JWT
- [x] Make the server thread safe
- [ ] Add pagination to the list endpoint
- [ ] Add unit tests
- [x] Clean Architecture

## Instructions to run

### Direct run

- Pre-requisites:

  - Go

- Installation:

```bash
git clone https://github.com/r-ush/posty.git
cd posty
go mod download
```

- Execution

```bash
DB_URL=<encoded db url> go run api/main.go
```

> loading env because godotenv package could not be used

<!-- ### Run using docker

```bash
docker
``` -->

## Architecture

- This project is built in `Clean Architecture`, it contains of two main modules, i.e. api and pkg.

- `service` acts as usecase layer

- `repository` as repository layer

- `api` contains all the necessary route handlings and backend supporting services(i.e. receiving requests and forwarding to proper handlers) it contains packages:
  - `main`: contains main.go.
  - `presenter`: contains structs for response conditions
  - `handlers`: contains all the necessaey handlers and linking with services, which in turn response using views.
- `pkg` contains the business logic divided into couple of packages
  - `pkg`: contains centralized errors.go file defining all the necessary errors which will thrown from backend and pkg
  - `user`: contains all the neccessary files for user business logic
  - `post`: contains all the neccessary files for post business logic
  - `entities`: necessary middle man structs for holding participants and meetings data from db and so forth
- `utils`: contains all necessary files for helping functions

## Other notable features

- Authentication
  - Uses SHA256 hashing
  - Every post made has to be authenticated with the password of the user (similar to the case with pushing to github on an https remote)
- packages
  - made using only [std](https://pkg.go.dev/std) packages and [go-mongo driver](go.mongodb.org/mongo-driver)
