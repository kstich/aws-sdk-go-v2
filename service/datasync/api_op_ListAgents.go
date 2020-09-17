// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of agents owned by an AWS account in the AWS Region specified in
// the request. The returned list is ordered by agent Amazon Resource Name (ARN).
// By default, this operation returns a maximum of 100 agents. This operation
// supports pagination that enables you to optionally reduce the number of agents
// returned in a response. If you have more agents than are returned in a response
// (that is, the response returns only a truncated list of your agents), the
// response contains a marker that you can specify in your next request to fetch
// the next page of agents.
func (c *Client) ListAgents(ctx context.Context, params *ListAgentsInput, optFns ...func(*Options)) (*ListAgentsOutput, error) {
	stack := middleware.NewStack("ListAgents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListAgentsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAgents(options.Region), middleware.Before)
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
			OperationName: "ListAgents",
			Err:           err,
		}
	}
	out := result.(*ListAgentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// ListAgentsRequest
type ListAgentsInput struct {
	// The maximum number of agents to list.
	MaxResults *int32
	// An opaque string that indicates the position at which to begin the next list of
	// agents.
	NextToken *string
}

// ListAgentsResponse
type ListAgentsOutput struct {
	// A list of agents in your account.
	Agents []*types.AgentListEntry
	// An opaque string that indicates the position at which to begin returning the
	// next list of agents.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListAgentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListAgents{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAgents{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAgents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "ListAgents",
	}
}
