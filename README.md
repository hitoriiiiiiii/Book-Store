# 📚 Go Bookstore API

A simple Bookstore REST API built using Go (Golang) and MySQL.
Fully Dockerized so anyone can run it easily.

------------------------------------------------------------

# FEATURES
- Create book
- Get all books
- Get book by ID
- Update book
- Delete book

------------------------------------------------------------

# 🛠 TECH STACK
- Go (Golang)
- MySQL
- GORM
- Docker
- Docker Compose

------------------------------------------------------------

# PROJECT STRUCTURE

Go-bookstore/
|-- cmd/main.go
|-- pkg/
|   |-- controllers
|   |-- models
|   |-- routes
|   |-- utils
|-- Dockerfile
|-- docker-compose.yml
|-- .env.example
|-- README.md

------------------------------------------------------------

#RUN WITH DOCKER

1. Clone repository
git clone https://github.com/your-username/go-bookstore.git
cd go-bookstore

2. Create env file
cp .env.example .env

3. Run project
docker-compose up --build

API will run on:
http://localhost:8080

------------------------------------------------------------

#API ENDPOINTS

POST    /book/        -> Create book
GET     /book/        -> Get all books
GET     /book/{id}    -> Get book by ID
PUT     /book/{id}    -> Update book
DELETE  /book/{id}    -> Delete book

------------------------------------------------------------

# DOCKER IMAGE

docker pull prarthana25/go-bookstore:latest

------------------------------------------------------------
# 🧪 Test with Postman

Base URL: http://localhost:8080

Content-Type: application/json

#AUTHOR
Prarthana Gade

------------------------------------------------------------

If you like this project, give it a star ⭐


# ⚙️ Environment Setup

Create a `.env` file using this format:

```env
DB_USER=root
DB_PASSWORD=root
DB_HOST=db
DB_PORT=3306
DB_NAME=bookstor

Create a .env file:

DB_USER=root
DB_PASSWORD=root
DB_HOST=db
DB_PORT=3306
DB_NAME=bookstore 

NOTE:
- Do NOT push .env to GitHub
- Push .env.example only
------------------------------------------------------------
