// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the status of your service-linked role deletion. After you use the
// DeleteServiceLinkedRole () API operation to submit a service-linked role for
// deletion, you can use the DeletionTaskId parameter in
// GetServiceLinkedRoleDeletionStatus to check the status of the deletion. If the
// deletion fails, this operation returns the reason that it failed, if that
// information is returned by the service.
func (c *Client) GetServiceLinkedRoleDeletionStatus(ctx context.Context, params *GetServiceLinkedRoleDeletionStatusInput, optFns ...func(*Options)) (*GetServiceLinkedRoleDeletionStatusOutput, error) {
	stack := middleware.NewStack("GetServiceLinkedRoleDeletionStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetServiceLinkedRoleDeletionStatusMiddlewares(stack)
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
	addOpGetServiceLinkedRoleDeletionStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceLinkedRoleDeletionStatus(options.Region), middleware.Before)
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
			OperationName: "GetServiceLinkedRoleDeletionStatus",
			Err:           err,
		}
	}
	out := result.(*GetServiceLinkedRoleDeletionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceLinkedRoleDeletionStatusInput struct {
	// The deletion task identifier. This identifier is returned by the
	// DeleteServiceLinkedRole () operation in the format task/aws-service-role///.
	DeletionTaskId *string
}

type GetServiceLinkedRoleDeletionStatusOutput struct {
	// An object that contains details about the reason the deletion failed.
	Reason *types.DeletionTaskFailureReasonType
	// The status of the deletion.
	Status types.DeletionTaskStatusType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetServiceLinkedRoleDeletionStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetServiceLinkedRoleDeletionStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetServiceLinkedRoleDeletionStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetServiceLinkedRoleDeletionStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetServiceLinkedRoleDeletionStatus",
	}
}
