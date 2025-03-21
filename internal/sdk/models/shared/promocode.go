// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PromoCode struct {
	// The code of the promo code
	Code string `json:"code"`
	// Whether the promo code has a usage limit
	HasUsageLimit *bool `json:"has_usage_limit,omitempty"`
	// The id of the promo code
	ID string `json:"id"`
	// The usage limit of the promo code
	UsageLimit *float64 `json:"usage_limit,omitempty"`
}

func (o *PromoCode) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *PromoCode) GetHasUsageLimit() *bool {
	if o == nil {
		return nil
	}
	return o.HasUsageLimit
}

func (o *PromoCode) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PromoCode) GetUsageLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.UsageLimit
}
