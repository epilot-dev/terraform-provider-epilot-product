// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Region string

const (
	RegionDe Region = "DE"
	RegionAt Region = "AT"
	RegionCh Region = "CH"
)

func (e Region) ToPointer() *Region {
	return &e
}

func (e *Region) UnmarshalJSON(data []byte) error {
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
		*e = Region(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Region: %v", v)
	}
}

type TaxType string

const (
	TaxTypeVat    TaxType = "VAT"
	TaxTypeCustom TaxType = "Custom"
)

func (e TaxType) ToPointer() *TaxType {
	return &e
}

func (e *TaxType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VAT":
		fallthrough
	case "Custom":
		*e = TaxType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaxType: %v", v)
	}
}

type Tax struct {
	ID          string  `json:"_id"`
	Active      bool    `json:"active"`
	Description *string `json:"description,omitempty"`
	Rate        string  `json:"rate"`
	Region      Region  `json:"region"`
	Type        TaxType `json:"type"`
}

func (o *Tax) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Tax) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *Tax) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Tax) GetRate() string {
	if o == nil {
		return ""
	}
	return o.Rate
}

func (o *Tax) GetRegion() Region {
	if o == nil {
		return Region("")
	}
	return o.Region
}

func (o *Tax) GetType() TaxType {
	if o == nil {
		return TaxType("")
	}
	return o.Type
}
