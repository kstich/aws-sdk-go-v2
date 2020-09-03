// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data
// Analytics API V2 Documentation (). Provides a description of the application,
// including the application Amazon Resource Name (ARN), status, latest version,
// and input and output configuration.
type ApplicationDetail struct {
	// Describes the application input configuration. For more information, see
	// Configuring Application Input
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html).
	// </p>
	InputDescriptions []*InputDescription
	// Name of the application.
	ApplicationName *string
	// Time stamp when the application version was created.
	CreateTimestamp *time.Time
	// Provides the current application version.
	ApplicationVersionId *int64
	// Time stamp when the application was last updated.
	LastUpdateTimestamp *time.Time
	// Status of the application.
	ApplicationStatus ApplicationStatus
	// Returns the application code that you provided to perform data analysis on any
	// of the in-application streams in your application.
	ApplicationCode *string
	// Description of the application.
	ApplicationDescription *string
	// Describes reference data sources configured for the application.  For more
	// information, see <a
	// href="https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html">Configuring
	// Application Input</a>. </p>
	ReferenceDataSourceDescriptions []*ReferenceDataSourceDescription
	// ARN of the application.
	ApplicationARN *string
	// Describes the CloudWatch log streams that are configured to receive application
	// messages. For more information about using CloudWatch log streams with Amazon
	// Kinesis Analytics applications, see Working with Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/cloudwatch-logs.html).
	CloudWatchLoggingOptionDescriptions []*CloudWatchLoggingOptionDescription
	// Describes the application output configuration. For more information, see
	// Configuring Application Output
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html).
	// </p>
	OutputDescriptions []*OutputDescription
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data
// Analytics API V2 Documentation (). Provides application summary information,
// including the application Amazon Resource Name (ARN), name, and status.
type ApplicationSummary struct {
	// ARN of the application.
	ApplicationARN *string
	// Status of the application.
	ApplicationStatus ApplicationStatus
	// Name of the application.
	ApplicationName *string
}

// Describes updates to apply to an existing Amazon Kinesis Analytics application.
type ApplicationUpdate struct {
	// Describes application input configuration updates.
	InputUpdates []*InputUpdate
	// Describes application code updates.
	ApplicationCodeUpdate *string
	// Describes application CloudWatch logging option updates.
	CloudWatchLoggingOptionUpdates []*CloudWatchLoggingOptionUpdate
	// Describes application reference data source updates.
	ReferenceDataSourceUpdates []*ReferenceDataSourceUpdate
	// Describes application output configuration updates.
	OutputUpdates []*OutputUpdate
}

// Provides a description of CloudWatch logging options, including the log stream
// Amazon Resource Name (ARN) and the role ARN.
type CloudWatchLoggingOption struct {
	// ARN of the CloudWatch log to receive application messages.
	LogStreamARN *string
	// IAM ARN of the role to use to send application messages. Note: To write
	// application messages to CloudWatch, the IAM role that is used must have the
	// PutLogEvents policy action enabled.
	RoleARN *string
}

// Description of the CloudWatch logging option.
type CloudWatchLoggingOptionDescription struct {
	// ARN of the CloudWatch log to receive application messages.
	LogStreamARN *string
	// ID of the CloudWatch logging option description.
	CloudWatchLoggingOptionId *string
	// IAM ARN of the role to use to send application messages. Note: To write
	// application messages to CloudWatch, the IAM role used must have the PutLogEvents
	// policy action enabled.
	RoleARN *string
}

// Describes CloudWatch logging option updates.
type CloudWatchLoggingOptionUpdate struct {
	// IAM ARN of the role to use to send application messages. Note: To write
	// application messages to CloudWatch, the IAM role used must have the PutLogEvents
	// policy action enabled.
	RoleARNUpdate *string
	// ARN of the CloudWatch log to receive application messages.
	LogStreamARNUpdate *string
	// ID of the CloudWatch logging option to update
	CloudWatchLoggingOptionId *string
}

