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

// Generates a unique symmetric data key for client-side encryption. This operation
// returns a plaintext copy of the data key and a copy that is encrypted under a
// customer master key (CMK) that you specify. You can use the plaintext key to
// encrypt your data outside of AWS KMS and store the encrypted data key with the
// encrypted data.  <p> <code>GenerateDataKey</code> returns a unique data key for
// each request. The bytes in the plaintext key are not related to the caller or
// the CMK.</p> <p>To generate a data key, specify the symmetric CMK that will be
// used to encrypt the data key. You cannot use an asymmetric CMK to generate data
// keys. To get the type of your CMK, use the <a>DescribeKey</a> operation. You
// must also specify the length of the data key. Use either the
// <code>KeySpec</code> or <code>NumberOfBytes</code> parameters (but not both).
// For 128-bit and 256-bit data keys, use the <code>KeySpec</code> parameter. </p>
// <p>To get only an encrypted copy of the data key, use
// <a>GenerateDataKeyWithoutPlaintext</a>. To generate an asymmetric data key pair,
// use the <a>GenerateDataKeyPair</a> or <a>GenerateDataKeyPairWithoutPlaintext</a>
// operation. To get a cryptographically secure random byte string, use
// <a>GenerateRandom</a>.</p> <p>You can use the optional encryption context to add
// additional security to the encryption operation. If you specify an
// <code>EncryptionContext</code>, you must specify the same encryption context (a
// case-sensitive exact match) when decrypting the encrypted data key. Otherwise,
// the request to decrypt fails with an <code>InvalidCiphertextException</code>.
// For more information, see <a
// href="https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context">Encryption
// Context</a> in the <i>AWS Key Management Service Developer Guide</i>.</p> <p>The
// CMK that you use for this operation must be in a compatible key state. For
// details, see How Key State Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide. How to use your data key We
// recommend that you use the following pattern to encrypt data locally in your
// application. You can write your own code or use a client-side encryption
// library, such as the AWS Encryption SDK
// (https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/), the Amazon
// DynamoDB Encryption Client
// (https://docs.aws.amazon.com/dynamodb-encryption-client/latest/devguide/), or
// Amazon S3 client-side encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html)
// to do these tasks for you. To encrypt data outside of AWS KMS:
//
//     * Use the
// GenerateDataKey operation to get a data key.
//
//     * Use the plaintext data key
// (in the Plaintext field of the response) to encrypt your data outside of AWS
// KMS. Then erase the plaintext data key from memory.
//
//     * Store the encrypted
// data key (in the CiphertextBlob field of the response) with the encrypted
// data.
//
// To decrypt data outside of AWS KMS:
//
//     * Use the Decrypt () operation
// to decrypt the encrypted data key. The operation returns a plaintext copy of the
// data key.
//
//     * Use the plaintext data key to decrypt data outside of AWS KMS,
// then erase the plaintext data key from memory.
func (c *Client) GenerateDataKey(ctx context.Context, params *GenerateDataKeyInput, optFns ...func(*Options)) (*GenerateDataKeyOutput, error) {
	stack := middleware.NewStack("GenerateDataKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGenerateDataKeyMiddlewares(stack)
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
	addOpGenerateDataKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateDataKey(options.Region), middleware.Before)
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
			OperationName: "GenerateDataKey",
			Err:           err,
		}
	}
	out := result.(*GenerateDataKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateDataKeyInput struct {
	// Specifies the length of the data key in bytes. For example, use the value 64 to
	// generate a 512-bit data key (64 bytes is 512 bits). For 128-bit (16-byte) and
	// 256-bit (32-byte) data keys, use the KeySpec parameter. You must specify either
	// the KeySpec or the NumberOfBytes parameter (but not both) in every
	// GenerateDataKey request.
	NumberOfBytes *int32
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
	// Identifies the symmetric CMK that encrypts the data key.  <p>To specify a CMK,
	// use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. When using
	// an alias name, prefix it with <code>"alias/"</code>. To specify a CMK in a
	// different AWS account, you must use the key ARN or alias ARN.</p> <p>For
	// example:</p> <ul> <li> <p>Key ID:
	// <code>1234abcd-12ab-34cd-56ef-1234567890ab</code> </p> </li> <li> <p>Key ARN:
	// <code>arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab</code>
	// </p> </li> <li> <p>Alias name: <code>alias/ExampleAlias</code> </p> </li> <li>
	// <p>Alias ARN: <code>arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias</code>
	// </p> </li> </ul> <p>To get the key ID and key ARN for a CMK, use <a>ListKeys</a>
	// or <a>DescribeKey</a>. To get the alias name and alias ARN, use
	// <a>ListAliases</a>.</p>
	KeyId *string
	// A list of grant tokens. For more information, see Grant Tokens
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []*string
	// Specifies the length of the data key. Use AES_128 to generate a 128-bit
	// symmetric key, or AES_256 to generate a 256-bit symmetric key. You must specify
	// either the KeySpec or the NumberOfBytes parameter (but not both) in every
	// GenerateDataKey request.
	KeySpec types.DataKeySpec
}

type GenerateDataKeyOutput struct {
	// The plaintext data key. When you use the HTTP API or the AWS CLI, the value is
	// Base64-encoded. Otherwise, it is not Base64-encoded. Use this data key to
	// encrypt your data outside of KMS. Then, remove it from memory as soon as
	// possible.
	Plaintext []byte
	// The encrypted copy of the data key. When you use the HTTP API or the AWS CLI,
	// the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	CiphertextBlob []byte
	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the CMK that encrypted the data key.
	KeyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGenerateDataKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateDataKey{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateDataKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opGenerateDataKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "GenerateDataKey",
	}
}
