// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a database in Amazon Lightsail. The delete relational database operation
// supports tag-based access control via resource tags applied to the resource
// identified by relationalDatabaseName. For more information, see the Lightsail
// Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) DeleteRelationalDatabase(ctx context.Context, params *DeleteRelationalDatabaseInput, optFns ...func(*Options)) (*DeleteRelationalDatabaseOutput, error) {
	stack := middleware.NewStack("DeleteRelationalDatabase", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteRelationalDatabaseMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteRelationalDatabaseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRelationalDatabase(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeleteRelationalDatabase",
			Err:           err,
		}
	}
	out := result.(*DeleteRelationalDatabaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRelationalDatabaseInput struct {
	// The name of the database that you are deleting.
	RelationalDatabaseName *string
	// The name of the database snapshot created if skip final snapshot is false, which
	// is the default value for that parameter. Specifying this parameter and also
	// specifying the skip final snapshot parameter to true results in an error.
	// Constraints:
	//
	//     * Must contain from 2 to 255 alphanumeric characters, or
	// hyphens.
	//
	//     * The first and last character must be a letter or number.
	FinalRelationalDatabaseSnapshotName *string
	// Determines whether a final database snapshot is created before your database is
	// deleted. If true is specified, no database snapshot is created. If false is
	// specified, a database snapshot is created before your database is deleted. You
	// must specify the final relational database snapshot name parameter if the skip
	// final snapshot parameter is false. Default: false
	SkipFinalSnapshot *bool
}

type DeleteRelationalDatabaseOutput struct {
	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteRelationalDatabaseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRelationalDatabase{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRelationalDatabase{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteRelationalDatabase(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DeleteRelationalDatabase",
	}
}
