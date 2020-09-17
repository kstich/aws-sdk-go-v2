// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// that is associated with an AWS CloudHSM cluster
// (https://docs.aws.amazon.com/cloudhsm/latest/userguide/clusters.html) that you
// own and manage. This operation is part of the Custom Key Store feature
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in AWS KMS, which combines the convenience and extensive integration of
// AWS KMS with the isolation and control of a single-tenant key store. Before you
// create the custom key store, you must assemble the required elements, including
// an AWS CloudHSM cluster that fulfills the requirements for a custom key store.
// For details about the required elements, see Assemble the Prerequisites
// (https://docs.aws.amazon.com/kms/latest/developerguide/create-keystore.html#before-keystore)
// in the AWS Key Management Service Developer Guide. When the operation completes
// successfully, it returns the ID of the new custom key store. Before you can use
// your new custom key store, you need to use the ConnectCustomKeyStore ()
// operation to connect the new key store to its AWS CloudHSM cluster. Even if you
// are not going to use your custom key store immediately, you might want to
// connect it to verify that all settings are correct and then disconnect it until
// you are ready to use it. For help with failures, see Troubleshooting a Custom
// Key Store
// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) CreateCustomKeyStore(ctx context.Context, params *CreateCustomKeyStoreInput, optFns ...func(*Options)) (*CreateCustomKeyStoreOutput, error) {
	stack := middleware.NewStack("CreateCustomKeyStore", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateCustomKeyStoreMiddlewares(stack)
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
	addOpCreateCustomKeyStoreValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomKeyStore(options.Region), middleware.Before)
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
			OperationName: "CreateCustomKeyStore",
			Err:           err,
		}
	}
	out := result.(*CreateCustomKeyStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomKeyStoreInput struct {
	// Identifies the AWS CloudHSM cluster for the custom key store. Enter the cluster
	// ID of any active AWS CloudHSM cluster that is not already associated with a
	// custom key store. To find the cluster ID, use the DescribeClusters
	// (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html)
	// operation.
	CloudHsmClusterId *string
	// Specifies a friendly name for the custom key store. The name must be unique in
	// your AWS account.
	CustomKeyStoreName *string
	// Enter the content of the trust anchor certificate for the cluster. This is the
	// content of the customerCA.crt file that you created when you initialized the
	// cluster
	// (https://docs.aws.amazon.com/cloudhsm/latest/userguide/initialize-cluster.html).
	TrustAnchorCertificate *string
	// Enter the password of the kmsuser crypto user (CU) account
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-store-concepts.html#concept-kmsuser)
	// in the specified AWS CloudHSM cluster. AWS KMS logs into the cluster as this
	// user to manage key material on your behalf. The password must be a string of 7
	// to 32 characters. Its value is case sensitive. This parameter tells AWS KMS the
	// kmsuser account password; it does not change the password in the AWS CloudHSM
	// cluster.
	KeyStorePassword *string
}

type CreateCustomKeyStoreOutput struct {
	// A unique identifier for the new custom key store.
	CustomKeyStoreId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateCustomKeyStoreMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCustomKeyStore{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCustomKeyStore{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCustomKeyStore(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "CreateCustomKeyStore",
	}
}
