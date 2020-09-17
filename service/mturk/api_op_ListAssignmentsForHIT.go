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

// The ListAssignmentsForHIT operation retrieves completed assignments for a HIT.
// You can use this operation to retrieve the results for a HIT. You can get
// assignments for a HIT at any time, even if the HIT is not yet Reviewable. If a
// HIT requested multiple assignments, and has received some results but has not
// yet become Reviewable, you can still retrieve the partial results with this
// operation. Use the AssignmentStatus parameter to control which set of
// assignments for a HIT are returned. The ListAssignmentsForHIT operation can
// return submitted assignments awaiting approval, or it can return assignments
// that have already been approved or rejected. You can set
// AssignmentStatus=Approved,Rejected to get assignments that have already been
// approved and rejected together in one result set. Only the Requester who created
// the HIT can retrieve the assignments for that HIT. Results are sorted and
// divided into numbered pages and the operation returns a single page of results.
// You can use the parameters of the operation to control sorting and pagination.
func (c *Client) ListAssignmentsForHIT(ctx context.Context, params *ListAssignmentsForHITInput, optFns ...func(*Options)) (*ListAssignmentsForHITOutput, error) {
	stack := middleware.NewStack("ListAssignmentsForHIT", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListAssignmentsForHITMiddlewares(stack)
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
	addOpListAssignmentsForHITValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAssignmentsForHIT(options.Region), middleware.Before)
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
			OperationName: "ListAssignmentsForHIT",
			Err:           err,
		}
	}
	out := result.(*ListAssignmentsForHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssignmentsForHITInput struct {
	MaxResults *int32
	// The status of the assignments to return: Submitted | Approved | Rejected
	AssignmentStatuses []types.AssignmentStatus
	// The ID of the HIT.
	HITId *string
	// Pagination token
	NextToken *string
}

type ListAssignmentsForHITOutput struct {
	// The collection of Assignment data structures returned by this call.
	Assignments []*types.Assignment
	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string
	// The number of assignments on the page in the filtered results list, equivalent
	// to the number of assignments returned by this call.
	NumResults *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListAssignmentsForHITMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListAssignmentsForHIT{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAssignmentsForHIT{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAssignmentsForHIT(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListAssignmentsForHIT",
	}
}
