// code generated by go generate - look at supertype_attribute.go.tmpl for source file
package superschema

import (
	"context"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/orange-cloudavenue/terraform-plugin-framework-supertypes"
)

var _ Attribute = SuperNumberAttribute{}

type SuperNumberAttribute struct {
	Deprecated *Deprecated
	Common     *schemaR.NumberAttribute
	Resource   *schemaR.NumberAttribute
	DataSource *schemaD.NumberAttribute
}

// IsResource returns true if the attribute is a resource attribute.
func (s SuperNumberAttribute) IsResource() bool {
	return s.Resource != nil || s.Common != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s SuperNumberAttribute) IsDataSource() bool {
	return s.DataSource != nil || s.Common != nil
}

// GetCustomType returns the custom type of the attribute.
func (s SuperNumberAttribute) getCustomType() basetypes.NumberTypable {
	return supertypes.NumberType{
		NumberType: basetypes.NumberType{},
	}
}

//nolint:dupl
func (s SuperNumberAttribute) GetResource(ctx context.Context) schemaR.Attribute {
	var (
		common   schemaR.NumberAttribute
		resource schemaR.NumberAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.Resource != nil {
		resource = *s.Resource
	}

	a := schemaR.NumberAttribute{
		Required:            computeIsRequired(common, resource),
		Optional:            computeIsOptional(common, resource),
		Computed:            computeIsComputed(common, resource),
		Sensitive:           computeIsSensitive(common, resource),
		MarkdownDescription: computeMarkdownDescription(common, resource),
		Description:         computeDescription(common, resource),
		DeprecationMessage:  computeDeprecationMessage(common, resource),
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
		a.CustomType = s.getCustomType().(supertypes.NumberType)
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
func (s SuperNumberAttribute) GetDataSource(ctx context.Context) schemaD.Attribute {
	var (
		common     schemaR.NumberAttribute
		dataSource schemaD.NumberAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.DataSource != nil {
		dataSource = *s.DataSource
	}

	a := schemaD.NumberAttribute{
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
		a.CustomType = s.getCustomType().(supertypes.NumberType)
	}

	deprecationMessage := ""
	if s.Deprecated != nil {
		a.DeprecationMessage = s.Deprecated.DeprecationMessage
		deprecationMessage = s.Deprecated.computeDeprecatedDocumentation()
	}

	a.MarkdownDescription = genDataSourceAttrDescription(ctx, a.MarkdownDescription, deprecationMessage, a.Validators)
	return a
}
