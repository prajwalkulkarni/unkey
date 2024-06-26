// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

type V1ApsisDeleteAPIRequestBody struct {
	// The id of the api to delete
	APIID string `json:"apiId"`
}

func (o *V1ApsisDeleteAPIRequestBody) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

// V1ApsisDeleteAPIResponseBody - The api was successfully deleted, it may take up to 30s for this to take effect in all regions
type V1ApsisDeleteAPIResponseBody struct {
}
