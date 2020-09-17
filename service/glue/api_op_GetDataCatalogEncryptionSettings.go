// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the security configuration for a specified catalog.
func (c *Client) GetDataCatalogEncryptionSettings(ctx context.Context, params *GetDataCatalogEncryptionSettingsInput, optFns ...func(*Options)) (*GetDataCatalogEncryptionSettingsOutput, error) {
	stack := middleware.NewStack("GetDataCatalogEncryptionSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDataCatalogEncryptionSettingsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataCatalogEncryptionSettings(options.Region), middleware.Before)
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
			OperationName: "GetDataCatalogEncryptionSettings",
			Err:           err,
		}
	}
	out := result.(*GetDataCatalogEncryptionSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataCatalogEncryptionSettingsInput struct {
	// The ID of the Data Catalog to retrieve the security configuration for. If none
	// is provided, the AWS account ID is used by default.
	CatalogId *string
}

type GetDataCatalogEncryptionSettingsOutput struct {
	// The requested security configuration.
	DataCatalogEncryptionSettings *types.DataCatalogEncryptionSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDataCatalogEncryptionSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDataCatalogEncryptionSettings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDataCatalogEncryptionSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDataCatalogEncryptionSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetDataCatalogEncryptionSettings",
	}
}
