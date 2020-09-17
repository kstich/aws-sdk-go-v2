// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Previews the agents installed on the EC2 instances that are part of the
// specified assessment target.
func (c *Client) PreviewAgents(ctx context.Context, params *PreviewAgentsInput, optFns ...func(*Options)) (*PreviewAgentsOutput, error) {
	stack := middleware.NewStack("PreviewAgents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPreviewAgentsMiddlewares(stack)
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
	addOpPreviewAgentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPreviewAgents(options.Region), middleware.Before)
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
			OperationName: "PreviewAgents",
			Err:           err,
		}
	}
	out := result.(*PreviewAgentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PreviewAgentsInput struct {
	// The ARN of the assessment target whose agents you want to preview.
	PreviewAgentsArn *string
	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 10. The maximum value is 500.
	MaxResults *int32
	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the PreviewAgents action. Subsequent
	// calls to the action fill nextToken in the request with the value of NextToken
	// from the previous response to continue listing data.
	NextToken *string
}

type PreviewAgentsOutput struct {
	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to be
	// listed, this parameter is set to null.
	NextToken *string
	// The resulting list of agents.
	AgentPreviews []*types.AgentPreview

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPreviewAgentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPreviewAgents{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPreviewAgents{}, middleware.After)
}

func newServiceMetadataMiddleware_opPreviewAgents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "PreviewAgents",
	}
}
