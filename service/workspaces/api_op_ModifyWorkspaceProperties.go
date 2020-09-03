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

// Modifies the specified WorkSpace properties. For important information about how
// to modify the size of the root and user volumes, see  Modify a WorkSpace
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/modify-workspaces.html).
func (c *Client) ModifyWorkspaceProperties(ctx context.Context, params *ModifyWorkspacePropertiesInput, optFns ...func(*Options)) (*ModifyWorkspacePropertiesOutput, error) {
	stack := middleware.NewStack("ModifyWorkspaceProperties", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpModifyWorkspacePropertiesMiddlewares(stack)
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
	addOpModifyWorkspacePropertiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyWorkspaceProperties(options.Region), middleware.Before)

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
			OperationName: "ModifyWorkspaceProperties",
			Err:           err,
		}
	}
	out := result.(*ModifyWorkspacePropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyWorkspacePropertiesInput struct {
	// The identifier of the WorkSpace.
	WorkspaceId *string
	// The properties of the WorkSpace.
	WorkspaceProperties *types.WorkspaceProperties
}

type ModifyWorkspacePropertiesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpModifyWorkspacePropertiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpModifyWorkspaceProperties{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyWorkspaceProperties{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyWorkspaceProperties(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "ModifyWorkspaceProperties",
	}
}
