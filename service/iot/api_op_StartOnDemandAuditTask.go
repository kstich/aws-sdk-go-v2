// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an on-demand Device Defender audit.
func (c *Client) StartOnDemandAuditTask(ctx context.Context, params *StartOnDemandAuditTaskInput, optFns ...func(*Options)) (*StartOnDemandAuditTaskOutput, error) {
	stack := middleware.NewStack("StartOnDemandAuditTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartOnDemandAuditTaskMiddlewares(stack)
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
	addOpStartOnDemandAuditTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartOnDemandAuditTask(options.Region), middleware.Before)
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
			OperationName: "StartOnDemandAuditTask",
			Err:           err,
		}
	}
	out := result.(*StartOnDemandAuditTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartOnDemandAuditTaskInput struct {
	// Which checks are performed during the audit. The checks you specify must be
	// enabled for your account or an exception occurs. Use
	// DescribeAccountAuditConfiguration to see the list of all checks, including those
	// that are enabled or UpdateAccountAuditConfiguration to select which checks are
	// enabled.
	TargetCheckNames []*string
}

type StartOnDemandAuditTaskOutput struct {
	// The ID of the on-demand audit you started.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartOnDemandAuditTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartOnDemandAuditTask{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartOnDemandAuditTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartOnDemandAuditTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "StartOnDemandAuditTask",
	}
}
