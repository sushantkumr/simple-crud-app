# simple-crud-app
Simple CRUD Application written in Golang

## Golang

### Creating Go module for the directory
go mod init simple-crud-app/iteamstore-app


###  Steps to run the setup as a standalone Go program
go run main.go

### To create a build
go build


## Docker

### Build Docker image
docker build -t your-image-name:your-tag .

### Run Docker image
docker run -p 8080:8080 your-image-name:your-tag
