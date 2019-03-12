# Lana Go

[![Build Status](https://travis-ci.com/danteay/lanago.svg?branch=master)](https://travis-ci.com/danteay/lanago)

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/aa577ae45ebdeca472a1#?env%5BLanaGo%5D=W3sia2V5IjoiaG9zdCIsInZhbHVlIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwIiwiZGVzY3JpcHRpb24iOiIiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6ImJhc2tldElkIiwidmFsdWUiOiJjOWEwZDhmOS1jMmYxLTRlOWUtODZhMy0zZDFkMDBiMDk4MmUiLCJkZXNjcmlwdGlvbiI6IiIsImVuYWJsZWQiOnRydWV9LHsia2V5IjoicHJvZEhvc3QiLCJ2YWx1ZSI6Imh0dHBzOi8vbGFuYWdvLmhlcm9rdWFwcC5jb20iLCJkZXNjcmlwdGlvbiI6IiIsInR5cGUiOiJ0ZXh0IiwiZW5hYmxlZCI6dHJ1ZX1d)

## Dependencies

* Golang >= 1.11
* Docker
* [Heroku-Cli](https://devcenter.heroku.com/articles/heroku-cli)
* [Govendor](https://github.com/kardianos/govendor/)
* [Gin-gonic as REST framework](https://gin-gonic.com/)
* [Godotenv as env loader](https://github.com/joho/godotenv)
* [GinRest as response helper](https://github.com/danteay/ginrest)
* [GoRedis as redis client](https://github.com/go-redis/redis)

### Project structure

The project was made with Gin-gonic framework, because this framework brings a faster development, great
support and a fastest resolve of the router. Gin-gonic has nothing except the Http handling, and this is
one of the reasons to get this framework in the project, getting fully control of all the other project
dependencies like the DB managers and configuration managers.

Godotenv was selected to help with the environment management, having the possibility of change different configurations easily, like local configurations, production and testing.

Redis was selected to have a semi persistent track of the baskets, this allow us to auto clean the
incomplete process that never go to the checkout (this will be conbained with persistent baskets in
checkout process and cookies)

### Project directories

The project structure has the next configuration:

```text
+ api/
|  + config/
|  + handlers/
|  + libs/
|  + models/
|  + rotes/
|  + services/
+ build/
|  + build.sh
|  + Dockerfile
|  + Makefile
+ config/
|  + dev.env
|  + test.env
|  + prod.env
+ out/
+ vendor/
|  + vendor.json
+ main.go
+ exec
+ .travis.yml
```

* **api**: Holds all the related code with the endpoints, libraries and business logic.
  * **config**: Dependency configurations
  * **handler**: Route handlers logic
  * **libs**: Specific libraries created for the project
  * **models**: Structure definitions
  * **routes**: All the route configurations
  * **services**: Libraries that handles business logic and external resources like calls to other micro services
* **build**: Holds all the necesary files to build the proyect in production and other stage needed. This files are
  make files, docker files and bash files
* **config**: Holds the environment configurations like dev, test and production
* **out**: This specific folder is created when the build command is triggered and holds the builded file of the project
* **vendor**: This is the dependency folder that is managed by Govendor

#### Exec file

This is a bash file that have preconfigured options like test, run and build and is used in the Travis-CI pipeline to deploy the application. this are the availabe commands:

```bash
./exec test      # Run tests with it's corresponding configurations
./exec run       # Run the project localy with the command 'go run'
./exec run build # Run the last builded docker project and is necesary to build it first before run this command
./exec build     # Run build the docker project
./exec publish   # Push the last docker build to the heroku container registry
./exec release   # Release the last docker images pushed to heroku in production
```

## Local deployment

To run the project locally first you need to have installed the dependencies listed above, then clone this repository inside `GOPATH/src/github.com/danteay/lanago`:

``` bash
git clone https://github.com/danteay/lanago.git $GOPATH/src/github.com/danteay/lanago
```

Then you just need to run the next command and you will have a redy local server runin on port 8080:

```bash
cd $GOPATH/src/github.com/danteay/lanago
./exec run
```
