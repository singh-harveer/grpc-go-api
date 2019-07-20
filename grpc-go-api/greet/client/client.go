package main

import(
	"log"
	"google.golang.org/grpc"
	"github.com/singh-harveer/grpc-go-api/greet/greetpb"
)

func main(){
	log.Println("hello i am a client !!")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil{
		log.Fatalf("could not connect to %v", err)
	}
	defer cc.Close()//close this connection once function execution is done
	c := greetpb.NewGreetServiceClient(cc)
}