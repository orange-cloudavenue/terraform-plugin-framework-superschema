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
	"reflect"
	"testing"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	supertypes "github.com/orange-cloudavenue/terraform-plugin-framework-supertypes"
)

// TestType is a simple struct used for testing generic Super*NestedAttributeOf[T] types.
type TestType struct {
	Name string `tfsdk:"name"`
	ID   int    `tfsdk:"id"`
}

// TestSuperSetNestedAttributeOfIsResource verifies SuperSetNestedAttributeOf[T].IsResource logic.
func TestSuperSetNestedAttributeOfIsResource(t *testing.T) {
	tests := []struct {
		name     string
		attr     SuperSetNestedAttributeOf[TestType]
		expected bool
	}{
		{
			name:     "only Common set",
			attr:     SuperSetNestedAttributeOf[TestType]{Common: &schemaR.SetNestedAttribute{}, Resource: nil},
			expected: true,
		},
		{
			name:     "only Resource set",
			attr:     SuperSetNestedAttributeOf[TestType]{Common: nil, Resource: &schemaR.SetNestedAttribute{}},
			expected: true,
		},
		{
			name:     "neither set",
			attr:     SuperSetNestedAttributeOf[TestType]{Common: nil, Resource: nil},
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

// TestSuperSetNestedAttributeOfIsDataSource verifies SuperSetNestedAttributeOf[T].IsDataSource logic.
func TestSuperSetNestedAttributeOfIsDataSource(t *testing.T) {
	tests := []struct {
		name     string
		attr     SuperSetNestedAttributeOf[TestType]
		expected bool
	}{
		{
			name:     "only Common set",
			attr:     SuperSetNestedAttributeOf[TestType]{Common: &schemaR.SetNestedAttribute{}, DataSource: nil},
			expected: true,
		},
		{
			name:     "only DataSource set",
			attr:     SuperSetNestedAttributeOf[TestType]{Common: nil, DataSource: &schemaD.SetNestedAttribute{}},
			expected: true,
		},
		{
			name:     "neither set",
			attr:     SuperSetNestedAttributeOf[TestType]{Common: nil, DataSource: nil},
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

// TestSuperSetNestedAttributeOfGetResourceNestedObjectCustomType verifies NestedObject.CustomType is set.
// This is the core fix for issue #52 - path-matching validators need NestedObject.CustomType set.
func TestSuperSetNestedAttributeOfGetResourceNestedObjectCustomType(t *testing.T) {
	ctx := context.Background()

	attr := SuperSetNestedAttributeOf[TestType]{
		Common: &schemaR.SetNestedAttribute{
			Optional: true,
		},
		Attributes: Attributes{
			"name": StringAttribute{
				Common: &schemaR.StringAttribute{Optional: true},
			},
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.SetNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaR.SetNestedAttribute, got %T", result)
	}

	// The fix: NestedObject.CustomType must be set to supertypes.NewObjectTypeOf[TestType](ctx)
	if resAttr.NestedObject.CustomType == nil {
		t.Error("NestedObject.CustomType is nil - fix for issue #52 not applied")
	}

	// Verify it's the correct type
	expectedType := supertypes.NewObjectTypeOf[TestType](ctx)
	if reflect.TypeOf(resAttr.NestedObject.CustomType) != reflect.TypeOf(expectedType) {
		t.Errorf("NestedObject.CustomType type mismatch: got %T, want %T", resAttr.NestedObject.CustomType, expectedType)
	}
}

// TestSuperSetNestedAttributeOfGetDataSourceNestedObjectCustomType verifies NestedObject.CustomType is set for datasource.
func TestSuperSetNestedAttributeOfGetDataSourceNestedObjectCustomType(t *testing.T) {
	ctx := context.Background()

	attr := SuperSetNestedAttributeOf[TestType]{
		Common: &schemaR.SetNestedAttribute{
			Computed: true,
		},
		Attributes: Attributes{
			"id": StringAttribute{
				Common: &schemaR.StringAttribute{Computed: true},
			},
		},
	}

	result := attr.GetDataSource(ctx)
	dsAttr, ok := result.(schemaD.SetNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaD.SetNestedAttribute, got %T", result)
	}

	// The fix: NestedObject.CustomType must be set even for datasource
	if dsAttr.NestedObject.CustomType == nil {
		t.Error("NestedObject.CustomType is nil in datasource - fix for issue #52 not applied")
	}

	expectedType := supertypes.NewObjectTypeOf[TestType](ctx)
	if reflect.TypeOf(dsAttr.NestedObject.CustomType) != reflect.TypeOf(expectedType) {
		t.Errorf("NestedObject.CustomType type mismatch: got %T, want %T", dsAttr.NestedObject.CustomType, expectedType)
	}
}

// TestSuperListNestedAttributeOfGetResourceNestedObjectCustomType verifies NestedObject.CustomType for ListNestedAttributeOf[T].
func TestSuperListNestedAttributeOfGetResourceNestedObjectCustomType(t *testing.T) {
	ctx := context.Background()

	attr := SuperListNestedAttributeOf[TestType]{
		Common: &schemaR.ListNestedAttribute{
			Optional: true,
		},
		Attributes: Attributes{
			"field": StringAttribute{
				Common: &schemaR.StringAttribute{Optional: true},
			},
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.ListNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaR.ListNestedAttribute, got %T", result)
	}

	if resAttr.NestedObject.CustomType == nil {
		t.Error("ListNestedAttributeOf[T]: NestedObject.CustomType is nil")
	}
}

// TestSuperMapNestedAttributeOfGetResourceNestedObjectCustomType verifies NestedObject.CustomType for MapNestedAttributeOf[T].
func TestSuperMapNestedAttributeOfGetResourceNestedObjectCustomType(t *testing.T) {
	ctx := context.Background()

	attr := SuperMapNestedAttributeOf[TestType]{
		Common: &schemaR.MapNestedAttribute{
			Optional: true,
		},
		Attributes: Attributes{
			"value": StringAttribute{
				Common: &schemaR.StringAttribute{Optional: true},
			},
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.MapNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaR.MapNestedAttribute, got %T", result)
	}

	if resAttr.NestedObject.CustomType == nil {
		t.Error("MapNestedAttributeOf[T]: NestedObject.CustomType is nil")
	}
}

// TestSuperSetNestedAttributeOfMergeLogic verifies merge of Common/Resource/DataSource with Of[T].
func TestSuperSetNestedAttributeOfMergeLogic(t *testing.T) {
	ctx := context.Background()

	attr := SuperSetNestedAttributeOf[TestType]{
		Common: &schemaR.SetNestedAttribute{
			Required:    true,
			Description: "common desc",
		},
		Resource: &schemaR.SetNestedAttribute{
			Optional:    true,
			Description: "resource desc",
		},
		Attributes: Attributes{
			"field": StringAttribute{
				Common: &schemaR.StringAttribute{Optional: true},
			},
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.SetNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaR.SetNestedAttribute, got %T", result)
	}

	// Verify merge logic still works with Of[T]
	if !resAttr.Required {
		t.Error("Expected Required=true from Common")
	}
	if !resAttr.Optional {
		t.Error("Expected Optional=true from Resource")
	}
}

// TestSuperSetNestedAttributeOfWithDeprecation verifies deprecation works with Of[T] generics.
func TestSuperSetNestedAttributeOfWithDeprecation(t *testing.T) {
	ctx := context.Background()

	attr := SuperSetNestedAttributeOf[TestType]{
		Common: &schemaR.SetNestedAttribute{
			Optional: true,
		},
		Deprecated: &Deprecated{
			DeprecationMessage: "Use new_set_nested instead",
		},
		Attributes: Attributes{
			"field": StringAttribute{
				Common: &schemaR.StringAttribute{Optional: true},
			},
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.SetNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaR.SetNestedAttribute, got %T", result)
	}

	if resAttr.DeprecationMessage != "Use new_set_nested instead" {
		t.Errorf("expected deprecation message, got %q", resAttr.DeprecationMessage)
	}
}

// TestSuperSetNestedAttributeOfNestedAttributesProcessing verifies nested attributes are processed correctly.
func TestSuperSetNestedAttributeOfNestedAttributesProcessing(t *testing.T) {
	ctx := context.Background()

	attr := SuperSetNestedAttributeOf[TestType]{
		Common: &schemaR.SetNestedAttribute{Optional: true},
		Attributes: Attributes{
			"name": StringAttribute{
				Common: &schemaR.StringAttribute{Required: true},
			},
			"description": StringAttribute{
				Resource: &schemaR.StringAttribute{Optional: true},
			},
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.SetNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaR.SetNestedAttribute, got %T", result)
	}

	if len(resAttr.NestedObject.Attributes) != 2 {
		t.Errorf("expected 2 nested attributes, got %d", len(resAttr.NestedObject.Attributes))
	}

	if _, ok := resAttr.NestedObject.Attributes["name"]; !ok {
		t.Error("expected 'name' in nested attributes")
	}
	if _, ok := resAttr.NestedObject.Attributes["description"]; !ok {
		t.Error("expected 'description' in nested attributes")
	}
}

// TestAllOfTypeVariantsImplementAttribute verifies all Of[T] variants implement Attribute interface.
func TestAllOfTypeVariantsImplementAttribute(t *testing.T) {
	// These compile-time checks ensure the interface is implemented
	var _ Attribute = SuperSetNestedAttributeOf[TestType]{}
	var _ Attribute = SuperListNestedAttributeOf[TestType]{}
	var _ Attribute = SuperMapNestedAttributeOf[TestType]{}
	var _ Attribute = SuperSingleNestedAttributeOf[TestType]{}
}

// TestSuperListNestedAttributeOfGetDataSourceNestedObjectCustomType verifies ListNestedAttributeOf[T] datasource fix.
func TestSuperListNestedAttributeOfGetDataSourceNestedObjectCustomType(t *testing.T) {
	ctx := context.Background()

	attr := SuperListNestedAttributeOf[TestType]{
		Common: &schemaR.ListNestedAttribute{
			Computed: true,
		},
		Attributes: Attributes{
			"id": StringAttribute{
				Common: &schemaR.StringAttribute{Computed: true},
			},
		},
	}

	result := attr.GetDataSource(ctx)
	dsAttr, ok := result.(schemaD.ListNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaD.ListNestedAttribute, got %T", result)
	}

	if dsAttr.NestedObject.CustomType == nil {
		t.Error("ListNestedAttributeOf[T]: NestedObject.CustomType is nil in datasource")
	}
}

// TestSuperMapNestedAttributeOfGetDataSourceNestedObjectCustomType verifies MapNestedAttributeOf[T] datasource fix.
func TestSuperMapNestedAttributeOfGetDataSourceNestedObjectCustomType(t *testing.T) {
	ctx := context.Background()

	attr := SuperMapNestedAttributeOf[TestType]{
		Common: &schemaR.MapNestedAttribute{
			Computed: true,
		},
		Attributes: Attributes{
			"value": StringAttribute{
				Common: &schemaR.StringAttribute{Computed: true},
			},
		},
	}

	result := attr.GetDataSource(ctx)
	dsAttr, ok := result.(schemaD.MapNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaD.MapNestedAttribute, got %T", result)
	}

	if dsAttr.NestedObject.CustomType == nil {
		t.Error("MapNestedAttributeOf[T]: NestedObject.CustomType is nil in datasource")
	}
}

// TestSuperSetNestedAttributeOfNestedObjectTypeNotNil verifies type assertion safety.
func TestSuperSetNestedAttributeOfNestedObjectTypeNotNil(t *testing.T) {
	ctx := context.Background()

	attr := SuperSetNestedAttributeOf[TestType]{
		Common:     &schemaR.SetNestedAttribute{Optional: true},
		Attributes: Attributes{},
	}

	result := attr.GetResource(ctx)
	resAttr := result.(schemaR.SetNestedAttribute)

	// Verify the type can be used (not nil and not a zero value)
	if resAttr.NestedObject.CustomType == nil {
		t.Fatal("NestedObject.CustomType is nil")
	}

	// Try to call methods on it to ensure it's a valid type
	typeName := resAttr.NestedObject.CustomType.String()
	if typeName == "" {
		t.Error("NestedObject.CustomType.String() returned empty - type may be invalid")
	}
}

// TestSuperSingleNestedAttributeOfGetResource verifies SingleNestedAttributeOf[T] works.
func TestSuperSingleNestedAttributeOfGetResource(t *testing.T) {
	ctx := context.Background()

	attr := SuperSingleNestedAttributeOf[TestType]{
		Common: &schemaR.SingleNestedAttribute{
			Optional: true,
		},
		Attributes: Attributes{
			"field": StringAttribute{
				Common: &schemaR.StringAttribute{Optional: true},
			},
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.SingleNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaR.SingleNestedAttribute, got %T", result)
	}

	if len(resAttr.Attributes) != 1 {
		t.Errorf("expected 1 nested attribute, got %d", len(resAttr.Attributes))
	}
}
