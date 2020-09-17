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

// Creates a block storage disk that can be attached to an Amazon Lightsail
// instance in the same Availability Zone (e.g., us-east-2a). The create disk
// operation supports tag-based access control via request tags. For more
// information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateDisk(ctx context.Context, params *CreateDiskInput, optFns ...func(*Options)) (*CreateDiskOutput, error) {
	stack := middleware.NewStack("CreateDisk", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDiskMiddlewares(stack)
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
	addOpCreateDiskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDisk(options.Region), middleware.Before)
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
			OperationName: "CreateDisk",
			Err:           err,
		}
	}
	out := result.(*CreateDiskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDiskInput struct {
	// The size of the disk in GB (e.g., 32).
	SizeInGb *int32
	// An array of objects that represent the add-ons to enable for the new disk.
	AddOns []*types.AddOnRequest
	// The unique Lightsail disk name (e.g., my-disk).
	DiskName *string
	// The Availability Zone where you want to create the disk (e.g., us-east-2a). Use
	// the same Availability Zone as the Lightsail instance to which you want to attach
	// the disk. Use the get regions operation to list the Availability Zones where
	// Lightsail is currently available.
	AvailabilityZone *string
	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []*types.Tag
}

type CreateDiskOutput struct {
	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDiskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDisk{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDisk{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDisk(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateDisk",
	}
}