// Provides additional mapping information when the record format uses delimiters,
// such as CSV. For example, the following sample records use CSV format, where the
// records use the '\n' as the row delimiter and a comma (",") as the column
// delimiter:  <p> <code>"name1", "address1"</code> </p> <p> <code>"name2",
// "address2"</code> </p>
type CSVMappingParameters struct {
	// Column delimiter. For example, in a CSV format, a comma (",") is the typical
	// column delimiter.
	RecordColumnDelimiter *string
	// Row delimiter. For example, in a CSV format, '\n' is the typical row delimiter.
	RecordRowDelimiter *string
}

// Describes the data format when records are written to the destination. For more
// information, see Configuring Application Output
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html).
// </p>
type DestinationSchema struct {
	// Specifies the format of the records on the output stream.
	RecordFormatType RecordFormatType
}

// When you configure the application input, you specify the streaming source, the
// in-application stream name that is created, and the mapping between the two. For
// more information, see Configuring Application Input
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html).
type Input struct {
	// Describes the number of in-application streams to create. Data from your source
	// is routed to these in-application input streams. (see Configuring Application
	// Input
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html).
	InputParallelism *InputParallelism
	// If the streaming source is an Amazon Kinesis stream, identifies the stream's
	// Amazon Resource Name (ARN) and an IAM role that enables Amazon Kinesis Analytics
	// to access the stream on your behalf. Note: Either KinesisStreamsInput or
	// KinesisFirehoseInput is required.
	KinesisStreamsInput *KinesisStreamsInput
	// Name prefix to use when creating an in-application stream. Suppose that you
	// specify a prefix "MyInApplicationStream." Amazon Kinesis Analytics then creates
	// one or more (as per the InputParallelism count you specified) in-application
	// streams with names "MyInApplicationStream_001," "MyInApplicationStream_002," and
	// so on.
	NamePrefix *string
	// If the streaming source is an Amazon Kinesis Firehose delivery stream,
	// identifies the delivery stream's ARN and an IAM role that enables Amazon Kinesis
	// Analytics to access the stream on your behalf. Note: Either KinesisStreamsInput
	// or KinesisFirehoseInput is required.
	KinesisFirehoseInput *KinesisFirehoseInput
	// Describes the format of the data in the streaming source, and how each data
	// element maps to corresponding columns in the in-application stream that is being
	// created. Also used to describe the format of the reference data source.
	InputSchema *SourceSchema
	// The InputProcessingConfiguration
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputProcessingConfiguration.html)
	// for the input. An input processor transforms records as they are received from
	// the stream, before the application's SQL code executes. Currently, the only
	// input processing configuration available is InputLambdaProcessor
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputLambdaProcessor.html).
	InputProcessingConfiguration *InputProcessingConfiguration
}

// When you start your application, you provide this configuration, which
// identifies the input source and the point in the input source at which you want
// the application to start processing records.
type InputConfiguration struct {
	// Input source ID. You can get this ID by calling the DescribeApplication
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation.
	Id *string
	// Point at which you want the application to start processing records from the
	// streaming source.
	InputStartingPositionConfiguration *InputStartingPositionConfiguration
}

