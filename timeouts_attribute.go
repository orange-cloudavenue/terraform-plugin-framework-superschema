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

	timeoutsD "github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	timeoutsE "github.com/hashicorp/terraform-plugin-framework-timeouts/ephemeral/timeouts"
	timeoutsR "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
)

var _ Attribute = TimeoutAttribute{}

type ResourceTimeoutAttribute struct {
	Create bool
	Read   bool
	Delete bool
	Update bool
}

type DatasourceTimeoutAttribute struct {
	Read bool
}
type TimeoutAttribute struct {
	Resource   *ResourceTimeoutAttribute
	DataSource *DatasourceTimeoutAttribute
	Ephemeral  *ResourceTimeoutAttribute
}

// IsResource returns true if the attribute is a resource attribute.
func (s TimeoutAttribute) IsResource() bool {
	return s.Resource != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s TimeoutAttribute) IsDataSource() bool {
	return s.DataSource != nil
}

// IsEphemeral returns true if the attribute is a ephemeral attribute.
func (s TimeoutAttribute) IsEphemeral() bool {
	return s.Ephemeral != nil
}

func (s TimeoutAttribute) GetResource(ctx context.Context) schemaR.Attribute {
	var a schemaR.Attribute

	if s.Resource != nil {
		a = timeoutsR.Attributes(ctx, timeoutsR.Opts{
			Create: s.Resource.Create,
			Read:   s.Resource.Read,
			Delete: s.Resource.Delete,
			Update: s.Resource.Update,
		})
	}
	return a
}

func (s TimeoutAttribute) GetDataSource(ctx context.Context) schemaD.Attribute {
	var a schemaD.Attribute

	if s.DataSource != nil && s.DataSource.Read {
		a = timeoutsD.Attributes(ctx)
	}
	return a
}

func (s TimeoutAttribute) GetEphemeral(ctx context.Context) schemaE.Attribute {
	var a schemaE.Attribute

	if s.Ephemeral != nil && s.Ephemeral.Read {
		a = timeoutsE.Attributes(ctx)
	}
	return a
}
