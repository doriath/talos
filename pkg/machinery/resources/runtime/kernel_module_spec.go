// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package runtime

import (
	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta"
	"github.com/cosi-project/runtime/pkg/resource/typed"
)

// KernelModuleSpecType is type of KernelModuleSpec resource.
const KernelModuleSpecType = resource.Type("KernelModuleSpecs.runtime.talos.dev")

// KernelModuleSpec resource holds information about Linux kernel module to load.
type KernelModuleSpec = typed.Resource[KernelModuleSpecSpec, KernelModuleSpecRD]

// KernelModuleSpecSpec describes Linux kernel module to load.
//gotagsrewrite:gen
type KernelModuleSpecSpec struct {
	Name string `yaml:"string" protobuf:"1"`
	// more options in the future: args, aliases, etc.
}

// NewKernelModuleSpec initializes a KernelModuleSpec resource.
func NewKernelModuleSpec(namespace resource.Namespace, id resource.ID) *KernelModuleSpec {
	return typed.NewResource[KernelModuleSpecSpec, KernelModuleSpecRD](
		resource.NewMetadata(namespace, KernelModuleSpecType, id, resource.VersionUndefined),
		KernelModuleSpecSpec{},
	)
}

// KernelModuleSpecRD is auxiliary resource data for KernelModuleSpec.
type KernelModuleSpecRD struct{}

// ResourceDefinition implements meta.ResourceDefinitionProvider interface.
func (KernelModuleSpecRD) ResourceDefinition(resource.Metadata, KernelModuleSpecSpec) meta.ResourceDefinitionSpec {
	return meta.ResourceDefinitionSpec{
		Type:             KernelModuleSpecType,
		Aliases:          []resource.Type{"modules"},
		DefaultNamespace: NamespaceName,
	}
}
