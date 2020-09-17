// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates an Identity and Access Management (IAM) role from an Neptune DB
// cluster.
func (c *Client) AddRoleToDBCluster(ctx context.Context, params *AddRoleToDBClusterInput, optFns ...func(*Options)) (*AddRoleToDBClusterOutput, error) {
	stack := middleware.NewStack("AddRoleToDBCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpAddRoleToDBClusterMiddlewares(stack)
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
	addOpAddRoleToDBClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddRoleToDBCluster(options.Region), middleware.Before)
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
			OperationName: "AddRoleToDBCluster",
			Err:           err,
		}
	}
	out := result.(*AddRoleToDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddRoleToDBClusterInput struct {
	// The name of the DB cluster to associate the IAM role with.
	DBClusterIdentifier *string
	// The Amazon Resource Name (ARN) of the IAM role to associate with the Neptune DB
	// cluster, for example arn:aws:iam::123456789012:role/NeptuneAccessRole.
	RoleArn *string
}

type AddRoleToDBClusterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpAddRoleToDBClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpAddRoleToDBCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpAddRoleToDBCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddRoleToDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "AddRoleToDBCluster",
	}
}
