// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cluster

import (
	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta"
	"github.com/cosi-project/runtime/pkg/resource/typed"
	"inet.af/netaddr"

	"github.com/talos-systems/talos/pkg/machinery/config/types/v1alpha1/machine"
)

//go:generate deep-copy -type AffiliateSpec -type ConfigSpec -type IdentitySpec -type MemberSpec -header-file ../../../../hack/boilerplate.txt -o deep_copy.generated.go .

// AffiliateType is type of Affiliate resource.
const AffiliateType = resource.Type("Affiliates.cluster.talos.dev")

// Affiliate resource holds information about cluster affiliate: it is discovered potential cluster member and/or KubeSpan peer.
//
// Controller builds local Affiliate structure for the node itself, other Affiliates are pulled from the registry during the discovery process.
type Affiliate = typed.Resource[AffiliateSpec, AffiliateRD]

// KubeSpanAffiliateSpec describes additional information specific for the KubeSpan.
type KubeSpanAffiliateSpec struct {
	PublicKey           string             `yaml:"publicKey"`
	Address             netaddr.IP         `yaml:"address"`
	AdditionalAddresses []netaddr.IPPrefix `yaml:"additionalAddresses"`
	Endpoints           []netaddr.IPPort   `yaml:"endpoints"`
}

// NewAffiliate initializes the Affiliate resource.
func NewAffiliate(namespace resource.Namespace, id resource.ID) *Affiliate {
	return typed.NewResource[AffiliateSpec, AffiliateRD](
		resource.NewMetadata(namespace, AffiliateType, id, resource.VersionUndefined),
		AffiliateSpec{},
	)
}

// AffiliateRD provides auxiliary methods for Affiliate.
type AffiliateRD struct{}

// ResourceDefinition implements typed.ResourceDefinition interface.
func (r AffiliateRD) ResourceDefinition(resource.Metadata, AffiliateSpec) meta.ResourceDefinitionSpec {
	return meta.ResourceDefinitionSpec{
		Type:             AffiliateType,
		Aliases:          []resource.Type{},
		DefaultNamespace: NamespaceName,
		PrintColumns: []meta.PrintColumn{
			{
				Name:     "Hostname",
				JSONPath: `{.hostname}`,
			},
			{
				Name:     "Machine Type",
				JSONPath: `{.machineType}`,
			},
			{
				Name:     "Addresses",
				JSONPath: `{.addresses}`,
			},
		},
	}
}

// AffiliateSpec describes Affiliate state.
//gotagsrewrite:gen
type AffiliateSpec struct {
	NodeID          string                `yaml:"nodeId" protobuf:"1"`
	Addresses       []netaddr.IP          `yaml:"addresses" protobuf:"2"`
	Hostname        string                `yaml:"hostname" protobuf:"3"`
	Nodename        string                `yaml:"nodename,omitempty" protobuf:"4"`
	OperatingSystem string                `yaml:"operatingSystem" protobuf:"5"`
	MachineType     machine.Type          `yaml:"machineType" protobuf:"6"`
	KubeSpan        KubeSpanAffiliateSpec `yaml:"kubespan,omitempty" protobuf:"7"`
}

// Merge two AffiliateSpecs.
//
//nolint:gocyclo
func (spec *AffiliateSpec) Merge(other *AffiliateSpec) {
	for _, addr := range other.Addresses {
		found := false

		for _, specAddr := range spec.Addresses {
			if addr == specAddr {
				found = true

				break
			}
		}

		if !found {
			spec.Addresses = append(spec.Addresses, addr)
		}
	}

	if other.Hostname != "" {
		spec.Hostname = other.Hostname
	}

	if other.Nodename != "" {
		spec.Nodename = other.Nodename
	}

	if other.MachineType != machine.TypeUnknown {
		spec.MachineType = other.MachineType
	}

	if other.KubeSpan.PublicKey != "" {
		spec.KubeSpan.PublicKey = other.KubeSpan.PublicKey
	}

	if !other.KubeSpan.Address.IsZero() {
		spec.KubeSpan.Address = other.KubeSpan.Address
	}

	for _, addr := range other.KubeSpan.AdditionalAddresses {
		found := false

		for _, specAddr := range spec.KubeSpan.AdditionalAddresses {
			if addr == specAddr {
				found = true

				break
			}
		}

		if !found {
			spec.KubeSpan.AdditionalAddresses = append(spec.KubeSpan.AdditionalAddresses, addr)
		}
	}

	for _, addr := range other.KubeSpan.Endpoints {
		found := false

		for _, specAddr := range spec.KubeSpan.Endpoints {
			if addr == specAddr {
				found = true

				break
			}
		}

		if !found {
			spec.KubeSpan.Endpoints = append(spec.KubeSpan.Endpoints, addr)
		}
	}
}
