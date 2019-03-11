# Lana Go challenge

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/aa577ae45ebdeca472a1#?env%5BLanaGo%5D=W3sia2V5IjoiaG9zdCIsInZhbHVlIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwIiwiZGVzY3JpcHRpb24iOiIiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6ImJhc2tldElkIiwidmFsdWUiOiJjOWEwZDhmOS1jMmYxLTRlOWUtODZhMy0zZDFkMDBiMDk4MmUiLCJkZXNjcmlwdGlvbiI6IiIsImVuYWJsZWQiOnRydWV9LHsia2V5IjoicHJvZEhvc3QiLCJ2YWx1ZSI6Imh0dHBzOi8vbGFuYWdvLmhlcm9rdWFwcC5jb20iLCJkZXNjcmlwdGlvbiI6IiIsInR5cGUiOiJ0ZXh0IiwiZW5hYmxlZCI6dHJ1ZX1d)

## Dependencies

* Golang >= 1.11
* Docker
* [Heroku-Cli](https://devcenter.heroku.com/articles/heroku-cli)
* [Govendor](https://github.com/kardianos/govendor/)

### Specific project dependencies

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
