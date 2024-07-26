// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type ClientError struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
}

func (o *ClientError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *ClientError) GetStatus() int64 {
	if o == nil {
		return 0
	}
	return o.Status
}