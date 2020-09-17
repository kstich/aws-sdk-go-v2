// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays the template object (which includes the Subject line, HTML part and
// text part) for the template you specify. You can execute this operation no more
// than once per second.
func (c *Client) GetTemplate(ctx context.Context, params *GetTemplateInput, optFns ...func(*Options)) (*GetTemplateOutput, error) {
	stack := middleware.NewStack("GetTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetTemplateMiddlewares(stack)
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
	addOpGetTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTemplate(options.Region), middleware.Before)
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
			OperationName: "GetTemplate",
			Err:           err,
		}
	}
	out := result.(*GetTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTemplateInput struct {
	// The name of the template you want to retrieve.
	TemplateName *string
}

type GetTemplateOutput struct {
	// The content of the email, composed of a subject line, an HTML part, and a
	// text-only part.
	Template *types.Template

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetTemplate",
	}
}
