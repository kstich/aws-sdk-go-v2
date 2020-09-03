// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the current CAs for a group.
func (c *Client) ListGroupCertificateAuthorities(ctx context.Context, params *ListGroupCertificateAuthoritiesInput, optFns ...func(*Options)) (*ListGroupCertificateAuthoritiesOutput, error) {
	stack := middleware.NewStack("ListGroupCertificateAuthorities", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListGroupCertificateAuthoritiesMiddlewares(stack)
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
	addOpListGroupCertificateAuthoritiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListGroupCertificateAuthorities(options.Region), middleware.Before)

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
			OperationName: "ListGroupCertificateAuthorities",
			Err:           err,
		}
	}
	out := result.(*ListGroupCertificateAuthoritiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupCertificateAuthoritiesInput struct {
	// The ID of the Greengrass group.
	GroupId *string
}

type ListGroupCertificateAuthoritiesOutput struct {
	// A list of certificate authorities associated with the group.
	GroupCertificateAuthorities []*types.GroupCertificateAuthorityProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListGroupCertificateAuthoritiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListGroupCertificateAuthorities{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListGroupCertificateAuthorities{}, middleware.After)
}

func newServiceMetadataMiddleware_opListGroupCertificateAuthorities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ListGroupCertificateAuthorities",
	}
}
