// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a traffic policy instance and all of the resource record sets that
// Amazon Route 53 created when you created the instance. In the Route 53 console,
// traffic policy instances are known as policy records.
func (c *Client) DeleteTrafficPolicyInstance(ctx context.Context, params *DeleteTrafficPolicyInstanceInput, optFns ...func(*Options)) (*DeleteTrafficPolicyInstanceOutput, error) {
	stack := middleware.NewStack("DeleteTrafficPolicyInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteTrafficPolicyInstanceMiddlewares(stack)
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
	addOpDeleteTrafficPolicyInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTrafficPolicyInstance(options.Region), middleware.Before)
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
			OperationName: "DeleteTrafficPolicyInstance",
			Err:           err,
		}
	}
	out := result.(*DeleteTrafficPolicyInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete a specified traffic policy instance.
type DeleteTrafficPolicyInstanceInput struct {
	// The ID of the traffic policy instance that you want to delete. When you delete a
	// traffic policy instance, Amazon Route 53 also deletes all of the resource record
	// sets that were created when you created the traffic policy instance.
	Id *string
}

// An empty element.
type DeleteTrafficPolicyInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteTrafficPolicyInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteTrafficPolicyInstance{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteTrafficPolicyInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteTrafficPolicyInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeleteTrafficPolicyInstance",
	}
}
