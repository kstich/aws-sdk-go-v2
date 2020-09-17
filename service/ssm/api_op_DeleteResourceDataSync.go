// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a Resource Data Sync configuration. After the configuration is deleted,
// changes to data on managed instances are no longer synced to or from the target.
// Deleting a sync configuration does not delete data.
func (c *Client) DeleteResourceDataSync(ctx context.Context, params *DeleteResourceDataSyncInput, optFns ...func(*Options)) (*DeleteResourceDataSyncOutput, error) {
	stack := middleware.NewStack("DeleteResourceDataSync", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteResourceDataSyncMiddlewares(stack)
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
	addOpDeleteResourceDataSyncValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteResourceDataSync(options.Region), middleware.Before)
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
			OperationName: "DeleteResourceDataSync",
			Err:           err,
		}
	}
	out := result.(*DeleteResourceDataSyncOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteResourceDataSyncInput struct {
	// Specify the type of resource data sync to delete.
	SyncType *string
	// The name of the configuration to delete.
	SyncName *string
}

type DeleteResourceDataSyncOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteResourceDataSyncMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteResourceDataSync{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteResourceDataSync{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteResourceDataSync(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DeleteResourceDataSync",
	}
}
