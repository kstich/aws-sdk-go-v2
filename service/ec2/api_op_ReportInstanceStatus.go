// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Submits feedback about the status of an instance. The instance must be in the
// running state. If your experience with the instance differs from the instance
// status returned by DescribeInstanceStatus (), use ReportInstanceStatus () to
// report your experience with the instance. Amazon EC2 collects this information
// to improve the accuracy of status checks. Use of this action does not change the
// value returned by DescribeInstanceStatus ().
func (c *Client) ReportInstanceStatus(ctx context.Context, params *ReportInstanceStatusInput, optFns ...func(*Options)) (*ReportInstanceStatusOutput, error) {
	stack := middleware.NewStack("ReportInstanceStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpReportInstanceStatusMiddlewares(stack)
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
	addOpReportInstanceStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReportInstanceStatus(options.Region), middleware.Before)
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
			OperationName: "ReportInstanceStatus",
			Err:           err,
		}
	}
	out := result.(*ReportInstanceStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReportInstanceStatusInput struct {
	// The time at which the reported instance health state began.
	StartTime *time.Time
	// The time at which the reported instance health state ended.
	EndTime *time.Time
	// The instances.
	Instances []*string
	// The reason codes that describe the health state of your instance.
	//
	//     *
	// instance-stuck-in-state: My instance is stuck in a state.
	//
	//     * unresponsive:
	// My instance is unresponsive.
	//
	//     * not-accepting-credentials: My instance is
	// not accepting my credentials.
	//
	//     * password-not-available: A password is not
	// available for my instance.
	//
	//     * performance-network: My instance is
	// experiencing performance problems that I believe are network related.
	//
	//     *
	// performance-instance-store: My instance is experiencing performance problems
	// that I believe are related to the instance stores.
	//
	//     *
	// performance-ebs-volume: My instance is experiencing performance problems that I
	// believe are related to an EBS volume.
	//
	//     * performance-other: My instance is
	// experiencing performance problems.
	//
	//     * other: [explain using the description
	// parameter]
	ReasonCodes []types.ReportInstanceReasonCodes
	// The status of all instances listed.
	Status types.ReportStatusType
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// Descriptive text about the health state of your instance.
	Description *string
}

type ReportInstanceStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpReportInstanceStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpReportInstanceStatus{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpReportInstanceStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opReportInstanceStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReportInstanceStatus",
	}
}
