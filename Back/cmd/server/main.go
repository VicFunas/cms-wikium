package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	// Import the generated Go code
	"github.com/VicFunas/cms-wikium/internal/handler"
	"github.com/VicFunas/cms-wikium/internal/repository"
	"github.com/VicFunas/cms-wikium/internal/service"
	pb "github.com/VicFunas/cms-wikium/proto"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
)

// server is used to implement greeter.GreeterServer.
type server struct {
	// You must embed this to have forward compatible implementations.
	pb.UnimplementedGreeterServer
	mongoClient *mongo.Client
}

// SayHello implements greeter.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received SayHello request for: %v", in.GetName())

	// Check MongoDB connection
	err := s.mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping MongoDB: %v", err)
		return &pb.HelloReply{Message: "Hello " + in.GetName() + " (but failed to connect to DB)"}, nil
	}

	log.Println("Successfully pinged MongoDB!")
	return &pb.HelloReply{Message: "Hello " + in.GetName() + " from your Go server! DB connection is healthy."}, nil
}

func main() {
	// --- Connect to MongoDB ---
	// IMPORTANT: Replace this with your own MongoDB Atlas connection string!
	// You can get this from your Atlas dashboard.
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		// Set a default for local testing if the env var isn't present
		mongoURI = "mongodb://localhost:27017"
		log.Println("MONGO_URI environment variable not set, using default.")
	}

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB!")

	// --- Initialize layers ---
	modRepo := repository.NewModRepository(client.Database("cms"))
	modService := service.NewModService(modRepo)
	modHandler := handler.NewModHandler(modService)

	// --- Set up gRPC Server ---
	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// Register our service with the gRPC server.
	pb.RegisterGreeterServer(s, &server{mongoClient: client})
	pb.RegisterModServiceServer(s, modHandler)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
