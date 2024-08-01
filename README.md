## Scripts Used

This script will install the package and create a "go.mod" file
```
go mod init github.com/radikaledward1/golang-rest-api-postgresql
```

This script will create/download the package in the c:/go/pkg/mod/github.com/gorilla/mux@v1.8.1
```
go get -u github.com/gorilla/mux@latest
```

This package install AIR Live Reload for Go on the Computer, (Is not a package for the project) similar as "CompileDaemon" command line Utility
````
go install github.com/air-verse/air@latest
````
Next, Run the next script to create the Air Config File
````
air init 
````
Use this script to use the Air Auto Reloaded and Detect any new change in the code
````
air
````
This script check, update and prune unneccesary dependencies
````
go mod tidy
````

Run the project
```
go run .
```

## Database

Run Dockerfile on the data folder
````
docker build -t gorestapi-pgsql .
````

Run the container created by the image and add the env variables.
````
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=secretpassword -e POSTGRES_DB=GO_CRUD --name gorestapi-db gorestapi-pgsql
````

Create a connection with the PostgreSql Database using your favorite DB Client and use the configurations:
````
host: localhost
database: <no default database selected>
user: postgres
password: secretpassword
````

The project uses GORM as ORM, to install run the next script:
````
go get -u gorm.io/gorm
````

Run the next script to install the driver for PostgreSQL
````
go get -u gorm.io/driver/postgres 
````

## Packages
https://github.com/gorilla/mux

https://gorm.io/docs/connecting_to_the_database.html

https://github.com/air-verse/air