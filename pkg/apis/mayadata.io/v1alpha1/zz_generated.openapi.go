// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/mayadata-io/gittrack/pkg/apis/mayadata.io/v1alpha1.GitTrack":       schema_pkg_apis_mayadata_v1alpha1_GitTrack(ref),
		"github.com/mayadata-io/gittrack/pkg/apis/mayadata.io/v1alpha1.GitTrackSpec":   schema_pkg_apis_mayadata_v1alpha1_GitTrackSpec(ref),
		"github.com/mayadata-io/gittrack/pkg/apis/mayadata.io/v1alpha1.GitTrackStatus": schema_pkg_apis_mayadata_v1alpha1_GitTrackStatus(ref),
	}
}

func schema_pkg_apis_mayadata_v1alpha1_GitTrack(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GitTrack is the Schema for the gittracks API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mayadata-io/gittrack/pkg/apis/mayadata.io/v1alpha1.GitTrackSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mayadata-io/gittrack/pkg/apis/mayadata.io/v1alpha1.GitTrackStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/mayadata-io/gittrack/pkg/apis/mayadata.io/v1alpha1.GitTrackSpec", "github.com/mayadata-io/gittrack/pkg/apis/mayadata.io/v1alpha1.GitTrackStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_mayadata_v1alpha1_GitTrackSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GitTrackSpec defines the desired state of GitTrack",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_mayadata_v1alpha1_GitTrackStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GitTrackStatus defines the observed state of GitTrack",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
