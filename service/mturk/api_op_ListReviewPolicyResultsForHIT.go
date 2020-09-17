// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ListReviewPolicyResultsForHIT operation retrieves the computed results and
// the actions taken in the course of executing your Review Policies for a given
// HIT. For information about how to specify Review Policies when you call
// CreateHIT, see Review Policies. The ListReviewPolicyResultsForHIT operation can
// return results for both Assignment-level and HIT-level review results.
func (c *Client) ListReviewPolicyResultsForHIT(ctx context.Context, params *ListReviewPolicyResultsForHITInput, optFns ...func(*Options)) (*ListReviewPolicyResultsForHITOutput, error) {
	stack := middleware.NewStack("ListReviewPolicyResultsForHIT", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListReviewPolicyResultsForHITMiddlewares(stack)
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
	addOpListReviewPolicyResultsForHITValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListReviewPolicyResultsForHIT(options.Region), middleware.Before)
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
			OperationName: "ListReviewPolicyResultsForHIT",
			Err:           err,
		}
	}
	out := result.(*ListReviewPolicyResultsForHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReviewPolicyResultsForHITInput struct {
	// Limit the number of results returned.
	MaxResults *int32
	// Pagination token
	NextToken *string
	// Specify if the operation should retrieve a list of the results computed by the
	// Review Policies.
	RetrieveResults *bool
	// The Policy Level(s) to retrieve review results for - HIT or Assignment. If
	// omitted, the default behavior is to retrieve all data for both policy levels.
	// For a list of all the described policies, see Review Policies.
	PolicyLevels []types.ReviewPolicyLevel
	// The unique identifier of the HIT to retrieve review results for.
	HITId *string
	// Specify if the operation should retrieve a list of the actions taken executing
	// the Review Policies and their outcomes.
	RetrieveActions *bool
}

type ListReviewPolicyResultsForHITOutput struct {
	// Contains both ReviewResult and ReviewAction elements for a particular HIT.
	HITReviewReport *types.ReviewReport
	// The name of the HIT-level Review Policy. This contains only the PolicyName
	// element.
	HITReviewPolicy *types.ReviewPolicy
	// The name of the Assignment-level Review Policy. This contains only the
	// PolicyName element.
	AssignmentReviewPolicy *types.ReviewPolicy
	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string
	// The HITId of the HIT for which results have been returned.
	HITId *string
	// Contains both ReviewResult and ReviewAction elements for an Assignment.
	AssignmentReviewReport *types.ReviewReport

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListReviewPolicyResultsForHITMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListReviewPolicyResultsForHIT{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListReviewPolicyResultsForHIT{}, middleware.After)
}

func newServiceMetadataMiddleware_opListReviewPolicyResultsForHIT(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListReviewPolicyResultsForHIT",
	}
}
