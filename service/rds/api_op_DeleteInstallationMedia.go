// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the installation medium for a DB engine that requires an on-premises
// customer provided license, such as Microsoft SQL Server.
func (c *Client) DeleteInstallationMedia(ctx context.Context, params *DeleteInstallationMediaInput, optFns ...func(*Options)) (*DeleteInstallationMediaOutput, error) {
	stack := middleware.NewStack("DeleteInstallationMedia", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteInstallationMediaMiddlewares(stack)
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
	addOpDeleteInstallationMediaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInstallationMedia(options.Region), middleware.Before)
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
			OperationName: "DeleteInstallationMedia",
			Err:           err,
		}
	}
	out := result.(*DeleteInstallationMediaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteInstallationMediaInput struct {
	// The installation medium ID.
	InstallationMediaId *string
}

// Contains the installation media for a DB engine that requires an on-premises
// customer provided license, such as Microsoft SQL Server.
type DeleteInstallationMediaOutput struct {
	// The status of the installation medium.
	Status *string
	// The engine version of the DB engine.
	EngineVersion *string
	// The custom Availability Zone (AZ) that contains the installation media.
	CustomAvailabilityZoneId *string
	// The path to the installation medium for the operating system associated with the
	// DB engine.
	OSInstallationMediaPath *string
	// The DB engine.
	Engine *string
	// If an installation media failure occurred, the cause of the failure.
	FailureCause *types.InstallationMediaFailureCause
	// The installation medium ID.
	InstallationMediaId *string
	// The path to the installation medium for the DB engine.
	EngineInstallationMediaPath *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteInstallationMediaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteInstallationMedia{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteInstallationMedia{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteInstallationMedia(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteInstallationMedia",
	}
}
