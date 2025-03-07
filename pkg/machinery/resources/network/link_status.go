// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package network

import (
	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta"
	"github.com/cosi-project/runtime/pkg/resource/typed"

	"github.com/talos-systems/talos/pkg/machinery/nethelpers"
)

// LinkStatusType is type of LinkStatus resource.
const LinkStatusType = resource.Type("LinkStatuses.net.talos.dev")

// LinkStatus resource holds physical network link status.
type LinkStatus = typed.Resource[LinkStatusSpec, LinkStatusRD]

// LinkStatusSpec describes status of rendered secrets.
//gotagsrewrite:gen
type LinkStatusSpec struct {
	// Fields coming from rtnetlink API.
	Index            uint32                      `yaml:"index" protobuf:"1"`
	Type             nethelpers.LinkType         `yaml:"type" protobuf:"2"`
	LinkIndex        uint32                      `yaml:"linkIndex" protobuf:"3"`
	Flags            nethelpers.LinkFlags        `yaml:"flags" protobuf:"4"`
	HardwareAddr     nethelpers.HardwareAddr     `yaml:"hardwareAddr" protobuf:"5"`
	BroadcastAddr    nethelpers.HardwareAddr     `yaml:"broadcastAddr" protobuf:"6"`
	MTU              uint32                      `yaml:"mtu" protobuf:"7"`
	QueueDisc        string                      `yaml:"queueDisc" protobuf:"8"`
	MasterIndex      uint32                      `yaml:"masterIndex,omitempty" protobuf:"9"`
	OperationalState nethelpers.OperationalState `yaml:"operationalState" protobuf:"10"`
	Kind             string                      `yaml:"kind" protobuf:"11"`
	SlaveKind        string                      `yaml:"slaveKind" protobuf:"12"`
	BusPath          string                      `yaml:"busPath,omitempty" protobuf:"13"`
	PCIID            string                      `yaml:"pciID,omitempty" protobuf:"14"`
	Driver           string                      `yaml:"driver,omitempty" protobuf:"15"`
	DriverVersion    string                      `yaml:"driverVersion,omitempty" protobuf:"16"`
	FirmwareVersion  string                      `yaml:"firmwareVersion,omitempty" protobuf:"17"`
	ProductID        string                      `yaml:"productID,omitempty" protobuf:"18"`
	VendorID         string                      `yaml:"vendorID,omitempty" protobuf:"19"`
	Product          string                      `yaml:"product,omitempty" protobuf:"20"`
	Vendor           string                      `yaml:"vendor,omitempty" protobuf:"21"`
	// Fields coming from ethtool API.
	LinkState     bool              `yaml:"linkState" protobuf:"22"`
	SpeedMegabits int               `yaml:"speedMbit,omitempty" protobuf:"23"`
	Port          nethelpers.Port   `yaml:"port" protobuf:"24"`
	Duplex        nethelpers.Duplex `yaml:"duplex" protobuf:"25"`
	// Following fields are only populated with respective Kind.
	VLAN         VLANSpec         `yaml:"vlan,omitempty" protobuf:"26"`
	BridgeMaster BridgeMasterSpec `yaml:"bridgeMaster,omitempty" protobuf:"27"`
	BondMaster   BondMasterSpec   `yaml:"bondMaster,omitempty" protobuf:"28"`
	Wireguard    WireguardSpec    `yaml:"wireguard,omitempty" protobuf:"29"`
}

// Physical checks if the link is physical ethernet.
func (s LinkStatusSpec) Physical() bool {
	return s.Type == nethelpers.LinkEther && s.Kind == ""
}

// NewLinkStatus initializes a LinkStatus resource.
func NewLinkStatus(namespace resource.Namespace, id resource.ID) *LinkStatus {
	return typed.NewResource[LinkStatusSpec, LinkStatusRD](
		resource.NewMetadata(namespace, LinkStatusType, id, resource.VersionUndefined),
		LinkStatusSpec{},
	)
}

// LinkStatusRD provides auxiliary methods for LinkStatus.
type LinkStatusRD struct{}

// ResourceDefinition implements typed.ResourceDefinition interface.
func (LinkStatusRD) ResourceDefinition(resource.Metadata, LinkStatusSpec) meta.ResourceDefinitionSpec {
	return meta.ResourceDefinitionSpec{
		Type:             LinkStatusType,
		Aliases:          []resource.Type{"link", "links"},
		DefaultNamespace: NamespaceName,
		PrintColumns: []meta.PrintColumn{
			{
				Name:     "Type",
				JSONPath: `{.type}`,
			},
			{
				Name:     "Kind",
				JSONPath: `{.kind}`,
			},
			{
				Name:     "Hw Addr",
				JSONPath: `{.hardwareAddr}`,
			},
			{
				Name:     "Oper State",
				JSONPath: `{.operationalState}`,
			},
			{
				Name:     "Link State",
				JSONPath: `{.linkState}`,
			},
		},
		Sensitivity: meta.NonSensitive,
	}
}
