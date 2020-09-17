// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the specified provisioning artifact (also known as a
// version) for the specified product.
func (c *Client) DescribeProvisioningArtifact(ctx context.Context, params *DescribeProvisioningArtifactInput, optFns ...func(*Options)) (*DescribeProvisioningArtifactOutput, error) {
	stack := middleware.NewStack("DescribeProvisioningArtifact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeProvisioningArtifactMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProvisioningArtifact(options.Region), middleware.Before)
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
			OperationName: "DescribeProvisioningArtifact",
			Err:           err,
		}
	}
	out := result.(*DescribeProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProvisioningArtifactInput struct {
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string
	// The product identifier.
	ProductId *string
	// The product name.
	ProductName *string
	// Indicates whether a verbose level of detail is enabled.
	Verbose *bool
	// The provisioning artifact name.
	ProvisioningArtifactName *string
}

type DescribeProvisioningArtifactOutput struct {
	// The status of the current request.
	Status types.Status
	// The URL of the CloudFormation template in Amazon S3.
	Info map[string]*string
	// Information about the provisioning artifact.
	ProvisioningArtifactDetail *types.ProvisioningArtifactDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeProvisioningArtifactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProvisioningArtifact{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProvisioningArtifact{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeProvisioningArtifact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DescribeProvisioningArtifact",
	}
}
