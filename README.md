create project with:

~~~cmd
    go mod init {your-package name}
~~~

I used:

~~~cmd
    go mod go-echo
~~~


install echo

~~~cmd
    go get github.com/labstack/echo/v4
~~~

copy paste the server.go content


run:

~~~cmd
    go run server.go
~~~


http://localhost:3000/hello


---

## Testing

The test goes in the same folder of the file to test

do not use the TestMain method in the same file of the other test methods

got test to run all the tests

---

do not do:

copy the downloaded dependencies to our project's folder that we are going to call vendor with command

~~~cmd
go mod vendor
~~~