// Describes the application input configuration. For more information, see
// Configuring Application Input
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html).
type InputDescription struct {
	// If an Amazon Kinesis stream is configured as streaming source, provides Amazon
	// Kinesis stream's Amazon Resource Name (ARN) and an IAM role that enables Amazon
	// Kinesis Analytics to access the stream on your behalf.
	KinesisStreamsInputDescription *KinesisStreamsInputDescription
	// Point at which the application is configured to read from the input stream.
	InputStartingPositionConfiguration *InputStartingPositionConfiguration
	// If an Amazon Kinesis Firehose delivery stream is configured as a streaming
	// source, provides the delivery stream's ARN and an IAM role that enables Amazon
	// Kinesis Analytics to access the stream on your behalf.
	KinesisFirehoseInputDescription *KinesisFirehoseInputDescription
	// Returns the in-application stream names that are mapped to the stream source.
	InAppStreamNames []*string
	// In-application name prefix.
	NamePrefix *string
	// The description of the preprocessor that executes on records in this input
	// before the application's code is run.
	InputProcessingConfigurationDescription *InputProcessingConfigurationDescription
	// Describes the format of the data in the streaming source, and how each data
	// element maps to corresponding columns in the in-application stream that is being
	// created.
	InputSchema *SourceSchema
	// Input ID associated with the application input. This is the ID that Amazon
	// Kinesis Analytics assigns to each input configuration you add to your
	// application.
	InputId *string
	// Describes the configured parallelism (number of in-application streams mapped to
	// the streaming source).
	InputParallelism *InputParallelism
}

// An object that contains the Amazon Resource Name (ARN) of the AWS Lambda
// (https://docs.aws.amazon.com/lambda/) function that is used to preprocess
// records in the stream, and the ARN of the IAM role that is used to access the
// AWS Lambda function.
type InputLambdaProcessor struct {
	// The ARN of the AWS Lambda (https://docs.aws.amazon.com/lambda/) function that
	// operates on records in the stream. To specify an earlier version of the Lambda
	// function than the latest, include the Lambda function version in the Lambda
	// function ARN. For more information about Lambda ARNs, see Example ARNs: AWS
	// Lambda ()
	ResourceARN *string
	// The ARN of the IAM role that is used to access the AWS Lambda function.
	RoleARN *string
}

// An object that contains the Amazon Resource Name (ARN) of the AWS Lambda
// (https://docs.aws.amazon.com/lambda/) function that is used to preprocess
// records in the stream, and the ARN of the IAM role that is used to access the
// AWS Lambda expression.
type InputLambdaProcessorDescription struct {
	// The ARN of the AWS Lambda (https://docs.aws.amazon.com/lambda/) function that is
	// used to preprocess the records in the stream.
	ResourceARN *string
	// The ARN of the IAM role that is used to access the AWS Lambda function.
	RoleARN *string
}

// Represents an update to the InputLambdaProcessor
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputLambdaProcessor.html)
// that is used to preprocess the records in the stream.
type InputLambdaProcessorUpdate struct {
	// The ARN of the new IAM role that is used to access the AWS Lambda function.
	RoleARNUpdate *string
	// The Amazon Resource Name (ARN) of the new AWS Lambda
	// (https://docs.aws.amazon.com/lambda/) function that is used to preprocess the
	// records in the stream. To specify an earlier version of the Lambda function than
	// the latest, include the Lambda function version in the Lambda function ARN. For
	// more information about Lambda ARNs, see Example ARNs: AWS Lambda ()
	ResourceARNUpdate *string
}

// Describes the number of in-application streams to create for a given streaming
// source. For information about parallelism, see Configuring Application Input
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html).
type InputParallelism struct {
	// Number of in-application streams to create. For more information, see Limits
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html).
	Count *int32
}

// Provides updates to the parallelism count.
type InputParallelismUpdate struct {
	// Number of in-application streams to create for the specified streaming source.
	CountUpdate *int32
}

// Provides a description of a processor that is used to preprocess the records in
// the stream before being processed by your application code. Currently, the only
// input processor available is AWS Lambda (https://docs.aws.amazon.com/lambda/).
type InputProcessingConfiguration struct {
	// The InputLambdaProcessor
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputLambdaProcessor.html)
	// that is used to preprocess the records in the stream before being processed by
	// your application code.
	InputLambdaProcessor *InputLambdaProcessor
}

