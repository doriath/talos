// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package kubespan

import (
	"time"

	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta"
	"github.com/cosi-project/runtime/pkg/resource/typed"
	"inet.af/netaddr"
)

// PeerStatusType is type of PeerStatus resource.
const PeerStatusType = resource.Type("KubeSpanPeerStatuses.kubespan.talos.dev")

// PeerStatus the Wireguard peer state for KubeSpan.
//
// PeerStatus is identified by the public key.
type PeerStatus = typed.Resource[PeerStatusSpec, PeerStatusRD]

// PeerStatusSpec describes PeerStatus state.
//gotagsrewrite:gen
type PeerStatusSpec struct {
	// Active endpoint as seen by the Wireguard.
	Endpoint netaddr.IPPort `yaml:"endpoint" protobuf:"1"`
	// Label derived from the peer spec.
	Label string `yaml:"label" protobuf:"2"`
	// Calculated state.
	State PeerState `yaml:"state" protobuf:"3"`
	// Tx/Rx bytes.
	ReceiveBytes  int64 `yaml:"receiveBytes" protobuf:"4"`
	TransmitBytes int64 `yaml:"transmitBytes" protobuf:"5"`
	// Handshake.
	LastHandshakeTime time.Time `yaml:"lastHandshakeTime" protobuf:"6"`
	// Endpoint selection input.
	LastUsedEndpoint   netaddr.IPPort `yaml:"lastUsedEndpoint" protobuf:"7"`
	LastEndpointChange time.Time      `yaml:"lastEndpointChange" protobuf:"8"`
}

// NewPeerStatus initializes a PeerStatus resource.
func NewPeerStatus(namespace resource.Namespace, id resource.ID) *PeerStatus {
	return typed.NewResource[PeerStatusSpec, PeerStatusRD](
		resource.NewMetadata(namespace, PeerStatusType, id, resource.VersionUndefined),
		PeerStatusSpec{},
	)
}

// PeerStatusRD provides auxiliary methods for PeerStatus.
type PeerStatusRD struct{}

// ResourceDefinition implements typed.ResourceDefinition interface.
func (PeerStatusRD) ResourceDefinition(resource.Metadata, PeerStatusSpec) meta.ResourceDefinitionSpec {
	return meta.ResourceDefinitionSpec{
		Type:             PeerStatusType,
		Aliases:          []resource.Type{},
		DefaultNamespace: NamespaceName,
		PrintColumns: []meta.PrintColumn{
			{
				Name:     "Label",
				JSONPath: `{.label}`,
			},
			{
				Name:     "Endpoint",
				JSONPath: `{.endpoint}`,
			},
			{
				Name:     "State",
				JSONPath: `{.state}`,
			},
			{
				Name:     "Rx",
				JSONPath: `{.receiveBytes}`,
			},
			{
				Name:     "Tx",
				JSONPath: `{.transmitBytes}`,
			},
		},
	}
}
