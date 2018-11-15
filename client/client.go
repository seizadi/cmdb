package client

import (
	"google.golang.org/grpc"
	context "golang.org/x/net/context"
	"github.com/seizadi/cmdb/pkg/pb"
	"google.golang.org/grpc/metadata"
	infoblox_api "github.com/infobloxopen/atlas-app-toolkit/query"
	atlas_rpc "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"

)

func GetConn(host string) (*grpc.ClientConn, error){
	// First we create the connection:
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	
	return conn, nil
}

func GetRegions(conn *grpc.ClientConn) (*pb.ListRegionsResponse, error) {
	
	// We can now create stubs that wrap conn:
	stub := pb.NewRegionsClient(conn)
	
	// Now we can use the stub to make RPCs
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjF9.GsXyFDDARjXe1t9DPo2LIBKHEal3O7t3vLI3edA7dGU"))
	reqList := &pb.ListRegionRequest{}
	respList, err := stub.List(ctx, reqList)
	if err != nil {
		return nil, err
	}
	
	return respList, nil
}

func GetRegion(conn *grpc.ClientConn, id string) (*pb.ReadRegionResponse, error) {
	
	stub := pb.NewRegionsClient(conn)
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjF9.GsXyFDDARjXe1t9DPo2LIBKHEal3O7t3vLI3edA7dGU"))

	resourceId := atlas_rpc.Identifier{ResourceId: id}
	fields := infoblox_api.FieldSelection{}
	
	// We'll make another request and also print the response metadata
	req := &pb.ReadRegionRequest{&resourceId, &fields}
	var respHdrs, respTrlrs metadata.MD
	resp, err := stub.Read(ctx, req,
		grpc.Header(&respHdrs), grpc.Trailer(&respTrlrs))
	if err != nil {
		return nil, err
	}
	
	return resp, nil
}