// Provides configuration information about an input processor. Currently, the only
// input processor available is AWS Lambda (https://docs.aws.amazon.com/lambda/).
type InputProcessingConfigurationDescription struct {
	// Provides configuration information about the associated
	// InputLambdaProcessorDescription
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputLambdaProcessorDescription.html).
	InputLambdaProcessorDescription *InputLambdaProcessorDescription
}

// Describes updates to an InputProcessingConfiguration
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputProcessingConfiguration.html).
type InputProcessingConfigurationUpdate struct {
	// Provides update information for an InputLambdaProcessor
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputLambdaProcessor.html).
	InputLambdaProcessorUpdate *InputLambdaProcessorUpdate
}

// Describes updates for the application's input schema.
type InputSchemaUpdate struct {
	// Specifies the encoding of the records in the streaming source. For example,
	// UTF-8.
	RecordEncodingUpdate *string
	// Specifies the format of the records on the streaming source.
	RecordFormatUpdate *RecordFormat
	// A list of RecordColumn objects. Each object describes the mapping of the
	// streaming source element to the corresponding column in the in-application
	// stream.
	RecordColumnUpdates []*RecordColumn
}

// Describes the point at which the application reads from the streaming source.
type InputStartingPositionConfiguration struct {
	// The starting position on the stream.
	//
	//     * NOW - Start reading just after the
	// most recent record in the stream, start at the request time stamp that the
	// customer issued.  </li> <li> <p> <code>TRIM_HORIZON</code> - Start reading at
	// the last untrimmed record in the stream, which is the oldest record available in
	// the stream. This option is not available for an Amazon Kinesis Firehose delivery
	// stream.</p> </li> <li> <p> <code>LAST_STOPPED_POINT</code> - Resume reading from
	// where the application last stopped reading.</p> </li> </ul>
	InputStartingPosition InputStartingPosition
}

// Describes updates to a specific input configuration (identified by the InputId
// of an application).
type InputUpdate struct {
	// Name prefix for in-application streams that Amazon Kinesis Analytics creates for
	// the specific streaming source.
	NamePrefixUpdate *string
	// Input ID of the application input to be updated.
	InputId *string
	// If an Amazon Kinesis stream is the streaming source to be updated, provides an
	// updated stream Amazon Resource Name (ARN) and IAM role ARN.
	KinesisStreamsInputUpdate *KinesisStreamsInputUpdate
	// Describes updates for an input processing configuration.
	InputProcessingConfigurationUpdate *InputProcessingConfigurationUpdate
	// Describes the data format on the streaming source, and how record elements on
	// the streaming source map to columns of the in-application stream that is
	// created.
	InputSchemaUpdate *InputSchemaUpdate
	// If an Amazon Kinesis Firehose delivery stream is the streaming source to be
	// updated, provides an updated stream ARN and IAM role ARN.
	KinesisFirehoseInputUpdate *KinesisFirehoseInputUpdate
	// Describes the parallelism updates (the number in-application streams Amazon
	// Kinesis Analytics creates for the specific streaming source).
	InputParallelismUpdate *InputParallelismUpdate
}

// Provides additional mapping information when JSON is the record format on the
// streaming source.
type JSONMappingParameters struct {
	// Path to the top-level parent that contains the records.
	RecordRowPath *string
}

// Identifies an Amazon Kinesis Firehose delivery stream as the streaming source.
// You provide the delivery stream's Amazon Resource Name (ARN) and an IAM role ARN
// that enables Amazon Kinesis Analytics to access the stream on your behalf.
type KinesisFirehoseInput struct {
	// ARN of the input delivery stream.
	ResourceARN *string
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the
	// stream on your behalf. You need to make sure that the role has the necessary
	// permissions to access the stream.
	RoleARN *string
}

// Describes the Amazon Kinesis Firehose delivery stream that is configured as the
// streaming source in the application input configuration.
type KinesisFirehoseInputDescription struct {
	// Amazon Resource Name (ARN) of the Amazon Kinesis Firehose delivery stream.
	ResourceARN *string
	// ARN of the IAM role that Amazon Kinesis Analytics assumes to access the stream.
	RoleARN *string
}

