// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package perf

import (
	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta"
	"github.com/cosi-project/runtime/pkg/resource/typed"
)

//nolint:lll
//go:generate deep-copy -type CPUSpec -type MemorySpec -header-file ../../../../hack/boilerplate.txt -o deep_copy.generated.go .

// CPUType is type of Etcd resource.
const CPUType = resource.Type("CPUStats.perf.talos.dev")

// CPUID is a resource ID of singleton instance.
const CPUID = resource.ID("latest")

// CPU represents the last CPU stats snapshot.
type CPU = typed.Resource[CPUSpec, CPURD]

// CPUSpec represents the last CPU stats snapshot.
//gotagsrewrite:gen
type CPUSpec struct {
	CPU             []CPUStat `yaml:"cpu" protobuf:"1"`
	CPUTotal        CPUStat   `yaml:"cpuTotal" protobuf:"2"`
	IRQTotal        uint64    `yaml:"irqTotal" protobuf:"3"`
	ContextSwitches uint64    `yaml:"contextSwitches" protobuf:"4"`
	ProcessCreated  uint64    `yaml:"processCreated" protobuf:"5"`
	ProcessRunning  uint64    `yaml:"processRunning" protobuf:"6"`
	ProcessBlocked  uint64    `yaml:"processBlocked" protobuf:"7"`
	SoftIrqTotal    uint64    `yaml:"softIrqTotal" protobuf:"8"`
}

// CPUStat represents a single cpu stat.
//gotagsrewrite:gen
type CPUStat struct {
	User      float64 `yaml:"user" protobuf:"1"`
	Nice      float64 `yaml:"nice" protobuf:"2"`
	System    float64 `yaml:"system" protobuf:"3"`
	Idle      float64 `yaml:"idle" protobuf:"4"`
	Iowait    float64 `yaml:"iowait" protobuf:"5"`
	Irq       float64 `yaml:"irq" protobuf:"6"`
	SoftIrq   float64 `yaml:"softIrq" protobuf:"7"`
	Steal     float64 `yaml:"steal" protobuf:"8"`
	Guest     float64 `yaml:"guest" protobuf:"9"`
	GuestNice float64 `yaml:"guestNice" protobuf:"10"`
}

// NewCPU creates new default CPU stats object.
func NewCPU() *CPU {
	return typed.NewResource[CPUSpec, CPURD](
		resource.NewMetadata(NamespaceName, CPUType, CPUID, resource.VersionUndefined),
		CPUSpec{},
	)
}

// CPURD is an auxiliary type for CPU resource.
type CPURD struct{}

// ResourceDefinition implements meta.ResourceDefinitionProvider interface.
func (CPURD) ResourceDefinition(resource.Metadata, CPUSpec) meta.ResourceDefinitionSpec {
	return meta.ResourceDefinitionSpec{
		Type:             CPUType,
		Aliases:          []resource.Type{},
		DefaultNamespace: NamespaceName,
		PrintColumns: []meta.PrintColumn{
			{
				Name:     "User",
				JSONPath: "{.cpuTotal.user}",
			},
			{
				Name:     "System",
				JSONPath: "{.cpuTotal.system}",
			},
		},
	}
}
