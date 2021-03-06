// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateWebhookInput struct {
	_ struct{} `type:"structure"`

	// A regular expression used to determine which repository branches are built
	// when a webhook is triggered. If the name of a branch matches the regular
	// expression, then it is built. If branchFilter is empty, then all branches
	// are built.
	//
	// It is recommended that you use filterGroups instead of branchFilter.
	BranchFilter *string `locationName:"branchFilter" type:"string"`

	// An array of arrays of WebhookFilter objects used to determine which webhooks
	// are triggered. At least one WebhookFilter in the array must specify EVENT
	// as its type.
	//
	// For a build to be triggered, at least one filter group in the filterGroups
	// array must pass. For a filter group to pass, each of its filters must pass.
	FilterGroups [][]WebhookFilter `locationName:"filterGroups" type:"list"`

	// The name of the AWS CodeBuild project.
	//
	// ProjectName is a required field
	ProjectName *string `locationName:"projectName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateWebhookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateWebhookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateWebhookInput"}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateWebhookOutput struct {
	_ struct{} `type:"structure"`

	// Information about a webhook that connects repository events to a build project
	// in AWS CodeBuild.
	Webhook *Webhook `locationName:"webhook" type:"structure"`
}

// String returns the string representation
func (s CreateWebhookOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateWebhook = "CreateWebhook"

// CreateWebhookRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// For an existing AWS CodeBuild build project that has its source code stored
// in a GitHub or Bitbucket repository, enables AWS CodeBuild to start rebuilding
// the source code every time a code change is pushed to the repository.
//
// If you enable webhooks for an AWS CodeBuild project, and the project is used
// as a build step in AWS CodePipeline, then two identical builds are created
// for each commit. One build is triggered through webhooks, and one through
// AWS CodePipeline. Because billing is on a per-build basis, you are billed
// for both builds. Therefore, if you are using AWS CodePipeline, we recommend
// that you disable webhooks in AWS CodeBuild. In the AWS CodeBuild console,
// clear the Webhook box. For more information, see step 5 in Change a Build
// Project's Settings (https://docs.aws.amazon.com/codebuild/latest/userguide/change-project.html#change-project-console).
//
//    // Example sending a request using CreateWebhookRequest.
//    req := client.CreateWebhookRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/CreateWebhook
func (c *Client) CreateWebhookRequest(input *CreateWebhookInput) CreateWebhookRequest {
	op := &aws.Operation{
		Name:       opCreateWebhook,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateWebhookInput{}
	}

	req := c.newRequest(op, input, &CreateWebhookOutput{})

	return CreateWebhookRequest{Request: req, Input: input, Copy: c.CreateWebhookRequest}
}

// CreateWebhookRequest is the request type for the
// CreateWebhook API operation.
type CreateWebhookRequest struct {
	*aws.Request
	Input *CreateWebhookInput
	Copy  func(*CreateWebhookInput) CreateWebhookRequest
}

// Send marshals and sends the CreateWebhook API request.
func (r CreateWebhookRequest) Send(ctx context.Context) (*CreateWebhookResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateWebhookResponse{
		CreateWebhookOutput: r.Request.Data.(*CreateWebhookOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateWebhookResponse is the response type for the
// CreateWebhook API operation.
type CreateWebhookResponse struct {
	*CreateWebhookOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateWebhook request.
func (r *CreateWebhookResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
