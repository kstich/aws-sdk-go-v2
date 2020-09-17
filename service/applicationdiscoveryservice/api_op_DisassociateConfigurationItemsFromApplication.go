// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates one or more configuration items from an application.
func (c *Client) DisassociateConfigurationItemsFromApplication(ctx context.Context, params *DisassociateConfigurationItemsFromApplicationInput, optFns ...func(*Options)) (*DisassociateConfigurationItemsFromApplicationOutput, error) {
	stack := middleware.NewStack("DisassociateConfigurationItemsFromApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisassociateConfigurationItemsFromApplicationMiddlewares(stack)
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
	addOpDisassociateConfigurationItemsFromApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateConfigurationItemsFromApplication(options.Region), middleware.Before)
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
			OperationName: "DisassociateConfigurationItemsFromApplication",
			Err:           err,
		}
	}
	out := result.(*DisassociateConfigurationItemsFromApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateConfigurationItemsFromApplicationInput struct {
	// Configuration ID of an application from which each item is disassociated.
	ApplicationConfigurationId *string
	// Configuration ID of each item to be disassociated from an application.
	ConfigurationIds []*string
}

type DisassociateConfigurationItemsFromApplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisassociateConfigurationItemsFromApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateConfigurationItemsFromApplication{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateConfigurationItemsFromApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateConfigurationItemsFromApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DisassociateConfigurationItemsFromApplication",
	}
}
