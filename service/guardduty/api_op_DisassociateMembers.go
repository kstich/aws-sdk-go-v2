// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates GuardDuty member accounts (to the current GuardDuty master
// account) specified by the account IDs.
func (c *Client) DisassociateMembers(ctx context.Context, params *DisassociateMembersInput, optFns ...func(*Options)) (*DisassociateMembersOutput, error) {
	stack := middleware.NewStack("DisassociateMembers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisassociateMembersMiddlewares(stack)
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
	addOpDisassociateMembersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateMembers(options.Region), middleware.Before)
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
			OperationName: "DisassociateMembers",
			Err:           err,
		}
	}
	out := result.(*DisassociateMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateMembersInput struct {
	// A list of account IDs of the GuardDuty member accounts that you want to
	// disassociate from the master account.
	AccountIds []*string
	// The unique ID of the detector of the GuardDuty account whose members you want to
	// disassociate from the master account.
	DetectorId *string
}

type DisassociateMembersOutput struct {
	// A list of objects that contain the unprocessed account and a result string that
	// explains why it was unprocessed.
	UnprocessedAccounts []*types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisassociateMembersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateMembers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateMembers{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateMembers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "DisassociateMembers",
	}
}
