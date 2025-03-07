// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package hardware

import (
	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta"
	"github.com/cosi-project/runtime/pkg/resource/typed"
)

// MemoryModuleType is type of MemoryModule resource.
const MemoryModuleType = resource.Type("MemoryModules.hardware.talos.dev")

// MemoryModule resource holds node MemoryModule information.
type MemoryModule = typed.Resource[MemoryModuleSpec, MemoryModuleRD]

// MemoryModuleSpec represents a single Memory.
//gotagsrewrite:gen
type MemoryModuleSpec struct {
	Size          uint32 `yaml:"sizeMiB,omitempty" protobuf:"1"`
	DeviceLocator string `yaml:"deviceLocator,omitempty" protobuf:"2"`
	BankLocator   string `yaml:"bankLocator,omitempty" protobuf:"3"`
	Speed         uint32 `yaml:"speed,omitempty" protobuf:"4"`
	Manufacturer  string `yaml:"manufacturer,omitempty" protobuf:"5"`
	SerialNumber  string `yaml:"serialNumber,omitempty" protobuf:"6"`
	AssetTag      string `yaml:"assetTag,omitempty" protobuf:"7"`
	ProductName   string `yaml:"productName,omitempty" protobuf:"8"`
}

// NewMemoryModuleInfo initializes a MemoryModuleInfo resource.
func NewMemoryModuleInfo(id string) *MemoryModule {
	return typed.NewResource[MemoryModuleSpec, MemoryModuleRD](
		resource.NewMetadata(NamespaceName, MemoryModuleType, id, resource.VersionUndefined),
		MemoryModuleSpec{},
	)
}

// MemoryModuleRD provides auxiliary methods for Memory info.
type MemoryModuleRD struct{}

// ResourceDefinition implements typed.ResourceDefinition interface.
func (c MemoryModuleRD) ResourceDefinition(resource.Metadata, MemoryModuleSpec) meta.ResourceDefinitionSpec {
	return meta.ResourceDefinitionSpec{
		Type: MemoryModuleType,
		Aliases: []resource.Type{
			"memorymodules",
			"ram",
		},
		DefaultNamespace: NamespaceName,
		PrintColumns: []meta.PrintColumn{
			{
				Name:     "Manufacturer",
				JSONPath: `{.manufacturer}`,
			},
			{
				Name:     "Model",
				JSONPath: `{.productName}`,
			},
			{
				Name:     "SizeMiB",
				JSONPath: `{.sizeMiB}`,
			},
		},
	}
}
