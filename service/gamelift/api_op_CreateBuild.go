// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateBuildInput
type CreateBuildInput struct {
	_ struct{} `type:"structure"`

	// Descriptive label that is associated with a build. Build names do not need
	// to be unique. You can use UpdateBuild to change this value later.
	Name *string `min:"1" type:"string"`

	// Operating system that the game server binaries are built to run on. This
	// value determines the type of fleet resources that you can use for this build.
	// If your game build contains multiple executables, they all must run on the
	// same operating system. If an operating system is not specified when creating
	// a build, Amazon GameLift uses the default value (WINDOWS_2012). This value
	// cannot be changed later.
	OperatingSystem OperatingSystem `type:"string" enum:"true"`

	// Information indicating where your game build files are stored. Use this parameter
	// only when creating a build with files stored in an Amazon S3 bucket that
	// you own. The storage location must specify an Amazon S3 bucket name and key,
	// as well as a the ARN for a role that you set up to allow Amazon GameLift
	// to access your Amazon S3 bucket. The S3 bucket must be in the same region
	// that you want to create a new build in.
	StorageLocation *S3Location `type:"structure"`

	// Version that is associated with a build or script. Version strings do not
	// need to be unique. You can use UpdateBuild to change this value later.
	Version *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateBuildInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBuildInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBuildInput"}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}
	if s.StorageLocation != nil {
		if err := s.StorageLocation.Validate(); err != nil {
			invalidParams.AddNested("StorageLocation", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateBuildOutput
type CreateBuildOutput struct {
	_ struct{} `type:"structure"`

	// The newly created build record, including a unique build ID and status.
	Build *Build `type:"structure"`

	// Amazon S3 location for your game build file, including bucket name and key.
	StorageLocation *S3Location `type:"structure"`

	// This element is returned only when the operation is called without a storage
	// location. It contains credentials to use when you are uploading a build file
	// to an Amazon S3 bucket that is owned by Amazon GameLift. Credentials have
	// a limited life span. To refresh these credentials, call RequestUploadCredentials.
	UploadCredentials *AwsCredentials `type:"structure"`
}

// String returns the string representation
func (s CreateBuildOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateBuild = "CreateBuild"

// CreateBuildRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Creates a new Amazon GameLift build record for your game server binary files
// and points to the location of your game server build files in an Amazon Simple
// Storage Service (Amazon S3) location.
//
// Game server binaries must be combined into a .zip file for use with Amazon
// GameLift.
//
// To create new builds quickly and easily, use the AWS CLI command upload-build
// (https://docs.aws.amazon.com/cli/latest/reference/gamelift/upload-build.html).
// This helper command uploads your build and creates a new build record in
// one step, and automatically handles the necessary permissions.
//
// The CreateBuild operation should be used only when you need to manually upload
// your build files, as in the following scenarios:
//
//    * Store a build file in an Amazon S3 bucket under your own AWS account.
//    To use this option, you must first give Amazon GameLift access to that
//    Amazon S3 bucket. To create a new build record using files in your Amazon
//    S3 bucket, call CreateBuild and specify a build name, operating system,
//    and the storage location of your game build.
//
//    * Upload a build file directly to Amazon GameLift's Amazon S3 account.
//    To use this option, you first call CreateBuild with a build name and operating
//    system. This action creates a new build record and returns an Amazon S3
//    storage location (bucket and key only) and temporary access credentials.
//    Use the credentials to manually upload your build file to the storage
//    location (see the Amazon S3 topic Uploading Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/UploadingObjects.html)).
//    You can upload files to a location only once.
//
// If successful, this operation creates a new build record with a unique build
// ID and places it in INITIALIZED status. You can use DescribeBuild to check
// the status of your build. A build must be in READY status before it can be
// used to create fleets.
//
// Learn more
//
// Uploading Your Game (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html)
//
// Create a Build with Files in Amazon S3 (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-cli-uploading.html#gamelift-build-cli-uploading-create-build)
//
// Related operations
//
//    * CreateBuild
//
//    * ListBuilds
//
//    * DescribeBuild
//
//    * UpdateBuild
//
//    * DeleteBuild
//
//    // Example sending a request using CreateBuildRequest.
//    req := client.CreateBuildRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateBuild
func (c *Client) CreateBuildRequest(input *CreateBuildInput) CreateBuildRequest {
	op := &aws.Operation{
		Name:       opCreateBuild,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateBuildInput{}
	}

	req := c.newRequest(op, input, &CreateBuildOutput{})
	return CreateBuildRequest{Request: req, Input: input, Copy: c.CreateBuildRequest}
}

// CreateBuildRequest is the request type for the
// CreateBuild API operation.
type CreateBuildRequest struct {
	*aws.Request
	Input *CreateBuildInput
	Copy  func(*CreateBuildInput) CreateBuildRequest
}

// Send marshals and sends the CreateBuild API request.
func (r CreateBuildRequest) Send(ctx context.Context) (*CreateBuildResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateBuildResponse{
		CreateBuildOutput: r.Request.Data.(*CreateBuildOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateBuildResponse is the response type for the
// CreateBuild API operation.
type CreateBuildResponse struct {
	*CreateBuildOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateBuild request.
func (r *CreateBuildResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}