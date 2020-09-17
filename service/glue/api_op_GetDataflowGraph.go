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

// Transforms a Python script into a directed acyclic graph (DAG).
func (c *Client) GetDataflowGraph(ctx context.Context, params *GetDataflowGraphInput, optFns ...func(*Options)) (*GetDataflowGraphOutput, error) {
	stack := middleware.NewStack("GetDataflowGraph", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDataflowGraphMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataflowGraph(options.Region), middleware.Before)
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
			OperationName: "GetDataflowGraph",
			Err:           err,
		}
	}
	out := result.(*GetDataflowGraphOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataflowGraphInput struct {
	// The Python script to transform.
	PythonScript *string
}

type GetDataflowGraphOutput struct {
	// A list of the edges in the resulting DAG.
	DagEdges []*types.CodeGenEdge
	// A list of the nodes in the resulting DAG.
	DagNodes []*types.CodeGenNode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDataflowGraphMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDataflowGraph{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDataflowGraph{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDataflowGraph(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetDataflowGraph",
	}
}
