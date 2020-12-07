// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GroupNamespaceRes group namespace res
// swagger:model GroupNamespaceRes
type GroupNamespaceRes struct {

	// Whether the data is valid
	EnableFlag int64 `json:"enableFlag,omitempty"`

	// the groupId
	GroupID int64 `json:"groupId,omitempty"`

	// the groupId
	GroupName string `json:"groupName,omitempty"`

	// the groupId
	GroupType string `json:"groupType,omitempty"`

	// the id of GroupStorage
	ID int64 `json:"id,omitempty"`

	// the name of namespace
	Namespace string `json:"namespace,omitempty"`

	// the namespaceId
	NamespaceID int64 `json:"namespaceId,omitempty"`

	// the remarks UserGroup
	Remarks string `json:"remarks,omitempty"`
}

// Validate validates this group namespace res
func (m *GroupNamespaceRes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupNamespaceRes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupNamespaceRes) UnmarshalBinary(b []byte) error {
	var res GroupNamespaceRes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}