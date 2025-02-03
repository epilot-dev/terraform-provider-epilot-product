// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/shared"
	"net/http"
)

type PatchCouponRequest struct {
	// Coupon to patch
	CouponPatch shared.CouponPatch `request:"mediaType=application/json"`
	// The coupon id
	CouponID string `pathParam:"style=simple,explode=false,name=couponId"`
}

func (o *PatchCouponRequest) GetCouponPatch() shared.CouponPatch {
	if o == nil {
		return shared.CouponPatch{}
	}
	return o.CouponPatch
}

func (o *PatchCouponRequest) GetCouponID() string {
	if o == nil {
		return ""
	}
	return o.CouponID
}

type PatchCouponResponse struct {
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

func (o *PatchCouponResponse) GetClientError() *shared.ClientError {
	if o == nil {
		return nil
	}
	return o.ClientError
}

func (o *PatchCouponResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCouponResponse) GetCoupon() *shared.Coupon {
	if o == nil {
		return nil
	}
	return o.Coupon
}

func (o *PatchCouponResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCouponResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchCouponResponse) GetServerError() *shared.ServerError {
	if o == nil {
		return nil
	}
	return o.ServerError
}
