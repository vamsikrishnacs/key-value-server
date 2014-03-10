package main


import (
	"fmt"
	 zmq "github.com/pebbe/zmq4"
	 "encoding/json"
)

type client struct{
socket *zmq.Socket
}

type request struct{
Op int
Id int
Roll int 
Addr string
Ph string
}

type reply struct{
Success bool
Id int
Roll int
Addr string
Ph string
}

func (c *client) send(msg request) (reply){

            fmt.Println(msg)
        message,_:=json.Marshal(msg)
	c.socket.SendBytes(message, 0)
	fmt.Println("Sending")
        resp,err:=c.socket.RecvBytes(0)
	 var rep reply
	 json.Unmarshal(resp,&rep)
	if err != nil {
		fmt.Println(err)
        }
          return rep

}
 ///specially written to write test code,..can be written in main as well
func (c *client) prepare(op int,id int,r int,a string,p string)(reply){
var m request

	    m=request{op,id,r,a,p}
	    rep:=c.send(m)
return rep
}






func main() {
        
        //connect
        var cl client
        cl.socket, _ = zmq.NewSocket(zmq.REQ)
        err:=cl.socket.Connect("tcp://127.0.0.1:5001")
        fmt.Println(err)
	fmt.Println("Initialisng the client........")
for{	
fmt.Println("User options 1.register 2.filldetails(roll,address,phone) 3.change_rollno 4.change_adress 5.change_ph 6.getdetails 7.deleteme")
	var op,id,r int
        var a,p string   
                fmt.Scanf("%d",&op)
	 
        switch(op){	
	       case 1:
	       //createuse
		fmt.Scanf("%d",&id)
	        rep:=cl.prepare(op,id,0,"","")
                if(rep.Success==true){
                 fmt.Println("User already exist")
                 }
	    
	
	        case 2:
	  //putdetails  
		 fmt.Scanf("%d %d %s %s",&id,&r,&a,&p)
                cl.prepare(op,id,r,a,p)
		
		case 3:	
			fmt.Scanf("%d %d",&id,&r)
                 
		        cl.prepare(op,id,r,"","")
		case 4:    
			fmt.Scanf("%d %s",&id,&a)
		        cl.prepare(op,id,0,a,"")
		case 5:
			fmt.Scanf("%d %s",&id,&p)
		        cl.prepare(op,id,0,"",p)
		case 6:
		   //getdetails
			fmt.Scanf("%d",&id)
		        rep:=cl.prepare(op,id,0,"","")
                      if(rep.Success==true){
			fmt.Println("rollno")
			fmt.Println(rep.Roll)
			fmt.Println("Address"+rep.Addr)
			fmt.Println("phone"+rep.Ph)
			}else{
			fmt.Println("you havenot registered")

			}
		case 7:
		    //deletekey
	                fmt.Scanf("%d",&id)
		        cl.prepare(op,id,0,"","")	    	
	}
          
      } 
}
