// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about all available AWS Trusted Advisor checks, including
// the name, ID, category, description, and metadata. You must specify a language
// code. The AWS Support API currently supports English ("en") and Japanese ("ja").
// The response contains a TrustedAdvisorCheckDescription () object for each check.
// You must set the AWS Region to us-east-1.
//
//     * You must have a Business or
// Enterprise support plan to use the AWS Support API.
//
//     * If you call the AWS
// Support API from an account that does not have a Business or Enterprise support
// plan, the SubscriptionRequiredException error message appears. For information
// about changing your support plan, see AWS Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) DescribeTrustedAdvisorChecks(ctx context.Context, params *DescribeTrustedAdvisorChecksInput, optFns ...func(*Options)) (*DescribeTrustedAdvisorChecksOutput, error) {
	stack := middleware.NewStack("DescribeTrustedAdvisorChecks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeTrustedAdvisorChecksMiddlewares(stack)
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
	addOpDescribeTrustedAdvisorChecksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrustedAdvisorChecks(options.Region), middleware.Before)
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
			OperationName: "DescribeTrustedAdvisorChecks",
			Err:           err,
		}
	}
	out := result.(*DescribeTrustedAdvisorChecksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrustedAdvisorChecksInput struct {
	// The ISO 639-1 code for the language in which AWS provides support. AWS Support
	// currently supports English ("en") and Japanese ("ja"). Language parameters must
	// be passed explicitly for operations that take them.
	Language *string
}

// Information about the Trusted Advisor checks returned by the
// DescribeTrustedAdvisorChecks () operation.
type DescribeTrustedAdvisorChecksOutput struct {
	// Information about all available Trusted Advisor checks.
	Checks []*types.TrustedAdvisorCheckDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeTrustedAdvisorChecksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrustedAdvisorChecks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrustedAdvisorChecks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTrustedAdvisorChecks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeTrustedAdvisorChecks",
	}
}
