key-value-server
================

A single node key value server

1.A single node key value server that maintains a map of student details with a chosen key by the user.
2.Database 1.Rollono 2.Address 3.phone

Features:
1.Implemented using ZEROMQ for communication between the clent and the server.
2.Supports mutiple clients that can connect a central server which maintains a key-value database
3.Used mutex based synchoronization for read-write access to the database.

Files include
1.client.go
2.server.go
3.server_test.go

Usage:
Uses ZEROMQ.install the zeromq 3.0
-----------

1. go get github.com/vamsikrishnacs/key-value-server
2. go install github.com/vamsikrishnacs/client.go
3. go install github.com/vamsikrishnacs/server.go

a.go run server.go&
b.go run client.go(1)
c.go run client.go(2)

Operations:
1.createuser(id)
2.filldetails(id,rollno,address,phoneno)
3/4/5 change detail(id,roll/address/phone)
6.getMydetails(id)
7.DeleteUSer

Testing:
Supports the tests for
1.Eliminating duplicate Users
2.test for deletion and update and search
3.tested with multiple clients

a.go test server_test.go
