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
	"testing"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// TestAttributesProcessResource verifies Attributes.process() filters for resource context.
func TestAttributesProcessResource(t *testing.T) {
	ctx := t.Context()

	attrs := Attributes{
		"common_only": StringAttribute{
			Common: &schemaR.StringAttribute{Description: "common only"},
		},
		"resource_only": StringAttribute{
			Resource: &schemaR.StringAttribute{Description: "resource only"},
		},
		"both": StringAttribute{
			Common:   &schemaR.StringAttribute{Description: "common"},
			Resource: &schemaR.StringAttribute{Description: "resource"},
		},
		"datasource_only": StringAttribute{
			DataSource: &schemaD.StringAttribute{Description: "datasource only"},
		},
	}

	result := attrs.process(ctx, resourceT)
	resAttrs, ok := result.(map[string]schemaR.Attribute)
	if !ok {
		t.Fatalf("expected map[string]schemaR.Attribute, got %T", result)
	}

	// Should include: common_only, resource_only, both
	// Should exclude: datasource_only (no Resource or Common)
	if len(resAttrs) != 3 {
		t.Errorf("expected 3 attributes, got %d: %v", len(resAttrs), resAttrs)
	}

	if _, ok := resAttrs["common_only"]; !ok {
		t.Error("expected 'common_only' in result")
	}
	if _, ok := resAttrs["resource_only"]; !ok {
		t.Error("expected 'resource_only' in result")
	}
	if _, ok := resAttrs["both"]; !ok {
		t.Error("expected 'both' in result")
	}
	if _, ok := resAttrs["datasource_only"]; ok {
		t.Error("did not expect 'datasource_only' in result")
	}
}

// TestAttributesProcessDataSource verifies Attributes.process() filters for datasource context.
func TestAttributesProcessDataSource(t *testing.T) {
	ctx := t.Context()

	attrs := Attributes{
		"common_only": StringAttribute{
			Common: &schemaR.StringAttribute{Description: "common only"},
		},
		"datasource_only": StringAttribute{
			DataSource: &schemaD.StringAttribute{Description: "datasource only"},
		},
		"both": StringAttribute{
			Common:     &schemaR.StringAttribute{Description: "common"},
			DataSource: &schemaD.StringAttribute{Description: "datasource"},
		},
		"resource_only": StringAttribute{
			Resource: &schemaR.StringAttribute{Description: "resource only"},
		},
	}

	result := attrs.process(ctx, dataSourceT)
	dsAttrs, ok := result.(map[string]schemaD.Attribute)
	if !ok {
		t.Fatalf("expected map[string]schemaD.Attribute, got %T", result)
	}

	// Should include: common_only, datasource_only, both
	// Should exclude: resource_only (no DataSource or Common)
	if len(dsAttrs) != 3 {
		t.Errorf("expected 3 attributes, got %d", len(dsAttrs))
	}

	if _, ok := dsAttrs["common_only"]; !ok {
		t.Error("expected 'common_only' in result")
	}
	if _, ok := dsAttrs["datasource_only"]; !ok {
		t.Error("expected 'datasource_only' in result")
	}
	if _, ok := dsAttrs["both"]; !ok {
		t.Error("expected 'both' in result")
	}
	if _, ok := dsAttrs["resource_only"]; ok {
		t.Error("did not expect 'resource_only' in result")
	}
}

// TestAttributesProcessEmpty verifies Attributes.process() handles empty map.
func TestAttributesProcessEmpty(t *testing.T) {
	ctx := t.Context()
	attrs := Attributes{}

	resultRes := attrs.process(ctx, resourceT)
	resAttrs, ok := resultRes.(map[string]schemaR.Attribute)
	if !ok {
		t.Fatalf("expected map[string]schemaR.Attribute, got %T", resultRes)
	}

	if len(resAttrs) != 0 {
		t.Errorf("expected 0 attributes, got %d", len(resAttrs))
	}

	resultDS := attrs.process(ctx, dataSourceT)
	dsAttrs, ok := resultDS.(map[string]schemaD.Attribute)
	if !ok {
		t.Fatalf("expected map[string]schemaD.Attribute, got %T", resultDS)
	}

	if len(dsAttrs) != 0 {
		t.Errorf("expected 0 attributes, got %d", len(dsAttrs))
	}
}

// TestSchemaGetResource verifies Schema.GetResource() processes all attributes for resource.
func TestSchemaGetResource(t *testing.T) {
	ctx := t.Context()

	schema := Schema{
		Common: SchemaDetails{
			MarkdownDescription: "base schema",
		},
		Resource: SchemaDetails{
			MarkdownDescription: "resource specific",
		},
		Attributes: Attributes{
			"name": StringAttribute{
				Resource: &schemaR.StringAttribute{Required: true},
			},
			"id": StringAttribute{
				Common: &schemaR.StringAttribute{Computed: true},
			},
		},
	}

	result := schema.GetResource(ctx)

	// Verify it returns a schemaR.Schema
	if len(result.Attributes) != 2 {
		t.Errorf("expected 2 attributes, got %d", len(result.Attributes))
	}

	if _, ok := result.Attributes["name"]; !ok {
		t.Error("expected 'name' attribute")
	}
	if _, ok := result.Attributes["id"]; !ok {
		t.Error("expected 'id' attribute")
	}
}

