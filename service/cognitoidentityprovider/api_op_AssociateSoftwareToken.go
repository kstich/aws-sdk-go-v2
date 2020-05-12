// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateSoftwareTokenInput struct {
	_ struct{} `type:"structure"`

	// The access token.
	AccessToken *string `type:"string" sensitive:"true"`

	// The session which should be passed both ways in challenge-response calls
	// to the service. This allows authentication of the user as part of the MFA
	// setup process.
	Session *string `min:"20" type:"string"`
}

// String returns the string representation
func (s AssociateSoftwareTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateSoftwareTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateSoftwareTokenInput"}
	if s.Session != nil && len(*s.Session) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("Session", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateSoftwareTokenOutput struct {
	_ struct{} `type:"structure"`

	// A unique generated shared secret code that is used in the TOTP algorithm
	// to generate a one time code.
	SecretCode *string `min:"16" type:"string" sensitive:"true"`

	// The session which should be passed both ways in challenge-response calls
	// to the service. This allows authentication of the user as part of the MFA
	// setup process.
	Session *string `min:"20" type:"string"`
}

// String returns the string representation
func (s AssociateSoftwareTokenOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateSoftwareToken = "AssociateSoftwareToken"

// AssociateSoftwareTokenRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Returns a unique generated shared secret key code for the user account. The
// request takes an access token or a session string, but not both.
//
//    // Example sending a request using AssociateSoftwareTokenRequest.
//    req := client.AssociateSoftwareTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AssociateSoftwareToken
func (c *Client) AssociateSoftwareTokenRequest(input *AssociateSoftwareTokenInput) AssociateSoftwareTokenRequest {
	op := &aws.Operation{
		Name:       opAssociateSoftwareToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateSoftwareTokenInput{}
	}

	req := c.newRequest(op, input, &AssociateSoftwareTokenOutput{})

	return AssociateSoftwareTokenRequest{Request: req, Input: input, Copy: c.AssociateSoftwareTokenRequest}
}

// AssociateSoftwareTokenRequest is the request type for the
// AssociateSoftwareToken API operation.
type AssociateSoftwareTokenRequest struct {
	*aws.Request
	Input *AssociateSoftwareTokenInput
	Copy  func(*AssociateSoftwareTokenInput) AssociateSoftwareTokenRequest
}

// Send marshals and sends the AssociateSoftwareToken API request.
func (r AssociateSoftwareTokenRequest) Send(ctx context.Context) (*AssociateSoftwareTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateSoftwareTokenResponse{
		AssociateSoftwareTokenOutput: r.Request.Data.(*AssociateSoftwareTokenOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateSoftwareTokenResponse is the response type for the
// AssociateSoftwareToken API operation.
type AssociateSoftwareTokenResponse struct {
	*AssociateSoftwareTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateSoftwareToken request.
func (r *AssociateSoftwareTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}