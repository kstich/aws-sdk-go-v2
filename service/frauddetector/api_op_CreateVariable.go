// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a variable.
func (c *Client) CreateVariable(ctx context.Context, params *CreateVariableInput, optFns ...func(*Options)) (*CreateVariableOutput, error) {
	stack := middleware.NewStack("CreateVariable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateVariableMiddlewares(stack)
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
	addOpCreateVariableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVariable(options.Region), middleware.Before)
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
			OperationName: "CreateVariable",
			Err:           err,
		}
	}
	out := result.(*CreateVariableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVariableInput struct {
	// The default value for the variable when no value is received.
	DefaultValue *string
	// The source of the data.
	DataSource types.DataSource
	// The variable type. For more information see Variable types
	// (https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types).
	// Valid Values: AUTH_CODE | AVS | BILLING_ADDRESS_L1 | BILLING_ADDRESS_L2 |
	// BILLING_CITY | BILLING_COUNTRY | BILLING_NAME | BILLING_PHONE | BILLING_STATE |
	// BILLING_ZIP | CARD_BIN | CATEGORICAL | CURRENCY_CODE | EMAIL_ADDRESS |
	// FINGERPRINT | FRAUD_LABEL | FREE_FORM_TEXT | IP_ADDRESS | NUMERIC | ORDER_ID |
	// PAYMENT_TYPE | PHONE_NUMBER | PRICE | PRODUCT_CATEGORY | SHIPPING_ADDRESS_L1 |
	// SHIPPING_ADDRESS_L2 | SHIPPING_CITY | SHIPPING_COUNTRY | SHIPPING_NAME |
	// SHIPPING_PHONE | SHIPPING_STATE | SHIPPING_ZIP | USERAGENT | SHIPPING_ZIP |
	// USERAGENT
	VariableType *string
	// A collection of key and value pairs.
	Tags []*types.Tag
	// The data type.
	DataType types.DataType
	// The name of the variable.
	Name *string
	// The description.
	Description *string
}

type CreateVariableOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateVariableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateVariable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateVariable{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateVariable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "CreateVariable",
	}
}
