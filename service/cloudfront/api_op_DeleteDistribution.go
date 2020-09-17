// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete a distribution.
func (c *Client) DeleteDistribution(ctx context.Context, params *DeleteDistributionInput, optFns ...func(*Options)) (*DeleteDistributionOutput, error) {
	stack := middleware.NewStack("DeleteDistribution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteDistributionMiddlewares(stack)
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
	addOpDeleteDistributionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDistribution(options.Region), middleware.Before)
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
			OperationName: "DeleteDistribution",
			Err:           err,
		}
	}
	out := result.(*DeleteDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This action deletes a web distribution. To delete a web distribution using the
// CloudFront API, perform the following steps. To delete a web distribution using
// the CloudFront API:
//
//     * Disable the web distribution
//
//     * Submit a GET
// Distribution Config request to get the current configuration and the Etag header
// for the distribution.
//
//     * Update the XML document that was returned in the
// response to your GET Distribution Config request to change the value of Enabled
// to false.
//
//     * Submit a PUT Distribution Config request to update the
// configuration for your distribution. In the request body, include the XML
// document that you updated in Step 3. Set the value of the HTTP If-Match header
// to the value of the ETag header that CloudFront returned when you submitted the
// GET Distribution Config request in Step 2.
//
//     * Review the response to the PUT
// Distribution Config request to confirm that the distribution was successfully
// disabled.
//
//     * Submit a GET Distribution request to confirm that your changes
// have propagated. When propagation is complete, the value of Status is
// Deployed.
//
//     * Submit a DELETE Distribution request. Set the value of the HTTP
// If-Match header to the value of the ETag header that CloudFront returned when
// you submitted the GET Distribution Config request in Step 6.
//
//     * Review the
// response to your DELETE Distribution request to confirm that the distribution
// was successfully deleted.
//
// For information about deleting a distribution using
// the CloudFront console, see Deleting a Distribution
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/HowToDeleteDistribution.html)
// in the Amazon CloudFront Developer Guide.
type DeleteDistributionInput struct {
	// The value of the ETag header that you received when you disabled the
	// distribution. For example: E2QWRUHAPOMQZL.
	IfMatch *string
	// The distribution ID.
	Id *string
}

type DeleteDistributionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteDistributionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteDistribution{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteDistribution{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDistribution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteDistribution",
	}
}
