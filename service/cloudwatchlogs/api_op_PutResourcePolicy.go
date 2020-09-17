// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates a resource policy allowing other AWS services to put log
// events to this account, such as Amazon Route 53. An account can have up to 10
// resource policies per region.
func (c *Client) PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) {
	stack := middleware.NewStack("PutResourcePolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutResourcePolicyMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourcePolicy(options.Region), middleware.Before)
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
			OperationName: "PutResourcePolicy",
			Err:           err,
		}
	}
	out := result.(*PutResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourcePolicyInput struct {
	// Details of the new policy, including the identity of the principal that is
	// enabled to put logs to this account. This is formatted as a JSON string. This
	// parameter is required. The following example creates a resource policy enabling
	// the Route 53 service to put DNS query logs in to the specified log group.
	// Replace "logArn" with the ARN of your CloudWatch Logs resource, such as a log
	// group or log stream. { "Version": "2012-10-17", "Statement": [ { "Sid":
	// "Route53LogsToCloudWatchLogs", "Effect": "Allow", "Principal": { "Service": [
	// "route53.amazonaws.com" ] }, "Action":"logs:PutLogEvents", "Resource": "logArn"
	// } ] }
	//     </p>
	PolicyDocument *string
	// Name of the new policy. This parameter is required.
	PolicyName *string
}

type PutResourcePolicyOutput struct {
	// The new policy.
	ResourcePolicy *types.ResourcePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutResourcePolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutResourcePolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResourcePolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutResourcePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutResourcePolicy",
	}
}
