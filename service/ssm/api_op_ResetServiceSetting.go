// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// ServiceSetting is an account-level setting for an AWS service. This setting
// defines how a user interacts with or uses a service or a feature of a service.
// For example, if an AWS service charges money to the account based on feature or
// service usage, then the AWS service team might create a default setting of
// "false". This means the user can't use this feature unless they change the
// setting to "true" and intentionally opt in for a paid feature. Services map a
// SettingId object to a setting value. AWS services teams define the default value
// for a SettingId. You can't create a new SettingId, but you can overwrite the
// default value if you have the ssm:UpdateServiceSetting permission for the
// setting. Use the GetServiceSetting () API action to view the current value. Use
// the UpdateServiceSetting () API action to change the default setting.  <p>Reset
// the service setting for the account to the default value as provisioned by the
// AWS  service team.
func (c *Client) ResetServiceSetting(ctx context.Context, params *ResetServiceSettingInput, optFns ...func(*Options)) (*ResetServiceSettingOutput, error) {
	stack := middleware.NewStack("ResetServiceSetting", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpResetServiceSettingMiddlewares(stack)
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
	addOpResetServiceSettingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetServiceSetting(options.Region), middleware.Before)
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
			OperationName: "ResetServiceSetting",
			Err:           err,
		}
	}
	out := result.(*ResetServiceSettingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body of the ResetServiceSetting API action.
type ResetServiceSettingInput struct {
	// The Amazon Resource Name (ARN) of the service setting to reset. The setting ID
	// can be /ssm/parameter-store/default-parameter-tier,
	// /ssm/parameter-store/high-throughput-enabled, or
	// /ssm/managed-instance/activation-tier. For example,
	// arn:aws:ssm:us-east-1:111122223333:servicesetting/ssm/parameter-store/high-throughput-enabled.
	SettingId *string
}

// The result body of the ResetServiceSetting API action.
type ResetServiceSettingOutput struct {
	// The current, effective service setting after calling the ResetServiceSetting API
	// action.
	ServiceSetting *types.ServiceSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpResetServiceSettingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpResetServiceSetting{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpResetServiceSetting{}, middleware.After)
}

func newServiceMetadataMiddleware_opResetServiceSetting(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ResetServiceSetting",
	}
}
