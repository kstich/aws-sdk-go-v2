// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deprecated. Use DescribeLoa () instead. Gets the LOA-CFA for the specified
// interconnect. The Letter of Authorization - Connecting Facility Assignment
// (LOA-CFA) is a document that is used when establishing your cross connect to AWS
// at the colocation facility. For more information, see Requesting Cross Connects
// at AWS Direct Connect Locations
// (https://docs.aws.amazon.com/directconnect/latest/UserGuide/Colocation.html) in
// the AWS Direct Connect User Guide.
func (c *Client) DescribeInterconnectLoa(ctx context.Context, params *DescribeInterconnectLoaInput, optFns ...func(*Options)) (*DescribeInterconnectLoaOutput, error) {
	stack := middleware.NewStack("DescribeInterconnectLoa", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeInterconnectLoaMiddlewares(stack)
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
	addOpDescribeInterconnectLoaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInterconnectLoa(options.Region), middleware.Before)
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
			OperationName: "DescribeInterconnectLoa",
			Err:           err,
		}
	}
	out := result.(*DescribeInterconnectLoaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInterconnectLoaInput struct {
	// The ID of the interconnect.
	InterconnectId *string
	// The name of the service provider who establishes connectivity on your behalf. If
	// you supply this parameter, the LOA-CFA lists the provider name alongside your
	// company name as the requester of the cross connect.
	ProviderName *string
	// The standard media type for the LOA-CFA document. The only supported value is
	// application/pdf.
	LoaContentType types.LoaContentType
}

type DescribeInterconnectLoaOutput struct {
	// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA).
	Loa *types.Loa

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeInterconnectLoaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInterconnectLoa{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInterconnectLoa{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInterconnectLoa(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DescribeInterconnectLoa",
	}
}
