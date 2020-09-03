// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve information about an environment. An environment is a logical
// deployment group of AppConfig applications, such as applications in a Production
// environment or in an EU_Region environment. Each configuration deployment
// targets an environment. You can enable one or more Amazon CloudWatch alarms for
// an environment. If an alarm is triggered during a deployment, AppConfig roles
// back the configuration.
func (c *Client) GetEnvironment(ctx context.Context, params *GetEnvironmentInput, optFns ...func(*Options)) (*GetEnvironmentOutput, error) {
	stack := middleware.NewStack("GetEnvironment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetEnvironmentMiddlewares(stack)
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
	addOpGetEnvironmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnvironment(options.Region), middleware.Before)

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
			OperationName: "GetEnvironment",
			Err:           err,
		}
	}
	out := result.(*GetEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnvironmentInput struct {
	// The ID of the environment you wnat to get.
	EnvironmentId *string
	// The ID of the application that includes the environment you want to get.
	ApplicationId *string
}

type GetEnvironmentOutput struct {
	// The name of the environment.
	Name *string
	// The application ID.
	ApplicationId *string
	// The description of the environment.
	Description *string
	// The state of the environment. An environment can be in one of the following
	// states: READY_FOR_DEPLOYMENT, DEPLOYING, ROLLING_BACK, or ROLLED_BACK
	State types.EnvironmentState
	// The environment ID.
	Id *string
	// Amazon CloudWatch alarms monitored during the deployment.
	Monitors []*types.Monitor

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetEnvironmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetEnvironment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEnvironment{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetEnvironment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "GetEnvironment",
	}
}
