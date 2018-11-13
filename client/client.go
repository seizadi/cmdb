package main

import (
	"fmt"
	"google.golang.org/grpc"
	context "golang.org/x/net/context"
	"github.com/seizadi/cmdb/pkg/pb"
	"google.golang.org/grpc/metadata"
	"os"
	infoblox_api "github.com/infobloxopen/atlas-app-toolkit/query"
	atlas_rpc "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"

)

func main() {
	// First we create the connection:
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	
	// We can now create stubs that wrap conn:
	stub := pb.NewRegionsClient(conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjF9.GsXyFDDARjXe1t9DPo2LIBKHEal3O7t3vLI3edA7dGU"))
	reqList := &pb.ListRegionRequest{}
	respList, err := stub.List(ctx, reqList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "RPC failed: %v\n", err)
	} else {
		fmt.Println(respList)
	}

	id := atlas_rpc.Identifier{ResourceId: "1"}
	fields := infoblox_api.FieldSelection{}
	
	// We'll make another request and also print the response metadata
	req := &pb.ReadRegionRequest{&id, &fields}
	var respHdrs, respTrlrs metadata.MD
	resp, err := stub.Read(ctx, req,
		grpc.Header(&respHdrs), grpc.Trailer(&respTrlrs))
	if err != nil {
		fmt.Fprintf(os.Stderr, "RPC failed: %v\n", err)
	} else {
		fmt.Println(resp)
	}
	fmt.Printf("Server sent headers: %v\n", respHdrs)
	fmt.Printf("Server sent trailers: %v\n", respTrlrs)
}