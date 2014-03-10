package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"sync"
)

type database struct {
	mutex sync.Mutex
	Map   map[int]*mydetails
}

type mydetails struct {
	Roll    int
	Address string
	phone   string
}

func (d *database) put(id int) bool {
	fmt.Println("put")
	_, ok := d.Map[id]
	fmt.Println(ok)
	return ok
}

func (d *database) createuser(id int) bool {

	ok := d.put(id)
	if !ok {
		d.mutex.Lock()
		d.Map[id] = &mydetails{0, "", ""}
		d.mutex.Unlock()
	}
	return ok
}
func (d *database) putroll(id int, roll int) {
	ok := d.put(id)
	if ok {
		d.mutex.Lock()
		d.Map[id].Roll = roll
		d.mutex.Unlock()
	}
}

func (d *database) putAddress(id int, addr string) {
	ok := d.put(id)
	if ok {
		d.mutex.Lock()
		d.Map[id].Address = addr
		d.mutex.Unlock()
	}
}

func (d *database) putphone(id int, ph string) {
	ok := d.put(id)
	if ok {
		d.mutex.Lock()
		d.Map[id].phone = ph
		d.mutex.Unlock()
	}
}

func (d *database) getroll(id int) int {
	_, ok := d.Map[id]
	s := 0
	if ok {
		d.mutex.Lock()
		s = d.Map[id].Roll
		d.mutex.Unlock()
	} else {
		s = 0
	}
	return s
}

func (d *database) getAddress(id int) string {
	_, ok := d.Map[id]
	s := ""
	if ok {
		d.mutex.Lock()
		s = d.Map[id].Address
		d.mutex.Unlock()
	} else {
		s = ""
	}
	return s
}
func (d *database) getphone(id int) string {
	_, ok := d.Map[id]
	s := ""
	if ok {
		d.mutex.Lock()
		s = d.Map[id].phone
		d.mutex.Unlock()
	} else {
		s = ""
	}
	return s
}

func (d *database) getMydetails(id int) (*mydetails, bool) {
	fmt.Println(id)
	_, ok := d.Map[id]
	var s *mydetails
	if ok {
		d.mutex.Lock()
		s, _ = d.Map[id]
		d.mutex.Unlock()
	} else {
		s = &mydetails{0, "", ""}
	}
	return s, ok
}

func (d *database) deleteMydetails(id int) {
	d.mutex.Lock()
	delete(d.Map, id)
	d.mutex.Unlock()
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

func (d *database) parse(msg request) reply {
	var r reply
	switch msg.Op {
	case 1:
		//createuser
		ok := d.createuser(msg.Id)
		r = reply{ok, msg.Id, 0, "", ""}

	case 2:
		//putdetails  
		d.putroll(msg.Id, msg.Roll)
		d.putAddress(msg.Id, msg.Addr)
		d.putphone(msg.Id, msg.Ph)
		r = reply{true, msg.Id, 0, "", ""}

	case 3:
		d.putroll(msg.Id, msg.Roll)
		r = reply{true, msg.Id, 0, "", ""}
	case 4:
		d.putAddress(msg.Id, msg.Addr)
		r = reply{true, msg.Id, 0, "", ""}
	case 5:
		d.putphone(msg.Id, msg.Ph)
		r = reply{true, msg.Id, 0, "", ""}
	case 6:
		//getdetails
		s, ok := d.getMydetails(msg.Id)
		r = reply{ok, msg.Id, s.Roll, s.Address, s.phone}
	case 7:
		//deletekey
		d.deleteMydetails(msg.Id)
		r = reply{true, msg.Id, 0, "", ""}

	}
	return r

}

func main() {

	socketrep, _ := zmq.NewSocket(zmq.REP)
	socketrep.Bind("tcp://127.0.0.1:5001")
	var d database
	d.Map = make(map[int]*mydetails)

	for {

		msg, err := socketrep.RecvBytes(0)
		fmt.Println("recived")
		if err != nil {
			fmt.Println(err)
		}
		var req request
		json.Unmarshal(msg, &req)
		fmt.Println(req)
		resp := d.parse(req)
		fmt.Println(resp)
		message, _ := json.Marshal(resp)
		socketrep.SendBytes(message, 0)

	}

}
