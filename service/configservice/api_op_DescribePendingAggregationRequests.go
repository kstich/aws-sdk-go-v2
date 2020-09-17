// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all pending aggregation requests.
func (c *Client) DescribePendingAggregationRequests(ctx context.Context, params *DescribePendingAggregationRequestsInput, optFns ...func(*Options)) (*DescribePendingAggregationRequestsOutput, error) {
	stack := middleware.NewStack("DescribePendingAggregationRequests", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribePendingAggregationRequestsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePendingAggregationRequests(options.Region), middleware.Before)
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
			OperationName: "DescribePendingAggregationRequests",
			Err:           err,
		}
	}
	out := result.(*DescribePendingAggregationRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePendingAggregationRequestsInput struct {
	// The maximum number of evaluation results returned on each page. The default is
	// maximum. If you specify 0, AWS Config uses the default.
	Limit *int32
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

type DescribePendingAggregationRequestsOutput struct {
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// Returns a PendingAggregationRequests object.
	PendingAggregationRequests []*types.PendingAggregationRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribePendingAggregationRequestsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePendingAggregationRequests{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePendingAggregationRequests{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePendingAggregationRequests(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribePendingAggregationRequests",
	}
}
