package main;
import (
	"context"
	"log"
	"google.golang.org/grpc"
	"github.com/singh-harveer/grpc-go-api/calculator/calculatorpb"

)

func main(){
	log.Println("Hello, i am a client !")
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Could not connect to %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorClient(cc)
	//log.Printf("")
	doUnary(c)

}




func doUnary(c calculatorpb.CalculatorClient){

	req := &calculatorpb.SumRequest{
		FirstNumber :23,
		SecondNumber:23,
	}
res, err := c.Sum(context.Background(), req)

if err != nil{
	log.Fatalf("Error while calling Sum RPC %v", err)
}
log.Printf("Respone from Sum RPC:  %v", res.Result)

}