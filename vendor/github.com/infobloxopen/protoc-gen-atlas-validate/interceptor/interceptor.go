package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ValidationErrorMetaKey = "Atlas-Validation-Error"
)

// ValidationClientInterceptor extracts validation error from metadata
// and throws InvalidArgumentError if error is not empty.
func ValidationClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		if req == nil {
			return
		}

		if err := GetAtlasValidationError(ctx); err != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func GetAtlasValidationError(ctx context.Context) error {
	imd, _ := metadata.FromIncomingContext(ctx)
	omd, _ := metadata.FromOutgoingContext(ctx)

	md := metadata.Join(imd, omd)

	errors := md.Get(ValidationErrorMetaKey)

	if len(errors) > 0 && errors[0] != "" {
		return fmt.Errorf(errors[0])
	}

	return nil
}
