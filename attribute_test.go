/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package superschema

import (
	"context"
	"testing"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// TestStringAttributeIsResource verifies StringAttribute.IsResource logic.
func TestStringAttributeIsResource(t *testing.T) {
	tests := []struct {
		name     string
		attr     StringAttribute
		expected bool
	}{
		{
			name:     "nil Common and nil Resource",
			attr:     StringAttribute{Common: nil, Resource: nil},
			expected: false,
		},
		{
			name:     "Common set, Resource nil",
			attr:     StringAttribute{Common: &schemaR.StringAttribute{}, Resource: nil},
			expected: true,
		},
		{
			name:     "Common nil, Resource set",
			attr:     StringAttribute{Common: nil, Resource: &schemaR.StringAttribute{}},
			expected: true,
		},
		{
			name:     "Both set",
			attr:     StringAttribute{Common: &schemaR.StringAttribute{}, Resource: &schemaR.StringAttribute{}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.attr.IsResource()
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestStringAttributeIsDataSource verifies StringAttribute.IsDataSource logic.
func TestStringAttributeIsDataSource(t *testing.T) {
	tests := []struct {
		name     string
		attr     StringAttribute
		expected bool
	}{
		{
			name:     "nil Common and nil DataSource",
			attr:     StringAttribute{Common: nil, DataSource: nil},
			expected: false,
		},
		{
			name:     "Common set, DataSource nil",
			attr:     StringAttribute{Common: &schemaR.StringAttribute{}, DataSource: nil},
			expected: true,
		},
		{
			name:     "Common nil, DataSource set",
			attr:     StringAttribute{Common: nil, DataSource: &schemaD.StringAttribute{}},
			expected: true,
		},
		{
			name:     "Both set",
			attr:     StringAttribute{Common: &schemaR.StringAttribute{}, DataSource: &schemaD.StringAttribute{}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.attr.IsDataSource()
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestStringAttributeGetResource verifies StringAttribute.GetResource merges Common and Resource correctly.
func TestStringAttributeGetResource(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name         string
		attr         StringAttribute
		expectedReq  bool
		expectedOpt  bool
		expectedComp bool
		expectedDesc string
	}{
		{
			name: "Common with Required only",
			attr: StringAttribute{
				Common: &schemaR.StringAttribute{Required: true, Description: "required field"},
			},
			expectedReq:  true,
			expectedOpt:  false,
			expectedComp: false,
			expectedDesc: "required field",
		},
		{
			name: "Resource with Optional only",
			attr: StringAttribute{
				Resource: &schemaR.StringAttribute{Optional: true, Description: "optional field"},
			},
			expectedReq:  false,
			expectedOpt:  true,
			expectedComp: false,
			expectedDesc: "optional field",
		},
		{
			name: "Common + Resource merge",
			attr: StringAttribute{
				Common:   &schemaR.StringAttribute{Required: true, Description: "base"},
				Resource: &schemaR.StringAttribute{Optional: true, Description: "resource"},
			},
			expectedReq:  true, // Required from Common: true
			expectedOpt:  true, // Optional from Resource: true
			expectedComp: false,
			expectedDesc: "base resource", // merged descriptions
		},
		{
			name: "All three set - union logic",
			attr: StringAttribute{
				Common:   &schemaR.StringAttribute{Required: true, Description: "common"},
				Resource: &schemaR.StringAttribute{Computed: true, Description: "resource"},
			},
			expectedReq:  true, // from Common
			expectedOpt:  false,
			expectedComp: true, // from Resource
			expectedDesc: "common resource",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.attr.GetResource(ctx)
			attr, ok := result.(schemaR.StringAttribute)
			if !ok {
				t.Fatalf("expected schemaR.StringAttribute, got %T", result)
			}

			if attr.Required != tt.expectedReq {
				t.Errorf("Required: got %v, want %v", attr.Required, tt.expectedReq)
			}
			if attr.Optional != tt.expectedOpt {
				t.Errorf("Optional: got %v, want %v", attr.Optional, tt.expectedOpt)
			}
			if attr.Computed != tt.expectedComp {
				t.Errorf("Computed: got %v, want %v", attr.Computed, tt.expectedComp)
			}
		})
	}
}

// TestStringAttributeGetDataSource verifies StringAttribute.GetDataSource merges Common and DataSource correctly.
func TestStringAttributeGetDataSource(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name         string
		attr         StringAttribute
		expectedReq  bool
		expectedOpt  bool
		expectedComp bool
	}{
		{
			name: "Common with Computed only",
			attr: StringAttribute{
				Common: &schemaR.StringAttribute{Computed: true},
			},
			expectedReq:  false,
			expectedOpt:  false,
			expectedComp: true,
		},
		{
			name: "DataSource with Optional only",
			attr: StringAttribute{
				DataSource: &schemaD.StringAttribute{Optional: true},
			},
			expectedReq:  false,
			expectedOpt:  true,
			expectedComp: false,
		},
		{
			name: "Common + DataSource merge (OR logic)",
			attr: StringAttribute{
				Common:     &schemaR.StringAttribute{Computed: true},
				DataSource: &schemaD.StringAttribute{Optional: true},
			},
			expectedReq:  false,
			expectedOpt:  true,
			expectedComp: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.attr.GetDataSource(ctx)
			attr, ok := result.(schemaD.StringAttribute)
			if !ok {
				t.Fatalf("expected schemaD.StringAttribute, got %T", result)
			}

			if attr.Required != tt.expectedReq {
				t.Errorf("Required: got %v, want %v", attr.Required, tt.expectedReq)
			}
			if attr.Optional != tt.expectedOpt {
				t.Errorf("Optional: got %v, want %v", attr.Optional, tt.expectedOpt)
			}
			if attr.Computed != tt.expectedComp {
				t.Errorf("Computed: got %v, want %v", attr.Computed, tt.expectedComp)
			}
		})
	}
}

// TestBoolAttributeIsResource verifies BoolAttribute.IsResource logic.
func TestBoolAttributeIsResource(t *testing.T) {
	tests := []struct {
		name     string
		attr     BoolAttribute
		expected bool
	}{
		{
			name:     "only Common set",
			attr:     BoolAttribute{Common: &schemaR.BoolAttribute{}, Resource: nil},
			expected: true,
		},
		{
			name:     "only Resource set",
			attr:     BoolAttribute{Common: nil, Resource: &schemaR.BoolAttribute{}},
			expected: true,
		},
		{
			name:     "neither set",
			attr:     BoolAttribute{Common: nil, Resource: nil},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.attr.IsResource()
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestBoolAttributeIsDataSource verifies BoolAttribute.IsDataSource logic.
func TestBoolAttributeIsDataSource(t *testing.T) {
	tests := []struct {
		name     string
		attr     BoolAttribute
		expected bool
	}{
		{
			name:     "only Common set",
			attr:     BoolAttribute{Common: &schemaR.BoolAttribute{}, DataSource: nil},
			expected: true,
		},
		{
			name:     "only DataSource set",
			attr:     BoolAttribute{Common: nil, DataSource: &schemaD.BoolAttribute{}},
			expected: true,
		},
		{
			name:     "neither set",
			attr:     BoolAttribute{Common: nil, DataSource: nil},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.attr.IsDataSource()
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestInt64AttributeGetResource verifies Int64Attribute.GetResource works correctly.
func TestInt64AttributeGetResource(t *testing.T) {
	ctx := context.Background()

	attr := Int64Attribute{
		Common: &schemaR.Int64Attribute{
			Required:    true,
			Description: "common desc",
		},
		Resource: &schemaR.Int64Attribute{
			Description: "resource desc",
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.Int64Attribute)
	if !ok {
		t.Fatalf("expected schemaR.Int64Attribute, got %T", result)
	}

	if !resAttr.Required {
		t.Error("Expected Required=true")
	}
	if resAttr.Optional {
		t.Error("Expected Optional=false")
	}
}

// TestListAttributeGetResource verifies ListAttribute.GetResource works correctly.
func TestListAttributeGetResource(t *testing.T) {
	ctx := context.Background()

	attr := ListAttribute{
		Common: &schemaR.ListAttribute{
			Optional: true,
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.ListAttribute)
	if !ok {
		t.Fatalf("expected schemaR.ListAttribute, got %T", result)
	}

	if !resAttr.Optional {
		t.Error("Expected Optional=true")
	}
}

// TestListNestedAttributeIsResource verifies ListNestedAttribute has Attributes field.
func TestListNestedAttributeStructure(t *testing.T) {
	attr := ListNestedAttribute{
		Common: &schemaR.ListNestedAttribute{},
		Attributes: Attributes{
			"name": StringAttribute{
				Common: &schemaR.StringAttribute{Required: true},
			},
		},
	}

	if attr.IsResource() != true {
		t.Error("Expected IsResource() to return true")
	}

	if len(attr.Attributes) != 1 {
		t.Error("Expected Attributes to have 1 entry")
	}

	if _, ok := attr.Attributes["name"]; !ok {
		t.Error("Expected 'name' attribute in Attributes map")
	}
}

// TestSetNestedAttributeIsDataSource verifies SetNestedAttribute.IsDataSource logic.
func TestSetNestedAttributeIsDataSource(t *testing.T) {
	attr := SetNestedAttribute{
		Common: &schemaR.SetNestedAttribute{},
		Attributes: Attributes{
			"id": StringAttribute{
				Common: &schemaR.StringAttribute{Computed: true},
			},
		},
	}

	if !attr.IsDataSource() {
		t.Error("Expected IsDataSource() to return true")
	}
}

// TestSingleNestedAttributeHasAttributes verifies SingleNestedAttribute.GetResource works with nested attributes.
func TestSingleNestedAttributeGetResource(t *testing.T) {
	ctx := context.Background()

	attr := SingleNestedAttribute{
		Common: &schemaR.SingleNestedAttribute{
			Optional: true,
		},
		Attributes: Attributes{
			"field": StringAttribute{
				Resource: &schemaR.StringAttribute{Required: true},
			},
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.SingleNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaR.SingleNestedAttribute, got %T", result)
	}

	if !resAttr.Optional {
		t.Error("Expected Optional=true")
	}

	if len(resAttr.Attributes) != 1 {
		t.Error("Expected nested Attributes to have 1 entry")
	}
}

// TestAttributeImplementsInterface verifies generated attributes implement the Attribute interface.
func TestAttributeImplementsInterface(t *testing.T) {
	var _ Attribute = StringAttribute{}
	var _ Attribute = BoolAttribute{}
	var _ Attribute = Int64Attribute{}
	var _ Attribute = Float64Attribute{}
	var _ Attribute = ListAttribute{}
	var _ Attribute = SetAttribute{}
	var _ Attribute = MapAttribute{}
	var _ Attribute = ObjectAttribute{}
	var _ Attribute = ListNestedAttribute{}
	var _ Attribute = SetNestedAttribute{}
	var _ Attribute = MapNestedAttribute{}
	var _ Attribute = SingleNestedAttribute{}
	var _ Attribute = NumberAttribute{}
	var _ Attribute = Int32Attribute{}
}
