// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/braket/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches for tasks that match the specified filter values.
func (c *Client) SearchQuantumTasks(ctx context.Context, params *SearchQuantumTasksInput, optFns ...func(*Options)) (*SearchQuantumTasksOutput, error) {
	stack := middleware.NewStack("SearchQuantumTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSearchQuantumTasksMiddlewares(stack)
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
	addOpSearchQuantumTasksValidationMiddleware(stack)
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
			OperationName: "SearchQuantumTasks",
			Err:           err,
		}
	}
	out := result.(*SearchQuantumTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchQuantumTasksInput struct {
	// A token used for pagination of results returned in the response. Use the token
	// returned from the previous request continue results where the previous request
	// ended.
	NextToken *string
	// Maximum number of results to return in the response.
	MaxResults *int32
	// Array of SearchQuantumTasksFilter objects.
	Filters []*types.SearchQuantumTasksFilter
}

type SearchQuantumTasksOutput struct {
	// An array of QuantumTaskSummary objects for tasks that match the specified
	// filters.
	QuantumTasks []*types.QuantumTaskSummary
	// A token used for pagination of results, or null if there are no additional
	// results. Use the token value in a subsequent request to continue results where
	// the previous request ended.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSearchQuantumTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSearchQuantumTasks{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchQuantumTasks{}, middleware.After)
}
