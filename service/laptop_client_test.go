package service

import (
	"context"
	"io"
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
	store := NewInMemoryLaptopStore()
	laptopStore, serverAddress := startTestLaptopServer(t, store)
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

func TestClientSearchLaptop(t *testing.T) {
	t.Parallel()

	filter := &pb.Filter{
		MaxPriceRmb: 2000,
		MinCpuCores: 4,
		MinCpuGhz:   2.0,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}
	laptopStore := NewInMemoryLaptopStore()
	expectedIDs := make(map[string]bool)
	for i := 0; i < 6; i++ {
		laptop := sample.NewLaptop()
		switch i {
		case 0:
			laptop.PriceRmb = 2500
		case 1:
			laptop.Cpu.NumberCores = 2
		case 2:
			laptop.Cpu.MinGhz = 1.5
		case 3:
			laptop.Ram = &pb.Memory{Value: 4, Unit: pb.Memory_MEGABYTE}
		case 4:
			laptop.PriceRmb = 1000
			laptop.Cpu.NumberCores = 8
			laptop.Cpu.MinGhz = 2.5
			laptop.Cpu.MaxGhz = laptop.Cpu.MinGhz + 2.0
			laptop.Ram = &pb.Memory{Value: 16, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		case 5:
			laptop.PriceRmb = 2000
			laptop.Cpu.NumberCores = 8
			laptop.Cpu.MinGhz = 2.5
			laptop.Cpu.MaxGhz = laptop.Cpu.MinGhz + 2.0
			laptop.Ram = &pb.Memory{Value: 16, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		}
		err := laptopStore.Save(laptop)
		require.NoError(t, err)
	}
	_, serverAddress := startTestLaptopServer(t, laptopStore)
	laptopClient := newTestLaptopClient(t, serverAddress)

	req := &pb.SearchLaptopRequest{Filter: filter}
	stream, err := laptopClient.SearchLaptop(context.Background(), req)
	require.NoError(t, err)

	found := 0

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		require.NoError(t, err)
		require.Contains(t, expectedIDs, res.GetLaptop().Id)

		found += 1
	}
	require.Equal(t, len(expectedIDs), found)
}

func startTestLaptopServer(t *testing.T, store LaptopStore) (laptopServer *LaptopServer, address string) {
	laptopServer = NewLaptopServer(store)

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
