// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a recommendation filter. For more information, see Using Filters with
// Amazon Personalize.
func (c *Client) CreateFilter(ctx context.Context, params *CreateFilterInput, optFns ...func(*Options)) (*CreateFilterOutput, error) {
	stack := middleware.NewStack("CreateFilter", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateFilterMiddlewares(stack)
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
	addOpCreateFilterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFilter(options.Region), middleware.Before)
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
			OperationName: "CreateFilter",
			Err:           err,
		}
	}
	out := result.(*CreateFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFilterInput struct {
	// The filter expression that designates the interaction types that the filter will
	// filter out. A filter expression must follow the following format: EXCLUDE itemId
	// WHERE INTERACTIONS.event_type in ("EVENT_TYPE") Where "EVENT_TYPE" is the type
	// of event to filter out. To filter out all items with any interactions history,
	// set "*" as the EVENT_TYPE. For more information, see Using Filters with Amazon
	// Personalize.
	FilterExpression *string
	// The ARN of the dataset group that the filter will belong to.
	DatasetGroupArn *string
	// The name of the filter to create.
	Name *string
}

type CreateFilterOutput struct {
	// The ARN of the new filter.
	FilterArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateFilterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFilter{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFilter{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateFilter(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateFilter",
	}
}
