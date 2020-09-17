// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of certificate authority (CA) certificates provided by Amazon
// DocumentDB for this AWS account.
func (c *Client) DescribeCertificates(ctx context.Context, params *DescribeCertificatesInput, optFns ...func(*Options)) (*DescribeCertificatesOutput, error) {
	stack := middleware.NewStack("DescribeCertificates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeCertificatesMiddlewares(stack)
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
	addOpDescribeCertificatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCertificates(options.Region), middleware.Before)
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
			OperationName: "DescribeCertificates",
			Err:           err,
		}
	}
	out := result.(*DescribeCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCertificatesInput struct {
	// An optional pagination token provided by a previous DescribeCertificates
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string
	// The user-supplied certificate identifier. If this parameter is specified,
	// information for only the specified certificate is returned. If this parameter is
	// omitted, a list of up to MaxRecords certificates is returned. This parameter is
	// not case sensitive. Constraints
	//
	//     * Must match an existing
	// CertificateIdentifier.
	CertificateIdentifier *string
	// This parameter is not currently supported.
	Filters []*types.Filter
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints:
	//
	//     * Minimum: 20
	//
	//     * Maximum: 100
	MaxRecords *int32
}

type DescribeCertificatesOutput struct {
	// A list of certificates for this AWS account.
	Certificates []*types.Certificate
	// An optional pagination token provided if the number of records retrieved is
	// greater than MaxRecords. If this parameter is specified, the marker specifies
	// the next record in the list. Including the value of Marker in the next call to
	// DescribeCertificates results in the next page of certificates.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeCertificatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCertificates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCertificates{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCertificates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeCertificates",
	}
}
