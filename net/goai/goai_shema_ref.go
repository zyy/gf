// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package goai

import (
	"reflect"

	"github.com/gogf/gf/v2/internal/json"
)

type SchemaRefs []SchemaRef

type SchemaRef struct {
	Ref   string
	Value *Schema
}

func (oai *OpenApiV3) newSchemaRefWithGolangType(golangType reflect.Type, tagMap map[string]string) (*SchemaRef, error) {
	var (
		oaiType   = oai.golangTypeToOAIType(golangType)
		oaiFormat = oai.golangTypeToOAIFormat(golangType)
		schemaRef = &SchemaRef{}
		schema    = &Schema{
			Type:        oaiType,
			Format:      oaiFormat,
			XExtensions: make(XExtensions),
		}
	)
	if len(tagMap) > 0 {
		if err := oai.tagMapToSchema(tagMap, schema); err != nil {
			return nil, err
		}
	}
	schemaRef.Value = schema
	switch oaiType {
	case
		TypeInteger,
		TypeNumber,
		TypeString,
		TypeBoolean:
		// Nothing to do.

	case
		TypeArray:
		subSchemaRef, err := oai.newSchemaRefWithGolangType(golangType.Elem(), nil)
		if err != nil {
			return nil, err
		}
		schema.Items = subSchemaRef
		if len(schema.Enum) > 0 {
			schema.Items.Value.Enum = schema.Enum
			schema.Enum = nil
		}

	case
		TypeObject:
		for golangType.Kind() == reflect.Ptr {
			golangType = golangType.Elem()
		}
		switch golangType.Kind() {
		case reflect.Map:
			// Specially for map type.
			subSchemaRef, err := oai.newSchemaRefWithGolangType(golangType.Elem(), nil)
			if err != nil {
				return nil, err
			}
			schema.AdditionalProperties = subSchemaRef
			return schemaRef, nil

		case reflect.Interface:
			// Specially for interface type.
			var (
				structTypeName = oai.golangTypeToSchemaName(golangType)
			)
			if oai.Components.Schemas.Get(structTypeName) == nil {
				if err := oai.addSchema(reflect.New(golangType).Interface()); err != nil {
					return nil, err
				}
			}
			schemaRef.Ref = structTypeName
			schemaRef.Value = nil

		default:
			// Normal struct object.
			var structTypeName = oai.golangTypeToSchemaName(golangType)
			if oai.Components.Schemas.Get(structTypeName) == nil {
				if err := oai.addSchema(reflect.New(golangType).Elem().Interface()); err != nil {
					return nil, err
				}
			}
			schemaRef.Ref = structTypeName
			schemaRef.Value = nil
		}
	}
	return schemaRef, nil
}

func (r SchemaRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}
