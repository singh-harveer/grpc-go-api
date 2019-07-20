package main
import(
	"net"
	"log"
	"github.com/singh-harveer/grpc-go-api/greet/greetpb"
	"google.golang.org/grpc"
	"context"
)
type Server struct{}

func (*Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){

	log.Printf("Greet function is invoked with request Payload:  \n %v", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := "Hello "+firstName + " " + lastName
	res := &greetpb.GreetResponse{
		Result :result,
	}
	 return res, nil

}


func main(){ 
	log.Println("starting GRPC server on 50051 ...")
	// do port biding
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil{
		log.Fatalf("failed to listen %v", lis)
	} else{
		log.Println("GRPC server started on 50051 ...")
	}
	

	//create a new grpc server
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err!=nil{
		log.Fatalf("failed to serve %v", err)
	}
	

}