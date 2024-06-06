# Backend

## Beforehand

The Database chosen in the project was [Postgre](https://www.postgresql.org/) so installing it is necessary. Creating a database is the next step, you could use the terminal with the following command `CREATE DATABASE databse-name`, or use a DBMS like [pgAdmin 4](https://www.pgadmin.org/). In the development of this app, pgAdmin 4 was used.

After creating the database, write its name down on the .env with the root user, and the password, or a user with rights for that specific database.

### Docker

Windows 10 Pro or latest MacOs, or latest linux versions can run Dockers Desktop. Another requirement is HYPER-V support enable.

First see if you have access to hyper-v with the following command in power-shell: `systeminfo`.

If is not enable there is two ways of enabling it:
1. `Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V -All`

2. `DISM /Online /Enable-Feature /All /FeatureName:Microsoft-Hyper-V`

When you already done activating it, you can proceed and create the image of your app, you have to run it inside the folder project, with the following command: `docker-compose up --build`

If the creation of the image is successful you can go and execute the *container* on **Docker Desktop**.

## Create .env file

To run the project you need to create a new .env file. The content of this file is in the template.

## Running

In GO you run the project with the command `go run file-name`, typically you would run a main.go.

To run this project use `go run ./cmd/main.go`.
The reason this is different is because of the go module in which is define the directory path of the project, in this case the path for the project is `/backend`.

The server use `http://localhost:8080/`. The 8080 port have to be free for the project to compile correctly.


## Mod file
While creating your own go project, if your going to use libraries or define folders for your projects then you have to create one. Use `go mod init project-path`

## Build

Run `go build` to build the project. The build will produce a binary, in windows a .exe is produce.

## Testing

Testing in Go can be done in multiple [ways](https://theifedayo.medium.com/a-comprehensive-guide-to-unit-testing-in-go-c292f20670b0) . In this case the test was made with testify for testing functions.

`go get github.com/stretchr/testify/assert`

Database mocks were used for emulating the real database behavior.

`go get github.com/DATA-DOG/go-sqlmock`

Sqlite is used along with sql-mock for faked sql databases.
`go get github.com/mattn/go-sqlite3`

Now, with all that install you can run the test. For that, first, you have to be positioned in the folder in which the file is located and run the following command:

`go test -v -run TestGetRecords`

## Further help

To get more help on the CLI use `go help` or go check out the [GO Help](https://pkg.go.dev/cmd/go/internal/help) page.
