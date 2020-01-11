package main

import (
	"context"
	"google.golang.org/grpc/codes"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/infobloxopen/atlas-app-toolkit/errors"
	"github.com/infobloxopen/atlas-app-toolkit/errors/mappers/validationerrors"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var ErrorMappings = []errors.MapFunc{

	// Default Validation Mapping
	validationerrors.DefaultMapping(),

	errors.NewMapping(
		errors.CondEq("NOT_EXISTS"),
		errors.NewContainer(
			codes.InvalidArgument, "The resource does not exist.",
		).WithField(
			"path/id", "the specified object was not found."),
	),

	//pqerrors.NewUniqueMapping("emails_address_key", "Contacts", "Primary Email Address"),

	errors.NewMapping(
		gorm.ErrRecordNotFound,
		errors.NewContainer(codes.NotFound, "record not found"),
	),

	errors.NewMapping(
		// Here CondAnd without condition functions serves as 'default'.
		errors.CondAnd(),
		errors.MapFunc(func(ctx context.Context, err error) (error, bool) {
			ctxlogrus.AddFields(ctx, logrus.Fields{"internal-error": err})
			reqID, exist := requestid.FromContext(ctx)
			if exist {
				return errors.NewContainer(codes.Internal, "Internal error occured. For more details see log for request %s", reqID), true
			}
			return errors.NewContainer(codes.Internal, "Internal error occured."), true
		}),
	),
}
