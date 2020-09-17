// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// An SaaS partner can use this operation to display the AWS account ID that a
// particular partner event source name is associated with. This operation is not
// used by AWS customers.
func (c *Client) ListPartnerEventSourceAccounts(ctx context.Context, params *ListPartnerEventSourceAccountsInput, optFns ...func(*Options)) (*ListPartnerEventSourceAccountsOutput, error) {
	stack := middleware.NewStack("ListPartnerEventSourceAccounts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListPartnerEventSourceAccountsMiddlewares(stack)
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
	addOpListPartnerEventSourceAccountsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPartnerEventSourceAccounts(options.Region), middleware.Before)
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
			OperationName: "ListPartnerEventSourceAccounts",
			Err:           err,
		}
	}
	out := result.(*ListPartnerEventSourceAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPartnerEventSourceAccountsInput struct {
	// The token returned by a previous call to this operation. Specifying this
	// retrieves the next set of results.
	NextToken *string
	// The name of the partner event source to display account information about.
	EventSourceName *string
	// Specifying this limits the number of results returned by this operation. The
	// operation also returns a NextToken which you can use in a subsequent operation
	// to retrieve the next set of results.
	Limit *int32
}

type ListPartnerEventSourceAccountsOutput struct {
	// The list of partner event sources returned by the operation.
	PartnerEventSourceAccounts []*types.PartnerEventSourceAccount
	// A token you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListPartnerEventSourceAccountsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListPartnerEventSourceAccounts{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPartnerEventSourceAccounts{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPartnerEventSourceAccounts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "ListPartnerEventSourceAccounts",
	}
}
