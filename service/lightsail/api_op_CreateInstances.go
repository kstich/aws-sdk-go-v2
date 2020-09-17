// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates one or more Amazon Lightsail instances. The create instances operation
// supports tag-based access control via request tags. For more information, see
// the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateInstances(ctx context.Context, params *CreateInstancesInput, optFns ...func(*Options)) (*CreateInstancesOutput, error) {
	stack := middleware.NewStack("CreateInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateInstancesMiddlewares(stack)
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
	addOpCreateInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstances(options.Region), middleware.Before)
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
			OperationName: "CreateInstances",
			Err:           err,
		}
	}
	out := result.(*CreateInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstancesInput struct {
	// The ID for a virtual private server image (e.g., app_wordpress_4_4 or
	// app_lamp_7_0). Use the get blueprints operation to return a list of available
	// images (or blueprints). Use active blueprints when creating new instances.
	// Inactive blueprints are listed to support customers with existing instances and
	// are not necessarily available to create new instances. Blueprints are marked
	// inactive when they become outdated due to operating system updates or new
	// application releases.
	BlueprintId *string
	// The name of your key pair.
	KeyPairName *string
	// The bundle of specification information for your virtual private server (or
	// instance), including the pricing plan (e.g., micro_1_0).
	BundleId *string
	// The Availability Zone in which to create your instance. Use the following
	// format: us-east-2a (case sensitive). You can get a list of Availability Zones by
	// using the get regions
	// (http://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetRegions.html)
	// operation. Be sure to add the include Availability Zones parameter to your
	// request.
	AvailabilityZone *string
	// The names to use for your new Lightsail instances. Separate multiple values
	// using quotation marks and commas, for example:
	// ["MyFirstInstance","MySecondInstance"]
	InstanceNames []*string
	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []*types.Tag
	// A launch script you can create that configures a server with additional user
	// data. For example, you might want to run apt-get -y update. Depending on the
	// machine image you choose, the command to get software on your instance varies.
	// Amazon Linux and CentOS use yum, Debian and Ubuntu use apt-get, and FreeBSD uses
	// pkg. For a complete list, see the Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/getting-started/article/compare-options-choose-lightsail-instance-image).
	UserData *string
	// An array of objects representing the add-ons to enable for the new instance.
	AddOns []*types.AddOnRequest
	// (Deprecated) The name for your custom image. In releases prior to June 12, 2017,
	// this parameter was ignored by the API. It is now deprecated.
	CustomImageName *string
}

type CreateInstancesOutput struct {
	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateInstances",
	}
}
