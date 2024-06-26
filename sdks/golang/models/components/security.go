// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Security struct {
	BearerAuth *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *Security) GetBearerAuth() *string {
	if o == nil {
		return nil
	}
	return o.BearerAuth
}
