// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns metadata about a backup vault specified by its name.
func (c *Client) DescribeBackupVault(ctx context.Context, params *DescribeBackupVaultInput, optFns ...func(*Options)) (*DescribeBackupVaultOutput, error) {
	stack := middleware.NewStack("DescribeBackupVault", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeBackupVaultMiddlewares(stack)
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
	addOpDescribeBackupVaultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBackupVault(options.Region), middleware.Before)

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
			OperationName: "DescribeBackupVault",
			Err:           err,
		}
	}
	out := result.(*DescribeBackupVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBackupVaultInput struct {
	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	BackupVaultName *string
}

type DescribeBackupVaultOutput struct {
	// The number of recovery points that are stored in a backup vault.
	NumberOfRecoveryPoints *int64
	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Region where they are created. They consist of lowercase letters, numbers, and
	// hyphens.
	BackupVaultName *string
	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string
	// The date and time that a backup vault is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time
	// The server-side encryption key that is used to protect your backups; for
	// example,
	// arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	EncryptionKeyArn *string
	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of executing the operation twice.
	CreatorRequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeBackupVaultMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBackupVault{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBackupVault{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeBackupVault(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeBackupVault",
	}
}
