// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package network

import (
	"sort"
	"time"

	"inet.af/netaddr"

	"github.com/talos-systems/talos/pkg/machinery/nethelpers"
)

// VLANSpec describes VLAN settings if Kind == "vlan".
//gotagsrewrite:gen
type VLANSpec struct {
	// VID is the vlan ID.
	VID uint16 `yaml:"vlanID" protobuf:"1"`

	// Protocol is the vlan protocol.
	Protocol nethelpers.VLANProtocol `yaml:"vlanProtocol" protobuf:"2"`
}

// BondMasterSpec describes bond settings if Kind == "bond".
//gotagsrewrite:gen
type BondMasterSpec struct {
	Mode            nethelpers.BondMode           `yaml:"mode" protobuf:"1"`
	HashPolicy      nethelpers.BondXmitHashPolicy `yaml:"xmitHashPolicy" protobuf:"2"`
	LACPRate        nethelpers.LACPRate           `yaml:"lacpRate" protobuf:"3"`
	ARPValidate     nethelpers.ARPValidate        `yaml:"arpValidate" protobuf:"4"`
	ARPAllTargets   nethelpers.ARPAllTargets      `yaml:"arpAllTargets" protobuf:"5"`
	PrimaryIndex    uint32                        `yaml:"primary,omitempty" protobuf:"6"`
	PrimaryReselect nethelpers.PrimaryReselect    `yaml:"primaryReselect" protobuf:"7"`
	FailOverMac     nethelpers.FailOverMAC        `yaml:"failOverMac" protobuf:"8"`
	ADSelect        nethelpers.ADSelect           `yaml:"adSelect,omitempty" protobuf:"9"`
	MIIMon          uint32                        `yaml:"miimon,omitempty" protobuf:"10"`
	UpDelay         uint32                        `yaml:"updelay,omitempty" protobuf:"11"`
	DownDelay       uint32                        `yaml:"downdelay,omitempty" protobuf:"12"`
	ARPInterval     uint32                        `yaml:"arpInterval,omitempty" protobuf:"13"`
	ResendIGMP      uint32                        `yaml:"resendIgmp,omitempty" protobuf:"14"`
	MinLinks        uint32                        `yaml:"minLinks,omitempty" protobuf:"15"`
	LPInterval      uint32                        `yaml:"lpInterval,omitempty" protobuf:"16"`
	PacketsPerSlave uint32                        `yaml:"packetsPerSlave,omitempty" protobuf:"17"`
	NumPeerNotif    uint8                         `yaml:"numPeerNotif,omitempty" protobuf:"18"`
	TLBDynamicLB    uint8                         `yaml:"tlbLogicalLb,omitempty" protobuf:"19"`
	AllSlavesActive uint8                         `yaml:"allSlavesActive,omitempty" protobuf:"20"`
	UseCarrier      bool                          `yaml:"useCarrier,omitempty" protobuf:"21"`
	ADActorSysPrio  uint16                        `yaml:"adActorSysPrio,omitempty" protobuf:"22"`
	ADUserPortKey   uint16                        `yaml:"adUserPortKey,omitempty" protobuf:"23"`
	PeerNotifyDelay uint32                        `yaml:"peerNotifyDelay,omitempty" protobuf:"24"`
}

// BridgeMasterSpec describes bridge settings if Kind == "bridge".
type BridgeMasterSpec struct {
	STP STPSpec `yaml:"stp,omitempty"`
}

// STPSpec describes Spanning Tree Protocol (STP) settings of a bridge.
type STPSpec struct {
	Enabled bool `yaml:"enabled"`
}

// WireguardSpec describes Wireguard settings if Kind == "wireguard".
type WireguardSpec struct {
	// PrivateKey is used to configure the link, present only in the LinkSpec.
	PrivateKey string `yaml:"privateKey,omitempty"`
	// PublicKey is only used in LinkStatus to show the link status.
	PublicKey    string          `yaml:"publicKey,omitempty"`
	ListenPort   int             `yaml:"listenPort"`
	FirewallMark int             `yaml:"firewallMark"`
	Peers        []WireguardPeer `yaml:"peers"`
}

