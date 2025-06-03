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
	"context"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaE "github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type Attributes map[string]Attribute

func (a Attributes) process(ctx context.Context, s schemaType) any {
	switch s {
	case resourceT:
		attributes := make(map[string]schemaR.Attribute)

		for k, v := range a {
			if v.IsResource() {
				attributes[k] = v.GetResource(ctx)
			}
		}
		return attributes

	case dataSourceT:
		attributes := make(map[string]schemaD.Attribute)

		for k, v := range a {
			if v.IsDataSource() {
				attributes[k] = v.GetDataSource(ctx)
			}
		}
		return attributes

	case ephemeralT:
		attributes := make(map[string]schemaE.Attribute)
		for k, v := range a {
			if v.IsEphemeral() {
				attributes[k] = v.GetEphemeral(ctx).(schemaE.Attribute)
			}
		}
		return attributes
	}

	return nil
}

type Attribute interface {
	IsResource() bool
	IsDataSource() bool
	IsEphemeral() bool
	GetResource(ctx context.Context) schemaR.Attribute
	GetDataSource(ctx context.Context) schemaD.Attribute
	GetEphemeral(ctx context.Context) schemaE.Attribute
}
