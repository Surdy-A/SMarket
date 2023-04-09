<h1 align="center"> City Smile Realty Web Application </h1> <br>




## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Technology](#Technology-Used)
- [Running](#Running-the-application)
- [Layout](#Layout)




## Introduction

City Smile Realty: Is a Real Estate Market Place. You can Search millions of for-sale and rental listings.


## Features
Here are some of the features:
- Add Property for rent and sale
- Delete Property
- Edit Property

## Technology Used
- Golang
- Gonic Gin
- CSS
- Javascript
- HTML

## Running the App
Perform the following steps to run the application:

1- Clone the App

```
$ git clone https://github.com/Surdy-A/City-Smile-Realty.git
``` 

2- Setup the environment variables and database
``` bash
$ source .env
```

3- Change directory into the City-Smile-Realty and run

```
$ go run main.go
``` 

Server is listening on localhost:8010

## Test

```bash
$ go test -v
=== RUN   TestEmptyTable
--- PASS: TestEmptyTable (0.00s)
=== RUN   TestGetNonExistentProduct
--- PASS: TestGetNonExistentProduct (0.00s)
=== RUN   TestCreateProduct
--- PASS: TestCreateProduct (0.00s)
=== RUN   TestGetProduct
--- PASS: TestGetProduct (0.00s)
=== RUN   TestUpdateProduct
--- PASS: TestUpdateProduct (0.01s)
=== RUN   TestDeleteProduct
--- PASS: TestDeleteProduct (0.01s)
PASS
ok      _/home/tom/r/go-mux-api 0.034s
```

## License

Copyright (c) 2023 Rendered Text

Distributed under the MIT License. See the file LICENSE.
