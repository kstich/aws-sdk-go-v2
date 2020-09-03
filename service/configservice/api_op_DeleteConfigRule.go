// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified AWS Config rule and all of its evaluation results. AWS
// Config sets the state of a rule to DELETING until the deletion is complete. You
// cannot update a rule while it is in this state. If you make a PutConfigRule or
// DeleteConfigRule request for the rule, you will receive a
// ResourceInUseException. You can check the state of a rule by using the
// DescribeConfigRules request.
func (c *Client) DeleteConfigRule(ctx context.Context, params *DeleteConfigRuleInput, optFns ...func(*Options)) (*DeleteConfigRuleOutput, error) {
	stack := middleware.NewStack("DeleteConfigRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteConfigRuleMiddlewares(stack)
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
	addOpDeleteConfigRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigRule(options.Region), middleware.Before)

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
			OperationName: "DeleteConfigRule",
			Err:           err,
		}
	}
	out := result.(*DeleteConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteConfigRuleInput struct {
	// The name of the AWS Config rule that you want to delete.
	ConfigRuleName *string
}

type DeleteConfigRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteConfigRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteConfigRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteConfigRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteConfigRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteConfigRule",
	}
}