// When updating application input configuration, provides information about an
// Amazon Kinesis Firehose delivery stream as the streaming source.
type KinesisFirehoseInputUpdate struct {
	// Amazon Resource Name (ARN) of the input Amazon Kinesis Firehose delivery stream
	// to read.
	ResourceARNUpdate *string
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the
	// stream on your behalf. You need to grant the necessary permissions to this role.
	RoleARNUpdate *string
}

// When configuring application output, identifies an Amazon Kinesis Firehose
// delivery stream as the destination. You provide the stream Amazon Resource Name
// (ARN) and an IAM role that enables Amazon Kinesis Analytics to write to the
// stream on your behalf.
type KinesisFirehoseOutput struct {
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to write to the
	// destination stream on your behalf. You need to grant the necessary permissions
	// to this role.
	RoleARN *string
	// ARN of the destination Amazon Kinesis Firehose delivery stream to write to.
	ResourceARN *string
}

// For an application output, describes the Amazon Kinesis Firehose delivery stream
// configured as its destination.
type KinesisFirehoseOutputDescription struct {
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the
	// stream.
	RoleARN *string
	// Amazon Resource Name (ARN) of the Amazon Kinesis Firehose delivery stream.
	ResourceARN *string
}

// When updating an output configuration using the UpdateApplication
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_UpdateApplication.html)
// operation, provides information about an Amazon Kinesis Firehose delivery stream
// configured as the destination.
type KinesisFirehoseOutputUpdate struct {
	// Amazon Resource Name (ARN) of the Amazon Kinesis Firehose delivery stream to
	// write to.
	ResourceARNUpdate *string
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the
	// stream on your behalf. You need to grant the necessary permissions to this role.
	RoleARNUpdate *string
}

// Identifies an Amazon Kinesis stream as the streaming source. You provide the
// stream's Amazon Resource Name (ARN) and an IAM role ARN that enables Amazon
// Kinesis Analytics to access the stream on your behalf.
type KinesisStreamsInput struct {
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the
	// stream on your behalf. You need to grant the necessary permissions to this role.
	RoleARN *string
	// ARN of the input Amazon Kinesis stream to read.
	ResourceARN *string
}

// Describes the Amazon Kinesis stream that is configured as the streaming source
// in the application input configuration.
type KinesisStreamsInputDescription struct {
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the
	// stream.
	RoleARN *string
	// Amazon Resource Name (ARN) of the Amazon Kinesis stream.
	ResourceARN *string
}

// When updating application input configuration, provides information about an
// Amazon Kinesis stream as the streaming source.
type KinesisStreamsInputUpdate struct {
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the
	// stream on your behalf. You need to grant the necessary permissions to this role.
	RoleARNUpdate *string
	// Amazon Resource Name (ARN) of the input Amazon Kinesis stream to read.
	ResourceARNUpdate *string
}

// When configuring application output, identifies an Amazon Kinesis stream as the
// destination. You provide the stream Amazon Resource Name (ARN) and also an IAM
// role ARN that Amazon Kinesis Analytics can use to write to the stream on your
// behalf.
type KinesisStreamsOutput struct {
	// ARN of the destination Amazon Kinesis stream to write to.
	ResourceARN *string
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to write to the
	// destination stream on your behalf. You need to grant the necessary permissions
	// to this role.
	RoleARN *string
}

// For an application output, describes the Amazon Kinesis stream configured as its
// destination.
type KinesisStreamsOutputDescription struct {
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the
	// stream.
	RoleARN *string
	// Amazon Resource Name (ARN) of the Amazon Kinesis stream.
	ResourceARN *string
}

