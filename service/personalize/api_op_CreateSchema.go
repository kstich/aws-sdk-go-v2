// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Amazon Personalize schema from the specified schema string. The
// schema you create must be in Avro JSON format. Amazon Personalize recognizes
// three schema variants. Each schema is associated with a dataset type and has a
// set of required field and keywords. You specify a schema when you call
// CreateDataset ().  <p class="title"> <b>Related APIs</b> </p> <ul> <li> <p>
// <a>ListSchemas</a> </p> </li> <li> <p> <a>DescribeSchema</a> </p> </li> <li> <p>
// <a>DeleteSchema</a> </p> </li> </ul>
func (c *Client) CreateSchema(ctx context.Context, params *CreateSchemaInput, optFns ...func(*Options)) (*CreateSchemaOutput, error) {
	stack := middleware.NewStack("CreateSchema", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateSchemaMiddlewares(stack)
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
	addOpCreateSchemaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSchema(options.Region), middleware.Before)

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
			OperationName: "CreateSchema",
			Err:           err,
		}
	}
	out := result.(*CreateSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSchemaInput struct {
	// A schema in Avro JSON format.
	Schema *string
	// The name for the schema.
	Name *string
}

type CreateSchemaOutput struct {
	// The Amazon Resource Name (ARN) of the created schema.
	SchemaArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateSchemaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSchema{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSchema{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSchema(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateSchema",
	}
}
