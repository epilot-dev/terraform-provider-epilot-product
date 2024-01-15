// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/models/shared"
	"net/http"
)

type PatchPriceRequest struct {
	// Price to patch
	PricePatch shared.PricePatch `request:"mediaType=application/json"`
	// The price id
	PriceID string `pathParam:"style=simple,explode=false,name=priceId"`
}

func (o *PatchPriceRequest) GetPricePatch() shared.PricePatch {
	if o == nil {
		return shared.PricePatch{}
	}
	return o.PricePatch
}

func (o *PatchPriceRequest) GetPriceID() string {
	if o == nil {
		return ""
	}
	return o.PriceID
}

type PatchPriceResponse struct {
	// Any error based on client data errors
	ClientError *shared.ClientError
	// HTTP response content type for this operation
	ContentType string
	// Price entity with id
	Price *shared.Price
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Any error based on the server-side
	ServerError *shared.ServerError
}

func (o *PatchPriceResponse) GetClientError() *shared.ClientError {
	if o == nil {
		return nil
	}
	return o.ClientError
}

func (o *PatchPriceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchPriceResponse) GetPrice() *shared.Price {
	if o == nil {
		return nil
	}
	return o.Price
}

func (o *PatchPriceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchPriceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchPriceResponse) GetServerError() *shared.ServerError {
	if o == nil {
		return nil
	}
	return o.ServerError
}