// When updating an output configuration using the UpdateApplication
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_UpdateApplication.html)
// operation, provides information about an Amazon Kinesis stream configured as the
// destination.
type KinesisStreamsOutputUpdate struct {
	// Amazon Resource Name (ARN) of the Amazon Kinesis stream where you want to write
	// the output.
	ResourceARNUpdate *string
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the
	// stream on your behalf. You need to grant the necessary permissions to this role.
	RoleARNUpdate *string
}

// When configuring application output, identifies an AWS Lambda function as the
// destination. You provide the function Amazon Resource Name (ARN) and also an IAM
// role ARN that Amazon Kinesis Analytics can use to write to the function on your
// behalf.
type LambdaOutput struct {
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to write to the
	// destination function on your behalf. You need to grant the necessary permissions
	// to this role.
	RoleARN *string
	// Amazon Resource Name (ARN) of the destination Lambda function to write to. To
	// specify an earlier version of the Lambda function than the latest, include the
	// Lambda function version in the Lambda function ARN. For more information about
	// Lambda ARNs, see Example ARNs: AWS Lambda ()
	ResourceARN *string
}

// For an application output, describes the AWS Lambda function configured as its
// destination.
type LambdaOutputDescription struct {
	// Amazon Resource Name (ARN) of the destination Lambda function.
	ResourceARN *string
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to write to the
	// destination function.
	RoleARN *string
}

// When updating an output configuration using the UpdateApplication
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_UpdateApplication.html)
// operation, provides information about an AWS Lambda function configured as the
// destination.
type LambdaOutputUpdate struct {
	// Amazon Resource Name (ARN) of the destination Lambda function. To specify an
	// earlier version of the Lambda function than the latest, include the Lambda
	// function version in the Lambda function ARN. For more information about Lambda
	// ARNs, see Example ARNs: AWS Lambda ()
	ResourceARNUpdate *string
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to write to the
	// destination function on your behalf. You need to grant the necessary permissions
	// to this role.
	RoleARNUpdate *string
}

// When configuring application input at the time of creating or updating an
// application, provides additional mapping information specific to the record
// format (such as JSON, CSV, or record fields delimited by some delimiter) on the
// streaming source.
type MappingParameters struct {
	// Provides additional mapping information when JSON is the record format on the
	// streaming source.
	JSONMappingParameters *JSONMappingParameters
	// Provides additional mapping information when the record format uses delimiters
	// (for example, CSV).
	CSVMappingParameters *CSVMappingParameters
}

// Describes application output configuration in which you identify an
// in-application stream and a destination where you want the in-application stream
// data to be written. The destination can be an Amazon Kinesis stream or an Amazon
// Kinesis Firehose delivery stream.  <p></p> <p>For limits on how many
// destinations an application can write and other limitations, see <a
// href="https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html">Limits</a>.
// </p>
type Output struct {
	// Name of the in-application stream.
	Name *string
	// Identifies an Amazon Kinesis Firehose delivery stream as the destination.
	KinesisFirehoseOutput *KinesisFirehoseOutput
	// Identifies an AWS Lambda function as the destination.
	LambdaOutput *LambdaOutput
	// Identifies an Amazon Kinesis stream as the destination.
	KinesisStreamsOutput *KinesisStreamsOutput
	// Describes the data format when records are written to the destination. For more
	// information, see Configuring Application Output
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html).
	DestinationSchema *DestinationSchema
}

// Describes the application output configuration, which includes the
// in-application stream name and the destination where the stream data is written.
// The destination can be an Amazon Kinesis stream or an Amazon Kinesis Firehose
// delivery stream.
type OutputDescription struct {
	// Describes the Amazon Kinesis Firehose delivery stream configured as the
	// destination where output is written.
	KinesisFirehoseOutputDescription *KinesisFirehoseOutputDescription
	// Name of the in-application stream configured as output.
	Name *string
	// Describes the AWS Lambda function configured as the destination where output is
	// written.
	LambdaOutputDescription *LambdaOutputDescription
	// Data format used for writing data to the destination.
	DestinationSchema *DestinationSchema
	// A unique identifier for the output configuration.
	OutputId *string
	// Describes Amazon Kinesis stream configured as the destination where output is
	// written.
	KinesisStreamsOutputDescription *KinesisStreamsOutputDescription
}

