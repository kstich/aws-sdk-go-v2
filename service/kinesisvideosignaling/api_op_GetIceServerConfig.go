// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideosignaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the Interactive Connectivity Establishment (ICE) server configuration
// information, including URIs, username, and password which can be used to
// configure the WebRTC connection. The ICE component uses this configuration
// information to setup the WebRTC connection, including authenticating with the
// Traversal Using Relays around NAT (TURN) relay server. TURN is a protocol that
// is used to improve the connectivity of peer-to-peer applications. By providing a
// cloud-based relay service, TURN ensures that a connection can be established
// even when one or more peers are incapable of a direct peer-to-peer connection.
// For more information, see A REST API For Access To TURN Services
// (https://tools.ietf.org/html/draft-uberti-rtcweb-turn-rest-00). You can invoke
// this API to establish a fallback mechanism in case either of the peers is unable
// to establish a direct peer-to-peer connection over a signaling channel. You must
// specify either a signaling channel ARN or the client ID in order to invoke this
// API.
func (c *Client) GetIceServerConfig(ctx context.Context, params *GetIceServerConfigInput, optFns ...func(*Options)) (*GetIceServerConfigOutput, error) {
	stack := middleware.NewStack("GetIceServerConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetIceServerConfigMiddlewares(stack)
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
	addOpGetIceServerConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIceServerConfig(options.Region), middleware.Before)

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
			OperationName: "GetIceServerConfig",
			Err:           err,
		}
	}
	out := result.(*GetIceServerConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIceServerConfigInput struct {
	// Specifies the desired service. Currently, TURN is the only valid value.
	Service types.Service
	// Unique identifier for the viewer. Must be unique within the signaling channel.
	ClientId *string
	// An optional user ID to be associated with the credentials.
	Username *string
	// The ARN of the signaling channel to be used for the peer-to-peer connection
	// between configured peers.
	ChannelARN *string
}

type GetIceServerConfigOutput struct {
	// The list of ICE server information objects.
	IceServerList []*types.IceServer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetIceServerConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetIceServerConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIceServerConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetIceServerConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "GetIceServerConfig",
	}
}
