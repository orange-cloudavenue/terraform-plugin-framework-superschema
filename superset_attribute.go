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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	supertypes "github.com/orange-cloudavenue/terraform-plugin-framework-supertypes"
)

var _ Attribute = SuperSetAttribute{}

type SuperSetAttribute struct {
	Deprecated *Deprecated
	Common     *schemaR.SetAttribute
	Resource   *schemaR.SetAttribute
	DataSource *schemaD.SetAttribute
}

// IsResource returns true if the attribute is a resource attribute.
func (s SuperSetAttribute) IsResource() bool {
	return s.Resource != nil || s.Common != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s SuperSetAttribute) IsDataSource() bool {
	return s.DataSource != nil || s.Common != nil
}

// GetCustomType returns the custom type of the attribute.
func (s SuperSetAttribute) getCustomType(oB attr.Type) basetypes.SetTypable {
	return supertypes.SetType{
		SetType: basetypes.SetType{
			ElemType: oB,
		},
	}
}

//nolint:dupl
func (s SuperSetAttribute) GetResource(ctx context.Context) schemaR.Attribute {
	var (
		common   schemaR.SetAttribute
		resource schemaR.SetAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.Resource != nil {
		resource = *s.Resource
	}

	a := schemaR.SetAttribute{
		Required:            computeIsRequired(common, resource),
		Optional:            computeIsOptional(common, resource),
		Computed:            computeIsComputed(common, resource),
		Sensitive:           computeIsSensitive(common, resource),
		MarkdownDescription: computeMarkdownDescription(common, resource),
		Description:         computeDescription(common, resource),
		DeprecationMessage:  computeDeprecationMessage(common, resource),
		ElementType:         common.ElementType,
	}
	if s.Resource != nil {
		if s.Resource.ElementType != nil {
			a.ElementType = s.Resource.ElementType
		}
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
		a.CustomType = s.getCustomType(a.ElementType).(supertypes.SetType)
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
func (s SuperSetAttribute) GetDataSource(ctx context.Context) schemaD.Attribute {
	var (
		common     schemaR.SetAttribute
		dataSource schemaD.SetAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.DataSource != nil {
		dataSource = *s.DataSource
	}

	a := schemaD.SetAttribute{
		Required:            computeIsRequired(common, dataSource),
		Optional:            computeIsOptional(common, dataSource),
		Computed:            computeIsComputed(common, dataSource),
		Sensitive:           computeIsSensitive(common, dataSource),
		MarkdownDescription: computeMarkdownDescription(common, dataSource),
		Description:         computeDescription(common, dataSource),
		DeprecationMessage:  computeDeprecationMessage(common, dataSource),
		ElementType:         common.ElementType,
	}

	if s.DataSource != nil {
		if s.DataSource.ElementType != nil {
			a.ElementType = s.DataSource.ElementType
		}
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
		a.CustomType = s.getCustomType(a.ElementType).(supertypes.SetType)
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

var _ Attribute = SuperSetAttributeOf[struct{}]{}

type SuperSetAttributeOf[T any] struct {
	Deprecated *Deprecated
	Common     *schemaR.SetAttribute
	Resource   *schemaR.SetAttribute
	DataSource *schemaD.SetAttribute
}

// IsResource returns true if the attribute is a resource attribute.
func (s SuperSetAttributeOf[T]) IsResource() bool {
	return s.Resource != nil || s.Common != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s SuperSetAttributeOf[T]) IsDataSource() bool {
	return s.DataSource != nil || s.Common != nil
}

//nolint:dupl
func (s SuperSetAttributeOf[T]) GetResource(ctx context.Context) schemaR.Attribute {
	var (
		common   schemaR.SetAttribute
		resource schemaR.SetAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.Resource != nil {
		resource = *s.Resource
	}

	a := schemaR.SetAttribute{
		Required:            computeIsRequired(common, resource),
		Optional:            computeIsOptional(common, resource),
		Computed:            computeIsComputed(common, resource),
		Sensitive:           computeIsSensitive(common, resource),
		MarkdownDescription: computeMarkdownDescription(common, resource),
		Description:         computeDescription(common, resource),
		DeprecationMessage:  computeDeprecationMessage(common, resource),
		ElementType:         common.ElementType,
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
		a.CustomType = supertypes.NewSetTypeOf[T](ctx)
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
func (s SuperSetAttributeOf[T]) GetDataSource(ctx context.Context) schemaD.Attribute {
	var (
		common     schemaR.SetAttribute
		dataSource schemaD.SetAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.DataSource != nil {
		dataSource = *s.DataSource
	}

	a := schemaD.SetAttribute{
		Required:            computeIsRequired(common, dataSource),
		Optional:            computeIsOptional(common, dataSource),
		Computed:            computeIsComputed(common, dataSource),
		Sensitive:           computeIsSensitive(common, dataSource),
		MarkdownDescription: computeMarkdownDescription(common, dataSource),
		Description:         computeDescription(common, dataSource),
		DeprecationMessage:  computeDeprecationMessage(common, dataSource),
		ElementType:         common.ElementType,
	}

	if s.DataSource != nil {
		if s.DataSource.ElementType != nil {
			a.ElementType = s.DataSource.ElementType
		}
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
		a.CustomType = supertypes.NewSetTypeOf[T](ctx)
	}

	deprecationMessage := ""
	if s.Deprecated != nil {
		a.DeprecationMessage = s.Deprecated.DeprecationMessage
		deprecationMessage = s.Deprecated.computeDeprecatedDocumentation()
	}

	a.MarkdownDescription = genDataSourceAttrDescription(ctx, a.MarkdownDescription, deprecationMessage, a.Validators)
	return a
}