// WireguardPeer describes a single peer.
type WireguardPeer struct {
	PublicKey                   string             `yaml:"publicKey"`
	PresharedKey                string             `yaml:"presharedKey"`
	Endpoint                    string             `yaml:"endpoint"`
	PersistentKeepaliveInterval time.Duration      `yaml:"persistentKeepaliveInterval"`
	AllowedIPs                  []netaddr.IPPrefix `yaml:"allowedIPs"`
}

// ID Returns the VID for type VLANSpec.
func (vlan VLANSpec) ID() uint16 {
	return vlan.VID
}

// MTU Returns MTU=0 for type VLANSpec.
func (vlan VLANSpec) MTU() uint32 {
	return 0
}

// Equal checks two WireguardPeer structs for equality.
//
// `spec` is considered to be the result of getting current Wireguard configuration,
// while `other` is the new (updated configuration).
func (peer *WireguardPeer) Equal(other *WireguardPeer) bool {
	if peer.PublicKey != other.PublicKey {
		return false
	}

	if peer.PresharedKey != other.PresharedKey {
		return false
	}

	// if the Endpoint is not set in `other`, don't consider this to be a change
	if other.Endpoint != "" && peer.Endpoint != other.Endpoint {
		return false
	}

	if peer.PersistentKeepaliveInterval != other.PersistentKeepaliveInterval {
		return false
	}

	if len(peer.AllowedIPs) != len(other.AllowedIPs) {
		return false
	}

	for i := range peer.AllowedIPs {
		if peer.AllowedIPs[i].IP().Compare(other.AllowedIPs[i].IP()) != 0 {
			return false
		}

		if peer.AllowedIPs[i].Bits() != other.AllowedIPs[i].Bits() {
			return false
		}
	}

	return true
}

// IsZero checks if the WireguardSpec is zero value.
func (spec *WireguardSpec) IsZero() bool {
	return spec.PrivateKey == "" && spec.ListenPort == 0 && spec.FirewallMark == 0 && len(spec.Peers) == 0
}

// Equal checks two WireguardSpecs for equality.
//
// Both specs should be sorted before calling this method.
//
// `spec` is considered to be the result of getting current Wireguard configuration,
// while `other` is the new (updated configuration).
func (spec *WireguardSpec) Equal(other *WireguardSpec) bool {
	if spec.PrivateKey != other.PrivateKey {
		return false
	}

	// listenPort of '0' means use any available port, so we shouldn't consider this to be a "change"
	if spec.ListenPort != other.ListenPort && other.ListenPort != 0 {
		return false
	}

	if spec.FirewallMark != other.FirewallMark {
		return false
	}

	if len(spec.Peers) != len(other.Peers) {
		return false
	}

	for i := range spec.Peers {
		if !spec.Peers[i].Equal(&other.Peers[i]) {
			return false
		}
	}

	return true
}

// Sort the spec so that comparison is possible.
func (spec *WireguardSpec) Sort() {
	sort.Slice(spec.Peers, func(i, j int) bool {
		return spec.Peers[i].PublicKey < spec.Peers[j].PublicKey
	})

	for k := range spec.Peers {
		k := k

		sort.Slice(spec.Peers[k].AllowedIPs, func(i, j int) bool {
			left := spec.Peers[k].AllowedIPs[i]
			right := spec.Peers[k].AllowedIPs[j]

			switch left.IP().Compare(right.IP()) {
			case -1:
				return true
			case 0:
				return left.Bits() < right.Bits()
			default:
				return false
			}
		})
	}
}

// Merge with other Wireguard spec overwriting non-zero values.
func (spec *WireguardSpec) Merge(other WireguardSpec) {
	if other.ListenPort != 0 {
		spec.ListenPort = other.ListenPort
	}

	if other.FirewallMark != 0 {
		spec.FirewallMark = other.FirewallMark
	}

	if other.PrivateKey != "" {
		spec.PrivateKey = other.PrivateKey
	}

	// avoid adding same peer twice, no real peer information merging for now
	for _, peer := range other.Peers {
		exists := false

		for _, p := range spec.Peers {
			if p.PublicKey == peer.PublicKey {
				exists = true

				break
			}
		}

		if !exists {
			spec.Peers = append(spec.Peers, peer)
		}
	}
}
