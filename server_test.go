package main

import "testing"

import (
	"encoding/json"
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

type client struct {
	socket *zmq.Socket
}

type request struct {
	Op   int
	Id   int
	Roll int
	Addr string
	Ph   string
}

type reply struct {
	Success bool
	Id      int
	Roll    int
	Addr    string
	Ph      string
}

func (c *client) send(msg request) reply {

	fmt.Println(msg)
	message, _ := json.Marshal(msg)
	c.socket.SendBytes(message, 0)
	fmt.Println("Sending")
	resp, err := c.socket.RecvBytes(0)
	var rep reply
	json.Unmarshal(resp, &rep)
	if err != nil {
		fmt.Println(err)
	}
	return rep

}

///specially written to write test code,..can be written in main as well
func (c *client) prepare(op int, id int, r int, a string, p string) reply {
	var m request

	m = request{op, id, r, a, p}
	rep := c.send(m)
	return rep
}

func Test(t *testing.T) {

	var cl client
	var cl1 client
	cl.socket, _ = zmq.NewSocket(zmq.REQ)

	err := cl.socket.Connect("tcp://127.0.0.1:5001")

	fmt.Println(err)
	fmt.Println("Initialisng the client........")

	fmt.Println("User options 1.register 2.filldetails(roll,address,phone) 3.change_rollno 4.change_adress 5.change_ph 6.getdetails 7.deleteme")
	cl1.socket, _ = zmq.NewSocket(zmq.REQ)

	cl1.socket.Connect("tcp://127.0.0.1:5001")

	var op, id, r int
	var a, p string

	//duplicate user detection test
	var rep reply
	op, id, r, a, p = 1, 12, r, a, p
	cl.prepare(op, id, 0, "", "")
	rep = cl.prepare(1, 12, 0, "", "")
	if rep.Success != true {

		t.Errorf("Same id inserted twice--error")
	}
	//retrieve user test
	op, id, r, a, p = 2, 12, 123, "iit", "929393939"
	cl.prepare(op, id, r, a, p)
	op, id, r, a, p = 6, 12, 0, "", ""
	rep = cl.prepare(op, id, r, a, p)

	if rep.Success != true {

		t.Errorf("Unable to retrieve the user")
	}
	//Non user dtection--shouldnot reply to non registered users
	op, id, r, a, p = 6, 13, 0, "", ""
	rep = cl.prepare(op, id, r, a, p)
	if rep.Success == true {

		t.Errorf("Error--Non key retrieved")
	}

	//delete followed by search
	rep = cl.prepare(7, 12, 0, "", "")
	rep = cl.prepare(6, 12, 0, "", "")
	if rep.Success == true {

		t.Errorf("Error-the user was just deleted")
	}

	//multiple clients
	cl.prepare(1, 12, 0, "", "")
	rep = cl1.prepare(1, 12, 0, "", "")

	if rep.Success != true {

		t.Errorf("Conflict-User was created by client 1--should output user already exist")
	}

}
