// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of task definition families that are registered to your account
// (which may include task definition families that no longer have any ACTIVE task
// definition revisions). You can filter out task definition families that do not
// contain any ACTIVE task definition revisions by setting the status parameter to
// ACTIVE. You can also filter the results with the familyPrefix parameter.
func (c *Client) ListTaskDefinitionFamilies(ctx context.Context, params *ListTaskDefinitionFamiliesInput, optFns ...func(*Options)) (*ListTaskDefinitionFamiliesOutput, error) {
	stack := middleware.NewStack("ListTaskDefinitionFamilies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListTaskDefinitionFamiliesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTaskDefinitionFamilies(options.Region), middleware.Before)
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
			OperationName: "ListTaskDefinitionFamilies",
			Err:           err,
		}
	}
	out := result.(*ListTaskDefinitionFamiliesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTaskDefinitionFamiliesInput struct {
	// The nextToken value returned from a ListTaskDefinitionFamilies request
	// indicating that more results are available to fulfill the request and further
	// calls will be needed. If maxResults was provided, it is possible the number of
	// results to be fewer than maxResults. This token should be treated as an opaque
	// identifier that is only used to retrieve the next items in a list and not for
	// other programmatic purposes.
	NextToken *string
	// The familyPrefix is a string that is used to filter the results of
	// ListTaskDefinitionFamilies. If you specify a familyPrefix, only task definition
	// family names that begin with the familyPrefix string are returned.
	FamilyPrefix *string
	// The maximum number of task definition family results returned by
	// ListTaskDefinitionFamilies in paginated output. When this parameter is used,
	// ListTaskDefinitions only returns maxResults results in a single page along with
	// a nextToken response element. The remaining results of the initial request can
	// be seen by sending another ListTaskDefinitionFamilies request with the returned
	// nextToken value. This value can be between 1 and 100. If this parameter is not
	// used, then ListTaskDefinitionFamilies returns up to 100 results and a nextToken
	// value if applicable.
	MaxResults *int32
	// The task definition family status with which to filter the
	// ListTaskDefinitionFamilies results. By default, both ACTIVE and INACTIVE task
	// definition families are listed. If this parameter is set to ACTIVE, only task
	// definition families that have an ACTIVE task definition revision are returned.
	// If this parameter is set to INACTIVE, only task definition families that do not
	// have any ACTIVE task definition revisions are returned. If you paginate the
	// resulting output, be sure to keep the status value constant in each subsequent
	// request.
	Status types.TaskDefinitionFamilyStatus
}

type ListTaskDefinitionFamiliesOutput struct {
	// The nextToken value to include in a future ListTaskDefinitionFamilies request.
	// When the results of a ListTaskDefinitionFamilies request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string
	// The list of task definition family names that match the
	// ListTaskDefinitionFamilies request.
	Families []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListTaskDefinitionFamiliesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListTaskDefinitionFamilies{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTaskDefinitionFamilies{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTaskDefinitionFamilies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "ListTaskDefinitionFamilies",
	}
}
