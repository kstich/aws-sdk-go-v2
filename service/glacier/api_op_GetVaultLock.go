// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation retrieves the following attributes from the lock-policy
// subresource set on the specified vault:
//
//     * The vault lock policy set on the
// vault.
//
//     * The state of the vault lock, which is either InProgess or
// Locked.
//
//     * When the lock ID expires. The lock ID is used to complete the
// vault locking process.
//
//     * When the vault lock was initiated and put into the
// InProgress state.
//
//     <p>A vault lock is put into the <code>InProgress</code>
// state by calling <a>InitiateVaultLock</a>. A vault lock is put into the
// <code>Locked</code> state by calling <a>CompleteVaultLock</a>. You can abort the
// vault locking process by calling <a>AbortVaultLock</a>. For more information
// about the vault locking process, <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html">Amazon
// Glacier Vault Lock</a>. </p> <p>If there is no vault lock policy set on the
// vault, the operation returns a <code>404 Not found</code> error. For more
// information about vault lock policies, <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock-policy.html">Amazon
// Glacier Access Control with Vault Lock Policies</a>. </p>
func (c *Client) GetVaultLock(ctx context.Context, params *GetVaultLockInput, optFns ...func(*Options)) (*GetVaultLockOutput, error) {
	stack := middleware.NewStack("GetVaultLock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetVaultLockMiddlewares(stack)
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
	addOpGetVaultLockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetVaultLock(options.Region), middleware.Before)

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
			OperationName: "GetVaultLock",
			Err:           err,
		}
	}
	out := result.(*GetVaultLockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input values for GetVaultLock.
type GetVaultLockInput struct {
	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	AccountId *string
	// The name of the vault.
	VaultName *string
}

// Contains the Amazon S3 Glacier response to your request.
type GetVaultLockOutput struct {
	// The state of the vault lock. InProgress or Locked.
	State *string
	// The vault lock policy as a JSON string, which uses "\" as an escape character.
	Policy *string
	// The UTC date and time at which the lock ID expires. This value can be null if
	// the vault lock is in a Locked state.
	ExpirationDate *string
	// The UTC date and time at which the vault lock was put into the InProgress state.
	CreationDate *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetVaultLockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetVaultLock{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVaultLock{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetVaultLock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "GetVaultLock",
	}
}
