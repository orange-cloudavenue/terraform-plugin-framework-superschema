/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package superschema

import (
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaE "github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type superAttribute interface {
	schemaR.Attribute
	schemaD.Attribute
	schemaE.Attribute
}

func computeIsComputed[X, T superAttribute](common X, resourceOrDatasource T) bool {
	return common.IsComputed() || resourceOrDatasource.IsComputed()
}

func computeIsOptional[X, T superAttribute](common X, resourceOrDatasource T) bool {
	return common.IsOptional() || resourceOrDatasource.IsOptional()
}

func computeIsRequired[X, T superAttribute](common X, resourceOrDatasource T) bool {
	return common.IsRequired() || resourceOrDatasource.IsRequired()
}

func computeIsSensitive[X, T superAttribute](common X, resourceOrDatasource T) bool {
	return common.IsSensitive() || resourceOrDatasource.IsSensitive()
}

func computeIsWriteOnly[X, T superAttribute](common X, resourceOrDatasource T) bool {
	return common.IsWriteOnly() || resourceOrDatasource.IsWriteOnly()
}

func computeMarkdownDescription[X, T superAttribute](common X, resourceOrDatasource T) string {
	return addToDescription(common.GetMarkdownDescription(), resourceOrDatasource.GetMarkdownDescription())
}

func computeDescription[X, T superAttribute](common X, resourceOrDatasource T) string {
	return addToDescription(common.GetDescription(), resourceOrDatasource.GetDescription())
}

func computeDeprecationMessage[X, T superAttribute](common X, resourceOrDatasource T) string {
	return addToDescription(common.GetDeprecationMessage(), resourceOrDatasource.GetDeprecationMessage())
}
