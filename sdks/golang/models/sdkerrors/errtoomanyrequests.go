// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// ErrTooManyRequestsCode - A machine readable error code.
type ErrTooManyRequestsCode string

const (
	ErrTooManyRequestsCodeTooManyRequests ErrTooManyRequestsCode = "TOO_MANY_REQUESTS"
)

func (e ErrTooManyRequestsCode) ToPointer() *ErrTooManyRequestsCode {
	return &e
}

func (e *ErrTooManyRequestsCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TOO_MANY_REQUESTS":
		*e = ErrTooManyRequestsCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ErrTooManyRequestsCode: %v", v)
	}
}

type ErrTooManyRequestsError struct {
	// A machine readable error code.
	Code ErrTooManyRequestsCode `json:"code"`
	// A link to our documentation with more details about this error code
	Docs string `json:"docs"`
	// A human readable explanation of what went wrong
	Message string `json:"message"`
	// Please always include the requestId in your error report
	RequestID string `json:"requestId"`
}

func (o *ErrTooManyRequestsError) GetCode() ErrTooManyRequestsCode {
	if o == nil {
		return ErrTooManyRequestsCode("")
	}
	return o.Code
}

func (o *ErrTooManyRequestsError) GetDocs() string {
	if o == nil {
		return ""
	}
	return o.Docs
}

func (o *ErrTooManyRequestsError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *ErrTooManyRequestsError) GetRequestID() string {
	if o == nil {
		return ""
	}
	return o.RequestID
}

// ErrTooManyRequests - The user has sent too many requests in a given amount of time ("rate limiting")
type ErrTooManyRequests struct {
	Error_ ErrTooManyRequestsError `json:"error"`
}

var _ error = &ErrTooManyRequests{}

func (e *ErrTooManyRequests) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
