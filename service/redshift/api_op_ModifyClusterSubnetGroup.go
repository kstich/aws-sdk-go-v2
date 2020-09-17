// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies a cluster subnet group to include the specified list of VPC subnets.
// The operation replaces the existing list of subnets with the new list of
// subnets.
func (c *Client) ModifyClusterSubnetGroup(ctx context.Context, params *ModifyClusterSubnetGroupInput, optFns ...func(*Options)) (*ModifyClusterSubnetGroupOutput, error) {
	stack := middleware.NewStack("ModifyClusterSubnetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyClusterSubnetGroupMiddlewares(stack)
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
	addOpModifyClusterSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterSubnetGroup(options.Region), middleware.Before)
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
			OperationName: "ModifyClusterSubnetGroup",
			Err:           err,
		}
	}
	out := result.(*ModifyClusterSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyClusterSubnetGroupInput struct {
	// The name of the subnet group to be modified.
	ClusterSubnetGroupName *string
	// A text description of the subnet group to be modified.
	Description *string
	// An array of VPC subnet IDs. A maximum of 20 subnets can be modified in a single
	// request.
	SubnetIds []*string
}

type ModifyClusterSubnetGroupOutput struct {
	// Describes a subnet group.
	ClusterSubnetGroup *types.ClusterSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyClusterSubnetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterSubnetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterSubnetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyClusterSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyClusterSubnetGroup",
	}
}
