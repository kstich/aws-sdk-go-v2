// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the exclusions that are specified by the exclusions' ARNs.
func (c *Client) DescribeExclusions(ctx context.Context, params *DescribeExclusionsInput, optFns ...func(*Options)) (*DescribeExclusionsOutput, error) {
	stack := middleware.NewStack("DescribeExclusions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeExclusionsMiddlewares(stack)
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
	addOpDescribeExclusionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExclusions(options.Region), middleware.Before)
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
			OperationName: "DescribeExclusions",
			Err:           err,
		}
	}
	out := result.(*DescribeExclusionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExclusionsInput struct {
	// The list of ARNs that specify the exclusions that you want to describe.
	ExclusionArns []*string
	// The locale into which you want to translate the exclusion's title, description,
	// and recommendation.
	Locale types.Locale
}

type DescribeExclusionsOutput struct {
	// Information about the exclusions.
	Exclusions map[string]*types.Exclusion
	// Exclusion details that cannot be described. An error code is provided for each
	// failed item.
	FailedItems map[string]*types.FailedItemDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeExclusionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeExclusions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeExclusions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeExclusions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "DescribeExclusions",
	}
}