// Describes updates to the output configuration identified by the OutputId.
type OutputUpdate struct {
	// If you want to specify a different in-application stream for this output
	// configuration, use this field to specify the new in-application stream name.
	NameUpdate *string
	// Describes an Amazon Kinesis Firehose delivery stream as the destination for the
	// output.
	KinesisFirehoseOutputUpdate *KinesisFirehoseOutputUpdate
	// Describes the data format when records are written to the destination. For more
	// information, see Configuring Application Output
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html).
	DestinationSchemaUpdate *DestinationSchema
	// Describes an AWS Lambda function as the destination for the output.
	LambdaOutputUpdate *LambdaOutputUpdate
	// Describes an Amazon Kinesis stream as the destination for the output.
	KinesisStreamsOutputUpdate *KinesisStreamsOutputUpdate
	// Identifies the specific output configuration that you want to update.
	OutputId *string
}

// Describes the mapping of each data element in the streaming source to the
// corresponding column in the in-application stream. Also used to describe the
// format of the reference data source.
type RecordColumn struct {
	// Reference to the data element in the streaming input or the reference data
	// source. This element is required if the RecordFormatType
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_RecordFormat.html#analytics-Type-RecordFormat-RecordFormatTypel)
	// is JSON.
	Mapping *string
	// Name of the column created in the in-application input stream or reference
	// table.
	Name *string
	// Type of column created in the in-application input stream or reference table.
	SqlType *string
}

// Describes the record format and relevant mapping information that should be
// applied to schematize the records on the stream.
type RecordFormat struct {
	// When configuring application input at the time of creating or updating an
	// application, provides additional mapping information specific to the record
	// format (such as JSON, CSV, or record fields delimited by some delimiter) on the
	// streaming source.
	MappingParameters *MappingParameters
	// The type of record format.
	RecordFormatType RecordFormatType
}

// Describes the reference data source by providing the source information (S3
// bucket name and object key name), the resulting in-application table name that
// is created, and the necessary schema to map the data elements in the Amazon S3
// object to the in-application table.
type ReferenceDataSource struct {
	// Identifies the S3 bucket and object that contains the reference data. Also
	// identifies the IAM role Amazon Kinesis Analytics can assume to read this object
	// on your behalf.  An Amazon Kinesis Analytics application loads reference data
	// only once. If the data changes, you call the <code>UpdateApplication</code>
	// operation to trigger reloading of data into your application. </p>
	S3ReferenceDataSource *S3ReferenceDataSource
	// Name of the in-application table to create.
	TableName *string
	// Describes the format of the data in the streaming source, and how each data
	// element maps to corresponding columns created in the in-application stream.
	ReferenceSchema *SourceSchema
}

// Describes the reference data source configured for an application.
type ReferenceDataSourceDescription struct {
	// Describes the format of the data in the streaming source, and how each data
	// element maps to corresponding columns created in the in-application stream.
	ReferenceSchema *SourceSchema
	// The in-application table name created by the specific reference data source
	// configuration.
	TableName *string
	// ID of the reference data source. This is the ID that Amazon Kinesis Analytics
	// assigns when you add the reference data source to your application using the
	// AddApplicationReferenceDataSource
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_AddApplicationReferenceDataSource.html)
	// operation.
	ReferenceId *string
	// Provides the S3 bucket name, the object key name that contains the reference
	// data. It also provides the Amazon Resource Name (ARN) of the IAM role that
	// Amazon Kinesis Analytics can assume to read the Amazon S3 object and populate
	// the in-application reference table.
	S3ReferenceDataSourceDescription *S3ReferenceDataSourceDescription
}

