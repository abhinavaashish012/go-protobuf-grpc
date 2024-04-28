package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"pcbook/pb"
)

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	laptopStore LaptopStore
}

// NewLaptopServer returns a new LaptopServer
func NewLaptopServer(laptopStore LaptopStore) *LaptopServer {
	return &LaptopServer{
		laptopStore: laptopStore,
	}
}

// LaptopServiceClient is the client API for LaptopService service.
type LaptopServiceClient interface {
	CreateLaptop(ctx context.Context, in *pb.CreateLaptopRequest, opts ...grpc.CallOption) (*pb.CreateLaptopResponse,
		error)
}

type laptopServiceClient struct {
	cc *grpc.ClientConn
}

func NewLaptopServiceClient(cc *grpc.ClientConn) LaptopServiceClient {
	return &laptopServiceClient{cc}
}

func (c *laptopServiceClient) CreateLaptop(ctx context.Context, in *pb.CreateLaptopRequest,
	opts ...grpc.CallOption) (*pb.CreateLaptopResponse, error) {
	out := new(pb.CreateLaptopResponse)
	err := c.cc.Invoke(ctx, "/pcbook.LaptopService/CreateLaptop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type LaptopServiceServer interface {
	CreateLaptop(context.Context, *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error)
}

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	fmt.Println("REQUEST :: ", req)
	laptop := req.GetLaptop()
	//fmt.Println("laptop ==== ", laptop.String())
	//	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// check if it's a valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}

	// some heavy processing
	// time.Sleep(6 * time.Second)

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	// save the laptop to store
	err := server.laptopStore.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
