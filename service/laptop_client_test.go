package service

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"github.jiaget.com/pc-book/pb"
	"github.jiaget.com/pc-book/sample"
	"github.jiaget.com/pc-book/serializer"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopStore, serverAddress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	// check the laptop is stored in the store
	laptop2, err := laptopStore.Store.Find(laptop.Id)
	requireEqualLaptop(t, laptop, laptop2)
}

func startTestLaptopServer(t *testing.T) (laptopServer *LaptopServer, address string) {
	laptopServer = NewLaptopServer(NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	// server will remain active ,so we need make it run in another goroutine
	go grpcServer.Serve(listener)

	address = listener.Addr().String()
	return
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func requireEqualLaptop(t *testing.T, laptop1, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
