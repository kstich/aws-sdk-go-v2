// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about device pools.
func (c *Client) ListDevicePools(ctx context.Context, params *ListDevicePoolsInput, optFns ...func(*Options)) (*ListDevicePoolsOutput, error) {
	stack := middleware.NewStack("ListDevicePools", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListDevicePoolsMiddlewares(stack)
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
	addOpListDevicePoolsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDevicePools(options.Region), middleware.Before)
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
			OperationName: "ListDevicePools",
			Err:           err,
		}
	}
	out := result.(*ListDevicePoolsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the result of a list device pools request.
type ListDevicePoolsInput struct {
	// The device pools' type. Allowed values include:
	//
	//     * CURATED: A device pool
	// that is created and managed by AWS Device Farm.
	//
	//     * PRIVATE: A device pool
	// that is created and managed by the device pool developer.
	Type types.DevicePoolType
	// The project ARN.
	Arn *string
	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

// Represents the result of a list device pools request.
type ListDevicePoolsOutput struct {
	// Information about the device pools.
	DevicePools []*types.DevicePool
	// If the number of items that are returned is significantly large, this is an
	// identifier that is also returned. It can be used in a subsequent call to this
	// operation to return the next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListDevicePoolsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListDevicePools{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDevicePools{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDevicePools(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListDevicePools",
	}
}
