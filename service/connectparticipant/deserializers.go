// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/connectparticipant/types"
	smithy "github.com/awslabs/smithy-go"
	smithyio "github.com/awslabs/smithy-go/io"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"strings"
)

type awsRestjson1_deserializeOpCreateParticipantConnection struct {
}

func (*awsRestjson1_deserializeOpCreateParticipantConnection) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpCreateParticipantConnection) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorCreateParticipantConnection(response)
	}
	output := &CreateParticipantConnectionOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentCreateParticipantConnectionOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorCreateParticipantConnection(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("AccessDeniedException", errorCode):
		return awsRestjson1_deserializeErrorAccessDeniedException(response, errorBody)

	case strings.EqualFold("InternalServerException", errorCode):
		return awsRestjson1_deserializeErrorInternalServerException(response, errorBody)

	case strings.EqualFold("ThrottlingException", errorCode):
		return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsRestjson1_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentCreateParticipantConnectionOutput(v **CreateParticipantConnectionOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *CreateParticipantConnectionOutput
	if *v == nil {
		sv = &CreateParticipantConnectionOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "ConnectionCredentials":
			if err := awsRestjson1_deserializeDocumentConnectionCredentials(&sv.ConnectionCredentials, value); err != nil {
				return err
			}

		case "Websocket":
			if err := awsRestjson1_deserializeDocumentWebsocket(&sv.Websocket, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpDisconnectParticipant struct {
}

func (*awsRestjson1_deserializeOpDisconnectParticipant) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpDisconnectParticipant) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorDisconnectParticipant(response)
	}
	output := &DisconnectParticipantOutput{}
	out.Result = output

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorDisconnectParticipant(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("AccessDeniedException", errorCode):
		return awsRestjson1_deserializeErrorAccessDeniedException(response, errorBody)

	case strings.EqualFold("InternalServerException", errorCode):
		return awsRestjson1_deserializeErrorInternalServerException(response, errorBody)

	case strings.EqualFold("ThrottlingException", errorCode):
		return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsRestjson1_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsRestjson1_deserializeOpGetTranscript struct {
}

func (*awsRestjson1_deserializeOpGetTranscript) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpGetTranscript) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorGetTranscript(response)
	}
	output := &GetTranscriptOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentGetTranscriptOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorGetTranscript(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("AccessDeniedException", errorCode):
		return awsRestjson1_deserializeErrorAccessDeniedException(response, errorBody)

	case strings.EqualFold("InternalServerException", errorCode):
		return awsRestjson1_deserializeErrorInternalServerException(response, errorBody)

	case strings.EqualFold("ThrottlingException", errorCode):
		return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsRestjson1_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentGetTranscriptOutput(v **GetTranscriptOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *GetTranscriptOutput
	if *v == nil {
		sv = &GetTranscriptOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "InitialContactId":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ContactId to be of type string, got %T instead", value)
				}
				sv.InitialContactId = &jtv
			}

		case "NextToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected NextToken to be of type string, got %T instead", value)
				}
				sv.NextToken = &jtv
			}

		case "Transcript":
			if err := awsRestjson1_deserializeDocumentTranscript(&sv.Transcript, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpSendEvent struct {
}

func (*awsRestjson1_deserializeOpSendEvent) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpSendEvent) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorSendEvent(response)
	}
	output := &SendEventOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentSendEventOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorSendEvent(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("AccessDeniedException", errorCode):
		return awsRestjson1_deserializeErrorAccessDeniedException(response, errorBody)

	case strings.EqualFold("InternalServerException", errorCode):
		return awsRestjson1_deserializeErrorInternalServerException(response, errorBody)

	case strings.EqualFold("ThrottlingException", errorCode):
		return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsRestjson1_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentSendEventOutput(v **SendEventOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *SendEventOutput
	if *v == nil {
		sv = &SendEventOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "AbsoluteTime":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Instant to be of type string, got %T instead", value)
				}
				sv.AbsoluteTime = &jtv
			}

		case "Id":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ChatItemId to be of type string, got %T instead", value)
				}
				sv.Id = &jtv
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpSendMessage struct {
}

