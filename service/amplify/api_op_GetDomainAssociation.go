// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the domain information for an Amplify app.
func (c *Client) GetDomainAssociation(ctx context.Context, params *GetDomainAssociationInput, optFns ...func(*Options)) (*GetDomainAssociationOutput, error) {
	stack := middleware.NewStack("GetDomainAssociation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDomainAssociationMiddlewares(stack)
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
	addOpGetDomainAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainAssociation(options.Region), middleware.Before)

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
			OperationName: "GetDomainAssociation",
			Err:           err,
		}
	}
	out := result.(*GetDomainAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the get domain association request.
type GetDomainAssociationInput struct {
	// The name of the domain.
	DomainName *string
	// The unique id for an Amplify app.
	AppId *string
}

// The result structure for the get domain association request.
type GetDomainAssociationOutput struct {
	// Describes the structure of a domain association, which associates a custom
	// domain with an Amplify app.
	DomainAssociation *types.DomainAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDomainAssociationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDomainAssociation{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomainAssociation{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDomainAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "GetDomainAssociation",
	}
}
