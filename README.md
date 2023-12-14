install  require golang
1. https://go.dev/  versi latest
2. https://gorm.io/  gorm  driver mysql
3. go mod tidy 

Copy the .env.example file and rename it to .env
Change the config for your local server

DB_HOST=      localhost
DB_PORT=      3306
DB_USER=      root
DB_PASSWORD=
DB_NAME=      go-course
SERVER_PORT=  8080

 
 Running Server
 go run server.go
