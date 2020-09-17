// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new subnet group. subnet groups must contain at least one subnet in at
// least two Availability Zones in the AWS Region.
func (c *Client) CreateDBSubnetGroup(ctx context.Context, params *CreateDBSubnetGroupInput, optFns ...func(*Options)) (*CreateDBSubnetGroupOutput, error) {
	stack := middleware.NewStack("CreateDBSubnetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateDBSubnetGroupMiddlewares(stack)
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
	addOpCreateDBSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBSubnetGroup(options.Region), middleware.Before)
	addResponseErrorWrapper(stack)

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
			OperationName: "CreateDBSubnetGroup",
			Err:           err,
		}
	}
	out := result.(*CreateDBSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to CreateDBSubnetGroup ().
type CreateDBSubnetGroupInput struct {
	// The description for the subnet group.
	DBSubnetGroupDescription *string
	// The tags to be assigned to the subnet group.
	Tags []*types.Tag
	// The name for the subnet group. This value is stored as a lowercase string.
	// Constraints: Must contain no more than 255 letters, numbers, periods,
	// underscores, spaces, or hyphens. Must not be default. Example: mySubnetgroup
	DBSubnetGroupName *string
	// The Amazon EC2 subnet IDs for the subnet group.
	SubnetIds []*string
}

type CreateDBSubnetGroupOutput struct {
	// Detailed information about a subnet group.
	DBSubnetGroup *types.DBSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateDBSubnetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBSubnetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBSubnetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDBSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBSubnetGroup",
	}
}
