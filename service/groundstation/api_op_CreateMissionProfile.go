// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a mission profile. dataflowEdges is a list of lists of strings. Each
// lower level list of strings has two elements: a from ARN and a to ARN.
func (c *Client) CreateMissionProfile(ctx context.Context, params *CreateMissionProfileInput, optFns ...func(*Options)) (*CreateMissionProfileOutput, error) {
	stack := middleware.NewStack("CreateMissionProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateMissionProfileMiddlewares(stack)
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
	addOpCreateMissionProfileValidationMiddleware(stack)
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
			OperationName: "CreateMissionProfile",
			Err:           err,
		}
	}
	out := result.(*CreateMissionProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateMissionProfileInput struct {
	// Amount of time after a contact ends that you’d like to receive a CloudWatch
	// event indicating the pass has finished.
	ContactPostPassDurationSeconds *int32
	// Amount of time prior to contact start you’d like to receive a CloudWatch event
	// indicating an upcoming pass.
	ContactPrePassDurationSeconds *int32
	// A list of lists of ARNs. Each list of ARNs is an edge, with a from Config and a
	// to Config.
	DataflowEdges [][]*string
	// Smallest amount of time in seconds that you’d like to see for an available
	// contact. AWS Ground Station will not present you with contacts shorter than this
	// duration.
	MinimumViableContactDurationSeconds *int32
	// Name of a mission profile.
	Name *string
	// Tags assigned to a mission profile.
	Tags map[string]*string
	// ARN of a tracking Config.
	TrackingConfigArn *string
}

//
type CreateMissionProfileOutput struct {
	// UUID of a mission profile.
	MissionProfileId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateMissionProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateMissionProfile{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMissionProfile{}, middleware.After)
}
