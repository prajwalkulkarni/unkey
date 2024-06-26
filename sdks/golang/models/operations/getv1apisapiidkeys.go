// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey/sdks/golang/internal/utils"
	"github.com/unkeyed/unkey/sdks/golang/models/components"
)

type GetV1ApisAPIIDKeysRequest struct {
	APIID   string   `pathParam:"style=simple,explode=false,name=apiId"`
	Limit   *int64   `default:"100" queryParam:"style=form,explode=true,name=limit"`
	Offset  *float64 `queryParam:"style=form,explode=true,name=offset"`
	OwnerID *string  `queryParam:"style=form,explode=true,name=ownerId"`
}

func (g GetV1ApisAPIIDKeysRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetV1ApisAPIIDKeysRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetV1ApisAPIIDKeysRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *GetV1ApisAPIIDKeysRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetV1ApisAPIIDKeysRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetV1ApisAPIIDKeysRequest) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

// GetV1ApisAPIIDKeysResponseBody - Keys belonging to the api
type GetV1ApisAPIIDKeysResponseBody struct {
	Keys []components.Key `json:"keys"`
	// The total number of keys for this api
	Total int64 `json:"total"`
}

func (o *GetV1ApisAPIIDKeysResponseBody) GetKeys() []components.Key {
	if o == nil {
		return []components.Key{}
	}
	return o.Keys
}

func (o *GetV1ApisAPIIDKeysResponseBody) GetTotal() int64 {
	if o == nil {
		return 0
	}
	return o.Total
}
