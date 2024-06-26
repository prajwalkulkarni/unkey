// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unkeyed/unkey/sdks/golang/internal/utils"
)

// PostV1KeysType - Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
//
// https://unkey.dev/docs/features/ratelimiting - Learn more
type PostV1KeysType string

const (
	PostV1KeysTypeFast       PostV1KeysType = "fast"
	PostV1KeysTypeConsistent PostV1KeysType = "consistent"
)

func (e PostV1KeysType) ToPointer() *PostV1KeysType {
	return &e
}

func (e *PostV1KeysType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fast":
		fallthrough
	case "consistent":
		*e = PostV1KeysType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostV1KeysType: %v", v)
	}
}

// PostV1KeysRatelimit - Unkey comes with per-key ratelimiting out of the box.
type PostV1KeysRatelimit struct {
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
	Type *PostV1KeysType `default:"fast" json:"type"`
	// The total amount of burstable requests.
	Limit int64 `json:"limit"`
	// How many tokens to refill during each refillInterval.
	RefillRate int64 `json:"refillRate"`
	// Determines the speed at which tokens are refilled, in milliseconds.
	RefillInterval int64 `json:"refillInterval"`
}

func (p PostV1KeysRatelimit) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostV1KeysRatelimit) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostV1KeysRatelimit) GetType() *PostV1KeysType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *PostV1KeysRatelimit) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *PostV1KeysRatelimit) GetRefillRate() int64 {
	if o == nil {
		return 0
	}
	return o.RefillRate
}

func (o *PostV1KeysRatelimit) GetRefillInterval() int64 {
	if o == nil {
		return 0
	}
	return o.RefillInterval
}

type PostV1KeysRequestBody struct {
	// Choose an `API` where this key should be created.
	APIID string `json:"apiId"`
	// To make it easier for your users to understand which product an api key belongs to, you can add prefix them.
	//
	// For example Stripe famously prefixes their customer ids with cus_ or their api keys with sk_live_.
	//
	// The underscore is automatically added if you are defining a prefix, for example: "prefix": "abc" will result in a key like abc_xxxxxxxxx
	//
	Prefix *string `json:"prefix,omitempty"`
	// The name for your Key. This is not customer facing.
	Name *string `json:"name,omitempty"`
	// The byte length used to generate your key determines its entropy as well as its length. Higher is better, but keys become longer and more annoying to handle. The default is 16 bytes, or 2^^128 possible combinations.
	ByteLength *int64 `default:"16" json:"byteLength"`
	// Your user’s Id. This will provide a link between Unkey and your customer record.
	// When validating a key, we will return this back to you, so you can clearly identify your user from their api key.
	OwnerID *string `json:"ownerId,omitempty"`
	// This is a place for dynamic meta data, anything that feels useful for you should go here
	Meta map[string]interface{} `json:"meta,omitempty"`
	// You can auto expire keys by providing a unix timestamp in milliseconds. Once Keys expire they will automatically be disabled and are no longer valid unless you enable them again.
	Expires *int64 `json:"expires,omitempty"`
	// You can limit the number of requests a key can make. Once a key reaches 0 remaining requests, it will automatically be disabled and is no longer valid unless you update it.
	Remaining *int64 `json:"remaining,omitempty"`
	// Unkey comes with per-key ratelimiting out of the box.
	Ratelimit *PostV1KeysRatelimit `json:"ratelimit,omitempty"`
}

func (p PostV1KeysRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostV1KeysRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostV1KeysRequestBody) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *PostV1KeysRequestBody) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *PostV1KeysRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PostV1KeysRequestBody) GetByteLength() *int64 {
	if o == nil {
		return nil
	}
	return o.ByteLength
}

func (o *PostV1KeysRequestBody) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *PostV1KeysRequestBody) GetMeta() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *PostV1KeysRequestBody) GetExpires() *int64 {
	if o == nil {
		return nil
	}
	return o.Expires
}

func (o *PostV1KeysRequestBody) GetRemaining() *int64 {
	if o == nil {
		return nil
	}
	return o.Remaining
}

func (o *PostV1KeysRequestBody) GetRatelimit() *PostV1KeysRatelimit {
	if o == nil {
		return nil
	}
	return o.Ratelimit
}

// PostV1KeysResponseBody - The configuration for an api
type PostV1KeysResponseBody struct {
	// The id of the key. This is not a secret and can be stored as a reference if you wish. You need the keyId to update or delete a key later.
	KeyID string `json:"keyId"`
	// The newly created api key, do not store this on your own system but pass it along to your user.
	Key string `json:"key"`
}

func (o *PostV1KeysResponseBody) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *PostV1KeysResponseBody) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}
