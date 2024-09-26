// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReplicationTargetLocation replication target location
//
// swagger:model ReplicationTargetLocation
type ReplicationTargetLocation struct {

	// regionZone of replication site
	Region string `json:"region,omitempty"`

	// the replication site is active / down
	// Enum: ["active","down"]
	Status string `json:"status,omitempty"`
}

// Validate validates this replication target location
func (m *ReplicationTargetLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var replicationTargetLocationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","down"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		replicationTargetLocationTypeStatusPropEnum = append(replicationTargetLocationTypeStatusPropEnum, v)
	}
}

const (

	// ReplicationTargetLocationStatusActive captures enum value "active"
	ReplicationTargetLocationStatusActive string = "active"

	// ReplicationTargetLocationStatusDown captures enum value "down"
	ReplicationTargetLocationStatusDown string = "down"
)

// prop value enum
func (m *ReplicationTargetLocation) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, replicationTargetLocationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReplicationTargetLocation) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this replication target location based on context it is used
func (m *ReplicationTargetLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationTargetLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationTargetLocation) UnmarshalBinary(b []byte) error {
	var res ReplicationTargetLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
