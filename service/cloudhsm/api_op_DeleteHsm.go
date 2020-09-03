// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see AWS
// CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/), the AWS
// CloudHSM Classic User Guide
// (https://docs.aws.amazon.com/cloudhsm/classic/userguide/), and the AWS CloudHSM
// Classic API Reference
// (https://docs.aws.amazon.com/cloudhsm/classic/APIReference/). For information
// about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide
// (https://docs.aws.amazon.com/cloudhsm/latest/userguide/), and the AWS CloudHSM
// API Reference (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
// Deletes an HSM. After completion, this operation cannot be undone and your key
// material cannot be recovered.
func (c *Client) DeleteHsm(ctx context.Context, params *DeleteHsmInput, optFns ...func(*Options)) (*DeleteHsmOutput, error) {
	stack := middleware.NewStack("DeleteHsm", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteHsmMiddlewares(stack)
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
	addOpDeleteHsmValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHsm(options.Region), middleware.Before)

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
			OperationName: "DeleteHsm",
			Err:           err,
		}
	}
	out := result.(*DeleteHsmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DeleteHsm () operation.
type DeleteHsmInput struct {
	// The ARN of the HSM to delete.
	HsmArn *string
}

// Contains the output of the DeleteHsm () operation.
type DeleteHsmOutput struct {
	// The status of the operation.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteHsmMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteHsm{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteHsm{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteHsm(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "DeleteHsm",
	}
}
