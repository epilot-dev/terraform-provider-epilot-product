// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TaxPatchRegion string

const (
	TaxPatchRegionDe TaxPatchRegion = "DE"
	TaxPatchRegionAt TaxPatchRegion = "AT"
	TaxPatchRegionCh TaxPatchRegion = "CH"
)

func (e TaxPatchRegion) ToPointer() *TaxPatchRegion {
	return &e
}
func (e *TaxPatchRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DE":
		fallthrough
	case "AT":
		fallthrough
	case "CH":
		*e = TaxPatchRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaxPatchRegion: %v", v)
	}
}

type TaxPatchType string

const (
	TaxPatchTypeVat    TaxPatchType = "VAT"
	TaxPatchTypeCustom TaxPatchType = "Custom"
)

func (e TaxPatchType) ToPointer() *TaxPatchType {
	return &e
}
func (e *TaxPatchType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VAT":
		fallthrough
	case "Custom":
		*e = TaxPatchType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaxPatchType: %v", v)
	}
}

type TaxPatch struct {
	Active      *bool           `json:"active,omitempty"`
	Description *string         `json:"description,omitempty"`
	Rate        *string         `json:"rate,omitempty"`
	Region      *TaxPatchRegion `json:"region,omitempty"`
	Type        *TaxPatchType   `json:"type,omitempty"`
}

func (o *TaxPatch) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *TaxPatch) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TaxPatch) GetRate() *string {
	if o == nil {
		return nil
	}
	return o.Rate
}

func (o *TaxPatch) GetRegion() *TaxPatchRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *TaxPatch) GetType() *TaxPatchType {
	if o == nil {
		return nil
	}
	return o.Type
}
