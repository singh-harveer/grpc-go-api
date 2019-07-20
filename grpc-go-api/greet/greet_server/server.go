package main
import(
	"net"
	"log"
	"github.com/singh-harveer/grpc-go-api/greet/greetpb"
	"google.golang.org/grpc"
)
type Server struct{}

func main(){ 
	log.Println("Running GRPC server ...")
	// do port biding
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil{
		log.Fatalf("failed to listen %v", lis)
	}

	//create a new grpc server
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err!=nil{
		log.Fatalf("failed to serve %v", err)
	}
	

}