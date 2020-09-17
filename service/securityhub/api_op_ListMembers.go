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

// Lists details about all member accounts for the current Security Hub master
// account.
func (c *Client) ListMembers(ctx context.Context, params *ListMembersInput, optFns ...func(*Options)) (*ListMembersOutput, error) {
	stack := middleware.NewStack("ListMembers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListMembersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMembers(options.Region), middleware.Before)
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
			OperationName: "ListMembers",
			Err:           err,
		}
	}
	out := result.(*ListMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMembersInput struct {
	// Specifies which member accounts to include in the response based on their
	// relationship status with the master account. The default value is TRUE. If
	// OnlyAssociated is set to TRUE, the response includes member accounts whose
	// relationship status with the master is set to ENABLED or DISABLED. If
	// OnlyAssociated is set to FALSE, the response includes all existing member
	// accounts.
	OnlyAssociated *bool
	// The token that is required for pagination. On your first call to the ListMembers
	// operation, set the value of this parameter to NULL. For subsequent calls to the
	// operation, to continue listing data, set the value of this parameter to the
	// value returned from the previous response.
	NextToken *string
	// The maximum number of items to return in the response.
	MaxResults *int32
}

type ListMembersOutput struct {
	// Member details returned by the operation.
	Members []*types.Member
	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListMembersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListMembers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListMembers{}, middleware.After)
}

func newServiceMetadataMiddleware_opListMembers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "ListMembers",
	}
}
