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

// Starts a job to create a one-time backup of the specified resource.
func (c *Client) StartBackupJob(ctx context.Context, params *StartBackupJobInput, optFns ...func(*Options)) (*StartBackupJobOutput, error) {
	stack := middleware.NewStack("StartBackupJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartBackupJobMiddlewares(stack)
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
	addOpStartBackupJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartBackupJob(options.Region), middleware.Before)

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
			OperationName: "StartBackupJob",
			Err:           err,
		}
	}
	out := result.(*StartBackupJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBackupJobInput struct {
	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. AWS Backup will transition and expire backups automatically
	// according to the lifecycle that you define. Backups transitioned to cold storage
	// must be stored in cold storage for a minimum of 90 days. Therefore, the “expire
	// after days” setting must be 90 days greater than the “transition to cold after
	// days” setting. The “transition to cold after days” setting cannot be changed
	// after a backup has been transitioned to cold.
	Lifecycle *types.Lifecycle
	// A value in minutes after a backup is scheduled before a job will be canceled if
	// it doesn't start successfully. This value is optional.
	StartWindowMinutes *int64
	// A value in minutes after a backup job is successfully started before it must be
	// completed or it will be canceled by AWS Backup. This value is optional.
	CompleteWindowMinutes *int64
	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string
	// To help organize your resources, you can assign your own metadata to the
	// resources that you create. Each tag is a key-value pair.
	RecoveryPointTags map[string]*string
	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	BackupVaultName *string
	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of
	// the ARN depends on the resource type.
	ResourceArn *string
	// A customer chosen string that can be used to distinguish between calls to
	// StartBackupJob.
	IdempotencyToken *string
}

type StartBackupJobOutput struct {
	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string
	// Uniquely identifies a request to AWS Backup to back up a resource.
	BackupJobId *string
	// The date and time that a backup job is started, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartBackupJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartBackupJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartBackupJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartBackupJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "StartBackupJob",
	}
}
