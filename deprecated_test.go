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

	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// TestDeprecatedAttribute verifies StringAttribute with Deprecated field.
func TestDeprecatedAttribute(t *testing.T) {
	ctx := context.Background()

	attr := StringAttribute{
		Common: &schemaR.StringAttribute{Optional: true, Description: "old field"},
		Deprecated: &Deprecated{
			DeprecationMessage: deprUseNewField,
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.StringAttribute)
	if !ok {
		t.Fatalf("expected schemaR.StringAttribute, got %T", result)
	}

	if resAttr.DeprecationMessage != deprUseNewField {
		t.Errorf("expected deprecation message, got %q", resAttr.DeprecationMessage)
	}
}

// TestNonDeprecatedAttribute verifies StringAttribute without Deprecated field.
func TestNonDeprecatedAttribute(t *testing.T) {
	ctx := context.Background()

	attr := StringAttribute{
		Common: &schemaR.StringAttribute{Optional: true},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.StringAttribute)
	if !ok {
		t.Fatalf("expected schemaR.StringAttribute, got %T", result)
	}

	if resAttr.DeprecationMessage != "" {
		t.Errorf("expected empty deprecation message, got %q", resAttr.DeprecationMessage)
	}
}

// TestDeprecatedWithDescription verifies deprecation works with description merging.
func TestDeprecatedWithDescription(t *testing.T) {
	ctx := context.Background()

	attr := StringAttribute{
		Common: &schemaR.StringAttribute{
			Optional:    true,
			Description: "base description",
		},
		Resource: &schemaR.StringAttribute{
			Description: "resource specific",
		},
		Deprecated: &Deprecated{
			DeprecationMessage: deprUseNewField,
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.StringAttribute)
	if !ok {
		t.Fatalf("expected schemaR.StringAttribute, got %T", result)
	}

	if resAttr.DeprecationMessage != "Use new_field instead" {
		t.Errorf("expected deprecation message, got %q", resAttr.DeprecationMessage)
	}

	// Description merge should still work
	if resAttr.Description == "" {
		t.Error("expected non-empty description")
	}
}

// TestDeprecatedInt64Attribute verifies deprecation on different attribute type.
func TestDeprecatedInt64Attribute(t *testing.T) {
	ctx := context.Background()

	attr := Int64Attribute{
		Common: &schemaR.Int64Attribute{Optional: true},
		Deprecated: &Deprecated{
			DeprecationMessage: "Use new_int_field instead",
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.Int64Attribute)
	if !ok {
		t.Fatalf("expected schemaR.Int64Attribute, got %T", result)
	}

	if resAttr.DeprecationMessage != "Use new_int_field instead" {
		t.Errorf("expected deprecation message, got %q", resAttr.DeprecationMessage)
	}
}

// TestDeprecatedListAttribute verifies deprecation on complex attribute type.
func TestDeprecatedListAttribute(t *testing.T) {
	ctx := context.Background()

	attr := ListAttribute{
		Common: &schemaR.ListAttribute{Optional: true},
		Deprecated: &Deprecated{
			DeprecationMessage: "List field is deprecated",
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.ListAttribute)
	if !ok {
		t.Fatalf("expected schemaR.ListAttribute, got %T", result)
	}

	if resAttr.DeprecationMessage != "List field is deprecated" {
		t.Errorf("expected deprecation message, got %q", resAttr.DeprecationMessage)
	}
}

// TestDeprecatedNestedAttribute verifies deprecation on nested attribute types.
func TestDeprecatedNestedAttribute(t *testing.T) {
	ctx := context.Background()

	attr := ListNestedAttribute{
		Common: &schemaR.ListNestedAttribute{Optional: true},
		Attributes: Attributes{
			attrField: StringAttribute{
				Common: &schemaR.StringAttribute{Optional: true},
			},
		},
		Deprecated: &Deprecated{
			DeprecationMessage: "Use new_list_nested instead",
		},
	}

	result := attr.GetResource(ctx)
	resAttr, ok := result.(schemaR.ListNestedAttribute)
	if !ok {
		t.Fatalf("expected schemaR.ListNestedAttribute, got %T", result)
	}

	if resAttr.DeprecationMessage != "Use new_list_nested instead" {
		t.Errorf("expected deprecation message, got %q", resAttr.DeprecationMessage)
	}
}

// TestSchemaWithDeprecation verifies Schema can include deprecated attributes.
func TestSchemaWithDeprecatedAttributes(t *testing.T) {
	ctx := context.Background()

	schema := Schema{
		Attributes: Attributes{
			"new_field": StringAttribute{
				Common: &schemaR.StringAttribute{Required: true},
			},
			"old_field": StringAttribute{
				Common: &schemaR.StringAttribute{Optional: true},
				Deprecated: &Deprecated{
					DeprecationMessage: deprUseNewField,
				},
			},
		},
	}

	result := schema.GetResource(ctx)

	if len(result.Attributes) != 2 {
		t.Errorf("expected 2 attributes, got %d", len(result.Attributes))
	}

	// Verify old_field is included but marked as deprecated
	oldField, ok := result.Attributes["old_field"]
	if !ok {
		t.Error("expected 'old_field' attribute")
	}

	oldFieldAttr, ok := oldField.(schemaR.StringAttribute)
	if !ok {
		t.Errorf("expected schemaR.StringAttribute, got %T", oldField)
	}

	if oldFieldAttr.DeprecationMessage != "Use new_field instead" {
		t.Errorf("expected deprecation message on old_field")
	}
}
