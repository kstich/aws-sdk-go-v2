// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeartifact

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetPackageVersionAssetInput struct {
	_ struct{} `type:"structure"`

	// The name of the requested asset.
	//
	// Asset is a required field
	Asset *string `location:"querystring" locationName:"asset" min:"1" type:"string" required:"true"`

	// The domain that contains the repository that contains the package version
	// with the requested asset.
	//
	// Domain is a required field
	Domain *string `location:"querystring" locationName:"domain" min:"2" type:"string" required:"true"`

	// The 12-digit account number of the AWS account that owns the domain. It does
	// not include dashes or spaces.
	DomainOwner *string `location:"querystring" locationName:"domain-owner" min:"12" type:"string"`

	// A format that specifies the type of the package version with the requested
	// asset file. The valid values are:
	//
	//    * npm
	//
	//    * pypi
	//
	//    * maven
	//
	// Format is a required field
	Format PackageFormat `location:"querystring" locationName:"format" type:"string" required:"true" enum:"true"`

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//    * The namespace of a Maven package is its groupId.
	//
	//    * The namespace of an npm package is its scope.
	//
	//    * A Python package does not contain a corresponding component, so Python
	//    packages do not have a namespace.
	Namespace *string `location:"querystring" locationName:"namespace" min:"1" type:"string"`

	// The name of the package that contains the requested asset.
	//
	// Package is a required field
	Package *string `location:"querystring" locationName:"package" min:"1" type:"string" required:"true"`

	// A string that contains the package version (for example, 3.5.2).
	//
	// PackageVersion is a required field
	PackageVersion *string `location:"querystring" locationName:"version" min:"1" type:"string" required:"true"`

	// The name of the package version revision that contains the requested asset.
	PackageVersionRevision *string `location:"querystring" locationName:"revision" min:"1" type:"string"`

	// The repository that contains the package version with the requested asset.
	//
	// Repository is a required field
	Repository *string `location:"querystring" locationName:"repository" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPackageVersionAssetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPackageVersionAssetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPackageVersionAssetInput"}

	if s.Asset == nil {
		invalidParams.Add(aws.NewErrParamRequired("Asset"))
	}
	if s.Asset != nil && len(*s.Asset) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Asset", 1))
	}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 2))
	}
	if s.DomainOwner != nil && len(*s.DomainOwner) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainOwner", 12))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}
	if s.Namespace != nil && len(*s.Namespace) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Namespace", 1))
	}

	if s.Package == nil {
		invalidParams.Add(aws.NewErrParamRequired("Package"))
	}
	if s.Package != nil && len(*s.Package) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Package", 1))
	}

	if s.PackageVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("PackageVersion"))
	}
	if s.PackageVersion != nil && len(*s.PackageVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PackageVersion", 1))
	}
	if s.PackageVersionRevision != nil && len(*s.PackageVersionRevision) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PackageVersionRevision", 1))
	}

	if s.Repository == nil {
		invalidParams.Add(aws.NewErrParamRequired("Repository"))
	}
	if s.Repository != nil && len(*s.Repository) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Repository", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPackageVersionAssetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Asset != nil {
		v := *s.Asset

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "asset", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Domain != nil {
		v := *s.Domain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainOwner != nil {
		v := *s.DomainOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain-owner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Package != nil {
		v := *s.Package

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "package", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PackageVersion != nil {
		v := *s.PackageVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PackageVersionRevision != nil {
		v := *s.PackageVersionRevision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "revision", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Repository != nil {
		v := *s.Repository

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "repository", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetPackageVersionAssetOutput struct {
	_ struct{} `type:"structure" payload:"Asset"`

	// The binary file, or asset, that is downloaded.
	Asset io.ReadCloser `locationName:"asset" type:"blob"`

	// The name of the asset that is downloaded.
	AssetName *string `location:"header" locationName:"X-AssetName" min:"1" type:"string"`

	// A string that contains the package version (for example, 3.5.2).
	PackageVersion *string `location:"header" locationName:"X-PackageVersion" min:"1" type:"string"`

	// The name of the package version revision that contains the downloaded asset.
	PackageVersionRevision *string `location:"header" locationName:"X-PackageVersionRevision" min:"1" type:"string"`
}

// String returns the string representation
func (s GetPackageVersionAssetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPackageVersionAssetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AssetName != nil {
		v := *s.AssetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-AssetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PackageVersion != nil {
		v := *s.PackageVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-PackageVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PackageVersionRevision != nil {
		v := *s.PackageVersionRevision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-PackageVersionRevision", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// Skipping Asset Output type's body not valid.
	return nil
}

const opGetPackageVersionAsset = "GetPackageVersionAsset"

// GetPackageVersionAssetRequest returns a request value for making API operation for
// CodeArtifact.
//
// Returns an asset (or file) that is in a package. For example, for a Maven
// package version, use GetPackageVersionAsset to download a JAR file, a POM
// file, or any other assets in the package version.
//
//    // Example sending a request using GetPackageVersionAssetRequest.
//    req := client.GetPackageVersionAssetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeartifact-2018-09-22/GetPackageVersionAsset
func (c *Client) GetPackageVersionAssetRequest(input *GetPackageVersionAssetInput) GetPackageVersionAssetRequest {
	op := &aws.Operation{
		Name:       opGetPackageVersionAsset,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/package/version/asset",
	}

	if input == nil {
		input = &GetPackageVersionAssetInput{}
	}

	req := c.newRequest(op, input, &GetPackageVersionAssetOutput{})

	return GetPackageVersionAssetRequest{Request: req, Input: input, Copy: c.GetPackageVersionAssetRequest}
}

// GetPackageVersionAssetRequest is the request type for the
// GetPackageVersionAsset API operation.
type GetPackageVersionAssetRequest struct {
	*aws.Request
	Input *GetPackageVersionAssetInput
	Copy  func(*GetPackageVersionAssetInput) GetPackageVersionAssetRequest
}

// Send marshals and sends the GetPackageVersionAsset API request.
func (r GetPackageVersionAssetRequest) Send(ctx context.Context) (*GetPackageVersionAssetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPackageVersionAssetResponse{
		GetPackageVersionAssetOutput: r.Request.Data.(*GetPackageVersionAssetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPackageVersionAssetResponse is the response type for the
// GetPackageVersionAsset API operation.
type GetPackageVersionAssetResponse struct {
	*GetPackageVersionAssetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPackageVersionAsset request.
func (r *GetPackageVersionAssetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
