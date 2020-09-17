// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches devices and lists the ones that meet a set of filter criteria.
func (c *Client) SearchDevices(ctx context.Context, params *SearchDevicesInput, optFns ...func(*Options)) (*SearchDevicesOutput, error) {
	stack := middleware.NewStack("SearchDevices", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSearchDevicesMiddlewares(stack)
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
	addOpSearchDevicesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchDevices(options.Region), middleware.Before)
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
			OperationName: "SearchDevices",
			Err:           err,
		}
	}
	out := result.(*SearchDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchDevicesInput struct {
	// The sort order to use in listing the specified set of devices. Supported sort
	// keys are DeviceName, DeviceStatus, RoomName, DeviceType, DeviceSerialNumber,
	// ConnectionStatus, NetworkProfileName, NetworkProfileArn, Feature, and
	// FailureCode.
	SortCriteria []*types.Sort
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32
	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string
	// The filters to use to list a specified set of devices. Supported filter keys are
	// DeviceName, DeviceStatus, DeviceStatusDetailCode, RoomName, DeviceType,
	// DeviceSerialNumber, UnassociatedOnly, ConnectionStatus (ONLINE and OFFLINE),
	// NetworkProfileName, NetworkProfileArn, Feature, and FailureCode.
	Filters []*types.Filter
}

type SearchDevicesOutput struct {
	// The devices that meet the specified set of filter criteria, in sort order.
	Devices []*types.DeviceData
	// The token returned to indicate that there is more data available.
	NextToken *string
	// The total number of devices returned.
	TotalCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSearchDevicesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSearchDevices{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchDevices{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchDevices(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchDevices",
	}
}