// TestSchemaGetDataSource verifies Schema.GetDataSource() processes all attributes for datasource.
func TestSchemaGetDataSource(t *testing.T) {
	ctx := t.Context()

	schema := Schema{
		Common: SchemaDetails{
			MarkdownDescription: "base schema",
		},
		DataSource: SchemaDetails{
			MarkdownDescription: "datasource specific",
		},
		Attributes: Attributes{
			"id": StringAttribute{
				DataSource: &schemaD.StringAttribute{Computed: true},
			},
			"name": StringAttribute{
				Common: &schemaR.StringAttribute{Optional: true},
			},
		},
	}

	result := schema.GetDataSource(ctx)

	// Verify it returns a schemaD.Schema
	if len(result.Attributes) != 2 {
		t.Errorf("expected 2 attributes, got %d", len(result.Attributes))
	}

	if _, ok := result.Attributes["id"]; !ok {
		t.Error("expected 'id' attribute")
	}
	if _, ok := result.Attributes["name"]; !ok {
		t.Error("expected 'name' attribute")
	}
}

// TestSchemaResourceOnly verifies attributes marked as Resource-only are excluded from DataSource.
func TestSchemaResourceOnly(t *testing.T) {
	ctx := t.Context()

	schema := Schema{
		Attributes: Attributes{
			"name": StringAttribute{
				Resource: &schemaR.StringAttribute{Required: true},
			},
		},
	}

	resSchema := schema.GetResource(ctx)
	if len(resSchema.Attributes) != 1 {
		t.Errorf("resource: expected 1 attribute, got %d", len(resSchema.Attributes))
	}

	dsSchema := schema.GetDataSource(ctx)
	if len(dsSchema.Attributes) != 0 {
		t.Errorf("datasource: expected 0 attributes, got %d", len(dsSchema.Attributes))
	}
}

// TestSchemaDataSourceOnly verifies attributes marked as DataSource-only are excluded from Resource.
func TestSchemaDataSourceOnly(t *testing.T) {
	ctx := t.Context()

	schema := Schema{
		Attributes: Attributes{
			"computed_field": StringAttribute{
				DataSource: &schemaD.StringAttribute{Computed: true},
			},
		},
	}

	resSchema := schema.GetResource(ctx)
	if len(resSchema.Attributes) != 0 {
		t.Errorf("resource: expected 0 attributes, got %d", len(resSchema.Attributes))
	}

	dsSchema := schema.GetDataSource(ctx)
	if len(dsSchema.Attributes) != 1 {
		t.Errorf("datasource: expected 1 attribute, got %d", len(dsSchema.Attributes))
	}
}

// TestSchemaCommonAttributes verifies Common attributes appear in both Resource and DataSource.
func TestSchemaCommonAttributes(t *testing.T) {
	ctx := t.Context()

	schema := Schema{
		Attributes: Attributes{
			"shared": StringAttribute{
				Common: &schemaR.StringAttribute{Optional: true},
			},
		},
	}

	resSchema := schema.GetResource(ctx)
	if len(resSchema.Attributes) != 1 {
		t.Errorf("resource: expected 1 attribute, got %d", len(resSchema.Attributes))
	}

	dsSchema := schema.GetDataSource(ctx)
	if len(dsSchema.Attributes) != 1 {
		t.Errorf("datasource: expected 1 attribute, got %d", len(dsSchema.Attributes))
	}
}

// TestComplexSchema verifies a realistic schema with mixed attribute types.
func TestComplexSchema(t *testing.T) {
	schema := Schema{
		Common: SchemaDetails{
			MarkdownDescription: "API Object",
		},
		Resource: SchemaDetails{
			MarkdownDescription: "for resources",
		},
		DataSource: SchemaDetails{
			MarkdownDescription: "for data sources",
		},
		Attributes: Attributes{
			"id": StringAttribute{
				DataSource: &schemaD.StringAttribute{Computed: true},
			},
			"name": StringAttribute{
				Common: &schemaR.StringAttribute{Required: true},
			},
			"tags": MapAttribute{
				Common: &schemaR.MapAttribute{Optional: true},
			},
			"config": SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{Optional: true},
				Attributes: Attributes{
					"enabled": BoolAttribute{
						Common: &schemaR.BoolAttribute{Optional: true},
					},
				},
			},
		},
	}

	resSchema := schema.GetResource(t.Context())
	// Should have: name, tags, config (NOT id, which is DataSource-only)
	if len(resSchema.Attributes) != 3 {
		t.Errorf("resource: expected 3 attributes, got %d", len(resSchema.Attributes))
	}

	dsSchema := schema.GetDataSource(t.Context())
	// Should have: id, name, tags, config (all present because id has DataSource, others have Common)
	if len(dsSchema.Attributes) != 4 {
		t.Errorf("datasource: expected 4 attributes, got %d", len(dsSchema.Attributes))
	}
}
