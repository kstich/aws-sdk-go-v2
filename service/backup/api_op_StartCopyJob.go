// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Starts a job to create a one-time copy of the specified resource.
func (c *Client) StartCopyJob(ctx context.Context, params *StartCopyJobInput, optFns ...func(*Options)) (*StartCopyJobOutput, error) {
	stack := middleware.NewStack("StartCopyJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartCopyJobMiddlewares(stack)
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
	addOpStartCopyJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartCopyJob(options.Region), middleware.Before)
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
			OperationName: "StartCopyJob",
			Err:           err,
		}
	}
	out := result.(*StartCopyJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartCopyJobInput struct {
	// The name of a logical source container where backups are stored. Backup vaults
	// are identified by names that are unique to the account used to create them and
	// the AWS Region where they are created. They consist of lowercase letters,
	// numbers, and hyphens.
	SourceBackupVaultName *string
	// A customer chosen string that can be used to distinguish between calls to
	// StartCopyJob.
	IdempotencyToken *string
	// An Amazon Resource Name (ARN) that uniquely identifies a destination backup
	// vault to copy to; for example,
	// arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	DestinationBackupVaultArn *string
	// Specifies the IAM role ARN used to copy the target recovery point; for example,
	// arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string
	// An ARN that uniquely identifies a recovery point to use for the copy job; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string
	// Contains an array of Transition objects specifying how long in days before a
	// recovery point transitions to cold storage or is deleted. Backups transitioned
	// to cold storage must be stored in cold storage for a minimum of 90 days.
	// Therefore, on the console, the “expire after days” setting must be 90 days
	// greater than the “transition to cold after days” setting. The “transition to
	// cold after days” setting cannot be changed after a backup has been transitioned
	// to cold.
	Lifecycle *types.Lifecycle
}

type StartCopyJobOutput struct {
	// The date and time that a copy job is started, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time
	// Uniquely identifies a copy job.
	CopyJobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartCopyJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartCopyJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartCopyJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartCopyJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "StartCopyJob",
	}
}