// When you update a reference data source configuration for an application, this
// object provides all the updated values (such as the source bucket name and
// object key name), the in-application table name that is created, and updated
// mapping information that maps the data in the Amazon S3 object to the
// in-application reference table that is created.
type ReferenceDataSourceUpdate struct {
	// ID of the reference data source being updated. You can use the
	// DescribeApplication
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get this value.
	ReferenceId *string
	// In-application table name that is created by this update.
	TableNameUpdate *string
	// Describes the S3 bucket name, object key name, and IAM role that Amazon Kinesis
	// Analytics can assume to read the Amazon S3 object on your behalf and populate
	// the in-application reference table.
	S3ReferenceDataSourceUpdate *S3ReferenceDataSourceUpdate
	// Describes the format of the data in the streaming source, and how each data
	// element maps to corresponding columns created in the in-application stream.
	ReferenceSchemaUpdate *SourceSchema
}

// Provides a description of an Amazon S3 data source, including the Amazon
// Resource Name (ARN) of the S3 bucket, the ARN of the IAM role that is used to
// access the bucket, and the name of the Amazon S3 object that contains the data.
type S3Configuration struct {
	// ARN of the S3 bucket that contains the data.
	BucketARN *string
	// The name of the object that contains the data.
	FileKey *string
	// IAM ARN of the role used to access the data.
	RoleARN *string
}

// Identifies the S3 bucket and object that contains the reference data. Also
// identifies the IAM role Amazon Kinesis Analytics can assume to read this object
// on your behalf. An Amazon Kinesis Analytics application loads reference data
// only once. If the data changes, you call the UpdateApplication
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_UpdateApplication.html)
// operation to trigger reloading of data into your application.
type S3ReferenceDataSource struct {
	// Amazon Resource Name (ARN) of the S3 bucket.
	BucketARN *string
	// Object key name containing reference data.
	FileKey *string
	// ARN of the IAM role that the service can assume to read data on your behalf.
	// This role must have permission for the s3:GetObject action on the object and
	// trust policy that allows Amazon Kinesis Analytics service principal to assume
	// this role.
	ReferenceRoleARN *string
}

// Provides the bucket name and object key name that stores the reference data.
type S3ReferenceDataSourceDescription struct {
	// Amazon S3 object key name.
	FileKey *string
	// Amazon Resource Name (ARN) of the S3 bucket.
	BucketARN *string
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to read the Amazon
	// S3 object on your behalf to populate the in-application reference table.
	ReferenceRoleARN *string
}

// Describes the S3 bucket name, object key name, and IAM role that Amazon Kinesis
// Analytics can assume to read the Amazon S3 object on your behalf and populate
// the in-application reference table.
type S3ReferenceDataSourceUpdate struct {
	// Amazon Resource Name (ARN) of the S3 bucket.
	BucketARNUpdate *string
	// Object key name.
	FileKeyUpdate *string
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to read the Amazon
	// S3 object and populate the in-application.
	ReferenceRoleARNUpdate *string
}

// Describes the format of the data in the streaming source, and how each data
// element maps to corresponding columns created in the in-application stream.
type SourceSchema struct {
	// Specifies the encoding of the records in the streaming source. For example,
	// UTF-8.
	RecordEncoding *string
	// A list of RecordColumn objects.
	RecordColumns []*RecordColumn
	// Specifies the format of the records on the streaming source.
	RecordFormat *RecordFormat
}

// A key-value pair (the value is optional) that you can define and assign to AWS
// resources. If you specify a tag that already exists, the tag value is replaced
// with the value that you specify in the request. Note that the maximum number of
// application tags includes system tags. The maximum number of user-defined
// application tags is 50. For more information, see Using Tagging
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-tagging.html).
type Tag struct {
	// The key of the key-value tag.
	Key *string
	// The value of the key-value tag. The value is optional.
	Value *string
}
