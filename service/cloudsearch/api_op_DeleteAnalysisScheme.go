// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the DeleteAnalysisScheme operation. Specifies
// the name of the domain you want to update and the analysis scheme you want
// to delete.
type DeleteAnalysisSchemeInput struct {
	_ struct{} `type:"structure"`

	// The name of the analysis scheme you want to delete.
	//
	// AnalysisSchemeName is a required field
	AnalysisSchemeName *string `min:"1" type:"string" required:"true"`

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start
	// with a letter or number and can contain the following characters: a-z (lowercase),
	// 0-9, and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAnalysisSchemeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAnalysisSchemeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAnalysisSchemeInput"}

	if s.AnalysisSchemeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AnalysisSchemeName"))
	}
	if s.AnalysisSchemeName != nil && len(*s.AnalysisSchemeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AnalysisSchemeName", 1))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a DeleteAnalysisScheme request. Contains the status of the
// deleted analysis scheme.
type DeleteAnalysisSchemeOutput struct {
	_ struct{} `type:"structure"`

	// The status of the analysis scheme being deleted.
	//
	// AnalysisScheme is a required field
	AnalysisScheme *AnalysisSchemeStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteAnalysisSchemeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteAnalysisScheme = "DeleteAnalysisScheme"

// DeleteAnalysisSchemeRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Deletes an analysis scheme. For more information, see Configuring Analysis
// Schemes (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-analysis-schemes.html)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using DeleteAnalysisSchemeRequest.
//    req := client.DeleteAnalysisSchemeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteAnalysisSchemeRequest(input *DeleteAnalysisSchemeInput) DeleteAnalysisSchemeRequest {
	op := &aws.Operation{
		Name:       opDeleteAnalysisScheme,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAnalysisSchemeInput{}
	}

	req := c.newRequest(op, input, &DeleteAnalysisSchemeOutput{})

	return DeleteAnalysisSchemeRequest{Request: req, Input: input, Copy: c.DeleteAnalysisSchemeRequest}
}

// DeleteAnalysisSchemeRequest is the request type for the
// DeleteAnalysisScheme API operation.
type DeleteAnalysisSchemeRequest struct {
	*aws.Request
	Input *DeleteAnalysisSchemeInput
	Copy  func(*DeleteAnalysisSchemeInput) DeleteAnalysisSchemeRequest
}

// Send marshals and sends the DeleteAnalysisScheme API request.
func (r DeleteAnalysisSchemeRequest) Send(ctx context.Context) (*DeleteAnalysisSchemeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAnalysisSchemeResponse{
		DeleteAnalysisSchemeOutput: r.Request.Data.(*DeleteAnalysisSchemeOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAnalysisSchemeResponse is the response type for the
// DeleteAnalysisScheme API operation.
type DeleteAnalysisSchemeResponse struct {
	*DeleteAnalysisSchemeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAnalysisScheme request.
func (r *DeleteAnalysisSchemeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
