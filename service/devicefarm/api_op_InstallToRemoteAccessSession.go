// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to install an Android application (in .apk format)
// or an iOS application (in .ipa format) as part of a remote access session.
type InstallToRemoteAccessSessionInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the app about which you are requesting information.
	//
	// AppArn is a required field
	AppArn *string `locationName:"appArn" min:"32" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the remote access session about which you
	// are requesting information.
	//
	// RemoteAccessSessionArn is a required field
	RemoteAccessSessionArn *string `locationName:"remoteAccessSessionArn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s InstallToRemoteAccessSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InstallToRemoteAccessSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InstallToRemoteAccessSessionInput"}

	if s.AppArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppArn"))
	}
	if s.AppArn != nil && len(*s.AppArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("AppArn", 32))
	}

	if s.RemoteAccessSessionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RemoteAccessSessionArn"))
	}
	if s.RemoteAccessSessionArn != nil && len(*s.RemoteAccessSessionArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("RemoteAccessSessionArn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server after AWS Device Farm makes a request
// to install to a remote access session.
type InstallToRemoteAccessSessionOutput struct {
	_ struct{} `type:"structure"`

	// An app to upload or that has been uploaded.
	AppUpload *Upload `locationName:"appUpload" type:"structure"`
}

// String returns the string representation
func (s InstallToRemoteAccessSessionOutput) String() string {
	return awsutil.Prettify(s)
}

const opInstallToRemoteAccessSession = "InstallToRemoteAccessSession"

// InstallToRemoteAccessSessionRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Installs an application to the device in a remote access session. For Android
// applications, the file must be in .apk format. For iOS applications, the
// file must be in .ipa format.
//
//    // Example sending a request using InstallToRemoteAccessSessionRequest.
//    req := client.InstallToRemoteAccessSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/InstallToRemoteAccessSession
func (c *Client) InstallToRemoteAccessSessionRequest(input *InstallToRemoteAccessSessionInput) InstallToRemoteAccessSessionRequest {
	op := &aws.Operation{
		Name:       opInstallToRemoteAccessSession,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InstallToRemoteAccessSessionInput{}
	}

	req := c.newRequest(op, input, &InstallToRemoteAccessSessionOutput{})

	return InstallToRemoteAccessSessionRequest{Request: req, Input: input, Copy: c.InstallToRemoteAccessSessionRequest}
}

// InstallToRemoteAccessSessionRequest is the request type for the
// InstallToRemoteAccessSession API operation.
type InstallToRemoteAccessSessionRequest struct {
	*aws.Request
	Input *InstallToRemoteAccessSessionInput
	Copy  func(*InstallToRemoteAccessSessionInput) InstallToRemoteAccessSessionRequest
}

// Send marshals and sends the InstallToRemoteAccessSession API request.
func (r InstallToRemoteAccessSessionRequest) Send(ctx context.Context) (*InstallToRemoteAccessSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &InstallToRemoteAccessSessionResponse{
		InstallToRemoteAccessSessionOutput: r.Request.Data.(*InstallToRemoteAccessSessionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// InstallToRemoteAccessSessionResponse is the response type for the
// InstallToRemoteAccessSession API operation.
type InstallToRemoteAccessSessionResponse struct {
	*InstallToRemoteAccessSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// InstallToRemoteAccessSession request.
func (r *InstallToRemoteAccessSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
