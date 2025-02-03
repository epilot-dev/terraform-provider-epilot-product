// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/shared"
	"net/http"
)

type DeleteCouponRequest struct {
	// The coupon id
	CouponID string `pathParam:"style=simple,explode=false,name=couponId"`
}

func (o *DeleteCouponRequest) GetCouponID() string {
	if o == nil {
		return ""
	}
	return o.CouponID
}

type DeleteCouponResponse struct {
	// Any error based on client data errors
	ClientError *shared.ClientError
	// HTTP response content type for this operation
	ContentType string
	// Coupon entity response
	Coupon *shared.Coupon
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Any error based on the server-side
	ServerError *shared.ServerError
}

func (o *DeleteCouponResponse) GetClientError() *shared.ClientError {
	if o == nil {
		return nil
	}
	return o.ClientError
}

func (o *DeleteCouponResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteCouponResponse) GetCoupon() *shared.Coupon {
	if o == nil {
		return nil
	}
	return o.Coupon
}

func (o *DeleteCouponResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteCouponResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteCouponResponse) GetServerError() *shared.ServerError {
	if o == nil {
		return nil
	}
	return o.ServerError
}
