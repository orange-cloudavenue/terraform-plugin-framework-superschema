/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

// code generated by go generate - look at type_attribute.go.tmpl for source file
package superschema

import (
	"context"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ Attribute = Int32Attribute{}

type Int32Attribute struct {
	Deprecated *Deprecated
	Common     *schemaR.Int32Attribute
	Resource   *schemaR.Int32Attribute
	DataSource *schemaD.Int32Attribute
}

// IsResource returns true if the attribute is a resource attribute.
func (s Int32Attribute) IsResource() bool {
	return s.Resource != nil || s.Common != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s Int32Attribute) IsDataSource() bool {
	return s.DataSource != nil || s.Common != nil
}

//nolint:dupl
func (s Int32Attribute) GetResource(ctx context.Context) schemaR.Attribute {
	var (
		common   schemaR.Int32Attribute
		resource schemaR.Int32Attribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.Resource != nil {
		resource = *s.Resource
	}

	a := schemaR.Int32Attribute{
		Required:            computeIsRequired(common, resource),
		Optional:            computeIsOptional(common, resource),
		Computed:            computeIsComputed(common, resource),
		Sensitive:           computeIsSensitive(common, resource),
		WriteOnly:           computeIsWriteOnly(common, resource),
		MarkdownDescription: computeMarkdownDescription(common, resource),
		Description:         computeDescription(common, resource),
		DeprecationMessage:  computeDeprecationMessage(common, resource),
	}

	a.Validators = append(a.Validators, common.Validators...)
	a.Validators = append(a.Validators, resource.Validators...)
	a.PlanModifiers = append(a.PlanModifiers, common.PlanModifiers...)
	a.PlanModifiers = append(a.PlanModifiers, resource.PlanModifiers...)

	defaultVDescription := ""
	if s.Resource != nil {
		if s.Resource.Default != nil {
			a.Default = s.Resource.Default
			defaultVDescription = s.Resource.Default.MarkdownDescription(ctx)
		}
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
func (s Int32Attribute) GetDataSource(ctx context.Context) schemaD.Attribute {
	var (
		common     schemaR.Int32Attribute
		dataSource schemaD.Int32Attribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.DataSource != nil {
		dataSource = *s.DataSource
	}

	a := schemaD.Int32Attribute{
		Required:            computeIsRequired(common, dataSource),
		Optional:            computeIsOptional(common, dataSource),
		Computed:            computeIsComputed(common, dataSource),
		Sensitive:           computeIsSensitive(common, dataSource),
		MarkdownDescription: computeMarkdownDescription(common, dataSource),
		Description:         computeDescription(common, dataSource),
		DeprecationMessage:  computeDeprecationMessage(common, dataSource),
	}

	a.Validators = append(a.Validators, common.Validators...)
	a.Validators = append(a.Validators, dataSource.Validators...)

	deprecationMessage := ""
	if s.Deprecated != nil {
		a.DeprecationMessage = s.Deprecated.DeprecationMessage
		deprecationMessage = s.Deprecated.computeDeprecatedDocumentation()
	}

	a.MarkdownDescription = genDataSourceAttrDescription(ctx, a.MarkdownDescription, deprecationMessage, a.Validators)
	return a
}
