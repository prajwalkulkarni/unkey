// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
)

// Op - The operation you want to perform on the remaining count
type Op string

const (
	OpIncrement Op = "increment"
	OpDecrement Op = "decrement"
	OpSet       Op = "set"
)

func (e Op) ToPointer() *Op {
	return &e
}

func (e *Op) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "increment":
		fallthrough
	case "decrement":
		fallthrough
	case "set":
		*e = Op(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Op: %v", v)
	}
}

type V1KeysUpdateRemainingRequestBody struct {
	// The id of the key you want to modify
	KeyID string `json:"keyId"`
	// The operation you want to perform on the remaining count
	Op Op `json:"op"`
	// The value you want to set, add or subtract the remaining count by
	Value *int64 `json:"value"`
}

func (o *V1KeysUpdateRemainingRequestBody) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *V1KeysUpdateRemainingRequestBody) GetOp() Op {
	if o == nil {
		return Op("")
	}
	return o.Op
}

func (o *V1KeysUpdateRemainingRequestBody) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

// V1KeysUpdateRemainingResponseBody - The configuration for an api
type V1KeysUpdateRemainingResponseBody struct {
	// The number of remaining requests for this key after updating it. `null` means unlimited.
	Remaining *int64 `json:"remaining"`
}

func (o *V1KeysUpdateRemainingResponseBody) GetRemaining() *int64 {
	if o == nil {
		return nil
	}
	return o.Remaining
}
