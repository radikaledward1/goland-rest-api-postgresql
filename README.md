## Scripts Used

This script will install the package and create a "go.mod" file
```
go mod init gorestapipos
```

This script will create/download the package in the c:/go/pkg/mod/github.com/gorilla/mux@v1.8.1
```
go get -u github.com/gorilla/mux@latest
```

This script check, update and prune unneccesary dependencies
````
go mod tidy
````

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

## Packages
https://github.com/gorilla/mux

https://github.com/githubnemo/CompileDaemon

## Notes
- Use CompileDaemon -command="./gorestapi.exe", to auto-compile you code after save changes.