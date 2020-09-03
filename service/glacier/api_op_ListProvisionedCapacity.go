// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation lists the provisioned capacity units for the specified AWS
// account.
func (c *Client) ListProvisionedCapacity(ctx context.Context, params *ListProvisionedCapacityInput, optFns ...func(*Options)) (*ListProvisionedCapacityOutput, error) {
	stack := middleware.NewStack("ListProvisionedCapacity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListProvisionedCapacityMiddlewares(stack)
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
	addOpListProvisionedCapacityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProvisionedCapacity(options.Region), middleware.Before)

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
			OperationName: "ListProvisionedCapacity",
			Err:           err,
		}
	}
	out := result.(*ListProvisionedCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProvisionedCapacityInput struct {
	// The AWS account ID of the account that owns the vault. You can either specify an
	// AWS account ID or optionally a single '-' (hyphen), in which case Amazon S3
	// Glacier uses the AWS account ID associated with the credentials used to sign the
	// request. If you use an account ID, don't include any hyphens ('-') in the ID.
	AccountId *string
}

type ListProvisionedCapacityOutput struct {
	// The response body contains the following JSON fields.
	ProvisionedCapacityList []*types.ProvisionedCapacityDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListProvisionedCapacityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListProvisionedCapacity{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListProvisionedCapacity{}, middleware.After)
}

func newServiceMetadataMiddleware_opListProvisionedCapacity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "ListProvisionedCapacity",
	}
}
