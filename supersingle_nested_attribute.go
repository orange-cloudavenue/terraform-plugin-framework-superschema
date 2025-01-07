/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

// code generated by go generate - look at supertype_attribute.go.tmpl for source file
package superschema

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	supertypes "github.com/orange-cloudavenue/terraform-plugin-framework-supertypes"
)

var _ Attribute = SuperSingleNestedAttribute{}

type SuperSingleNestedAttribute struct {
	Deprecated *Deprecated
	Common     *schemaR.SingleNestedAttribute
	Resource   *schemaR.SingleNestedAttribute
	DataSource *schemaD.SingleNestedAttribute
	Attributes Attributes
}

// IsResource returns true if the attribute is a resource attribute.
func (s SuperSingleNestedAttribute) IsResource() bool {
	return s.Resource != nil || s.Common != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s SuperSingleNestedAttribute) IsDataSource() bool {
	return s.DataSource != nil || s.Common != nil
}

// GetCustomType returns the custom type of the attribute.
func (s SuperSingleNestedAttribute) getCustomType(aT map[string]attr.Type) basetypes.ObjectTypable {
	return supertypes.SingleNestedType{
		ObjectType: types.ObjectType{
			AttrTypes: aT,
		},
	}
}

//nolint:dupl
func (s SuperSingleNestedAttribute) GetResource(ctx context.Context) schemaR.Attribute {
	var (
		common   schemaR.SingleNestedAttribute
		resource schemaR.SingleNestedAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.Resource != nil {
		resource = *s.Resource
	}

	a := schemaR.SingleNestedAttribute{
		Required:            computeIsRequired(common, resource),
		Optional:            computeIsOptional(common, resource),
		Computed:            computeIsComputed(common, resource),
		Sensitive:           computeIsSensitive(common, resource),
		MarkdownDescription: computeMarkdownDescription(common, resource),
		Description:         computeDescription(common, resource),
		DeprecationMessage:  computeDeprecationMessage(common, resource),
		Attributes:          s.Attributes.process(ctx, resourceT).(map[string]schemaR.Attribute),
	}

	a.Validators = append(a.Validators, common.Validators...)
	a.Validators = append(a.Validators, resource.Validators...)
	a.PlanModifiers = append(a.PlanModifiers, common.PlanModifiers...)
	a.PlanModifiers = append(a.PlanModifiers, resource.PlanModifiers...)

	defaultVDescription := ""

	if s.Common != nil {
		if s.Common.CustomType != nil {
			a.CustomType = s.Common.CustomType
		}
	}

	if s.Resource != nil {
		if s.Resource.Default != nil {
			a.Default = s.Resource.Default
			defaultVDescription = s.Resource.Default.MarkdownDescription(ctx)
		}
		if s.Resource.CustomType != nil {
			a.CustomType = s.Resource.CustomType
		}
	}
	// * If user has not provided a custom type, we will use the default supertypes
	if a.CustomType == nil {
		attrTypes := make(map[string]attr.Type, len(a.Attributes))

		for name, attribute := range a.Attributes {
			attrTypes[name] = attribute.GetType()
		}

		a.CustomType = s.getCustomType(attrTypes).(supertypes.SingleNestedType)

	}

	deprecationMessage := ""
	if s.Deprecated != nil {
		a.DeprecationMessage = s.Deprecated.DeprecationMessage
		deprecationMessage = s.Deprecated.computeDeprecatedDocumentation()
	}

	a.MarkdownDescription = genResourceAttrDescription(ctx, a.MarkdownDescription, defaultVDescription, deprecationMessage, a.Validators, a.PlanModifiers)
	return a
}

//nolint:dupl
func (s SuperSingleNestedAttribute) GetDataSource(ctx context.Context) schemaD.Attribute {
	var (
		common     schemaR.SingleNestedAttribute
		dataSource schemaD.SingleNestedAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.DataSource != nil {
		dataSource = *s.DataSource
	}

	a := schemaD.SingleNestedAttribute{
		Required:            computeIsRequired(common, dataSource),
		Optional:            computeIsOptional(common, dataSource),
		Computed:            computeIsComputed(common, dataSource),
		Sensitive:           computeIsSensitive(common, dataSource),
		MarkdownDescription: computeMarkdownDescription(common, dataSource),
		Description:         computeDescription(common, dataSource),
		DeprecationMessage:  computeDeprecationMessage(common, dataSource),
		Attributes:          s.Attributes.process(ctx, dataSourceT).(map[string]schemaD.Attribute),
	}

	a.Validators = append(a.Validators, common.Validators...)
	a.Validators = append(a.Validators, dataSource.Validators...)

	if s.Common != nil {
		if s.Common.CustomType != nil {
			a.CustomType = s.Common.CustomType
		}
	}

	if s.DataSource != nil {
		if s.DataSource.CustomType != nil {
			a.CustomType = s.DataSource.CustomType
		}
	}
	// * If user has not provided a custom type, we will use the default supertypes
	if a.CustomType == nil {
		attrTypes := make(map[string]attr.Type, len(a.Attributes))

		for name, attribute := range a.Attributes {
			attrTypes[name] = attribute.GetType()
		}

		a.CustomType = s.getCustomType(attrTypes).(supertypes.SingleNestedType)
	}

	deprecationMessage := ""
	if s.Deprecated != nil {
		a.DeprecationMessage = s.Deprecated.DeprecationMessage
		deprecationMessage = s.Deprecated.computeDeprecatedDocumentation()
	}

	a.MarkdownDescription = genDataSourceAttrDescription(ctx, a.MarkdownDescription, deprecationMessage, a.Validators)
	return a
}

// * SuperTypeOf

var _ Attribute = SuperSingleNestedAttributeOf[struct{}]{}

type SuperSingleNestedAttributeOf[T any] struct {
	Deprecated *Deprecated
	Common     *schemaR.SingleNestedAttribute
	Resource   *schemaR.SingleNestedAttribute
	DataSource *schemaD.SingleNestedAttribute
	Attributes Attributes
}

// IsResource returns true if the attribute is a resource attribute.
func (s SuperSingleNestedAttributeOf[T]) IsResource() bool {
	return s.Resource != nil || s.Common != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s SuperSingleNestedAttributeOf[T]) IsDataSource() bool {
	return s.DataSource != nil || s.Common != nil
}

//nolint:dupl
func (s SuperSingleNestedAttributeOf[T]) GetResource(ctx context.Context) schemaR.Attribute {
	var (
		common   schemaR.SingleNestedAttribute
		resource schemaR.SingleNestedAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.Resource != nil {
		resource = *s.Resource
	}

	a := schemaR.SingleNestedAttribute{
		Required:            computeIsRequired(common, resource),
		Optional:            computeIsOptional(common, resource),
		Computed:            computeIsComputed(common, resource),
		Sensitive:           computeIsSensitive(common, resource),
		MarkdownDescription: computeMarkdownDescription(common, resource),
		Description:         computeDescription(common, resource),
		DeprecationMessage:  computeDeprecationMessage(common, resource),
		Attributes:          s.Attributes.process(ctx, resourceT).(map[string]schemaR.Attribute),
	}

	a.Validators = append(a.Validators, common.Validators...)
	a.Validators = append(a.Validators, resource.Validators...)
	a.PlanModifiers = append(a.PlanModifiers, common.PlanModifiers...)
	a.PlanModifiers = append(a.PlanModifiers, resource.PlanModifiers...)

	defaultVDescription := ""

	if s.Common != nil {
		if s.Common.CustomType != nil {
			a.CustomType = s.Common.CustomType
		}
	}

	if s.Resource != nil {
		if s.Resource.Default != nil {
			a.Default = s.Resource.Default
			defaultVDescription = s.Resource.Default.MarkdownDescription(ctx)
		}
		if s.Resource.CustomType != nil {
			a.CustomType = s.Resource.CustomType
		}
	}
	// * If user has not provided a custom type, we will use the default supertypes
	if a.CustomType == nil {
		a.CustomType = supertypes.NewSingleNestedObjectTypeOf[T](ctx)
	}

	deprecationMessage := ""
	if s.Deprecated != nil {
		a.DeprecationMessage = s.Deprecated.DeprecationMessage
		deprecationMessage = s.Deprecated.computeDeprecatedDocumentation()
	}

	a.MarkdownDescription = genResourceAttrDescription(ctx, a.MarkdownDescription, defaultVDescription, deprecationMessage, a.Validators, a.PlanModifiers)
	return a
}

//nolint:dupl
func (s SuperSingleNestedAttributeOf[T]) GetDataSource(ctx context.Context) schemaD.Attribute {
	var (
		common     schemaR.SingleNestedAttribute
		dataSource schemaD.SingleNestedAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.DataSource != nil {
		dataSource = *s.DataSource
	}

	a := schemaD.SingleNestedAttribute{
		Required:            computeIsRequired(common, dataSource),
		Optional:            computeIsOptional(common, dataSource),
		Computed:            computeIsComputed(common, dataSource),
		Sensitive:           computeIsSensitive(common, dataSource),
		MarkdownDescription: computeMarkdownDescription(common, dataSource),
		Description:         computeDescription(common, dataSource),
		DeprecationMessage:  computeDeprecationMessage(common, dataSource),
		Attributes:          s.Attributes.process(ctx, dataSourceT).(map[string]schemaD.Attribute),
	}

	a.Validators = append(a.Validators, common.Validators...)
	a.Validators = append(a.Validators, dataSource.Validators...)

	if s.Common != nil {
		if s.Common.CustomType != nil {
			a.CustomType = s.Common.CustomType
		}
	}

	if s.DataSource != nil {
		if s.DataSource.CustomType != nil {
			a.CustomType = s.DataSource.CustomType
		}
	}
	// * If user has not provided a custom type, we will use the default supertypes
	if a.CustomType == nil {
		a.CustomType = supertypes.NewSingleNestedObjectTypeOf[T](ctx)
	}

	deprecationMessage := ""
	if s.Deprecated != nil {
		a.DeprecationMessage = s.Deprecated.DeprecationMessage
		deprecationMessage = s.Deprecated.computeDeprecatedDocumentation()
	}

	a.MarkdownDescription = genDataSourceAttrDescription(ctx, a.MarkdownDescription, deprecationMessage, a.Validators)
	return a
}
