// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates user details within the UpdateUserRequestItem () object for up to 20
// users for the specified Amazon Chime account. Currently, only LicenseType
// updates are supported for this action.
func (c *Client) BatchUpdateUser(ctx context.Context, params *BatchUpdateUserInput, optFns ...func(*Options)) (*BatchUpdateUserOutput, error) {
	stack := middleware.NewStack("BatchUpdateUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchUpdateUserMiddlewares(stack)
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
	addOpBatchUpdateUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateUser(options.Region), middleware.Before)
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
			OperationName: "BatchUpdateUser",
			Err:           err,
		}
	}
	out := result.(*BatchUpdateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateUserInput struct {
	// The Amazon Chime account ID.
	AccountId *string
	// The request containing the user IDs and details to update.
	UpdateUserRequestItems []*types.UpdateUserRequestItem
}

type BatchUpdateUserOutput struct {
	// If the BatchUpdateUser () action fails for one or more of the user IDs in the
	// request, a list of the user IDs is returned, along with error codes and error
	// messages.
	UserErrors []*types.UserError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchUpdateUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateUser{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchUpdateUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchUpdateUser",
	}
}
