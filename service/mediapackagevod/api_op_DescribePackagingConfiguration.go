// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a description of a MediaPackage VOD PackagingConfiguration resource.
func (c *Client) DescribePackagingConfiguration(ctx context.Context, params *DescribePackagingConfigurationInput, optFns ...func(*Options)) (*DescribePackagingConfigurationOutput, error) {
	stack := middleware.NewStack("DescribePackagingConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribePackagingConfigurationMiddlewares(stack)
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
	addOpDescribePackagingConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePackagingConfiguration(options.Region), middleware.Before)
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
			OperationName: "DescribePackagingConfiguration",
			Err:           err,
		}
	}
	out := result.(*DescribePackagingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePackagingConfigurationInput struct {
	// The ID of a MediaPackage VOD PackagingConfiguration resource.
	Id *string
}

type DescribePackagingConfigurationOutput struct {
	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage
	// The ID of the PackagingConfiguration.
	Id *string
	// A collection of tags associated with a resource
	Tags map[string]*string
	// A CMAF packaging configuration.
	CmafPackage *types.CmafPackage
	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *types.MssPackage
	// The ID of a PackagingGroup.
	PackagingGroupId *string
	// The ARN of the PackagingConfiguration.
	Arn *string
	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribePackagingConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribePackagingConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePackagingConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePackagingConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "DescribePackagingConfiguration",
	}
}
