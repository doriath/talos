// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package kubespan

import (
	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta"
	"github.com/cosi-project/runtime/pkg/resource/typed"
	"inet.af/netaddr"
)

// PeerSpecType is type of PeerSpec resource.
const PeerSpecType = resource.Type("KubeSpanPeerSpecs.kubespan.talos.dev")

// PeerSpec is produced from cluster.Affiliate which has KubeSpan information attached.
//
// PeerSpec is identified by the public key.
type PeerSpec = typed.Resource[PeerSpecSpec, PeerSpecRD]

// PeerSpecSpec describes PeerSpec state.
//gotagsrewrite:gen
type PeerSpecSpec struct {
	Address    netaddr.IP         `yaml:"address" protobuf:"1"`
	AllowedIPs []netaddr.IPPrefix `yaml:"allowedIPs" protobuf:"2"`
	Endpoints  []netaddr.IPPort   `yaml:"endpoints" protobuf:"3"`
	Label      string             `yaml:"label" protobuf:"4"`
}

// NewPeerSpec initializes a PeerSpec resource.
func NewPeerSpec(namespace resource.Namespace, id resource.ID) *PeerSpec {
	return typed.NewResource[PeerSpecSpec, PeerSpecRD](
		resource.NewMetadata(namespace, PeerSpecType, id, resource.VersionUndefined),
		PeerSpecSpec{},
	)
}

// PeerSpecRD provides auxiliary methods for PeerSpec.
type PeerSpecRD struct{}

// ResourceDefinition implements typed.ResourceDefinition interface.
func (PeerSpecRD) ResourceDefinition(resource.Metadata, PeerSpecSpec) meta.ResourceDefinitionSpec {
	return meta.ResourceDefinitionSpec{
		Type:             PeerSpecType,
		Aliases:          []resource.Type{},
		DefaultNamespace: NamespaceName,
		PrintColumns: []meta.PrintColumn{
			{
				Name:     "Label",
				JSONPath: `{.label}`,
			},
			{
				Name:     "Endpoints",
				JSONPath: `{.endpoints}`,
			},
		},
	}
}
