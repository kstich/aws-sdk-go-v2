// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Manually starts server maintenance. This command can be useful if an earlier
// maintenance attempt failed, and the underlying cause of maintenance failure has
// been resolved. The server is in an UNDER_MAINTENANCE state while maintenance is
// in progress. Maintenance can only be started on servers in HEALTHY and UNHEALTHY
// states. Otherwise, an InvalidStateException is thrown. A
// ResourceNotFoundException is thrown when the server does not exist. A
// ValidationException is raised when parameters of the request are not valid.
func (c *Client) StartMaintenance(ctx context.Context, params *StartMaintenanceInput, optFns ...func(*Options)) (*StartMaintenanceOutput, error) {
	stack := middleware.NewStack("StartMaintenance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartMaintenanceMiddlewares(stack)
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
	addOpStartMaintenanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartMaintenance(options.Region), middleware.Before)
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
			OperationName: "StartMaintenance",
			Err:           err,
		}
	}
	out := result.(*StartMaintenanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMaintenanceInput struct {
	// Engine attributes that are specific to the server on which you want to run
	// maintenance. Attributes accepted in a StartMaintenance request for Chef
	//
	//     *
	// CHEF_MAJOR_UPGRADE: If a Chef Automate server is eligible for upgrade to Chef
	// Automate 2, add this engine attribute to a StartMaintenance request and set the
	// value to true to upgrade the server to Chef Automate 2. For more information,
	// see Upgrade an AWS OpsWorks for Chef Automate Server to Chef Automate 2
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/opscm-a2upgrade.html).
	EngineAttributes []*types.EngineAttribute
	// The name of the server on which to run maintenance.
	ServerName *string
}

type StartMaintenanceOutput struct {
	// Contains the response to a StartMaintenance request.
	Server *types.Server

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartMaintenanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartMaintenance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMaintenance{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartMaintenance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "StartMaintenance",
	}
}
