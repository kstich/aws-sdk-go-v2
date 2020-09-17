// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets all labels or a specific label if name is provided. This is a paginated
// API. If you provide a null maxResults, this action retrieves a maximum of 50
// records per page. If you provide a maxResults, the value must be between 10 and
// 50. To get the next page results, provide the pagination token from the
// GetGetLabelsResponse as part of your request. A null pagination token fetches
// the records from the beginning.
func (c *Client) GetLabels(ctx context.Context, params *GetLabelsInput, optFns ...func(*Options)) (*GetLabelsOutput, error) {
	stack := middleware.NewStack("GetLabels", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetLabelsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLabels(options.Region), middleware.Before)
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
			OperationName: "GetLabels",
			Err:           err,
		}
	}
	out := result.(*GetLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLabelsInput struct {
	// The name of the label or labels to get.
	Name *string
	// The maximum number of objects to return for the request.
	MaxResults *int32
	// The next token for the subsequent request.
	NextToken *string
}

type GetLabelsOutput struct {
	// The next page token.
	NextToken *string
	// An array of labels.
	Labels []*types.Label

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetLabelsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetLabels{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLabels{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetLabels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetLabels",
	}
}
