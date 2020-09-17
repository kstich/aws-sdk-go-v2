// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates a unique symmetric data key. This operation returns a data key that is
// encrypted under a customer master key (CMK) that you specify. To request an
// asymmetric data key pair, use the GenerateDataKeyPair () or
// GenerateDataKeyPairWithoutPlaintext () operations.
// GenerateDataKeyWithoutPlaintext is identical to the GenerateDataKey () operation
// except that returns only the encrypted copy of the data key. This operation is
// useful for systems that need to encrypt data at some point, but not immediately.
// When you need to encrypt the data, you call the Decrypt () operation on the
// encrypted copy of the key. It's also useful in distributed systems with
// different levels of trust. For example, you might store encrypted data in
// containers. One component of your system creates new containers and stores an
// encrypted data key with each container. Then, a different component puts the
// data into the containers. That component first decrypts the data key, uses the
// plaintext data key to encrypt data, puts the encrypted data into the container,
// and then destroys the plaintext data key. In this system, the component that
// creates the containers never sees the plaintext data key.
// GenerateDataKeyWithoutPlaintext returns a unique data key for each request. The
// bytes in the keys are not related to the caller or CMK that is used to encrypt
// the private key.  <p>To generate a data key, you must specify the symmetric
// customer master key (CMK) that is used to encrypt the data key. You cannot use
// an asymmetric CMK to generate a data key. To get the type of your CMK, use the
// <a>DescribeKey</a> operation.</p> <p>If the operation succeeds, you will find
// the encrypted copy of the data key in the <code>CiphertextBlob</code> field.</p>
// <p>You can use the optional encryption context to add additional security to the
// encryption operation. If you specify an <code>EncryptionContext</code>, you must
// specify the same encryption context (a case-sensitive exact match) when
// decrypting the encrypted data key. Otherwise, the request to decrypt fails with
// an <code>InvalidCiphertextException</code>. For more information, see <a
// href="https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context">Encryption
// Context</a> in the <i>AWS Key Management Service Developer Guide</i>.</p> <p>The
// CMK that you use for this operation must be in a compatible key state. For
// details, see How Key State Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) GenerateDataKeyWithoutPlaintext(ctx context.Context, params *GenerateDataKeyWithoutPlaintextInput, optFns ...func(*Options)) (*GenerateDataKeyWithoutPlaintextOutput, error) {
	stack := middleware.NewStack("GenerateDataKeyWithoutPlaintext", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGenerateDataKeyWithoutPlaintextMiddlewares(stack)
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
	addOpGenerateDataKeyWithoutPlaintextValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateDataKeyWithoutPlaintext(options.Region), middleware.Before)
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
			OperationName: "GenerateDataKeyWithoutPlaintext",
			Err:           err,
		}
	}
	out := result.(*GenerateDataKeyWithoutPlaintextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateDataKeyWithoutPlaintextInput struct {
	// Specifies the encryption context that will be used when encrypting the data key.
	// An encryption context is a collection of non-secret key-value pairs that
	// represents additional authenticated data. When you use an encryption context to
	// encrypt data, you must specify the same (an exact case-sensitive match)
	// encryption context to decrypt the data. An encryption context is optional when
	// encrypting with a symmetric CMK, but it is highly recommended. For more
	// information, see Encryption Context
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	EncryptionContext map[string]*string
	// The identifier of the symmetric customer master key (CMK) that encrypts the data
	// key. To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name,
	// or alias ARN. When using an alias name, prefix it with "alias/". To specify a
	// CMK in a different AWS account, you must use the key ARN or alias ARN. For
	// example:
	//
	//     * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//
	// * Alias name: alias/ExampleAlias
	//
	//     * Alias ARN:
	// arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key
	// ARN for a CMK, use ListKeys () or DescribeKey (). To get the alias name and
	// alias ARN, use ListAliases ().
	KeyId *string
	// A list of grant tokens. For more information, see Grant Tokens
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []*string
	// The length of the data key. Use AES_128 to generate a 128-bit symmetric key, or
	// AES_256 to generate a 256-bit symmetric key.
	KeySpec types.DataKeySpec
	// The length of the data key in bytes. For example, use the value 64 to generate a
	// 512-bit data key (64 bytes is 512 bits). For common key lengths (128-bit and
	// 256-bit symmetric keys), we recommend that you use the KeySpec field instead of
	// this one.
	NumberOfBytes *int32
}

type GenerateDataKeyWithoutPlaintextOutput struct {
	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the CMK that encrypted the data key.
	KeyId *string
	// The encrypted data key. When you use the HTTP API or the AWS CLI, the value is
	// Base64-encoded. Otherwise, it is not Base64-encoded.
	CiphertextBlob []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGenerateDataKeyWithoutPlaintextMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateDataKeyWithoutPlaintext{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateDataKeyWithoutPlaintext{}, middleware.After)
}

func newServiceMetadataMiddleware_opGenerateDataKeyWithoutPlaintext(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "GenerateDataKeyWithoutPlaintext",
	}
}
