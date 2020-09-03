// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the specified HSM client certificate. If no
// certificate ID is specified, returns information about all the HSM certificates
// owned by your AWS customer account. If you specify both tag keys and tag values
// in the same request, Amazon Redshift returns all HSM client certificates that
// match any combination of the specified keys and values. For example, if you have
// owner and environment for tag keys, and admin and test for tag values, all HSM
// client certificates that have any combination of those values are returned. If
// both tag keys and values are omitted from the request, HSM client certificates
// are returned regardless of whether they have tag keys or values associated with
// them.
func (c *Client) DescribeHsmClientCertificates(ctx context.Context, params *DescribeHsmClientCertificatesInput, optFns ...func(*Options)) (*DescribeHsmClientCertificatesOutput, error) {
	stack := middleware.NewStack("DescribeHsmClientCertificates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeHsmClientCertificatesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHsmClientCertificates(options.Region), middleware.Before)

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
			OperationName: "DescribeHsmClientCertificates",
			Err:           err,
		}
	}
	out := result.(*DescribeHsmClientCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeHsmClientCertificatesInput struct {
	// The identifier of a specific HSM client certificate for which you want
	// information. If no identifier is specified, information is returned for all HSM
	// client certificates owned by your AWS customer account.
	HsmClientCertificateIdentifier *string
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32
	// A tag value or values for which you want to return all matching HSM client
	// certificates that are associated with the specified tag value or values. For
	// example, suppose that you have HSM client certificates that are tagged with
	// values called admin and test. If you specify both of these tag values in the
	// request, Amazon Redshift returns a response with the HSM client certificates
	// that have either or both of these tag values associated with them.
	TagValues []*string
	// A tag key or keys for which you want to return all matching HSM client
	// certificates that are associated with the specified key or keys. For example,
	// suppose that you have HSM client certificates that are tagged with keys called
	// owner and environment. If you specify both of these tag keys in the request,
	// Amazon Redshift returns a response with the HSM client certificates that have
	// either or both of these tag keys associated with them.
	TagKeys []*string
	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeHsmClientCertificates () request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records by
	// providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string
}

//
type DescribeHsmClientCertificatesOutput struct {
	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string
	// A list of the identifiers for one or more HSM client certificates used by Amazon
	// Redshift clusters to store and retrieve database encryption keys in an HSM.
	HsmClientCertificates []*types.HsmClientCertificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeHsmClientCertificatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeHsmClientCertificates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeHsmClientCertificates{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeHsmClientCertificates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeHsmClientCertificates",
	}
}
