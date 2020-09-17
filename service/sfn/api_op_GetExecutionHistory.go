// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the history of the specified execution as a list of events. By default,
// the results are returned in ascending order of the timeStamp of the events. Use
// the reverseOrder parameter to get the latest events first. If nextToken is
// returned, there are more results available. The value of nextToken is a unique
// pagination token for each page. Make the call again using the returned token to
// retrieve the next page. Keep all other arguments unchanged. Each pagination
// token expires after 24 hours. Using an expired pagination token will return an
// HTTP 400 InvalidToken error. This API action is not supported by EXPRESS state
// machines.
func (c *Client) GetExecutionHistory(ctx context.Context, params *GetExecutionHistoryInput, optFns ...func(*Options)) (*GetExecutionHistoryOutput, error) {
	stack := middleware.NewStack("GetExecutionHistory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpGetExecutionHistoryMiddlewares(stack)
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
	addOpGetExecutionHistoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetExecutionHistory(options.Region), middleware.Before)
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
			OperationName: "GetExecutionHistory",
			Err:           err,
		}
	}
	out := result.(*GetExecutionHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExecutionHistoryInput struct {
	// The maximum number of results that are returned per call. You can use nextToken
	// to obtain further pages of results. The default is 100 and the maximum allowed
	// page size is 1000. A value of 0 uses the default. This is only an upper limit.
	// The actual number of results returned per call might be fewer than the specified
	// maximum.
	MaxResults *int32
	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string
	// Lists events in descending order of their timeStamp.
	ReverseOrder *bool
	// The Amazon Resource Name (ARN) of the execution.
	ExecutionArn *string
}

type GetExecutionHistoryOutput struct {
	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string
	// The list of events that occurred in the execution.
	Events []*types.HistoryEvent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpGetExecutionHistoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpGetExecutionHistory{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetExecutionHistory{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetExecutionHistory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "GetExecutionHistory",
	}
}
