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

// Use UpdateRole () instead. Modifies only the description of a role. This
// operation performs the same function as the Description parameter in the
// UpdateRole operation.
func (c *Client) UpdateRoleDescription(ctx context.Context, params *UpdateRoleDescriptionInput, optFns ...func(*Options)) (*UpdateRoleDescriptionOutput, error) {
	stack := middleware.NewStack("UpdateRoleDescription", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUpdateRoleDescriptionMiddlewares(stack)
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
	addOpUpdateRoleDescriptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRoleDescription(options.Region), middleware.Before)
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
			OperationName: "UpdateRoleDescription",
			Err:           err,
		}
	}
	out := result.(*UpdateRoleDescriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoleDescriptionInput struct {
	// The name of the role that you want to modify.
	RoleName *string
	// The new description that you want to apply to the specified role.
	Description *string
}

type UpdateRoleDescriptionOutput struct {
	// A structure that contains details about the modified role.
	Role *types.Role

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUpdateRoleDescriptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUpdateRoleDescription{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateRoleDescription{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRoleDescription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateRoleDescription",
	}
}
