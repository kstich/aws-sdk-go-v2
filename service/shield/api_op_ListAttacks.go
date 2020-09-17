// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all ongoing DDoS attacks or all DDoS attacks during a specified time
// period.
func (c *Client) ListAttacks(ctx context.Context, params *ListAttacksInput, optFns ...func(*Options)) (*ListAttacksOutput, error) {
	stack := middleware.NewStack("ListAttacks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListAttacksMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAttacks(options.Region), middleware.Before)
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
			OperationName: "ListAttacks",
			Err:           err,
		}
	}
	out := result.(*ListAttacksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttacksInput struct {
	// The ListAttacksRequest.NextMarker value from a previous call to
	// ListAttacksRequest. Pass null if this is the first call.
	NextToken *string
	// The ARN (Amazon Resource Name) of the resource that was attacked. If this is
	// left blank, all applicable resources for this account will be included.
	ResourceArns []*string
	// The end of the time period for the attacks. This is a timestamp type. The sample
	// request above indicates a number type because the default used by WAF is Unix
	// time in seconds. However any valid timestamp format
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types)
	// is allowed.
	EndTime *types.TimeRange
	// The maximum number of AttackSummary () objects to be returned. If this is left
	// blank, the first 20 results will be returned. This is a maximum value; it is
	// possible that AWS WAF will return the results in smaller batches. That is, the
	// number of AttackSummary () objects returned could be less than MaxResults, even
	// if there are still more AttackSummary () objects yet to return. If there are
	// more AttackSummary () objects to return, AWS WAF will always also return a
	// NextToken.
	MaxResults *int32
	// The start of the time period for the attacks. This is a timestamp type. The
	// sample request above indicates a number type because the default used by WAF is
	// Unix time in seconds. However any valid timestamp format
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types)
	// is allowed.
	StartTime *types.TimeRange
}

type ListAttacksOutput struct {
	// The token returned by a previous call to indicate that there is more data
	// available. If not null, more results are available. Pass this value for the
	// NextMarker parameter in a subsequent call to ListAttacks to retrieve the next
	// set of items. AWS WAF might return the list of AttackSummary () objects in
	// batches smaller than the number specified by MaxResults. If there are more
	// AttackSummary () objects to return, AWS WAF will always also return a NextToken.
	NextToken *string
	// The attack information for the specified time range.
	AttackSummaries []*types.AttackSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListAttacksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListAttacks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAttacks{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAttacks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "ListAttacks",
	}
}
