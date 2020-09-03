// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a registered Elastic IP address's name. For more information, see
// Resource Management
// (https://docs.aws.amazon.com/opsworks/latest/userguide/resources.html). Required
// Permissions: To use this action, an IAM user must have a Manage permissions
// level for the stack, or an attached policy that explicitly grants permissions.
// For more information on user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) UpdateElasticIp(ctx context.Context, params *UpdateElasticIpInput, optFns ...func(*Options)) (*UpdateElasticIpOutput, error) {
	stack := middleware.NewStack("UpdateElasticIp", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateElasticIpMiddlewares(stack)
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
	addOpUpdateElasticIpValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateElasticIp(options.Region), middleware.Before)

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
			OperationName: "UpdateElasticIp",
			Err:           err,
		}
	}
	out := result.(*UpdateElasticIpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateElasticIpInput struct {
	// The new name.
	Name *string
	// The IP address for which you want to update the name.
	ElasticIp *string
}

type UpdateElasticIpOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateElasticIpMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateElasticIp{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateElasticIp{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateElasticIp(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "UpdateElasticIp",
	}
}