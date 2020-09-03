// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used to cancel an open change request. Must be sent before the status of the
// request changes to APPLYING, the final stage of completing your change request.
// You can describe a change during the 60-day request history retention period for
// API calls.
func (c *Client) CancelChangeSet(ctx context.Context, params *CancelChangeSetInput, optFns ...func(*Options)) (*CancelChangeSetOutput, error) {
	stack := middleware.NewStack("CancelChangeSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCancelChangeSetMiddlewares(stack)
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
	addOpCancelChangeSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelChangeSet(options.Region), middleware.Before)

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
			OperationName: "CancelChangeSet",
			Err:           err,
		}
	}
	out := result.(*CancelChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelChangeSetInput struct {
	// Required. The unique identifier of the StartChangeSet request that you want to
	// cancel.
	ChangeSetId *string
	// Required. The catalog related to the request. Fixed value: AWSMarketplace.
	Catalog *string
}

type CancelChangeSetOutput struct {
	// The unique identifier for the change set referenced in this request.
	ChangeSetId *string
	// The ARN associated with the change set referenced in this request.
	ChangeSetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCancelChangeSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCancelChangeSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelChangeSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelChangeSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "CancelChangeSet",
	}
}
