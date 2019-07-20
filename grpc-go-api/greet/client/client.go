package main

import(
	"log"
	"google.golang.org/grpc"
	"github.com/singh-harveer/grpc-go-api/greet/greetpb"
	"context"
)

func main(){
	log.Println("hello i am a client !!")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil{
		log.Fatalf("could not connect to %v", err)
	}
	defer cc.Close()//close this connection once function execution is done
	c := greetpb.NewGreetServiceClient(cc)
	// log.Printf("Created Client: %f", c)
	doUnary(c)// 

	

}


func doUnary(c greetpb.GreetServiceClient){

	log.Println("Starting to do a Unary RPC.")
	req := &greetpb.GreetRequest{
		Greeting :&greetpb.Greeting {
			FirstName :"Harry",
			LastName :"Singh",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil{
		log.Fatalf("Error while calling Greet RPC %v", err)
	}
	log.Printf("Response from Greet RPC:  %v", res.Result)


}