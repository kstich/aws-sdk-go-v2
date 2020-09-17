// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds tags to a project.
func (c *Client) TagProject(ctx context.Context, params *TagProjectInput, optFns ...func(*Options)) (*TagProjectOutput, error) {
	stack := middleware.NewStack("TagProject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTagProjectMiddlewares(stack)
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
	addOpTagProjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagProject(options.Region), middleware.Before)
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
			OperationName: "TagProject",
			Err:           err,
		}
	}
	out := result.(*TagProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagProjectInput struct {
	// The ID of the project you want to add a tag to.
	Id *string
	// The tags you want to add to the project.
	Tags map[string]*string
}

type TagProjectOutput struct {
	// The tags for the project.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTagProjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTagProject{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagProject{}, middleware.After)
}

func newServiceMetadataMiddleware_opTagProject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar",
		OperationName: "TagProject",
	}
}
