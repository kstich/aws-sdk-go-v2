// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the available standards in Security Hub. For each standard,
// the results include the standard ARN, the name, and a description.
func (c *Client) DescribeStandards(ctx context.Context, params *DescribeStandardsInput, optFns ...func(*Options)) (*DescribeStandardsOutput, error) {
	stack := middleware.NewStack("DescribeStandards", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeStandardsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStandards(options.Region), middleware.Before)
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
			OperationName: "DescribeStandards",
			Err:           err,
		}
	}
	out := result.(*DescribeStandardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStandardsInput struct {
	// The maximum number of standards to return.
	MaxResults *int32
	// The token that is required for pagination. On your first call to the
	// DescribeStandards operation, set the value of this parameter to NULL. For
	// subsequent calls to the operation, to continue listing data, set the value of
	// this parameter to the value returned from the previous response.
	NextToken *string
}

type DescribeStandardsOutput struct {
	// A list of available standards.
	Standards []*types.Standard
	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeStandardsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeStandards{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeStandards{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStandards(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DescribeStandards",
	}
}
