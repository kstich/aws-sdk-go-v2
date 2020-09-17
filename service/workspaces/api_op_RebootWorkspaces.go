// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Reboots the specified WorkSpaces. You cannot reboot a WorkSpace unless its state
// is AVAILABLE or UNHEALTHY. This operation is asynchronous and returns before the
// WorkSpaces have rebooted.
func (c *Client) RebootWorkspaces(ctx context.Context, params *RebootWorkspacesInput, optFns ...func(*Options)) (*RebootWorkspacesOutput, error) {
	stack := middleware.NewStack("RebootWorkspaces", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRebootWorkspacesMiddlewares(stack)
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
	addOpRebootWorkspacesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRebootWorkspaces(options.Region), middleware.Before)
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
			OperationName: "RebootWorkspaces",
			Err:           err,
		}
	}
	out := result.(*RebootWorkspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RebootWorkspacesInput struct {
	// The WorkSpaces to reboot. You can specify up to 25 WorkSpaces.
	RebootWorkspaceRequests []*types.RebootRequest
}

type RebootWorkspacesOutput struct {
	// Information about the WorkSpaces that could not be rebooted.
	FailedRequests []*types.FailedWorkspaceChangeRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRebootWorkspacesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRebootWorkspaces{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRebootWorkspaces{}, middleware.After)
}

func newServiceMetadataMiddleware_opRebootWorkspaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "RebootWorkspaces",
	}
}
