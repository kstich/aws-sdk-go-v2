// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new MSK cluster.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	stack := middleware.NewStack("CreateCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateClusterMiddlewares(stack)
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
	addOpCreateClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before)
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
			OperationName: "CreateCluster",
			Err:           err,
		}
	}
	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {
	// Includes all client authentication related information.
	ClientAuthentication *types.ClientAuthentication
	// The version of Apache Kafka.
	KafkaVersion *string
	// The name of the cluster.
	ClusterName *string
	// Specifies the level of monitoring for the MSK cluster. The possible values are
	// DEFAULT, PER_BROKER, and PER_TOPIC_PER_BROKER.
	EnhancedMonitoring types.EnhancedMonitoring
	// Create tags when creating the cluster.
	Tags map[string]*string
	// Represents the configuration that you want MSK to use for the brokers in a
	// cluster.
	ConfigurationInfo *types.ConfigurationInfo
	// The number of broker nodes in the cluster.
	NumberOfBrokerNodes *int32
	LoggingInfo         *types.LoggingInfo
	// Information about the broker nodes in the cluster.
	BrokerNodeGroupInfo *types.BrokerNodeGroupInfo
	// Includes all encryption-related information.
	EncryptionInfo *types.EncryptionInfo
	// The settings for open monitoring.
	OpenMonitoring *types.OpenMonitoringInfo
}

type CreateClusterOutput struct {
	// The name of the MSK cluster.
	ClusterName *string
	// The state of the cluster. The possible states are CREATING, ACTIVE, and FAILED.
	State types.ClusterState
	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateCluster{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "CreateCluster",
	}
}
