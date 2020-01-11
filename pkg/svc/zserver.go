package svc

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/errors"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"strconv"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"
	"github.com/seizadi/cmdb/pkg/pb"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// ~~~~~~~~~~~~~~~~~~~~~~~~~ A BRIEF DEVELOPMENT GUIDE ~~~~~~~~~~~~~~~~~~~~~~~~~
//
// TODO: Extend the Cmdb service by defining new RPCs and
// and message types in the pb/service.proto file. These RPCs and messages
// compose the API for your service. After modifying the proto schema in
// pb/service.proto, call "make protobuf" to regenerate the protobuf files.
//
// TODO: Create an implementation of the Cmdb server
// interface. This interface is generated by the protobuf compiler and exists
// inside the pb/service.pb.go file. The "server" struct already provides an
// implementation of Cmdb server interface service, but only
// for the GetVersion function. You will need to implement any new RPCs you
// add to your protobuf schema.
//
// TODO: Update the GetVersion function when newer versions of your service
// become available. Feel free to change GetVersion to better-suit how your
// versioning system, or get rid of it entirely. GetVersion helps make up a
// simple "starter" example that allows an end-to-end example. It is not
// required.
//
// TODO: Oh yeah, delete this guide when you no longer need it.
//
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~ FAREWELL AND GOOD LUCK ~~~~~~~~~~~~~~~~~~~~~~~~~~
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

const (
	// version is the current version of the service
	version = "0.0.1"
)

// Default implementation of the Cmdb server interface
type server struct{ db *gorm.DB }

// GetVersion returns the current version of the service
func (server) GetVersion(context.Context, *empty.Empty) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: version}, nil
}

// NewBasicServer returns an instance of the default server interface
func NewBasicServer(database *gorm.DB) (pb.CmdbServer, error) {
	return &server{db: database}, nil
}

// List wraps default RegionsDefaultServer.List implementation by adding
// application specific page token implementation.
// Actually the service supports "composite" pagination in a specific way:
// - limit and offset are still supported but without page token
// - if an user requests page token and provides limit then limit value will be
//	 used as a step for all further requests
//		page_toke = null & limit = 2 -> page_token=base64(offset=3:limit=2)
// - if an user requests page token and provides offset then only first time
//	 the provided offset is applied
//		page_token = null & offset = 2 & limit = 2 -> page_token=base64(offset=2+2:limit=2)
func (s *regionsServer) List(ctx context.Context, in *pb.ListRegionRequest) (*pb.ListRegionsResponse, error) {
	page := &query.Pagination{}
	err := gateway.GetCollectionOp(in, page)
	if err != nil {
		return nil, err
	}

	ptoken := page.GetPageToken()
	// do not handle page token
	if ptoken == "" {
		return s.RegionsDefaultServer.List(ctx, in)
	}

	// decode provided token (null means a client is requesting new token)
	// update context with new pagination request
	if ptoken != "null" {
		page.Offset, page.Limit, err = DecodePageToken(ptoken)
		if err != nil {
			return nil, err
		}
		if err := gateway.SetCollectionOps(in, page); err != nil {
			grpclog.Errorf("collection operator interceptor: failed to set pagination operator - %s", err)
			return nil, err
		}
	}

	// forward request to default implementation
	resp, err := s.RegionsDefaultServer.List(ctx, in)
	if err != nil {
		return nil, err
	}

	// prepare and set response page info
	var pinfo query.PageInfo
	if length := len(resp.Results); length == 0 {
		pinfo.SetLastToken()
	} else {
		pinfo.PageToken = EncodePageToken(page.GetOffset()+int32(length), page.DefaultLimit())
	}
	if err := gateway.SetPageInfo(ctx, &pinfo); err != nil {
		return nil, err
	}

	return resp, nil
}

// DecodePageToken decodes page token from the user's request.
// Return error if provided token is malformed or contains ivalid values,
// otherwise return offset, limit.
func DecodePageToken(ptoken string) (offset, limit int32, err error) {
	errC := errors.InitContainer()
	data, err := base64.StdEncoding.DecodeString(ptoken)
	if err != nil {
		return 0, 0, errC.New(codes.InvalidArgument, "Invalid page token %q.", err)
	}
	vals := strings.SplitN(string(data), ":", 2)
	if len(vals) != 2 {
		return 0, 0, errC.New(codes.InvalidArgument, "Malformed page token.")
	}

	o, err := strconv.Atoi(vals[0])
	if err != nil {
		errC.Set("page_token", codes.InvalidArgument, "invalid offset value %q.", vals[0])
		errC.WithField("offset", "Invalid offset value. The valid value is an unsigned integer.")
	}

	l, err := strconv.Atoi(vals[1])
	if err != nil {
		errC.Set("page_token", codes.InvalidArgument, "invalid limit value %q.", vals[1])
		errC.WithField("limit", "Invalid limit value. The valid value is an unsigned integer.")
	}

	limit = int32(l)
	offset = int32(o)

	if err := errC.IfSet(codes.InvalidArgument, "Page token validation failed."); err != nil {
		return 0, 0, errC
	}

	return limit, offset, nil
}

// EncodePageToken encodes offset and limit to a string in application specific
// format (offset:limit) in base64 encoding.
func EncodePageToken(offset, limit int32) string {
	data := fmt.Sprintf("%d:%d", offset, limit)
	return base64.StdEncoding.EncodeToString([]byte(data))
}
