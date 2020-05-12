// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datapipeline

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for ReportTaskProgress.
type ReportTaskProgressInput struct {
	_ struct{} `type:"structure"`

	// Key-value pairs that define the properties of the ReportTaskProgressInput
	// object.
	Fields []Field `locationName:"fields" type:"list"`

	// The ID of the task assigned to the task runner. This value is provided in
	// the response for PollForTask.
	//
	// TaskId is a required field
	TaskId *string `locationName:"taskId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ReportTaskProgressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReportTaskProgressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReportTaskProgressInput"}

	if s.TaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskId"))
	}
	if s.TaskId != nil && len(*s.TaskId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskId", 1))
	}
	if s.Fields != nil {
		for i, v := range s.Fields {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Fields", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of ReportTaskProgress.
type ReportTaskProgressOutput struct {
	_ struct{} `type:"structure"`

	// If true, the calling task runner should cancel processing of the task. The
	// task runner does not need to call SetTaskStatus for canceled tasks.
	//
	// Canceled is a required field
	Canceled *bool `locationName:"canceled" type:"boolean" required:"true"`
}

// String returns the string representation
func (s ReportTaskProgressOutput) String() string {
	return awsutil.Prettify(s)
}

const opReportTaskProgress = "ReportTaskProgress"

// ReportTaskProgressRequest returns a request value for making API operation for
// AWS Data Pipeline.
//
// Task runners call ReportTaskProgress when assigned a task to acknowledge
// that it has the task. If the web service does not receive this acknowledgement
// within 2 minutes, it assigns the task in a subsequent PollForTask call. After
// this initial acknowledgement, the task runner only needs to report progress
// every 15 minutes to maintain its ownership of the task. You can change this
// reporting time from 15 minutes by specifying a reportProgressTimeout field
// in your pipeline.
//
// If a task runner does not report its status after 5 minutes, AWS Data Pipeline
// assumes that the task runner is unable to process the task and reassigns
// the task in a subsequent response to PollForTask. Task runners should call
// ReportTaskProgress every 60 seconds.
//
//    // Example sending a request using ReportTaskProgressRequest.
//    req := client.ReportTaskProgressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ReportTaskProgress
func (c *Client) ReportTaskProgressRequest(input *ReportTaskProgressInput) ReportTaskProgressRequest {
	op := &aws.Operation{
		Name:       opReportTaskProgress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReportTaskProgressInput{}
	}

	req := c.newRequest(op, input, &ReportTaskProgressOutput{})

	return ReportTaskProgressRequest{Request: req, Input: input, Copy: c.ReportTaskProgressRequest}
}

// ReportTaskProgressRequest is the request type for the
// ReportTaskProgress API operation.
type ReportTaskProgressRequest struct {
	*aws.Request
	Input *ReportTaskProgressInput
	Copy  func(*ReportTaskProgressInput) ReportTaskProgressRequest
}

// Send marshals and sends the ReportTaskProgress API request.
func (r ReportTaskProgressRequest) Send(ctx context.Context) (*ReportTaskProgressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReportTaskProgressResponse{
		ReportTaskProgressOutput: r.Request.Data.(*ReportTaskProgressOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReportTaskProgressResponse is the response type for the
// ReportTaskProgress API operation.
type ReportTaskProgressResponse struct {
	*ReportTaskProgressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReportTaskProgress request.
func (r *ReportTaskProgressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}