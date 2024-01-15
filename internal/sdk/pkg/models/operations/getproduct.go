// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetProductRequest struct {
	// Hydrates entities in relations when passed true
	Hydrate *bool `queryParam:"style=form,explode=true,name=hydrate"`
	// The product id
	ProductID string `pathParam:"style=simple,explode=false,name=productId"`
}

func (o *GetProductRequest) GetHydrate() *bool {
	if o == nil {
		return nil
	}
	return o.Hydrate
}

func (o *GetProductRequest) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

type GetProductResponse struct {
	// Any error based on client data errors
	ClientError *shared.ClientError
	// HTTP response content type for this operation
	ContentType string
	// Product entity response
	Product *shared.Product
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Any error based on the server-side
	ServerError *shared.ServerError
}

func (o *GetProductResponse) GetClientError() *shared.ClientError {
	if o == nil {
		return nil
	}
	return o.ClientError
}

func (o *GetProductResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProductResponse) GetProduct() *shared.Product {
	if o == nil {
		return nil
	}
	return o.Product
}

func (o *GetProductResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProductResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetProductResponse) GetServerError() *shared.ServerError {
	if o == nil {
		return nil
	}
	return o.ServerError
}
