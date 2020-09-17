// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current status of the specified configuration recorder. If a
// configuration recorder is not specified, this action returns the status of all
// configuration recorders associated with the account. Currently, you can specify
// only one configuration recorder per region in your account.
func (c *Client) DescribeConfigurationRecorderStatus(ctx context.Context, params *DescribeConfigurationRecorderStatusInput, optFns ...func(*Options)) (*DescribeConfigurationRecorderStatusOutput, error) {
	stack := middleware.NewStack("DescribeConfigurationRecorderStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeConfigurationRecorderStatusMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationRecorderStatus(options.Region), middleware.Before)
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
			OperationName: "DescribeConfigurationRecorderStatus",
			Err:           err,
		}
	}
	out := result.(*DescribeConfigurationRecorderStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeConfigurationRecorderStatus () action.
type DescribeConfigurationRecorderStatusInput struct {
	// The name(s) of the configuration recorder. If the name is not specified, the
	// action returns the current status of all the configuration recorders associated
	// with the account.
	ConfigurationRecorderNames []*string
}

// The output for the DescribeConfigurationRecorderStatus () action, in JSON
// format.
type DescribeConfigurationRecorderStatusOutput struct {
	// A list that contains status of the specified recorders.
	ConfigurationRecordersStatus []*types.ConfigurationRecorderStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeConfigurationRecorderStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigurationRecorderStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigurationRecorderStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeConfigurationRecorderStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeConfigurationRecorderStatus",
	}
}
