// code generated by go generate - look at supertype_attribute.go.tmpl for source file
package superschema

import (
	"context"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/orange-cloudavenue/terraform-plugin-framework-supertypes"
)

var _ Attribute = SuperMapNestedAttribute{}

type SuperMapNestedAttribute struct {
	Deprecated *Deprecated
	Common     *schemaR.MapNestedAttribute
	Resource   *schemaR.MapNestedAttribute
	DataSource *schemaD.MapNestedAttribute
	Attributes Attributes
}

// IsResource returns true if the attribute is a resource attribute.
func (s SuperMapNestedAttribute) IsResource() bool {
	return s.Resource != nil || s.Common != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s SuperMapNestedAttribute) IsDataSource() bool {
	return s.DataSource != nil || s.Common != nil
}

// GetCustomType returns the custom type of the attribute.
func (s SuperMapNestedAttribute) getCustomType(oB basetypes.ObjectTypable) basetypes.MapTypable {
	return supertypes.MapNestedType{
		MapType: basetypes.MapType{
			ElemType: oB,
		},
	}
}

//nolint:dupl
func (s SuperMapNestedAttribute) GetResource(ctx context.Context) schemaR.Attribute {
	var (
		common   schemaR.MapNestedAttribute
		resource schemaR.MapNestedAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.Resource != nil {
		resource = *s.Resource
	}

	a := schemaR.MapNestedAttribute{
		Required:            computeIsRequired(common, resource),
		Optional:            computeIsOptional(common, resource),
		Computed:            computeIsComputed(common, resource),
		Sensitive:           computeIsSensitive(common, resource),
		MarkdownDescription: computeMarkdownDescription(common, resource),
		Description:         computeDescription(common, resource),
		DeprecationMessage:  computeDeprecationMessage(common, resource),
		NestedObject: schemaR.NestedAttributeObject{
			Attributes: s.Attributes.process(ctx, resourceT).(map[string]schemaR.Attribute),
		},
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
		a.CustomType = s.getCustomType(a.NestedObject.Type()).(supertypes.MapNestedType)
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
func (s SuperMapNestedAttribute) GetDataSource(ctx context.Context) schemaD.Attribute {
	var (
		common     schemaR.MapNestedAttribute
		dataSource schemaD.MapNestedAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.DataSource != nil {
		dataSource = *s.DataSource
	}

	a := schemaD.MapNestedAttribute{
		Required:            computeIsRequired(common, dataSource),
		Optional:            computeIsOptional(common, dataSource),
		Computed:            computeIsComputed(common, dataSource),
		Sensitive:           computeIsSensitive(common, dataSource),
		MarkdownDescription: computeMarkdownDescription(common, dataSource),
		Description:         computeDescription(common, dataSource),
		DeprecationMessage:  computeDeprecationMessage(common, dataSource),
		NestedObject: schemaD.NestedAttributeObject{
			Attributes: s.Attributes.process(ctx, dataSourceT).(map[string]schemaD.Attribute),
		},
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
		a.CustomType = s.getCustomType(a.NestedObject.Type()).(supertypes.MapNestedType)
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

var _ Attribute = SuperMapNestedAttributeOf[struct{}]{}

type SuperMapNestedAttributeOf[T any] struct {
	Deprecated *Deprecated
	Common     *schemaR.MapNestedAttribute
	Resource   *schemaR.MapNestedAttribute
	DataSource *schemaD.MapNestedAttribute
	Attributes Attributes
}

// IsResource returns true if the attribute is a resource attribute.
func (s SuperMapNestedAttributeOf[T]) IsResource() bool {
	return s.Resource != nil || s.Common != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s SuperMapNestedAttributeOf[T]) IsDataSource() bool {
	return s.DataSource != nil || s.Common != nil
}

//nolint:dupl
func (s SuperMapNestedAttributeOf[T]) GetResource(ctx context.Context) schemaR.Attribute {
	var (
		common   schemaR.MapNestedAttribute
		resource schemaR.MapNestedAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.Resource != nil {
		resource = *s.Resource
	}

	a := schemaR.MapNestedAttribute{
		Required:            computeIsRequired(common, resource),
		Optional:            computeIsOptional(common, resource),
		Computed:            computeIsComputed(common, resource),
		Sensitive:           computeIsSensitive(common, resource),
		MarkdownDescription: computeMarkdownDescription(common, resource),
		Description:         computeDescription(common, resource),
		DeprecationMessage:  computeDeprecationMessage(common, resource),
		NestedObject: schemaR.NestedAttributeObject{
			Attributes: s.Attributes.process(ctx, resourceT).(map[string]schemaR.Attribute),
		},
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
		a.CustomType = supertypes.NewMapNestedObjectTypeOf[T](ctx)
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
func (s SuperMapNestedAttributeOf[T]) GetDataSource(ctx context.Context) schemaD.Attribute {
	var (
		common     schemaR.MapNestedAttribute
		dataSource schemaD.MapNestedAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.DataSource != nil {
		dataSource = *s.DataSource
	}

	a := schemaD.MapNestedAttribute{
		Required:            computeIsRequired(common, dataSource),
		Optional:            computeIsOptional(common, dataSource),
		Computed:            computeIsComputed(common, dataSource),
		Sensitive:           computeIsSensitive(common, dataSource),
		MarkdownDescription: computeMarkdownDescription(common, dataSource),
		Description:         computeDescription(common, dataSource),
		DeprecationMessage:  computeDeprecationMessage(common, dataSource),
		NestedObject: schemaD.NestedAttributeObject{
			Attributes: s.Attributes.process(ctx, dataSourceT).(map[string]schemaD.Attribute),
		},
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
		a.CustomType = supertypes.NewMapNestedObjectTypeOf[T](ctx)
	}

	deprecationMessage := ""
	if s.Deprecated != nil {
		a.DeprecationMessage = s.Deprecated.DeprecationMessage
		deprecationMessage = s.Deprecated.computeDeprecatedDocumentation()
	}

	a.MarkdownDescription = genDataSourceAttrDescription(ctx, a.MarkdownDescription, deprecationMessage, a.Validators)
	return a
}
