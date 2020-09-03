// Code generated by smithy-go-codegen DO NOT EDIT.

package ssooidc

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCreateToken struct {
}

func (*awsRestjson1_serializeOpCreateToken) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateTokenInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/token")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeDocumentCreateTokenInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeHttpBindingsCreateTokenInput(v *CreateTokenInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeDocumentCreateTokenInput(v *CreateTokenInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientId != nil {
		ok := object.Key("clientId")
		ok.String(*v.ClientId)
	}

	if v.ClientSecret != nil {
		ok := object.Key("clientSecret")
		ok.String(*v.ClientSecret)
	}

	if v.Code != nil {
		ok := object.Key("code")
		ok.String(*v.Code)
	}

	if v.DeviceCode != nil {
		ok := object.Key("deviceCode")
		ok.String(*v.DeviceCode)
	}

	if v.GrantType != nil {
		ok := object.Key("grantType")
		ok.String(*v.GrantType)
	}

	if v.RedirectUri != nil {
		ok := object.Key("redirectUri")
		ok.String(*v.RedirectUri)
	}

	if v.RefreshToken != nil {
		ok := object.Key("refreshToken")
		ok.String(*v.RefreshToken)
	}

	if v.Scope != nil {
		ok := object.Key("scope")
		if err := awsRestjson1_serializeDocumentScopes(v.Scope, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpRegisterClient struct {
}

func (*awsRestjson1_serializeOpRegisterClient) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpRegisterClient) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*RegisterClientInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/client/register")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeDocumentRegisterClientInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeHttpBindingsRegisterClientInput(v *RegisterClientInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeDocumentRegisterClientInput(v *RegisterClientInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientName != nil {
		ok := object.Key("clientName")
		ok.String(*v.ClientName)
	}

	if v.ClientType != nil {
		ok := object.Key("clientType")
		ok.String(*v.ClientType)
	}

	if v.Scopes != nil {
		ok := object.Key("scopes")
		if err := awsRestjson1_serializeDocumentScopes(v.Scopes, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpStartDeviceAuthorization struct {
}

func (*awsRestjson1_serializeOpStartDeviceAuthorization) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStartDeviceAuthorization) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartDeviceAuthorizationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/device_authorization")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeDocumentStartDeviceAuthorizationInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeHttpBindingsStartDeviceAuthorizationInput(v *StartDeviceAuthorizationInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeDocumentStartDeviceAuthorizationInput(v *StartDeviceAuthorizationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientId != nil {
		ok := object.Key("clientId")
		ok.String(*v.ClientId)
	}

	if v.ClientSecret != nil {
		ok := object.Key("clientSecret")
		ok.String(*v.ClientSecret)
	}

	if v.StartUrl != nil {
		ok := object.Key("startUrl")
		ok.String(*v.StartUrl)
	}

	return nil
}

func awsRestjson1_serializeDocumentScopes(v []*string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		av.String(*v[i])
	}
	return nil
}
