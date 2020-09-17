// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves information about a dashboard.
func (c *Client) DescribeDashboard(ctx context.Context, params *DescribeDashboardInput, optFns ...func(*Options)) (*DescribeDashboardOutput, error) {
	stack := middleware.NewStack("DescribeDashboard", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeDashboardMiddlewares(stack)
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
	addOpDescribeDashboardValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDashboard(options.Region), middleware.Before)
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
			OperationName: "DescribeDashboard",
			Err:           err,
		}
	}
	out := result.(*DescribeDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDashboardInput struct {
	// The ID of the dashboard.
	DashboardId *string
}

type DescribeDashboardOutput struct {
	// The ID of the project that the dashboard is in.
	ProjectId *string
	// The date the dashboard was created, in Unix epoch time.
	DashboardCreationDate *time.Time
	// The ID of the dashboard.
	DashboardId *string
	// The dashboard's description.
	DashboardDescription *string
	// The date the dashboard was last updated, in Unix epoch time.
	DashboardLastUpdateDate *time.Time
	// The dashboard's definition JSON literal. For detailed information, see Creating
	// Dashboards (CLI)
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-dashboards-using-aws-cli.html)
	// in the AWS IoT SiteWise User Guide.
	DashboardDefinition *string
	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the dashboard, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:dashboard/${DashboardId}
	DashboardArn *string
	// The name of the dashboard.
	DashboardName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeDashboardMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDashboard{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDashboard{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDashboard(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeDashboard",
	}
}
