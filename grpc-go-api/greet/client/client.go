package main

import(
	"log"
	"google.golang.org/grpc"
	"github.com/singh-harveer/grpc-go-api/greet/greetpb"
	"context"
	"io"
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
	// doUnary(c)// 
	doServerStreaming(c)
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

func doServerStreaming(c greetpb.GreetServiceClient){
	log.Println("strating to do server streaming RPC ..")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName:"Harry",
			LastName: "Singh",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil{
		log.Fatalf("Error while calling GreetManyTimes RPC: \n %v",err.Error())
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF{
			// we are reached at end of stream
			break
		}
		if err != nil{
			log.Fatalf("Error while reading stream: %v", err)
		}
	
		log.Printf("Response from server: %v", msg.GetResult())
	}

}