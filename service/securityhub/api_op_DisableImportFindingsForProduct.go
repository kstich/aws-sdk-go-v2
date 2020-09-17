// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the integration of the specified product with Security Hub. After the
// integration is disabled, findings from that product are no longer sent to
// Security Hub.
func (c *Client) DisableImportFindingsForProduct(ctx context.Context, params *DisableImportFindingsForProductInput, optFns ...func(*Options)) (*DisableImportFindingsForProductOutput, error) {
	stack := middleware.NewStack("DisableImportFindingsForProduct", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisableImportFindingsForProductMiddlewares(stack)
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
	addOpDisableImportFindingsForProductValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableImportFindingsForProduct(options.Region), middleware.Before)
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
			OperationName: "DisableImportFindingsForProduct",
			Err:           err,
		}
	}
	out := result.(*DisableImportFindingsForProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableImportFindingsForProductInput struct {
	// The ARN of the integrated product to disable the integration for.
	ProductSubscriptionArn *string
}

type DisableImportFindingsForProductOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisableImportFindingsForProductMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisableImportFindingsForProduct{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisableImportFindingsForProduct{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableImportFindingsForProduct(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DisableImportFindingsForProduct",
	}
}
