<h1 align="center">SMarket E-Commerce API </h1> <br>




## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Technology](#Technology-Used)
- [Running](#Running-the-application)



## Introduction

SMarket: Is an E-Commerce REST API. It can be used to perform all CRUD operations on a product.


## Features
Here are some of the features:
- Create Product
- Get Product(s)
- Update Product
- Delete Product
- Authenticate Users
- 
## Technology Used
- Golang

## Running the App
Perform the following steps to run the application:

1- Clone the App

```
$ git clone https://github.com/Surdy-A/SMarket.git
``` 

2- Setup the environment variables and database
``` bash
$ source .env
```

3- Change directory into the SMarket and run

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
