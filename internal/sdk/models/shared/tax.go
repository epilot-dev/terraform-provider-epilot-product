// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/internal/utils"
	"time"
)

type TaxSchema string

const (
	TaxSchemaTax TaxSchema = "tax"
)

func (e TaxSchema) ToPointer() *TaxSchema {
	return &e
}
func (e *TaxSchema) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tax":
		*e = TaxSchema(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaxSchema: %v", v)
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
	// Additional fields that are not part of the schema
	Additional map[string]any `json:"__additional,omitempty"`
	// Access control list (ACL) for an entity. Defines sharing access to external orgs or users.
	ACL       *BaseEntityACL `json:"_acl,omitempty"`
	CreatedAt *time.Time     `json:"_created_at,omitempty"`
	Files     *BaseRelation  `json:"_files,omitempty"`
	ID        *string        `json:"_id,omitempty"`
	// Manifest ID used to create/update the entity
	Manifest []string `json:"_manifest,omitempty"`
	// Organization Id the entity belongs to
	Org         string            `json:"_org"`
	Owners      []BaseEntityOwner `json:"_owners,omitempty"`
	Purpose     []string          `json:"_purpose,omitempty"`
	Schema      TaxSchema         `json:"_schema"`
	Tags        []string          `json:"_tags,omitempty"`
	Title       *string           `json:"_title,omitempty"`
	UpdatedAt   *time.Time        `json:"_updated_at,omitempty"`
	Active      bool              `json:"active"`
	Description *string           `json:"description,omitempty"`
	Rate        string            `json:"rate"`
	Region      string            `json:"region"`
	Type        TaxType           `json:"type"`
}

func (t Tax) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *Tax) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Tax) GetAdditional() map[string]any {
	if o == nil {
		return nil
	}
	return o.Additional
}

func (o *Tax) GetACL() *BaseEntityACL {
	if o == nil {
		return nil
	}
	return o.ACL
}

func (o *Tax) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Tax) GetFiles() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.Files
}

func (o *Tax) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Tax) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *Tax) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *Tax) GetOwners() []BaseEntityOwner {
	if o == nil {
		return nil
	}
	return o.Owners
}

func (o *Tax) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *Tax) GetSchema() TaxSchema {
	if o == nil {
		return TaxSchema("")
	}
	return o.Schema
}

func (o *Tax) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Tax) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Tax) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
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

func (o *Tax) GetRegion() string {
	if o == nil {
		return ""
	}
	return o.Region
}

func (o *Tax) GetType() TaxType {
	if o == nil {
		return TaxType("")
	}
	return o.Type
}
