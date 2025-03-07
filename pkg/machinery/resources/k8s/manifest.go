// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package k8s

import (
	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta"
	"github.com/cosi-project/runtime/pkg/resource/typed"
)

// ManifestType is type of Manifest resource.
const ManifestType = resource.Type("Manifests.kubernetes.talos.dev")

// Manifest resource holds definition of kubelet static pod.
type Manifest = typed.Resource[ManifestSpec, ManifestRD]

// ManifestSpec holds the Kubernetes resources spec.
//gotagsrewrite:gen
type ManifestSpec struct {
	Items []map[string]interface{} `protobuf:"1"`
}

// MarshalYAML implements yaml.Marshaler.
func (spec ManifestSpec) MarshalYAML() (interface{}, error) {
	return spec.Items, nil
}

// NewManifest initializes an empty Manifest resource.
func NewManifest(namespace resource.Namespace, id resource.ID) *Manifest {
	return typed.NewResource[ManifestSpec, ManifestRD](
		resource.NewMetadata(namespace, ManifestType, id, resource.VersionUndefined),
		ManifestSpec{},
	)
}

// ManifestRD provides auxiliary methods for Manifest.
type ManifestRD struct{}

// ResourceDefinition implements typed.ResourceDefinition interface.
func (ManifestRD) ResourceDefinition(resource.Metadata, ManifestSpec) meta.ResourceDefinitionSpec {
	return meta.ResourceDefinitionSpec{
		Type:             ManifestType,
		Aliases:          []resource.Type{},
		DefaultNamespace: ControlPlaneNamespaceName,
	}
}
