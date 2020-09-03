// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The custom terminology applied to the input text by Amazon Translate for the
// translated text response. This is optional in the response and will only be
// present if you specified terminology input in the request. Currently, only one
// terminology can be applied per TranslateText request.
type AppliedTerminology struct {
	// The name of the custom terminology applied to the input text by Amazon Translate
	// for the translated text response.
	Name *string
	// The specific terms of the custom terminology applied to the input text by Amazon
	// Translate for the translated text response. A maximum of 250 terms will be
	// returned, and the specific terms applied will be the first 250 terms in the
	// source text.
	Terms []*Term
}

// The encryption key used to encrypt the custom terminologies used by Amazon
// Translate.
type EncryptionKey struct {
	// The Amazon Resource Name (ARN) of the encryption key being used to encrypt the
	// custom terminology.
	Id *string
	// The type of encryption key used by Amazon Translate to encrypt custom
	// terminologies.
	Type EncryptionKeyType
}

// The input configuration properties for requesting a batch translation job.
type InputDataConfig struct {
	// The URI of the AWS S3 folder that contains the input file. The folder must be in
	// the same Region as the API endpoint you are calling.
	S3Uri *string
	// The multipurpose internet mail extension (MIME) type of the input files. Valid
	// values are text/plain for plaintext files and text/html for HTML files.
	ContentType *string
}

// The number of documents successfully and unsuccessfully processed during a
// translation job.
type JobDetails struct {
	// The number of documents used as input in a translation job.
	InputDocumentsCount *int32
	// The number of documents that could not be processed during a translation job.
	DocumentsWithErrorsCount *int32
	// The number of documents successfully processed during a translation job.
	TranslatedDocumentsCount *int32
}

// The output configuration properties for a batch translation job.
type OutputDataConfig struct {
	// The URI of the S3 folder that contains a translation job's output file. The
	// folder must be in the same Region as the API endpoint that you are calling.
	S3Uri *string
}

// The term being translated by the custom terminology.
type Term struct {
	// The target text of the term being translated by the custom terminology.
	TargetText *string
	// The source text of the term being translated by the custom terminology.
	SourceText *string
}

// The data associated with the custom terminology.
type TerminologyData struct {
	// The data format of the custom terminology. Either CSV or TMX.
	Format TerminologyDataFormat
	// The file containing the custom terminology data. Your version of the AWS SDK
	// performs a Base64-encoding on this field before sending a request to the AWS
	// service. Users of the SDK should not perform Base64-encoding themselves.
	File []byte
}

// The location of the custom terminology data.
type TerminologyDataLocation struct {
	// The location of the custom terminology data.
	Location *string
	// The repository type for the custom terminology data.
	RepositoryType *string
}

// The properties of the custom terminology.
type TerminologyProperties struct {
	// The size of the file used when importing a custom terminology.
	SizeBytes *int32
	// The description of the custom terminology properties.
	Description *string
	// The name of the custom terminology.
	Name *string
	// The language code for the source text of the translation request for which the
	// custom terminology is being used.
	SourceLanguageCode *string
	// The Amazon Resource Name (ARN) of the custom terminology.
	Arn *string
	// The language codes for the target languages available with the custom
	// terminology file. All possible target languages are returned in array.
	TargetLanguageCodes []*string
	// The encryption key for the custom terminology.
	EncryptionKey *EncryptionKey
	// The number of terms included in the custom terminology.
	TermCount *int32
	// The time at which the custom terminology was created, based on the timestamp.
	CreatedAt *time.Time
	// The time at which the custom terminology was last update, based on the
	// timestamp.
	LastUpdatedAt *time.Time
}

// Provides information for filtering a list of translation jobs. For more
// information, see ListTextTranslationJobs ().
type TextTranslationJobFilter struct {
	// Filters the list of jobs by name.
	JobName *string
	// Filters the list of jobs based on the time that the job was submitted for
	// processing and returns only the jobs submitted after the specified time. Jobs
	// are returned in descending order, newest to oldest.
	SubmittedAfterTime *time.Time
	// Filters the list of jobs based on the time that the job was submitted for
	// processing and returns only the jobs submitted before the specified time. Jobs
	// are returned in ascending order, oldest to newest.
	SubmittedBeforeTime *time.Time
	// Filters the list of jobs based by job status.
	JobStatus JobStatus
}

// Provides information about a translation job.
type TextTranslationJobProperties struct {
	// The user-defined name of the translation job.
	JobName *string
	// A list containing the names of the terminologies applied to a translation job.
	// Only one terminology can be applied per StartTextTranslationJob () request at
	// this time.
	TerminologyNames []*string
	// The language code of the language of the target text. The language must be a
	// language supported by Amazon Translate.
	TargetLanguageCodes []*string
	// An explanation of any errors that may have occured during the translation job.
	Message *string
	// The status of the translation job.
	JobStatus JobStatus
	// The time at which the translation job ended.
	EndTime *time.Time
	// The time at which the translation job was submitted.
	SubmittedTime *time.Time
	// The output configuration properties that were specified when the job was
	// requested.
	OutputDataConfig *OutputDataConfig
	// The input configuration properties that were specified when the job was
	// requested.
	InputDataConfig *InputDataConfig
	// The number of documents successfully and unsuccessfully processed during the
	// translation job.
	JobDetails *JobDetails
	// The language code of the language of the source text. The language must be a
	// language supported by Amazon Translate.
	SourceLanguageCode *string
	// The ID of the translation job.
	JobId *string
	// The Amazon Resource Name (ARN) of an AWS Identity Access and Management (IAM)
	// role that granted Amazon Translate read access to the job's input data.
	DataAccessRoleArn *string
}
