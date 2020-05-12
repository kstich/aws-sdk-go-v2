// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetTableInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the table resides. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the database in the catalog in which the table resides. For Hive
	// compatibility, this name is entirely lowercase.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// The name of the table for which to retrieve the definition. For Hive compatibility,
	// this name is entirely lowercase.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTableInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTableOutput struct {
	_ struct{} `type:"structure"`

	// The Table object that defines the specified table.
	Table *Table `type:"structure"`
}

// String returns the string representation
func (s GetTableOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTable = "GetTable"

// GetTableRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the Table definition in a Data Catalog for a specified table.
//
//    // Example sending a request using GetTableRequest.
//    req := client.GetTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTable
func (c *Client) GetTableRequest(input *GetTableInput) GetTableRequest {
	op := &aws.Operation{
		Name:       opGetTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTableInput{}
	}

	req := c.newRequest(op, input, &GetTableOutput{})

	return GetTableRequest{Request: req, Input: input, Copy: c.GetTableRequest}
}

// GetTableRequest is the request type for the
// GetTable API operation.
type GetTableRequest struct {
	*aws.Request
	Input *GetTableInput
	Copy  func(*GetTableInput) GetTableRequest
}

// Send marshals and sends the GetTable API request.
func (r GetTableRequest) Send(ctx context.Context) (*GetTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTableResponse{
		GetTableOutput: r.Request.Data.(*GetTableOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTableResponse is the response type for the
// GetTable API operation.
type GetTableResponse struct {
	*GetTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTable request.
func (r *GetTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}