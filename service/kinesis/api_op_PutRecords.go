// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A PutRecords request.
type PutRecordsInput struct {
	_ struct{} `type:"structure"`

	// The records associated with the request.
	//
	// Records is a required field
	Records []PutRecordsRequestEntry `min:"1" type:"list" required:"true"`

	// The stream name associated with the request.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutRecordsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRecordsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRecordsInput"}

	if s.Records == nil {
		invalidParams.Add(aws.NewErrParamRequired("Records"))
	}
	if s.Records != nil && len(s.Records) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Records", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}
	if s.Records != nil {
		for i, v := range s.Records {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Records", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// PutRecords results.
type PutRecordsOutput struct {
	_ struct{} `type:"structure"`

	// The encryption type used on the records. This parameter can be one of the
	// following values:
	//
	//    * NONE: Do not encrypt the records.
	//
	//    * KMS: Use server-side encryption on the records using a customer-managed
	//    AWS KMS key.
	EncryptionType EncryptionType `type:"string" enum:"true"`

	// The number of unsuccessfully processed records in a PutRecords request.
	FailedRecordCount *int64 `min:"1" type:"integer"`

	// An array of successfully and unsuccessfully processed record results, correlated
	// with the request by natural ordering. A record that is successfully added
	// to a stream includes SequenceNumber and ShardId in the result. A record that
	// fails to be added to a stream includes ErrorCode and ErrorMessage in the
	// result.
	//
	// Records is a required field
	Records []PutRecordsResultEntry `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s PutRecordsOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutRecords = "PutRecords"

// PutRecordsRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Writes multiple data records into a Kinesis data stream in a single call
// (also referred to as a PutRecords request). Use this operation to send data
// into the stream for data ingestion and processing.
//
// Each PutRecords request can support up to 500 records. Each record in the
// request can be as large as 1 MB, up to a limit of 5 MB for the entire request,
// including partition keys. Each shard can support writes up to 1,000 records
// per second, up to a maximum data write total of 1 MB per second.
//
// You must specify the name of the stream that captures, stores, and transports
// the data; and an array of request Records, with each record in the array
// requiring a partition key and data blob. The record size limit applies to
// the total size of the partition key and data blob.
//
// The data blob can be any type of data; for example, a segment from a log
// file, geographic/location data, website clickstream data, and so on.
//
// The partition key is used by Kinesis Data Streams as input to a hash function
// that maps the partition key and associated data to a specific shard. An MD5
// hash function is used to map partition keys to 128-bit integer values and
// to map associated data records to shards. As a result of this hashing mechanism,
// all data records with the same partition key map to the same shard within
// the stream. For more information, see Adding Data to a Stream (http://docs.aws.amazon.com/kinesis/latest/dev/developing-producers-with-sdk.html#kinesis-using-sdk-java-add-data-to-stream)
// in the Amazon Kinesis Data Streams Developer Guide.
//
// Each record in the Records array may include an optional parameter, ExplicitHashKey,
// which overrides the partition key to shard mapping. This parameter allows
// a data producer to determine explicitly the shard where the record is stored.
// For more information, see Adding Multiple Records with PutRecords (http://docs.aws.amazon.com/kinesis/latest/dev/developing-producers-with-sdk.html#kinesis-using-sdk-java-putrecords)
// in the Amazon Kinesis Data Streams Developer Guide.
//
// The PutRecords response includes an array of response Records. Each record
// in the response array directly correlates with a record in the request array
// using natural ordering, from the top to the bottom of the request and response.
// The response Records array always includes the same number of records as
// the request array.
//
// The response Records array includes both successfully and unsuccessfully
// processed records. Kinesis Data Streams attempts to process all records in
// each PutRecords request. A single record failure does not stop the processing
// of subsequent records.
//
// A successfully processed record includes ShardId and SequenceNumber values.
// The ShardId parameter identifies the shard in the stream where the record
// is stored. The SequenceNumber parameter is an identifier assigned to the
// put record, unique to all records in the stream.
//
// An unsuccessfully processed record includes ErrorCode and ErrorMessage values.
// ErrorCode reflects the type of error and can be one of the following values:
// ProvisionedThroughputExceededException or InternalFailure. ErrorMessage provides
// more detailed information about the ProvisionedThroughputExceededException
// exception including the account ID, stream name, and shard ID of the record
// that was throttled. For more information about partially successful responses,
// see Adding Multiple Records with PutRecords (http://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-add-data-to-stream.html#kinesis-using-sdk-java-putrecords)
// in the Amazon Kinesis Data Streams Developer Guide.
//
// By default, data records are accessible for 24 hours from the time that they
// are added to a stream. You can use IncreaseStreamRetentionPeriod or DecreaseStreamRetentionPeriod
// to modify this retention period.
//
//    // Example sending a request using PutRecordsRequest.
//    req := client.PutRecordsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/PutRecords
func (c *Client) PutRecordsRequest(input *PutRecordsInput) PutRecordsRequest {
	op := &aws.Operation{
		Name:       opPutRecords,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutRecordsInput{}
	}

	req := c.newRequest(op, input, &PutRecordsOutput{})

	return PutRecordsRequest{Request: req, Input: input, Copy: c.PutRecordsRequest}
}

// PutRecordsRequest is the request type for the
// PutRecords API operation.
type PutRecordsRequest struct {
	*aws.Request
	Input *PutRecordsInput
	Copy  func(*PutRecordsInput) PutRecordsRequest
}

// Send marshals and sends the PutRecords API request.
func (r PutRecordsRequest) Send(ctx context.Context) (*PutRecordsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutRecordsResponse{
		PutRecordsOutput: r.Request.Data.(*PutRecordsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutRecordsResponse is the response type for the
// PutRecords API operation.
type PutRecordsResponse struct {
	*PutRecordsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutRecords request.
func (r *PutRecordsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
