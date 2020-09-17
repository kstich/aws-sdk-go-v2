// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes available AWS services in a prefix list format, which includes the
// prefix list name and prefix list ID of the service and the IP address range for
// the service. We recommend that you use DescribeManagedPrefixLists () instead.
func (c *Client) DescribePrefixLists(ctx context.Context, params *DescribePrefixListsInput, optFns ...func(*Options)) (*DescribePrefixListsOutput, error) {
	stack := middleware.NewStack("DescribePrefixLists", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribePrefixListsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePrefixLists(options.Region), middleware.Before)
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
			OperationName: "DescribePrefixLists",
			Err:           err,
		}
	}
	out := result.(*DescribePrefixListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePrefixListsInput struct {
	// The token for the next page of results.
	NextToken *string
	// One or more filters.
	//
	//     * prefix-list-id: The ID of a prefix list.
	//
	//     *
	// prefix-list-name: The name of a prefix list.
	Filters []*types.Filter
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32
	// One or more prefix list IDs.
	PrefixListIds []*string
}

type DescribePrefixListsOutput struct {
	// All available prefix lists.
	PrefixLists []*types.PrefixList
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribePrefixListsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribePrefixLists{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribePrefixLists{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePrefixLists(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribePrefixLists",
	}
}
