package main
import(
	"log"
	"net"
	"github.com/singh-harveer/grpc-go-api/calculator/calculatorpb"
	"google.golang.org/grpc"
	"context"
)

type Server struct{

}

func (*Server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error){

	log.Printf("Invoked Sum RPC with \n %v", req)
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()

	res := &calculatorpb.SumResponse{
		Result: firstNumber + secondNumber,
	}
	return res, nil

}


func main(){

	log.Printf("starting GRPC server on 50052 ..")
	//do port binding
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err !=nil{
		log.Fatalf("Failed to listen %v", err.Error())

	}else{
		log.Printf("GRPC Server started on 50052 ...")
	}
	defer lis.Close()
	server := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(server, &Server{})
	if err := server.Serve(lis); err !=nil{
		log.Fatalf("Failed to serve %v", err)
	}
}