func (*awsRestjson1_deserializeOpSendMessage) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpSendMessage) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorSendMessage(response)
	}
	output := &SendMessageOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentSendMessageOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorSendMessage(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("AccessDeniedException", errorCode):
		return awsRestjson1_deserializeErrorAccessDeniedException(response, errorBody)

	case strings.EqualFold("InternalServerException", errorCode):
		return awsRestjson1_deserializeErrorInternalServerException(response, errorBody)

	case strings.EqualFold("ThrottlingException", errorCode):
		return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsRestjson1_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentSendMessageOutput(v **SendMessageOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *SendMessageOutput
	if *v == nil {
		sv = &SendMessageOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "AbsoluteTime":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Instant to be of type string, got %T instead", value)
				}
				sv.AbsoluteTime = &jtv
			}

		case "Id":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ChatItemId to be of type string, got %T instead", value)
				}
				sv.Id = &jtv
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeErrorAccessDeniedException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.AccessDeniedException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentAccessDeniedException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorInternalServerException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InternalServerException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInternalServerException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorThrottlingException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ThrottlingException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentThrottlingException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorValidationException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ValidationException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentValidationException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeDocumentAccessDeniedException(v **types.AccessDeniedException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.AccessDeniedException
	if *v == nil {
		sv = &types.AccessDeniedException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = &jtv
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentConnectionCredentials(v **types.ConnectionCredentials, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ConnectionCredentials
	if *v == nil {
		sv = &types.ConnectionCredentials{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "ConnectionToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ParticipantToken to be of type string, got %T instead", value)
				}
				sv.ConnectionToken = &jtv
			}

		case "Expiry":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ISO8601Datetime to be of type string, got %T instead", value)
				}
				sv.Expiry = &jtv
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentInternalServerException(v **types.InternalServerException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InternalServerException
	if *v == nil {
		sv = &types.InternalServerException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = &jtv
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentItem(v **types.Item, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.Item
	if *v == nil {
		sv = &types.Item{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "AbsoluteTime":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Instant to be of type string, got %T instead", value)
				}
				sv.AbsoluteTime = &jtv
			}

		case "Content":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ChatContent to be of type string, got %T instead", value)
				}
				sv.Content = &jtv
			}

		case "ContentType":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ChatContentType to be of type string, got %T instead", value)
				}
				sv.ContentType = &jtv
			}

		case "DisplayName":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected DisplayName to be of type string, got %T instead", value)
				}
				sv.DisplayName = &jtv
			}

		case "Id":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ChatItemId to be of type string, got %T instead", value)
				}
				sv.Id = &jtv
			}

		case "ParticipantId":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ParticipantId to be of type string, got %T instead", value)
				}
				sv.ParticipantId = &jtv
			}

		case "ParticipantRole":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ParticipantRole to be of type string, got %T instead", value)
				}
				sv.ParticipantRole = types.ParticipantRole(jtv)
			}

		case "Type":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ChatItemType to be of type string, got %T instead", value)
				}
				sv.Type = types.ChatItemType(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentThrottlingException(v **types.ThrottlingException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ThrottlingException
	if *v == nil {
		sv = &types.ThrottlingException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Message to be of type string, got %T instead", value)
				}
				sv.Message = &jtv
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentTranscript(v *[]*types.Item, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []*types.Item
	if *v == nil {
		cv = []*types.Item{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col *types.Item
		if err := awsRestjson1_deserializeDocumentItem(&col, value); err != nil {
			return err
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentValidationException(v **types.ValidationException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ValidationException
	if *v == nil {
		sv = &types.ValidationException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Reason to be of type string, got %T instead", value)
				}
				sv.Message = &jtv
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentWebsocket(v **types.Websocket, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.Websocket
	if *v == nil {
		sv = &types.Websocket{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "ConnectionExpiry":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ISO8601Datetime to be of type string, got %T instead", value)
				}
				sv.ConnectionExpiry = &jtv
			}

		case "Url":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected PreSignedConnectionUrl to be of type string, got %T instead", value)
				}
				sv.Url = &jtv
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}
