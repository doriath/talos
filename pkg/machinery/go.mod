module github.com/talos-systems/talos/pkg/machinery

go 1.18

// forked go-yaml that introduces RawYAML interface, which can be used to populate YAML fields using bytes
// which are then encoded as a valid YAML blocks with proper indentiation
replace gopkg.in/yaml.v3 => github.com/unix4ever/yaml v0.0.0-20220527175918-f17b0f05cf2c

require (
	github.com/containerd/go-cni v1.1.7
	github.com/cosi-project/runtime v0.0.0-20220705131812-22c6aa1ca7ec
	github.com/dustin/go-humanize v1.0.0
	github.com/evanphx/json-patch v5.6.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/jsimonetti/rtnetlink v1.2.0
	github.com/mdlayher/ethtool v0.0.0-20220213132912-856bd6cb8a38
	github.com/opencontainers/runtime-spec v1.0.3-0.20200929063507-e6143ca7d51d
	github.com/siderolabs/go-pointer v1.0.0
	github.com/stretchr/testify v1.8.0
	github.com/talos-systems/crypto v0.3.6-0.20220622130438-e9df1b8ca74c
	github.com/talos-systems/go-blockdevice v0.3.4
	github.com/talos-systems/go-debug v0.2.1
	github.com/talos-systems/net v0.3.2
	google.golang.org/genproto v0.0.0-20220722212130-b98a9ff5e252
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v3 v3.0.1
	inet.af/netaddr v0.0.0-20220617031823-097006376321
)

require (
	github.com/containernetworking/cni v1.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/josharian/native v1.0.0 // indirect
	github.com/mdlayher/genetlink v1.2.0 // indirect
	github.com/mdlayher/netlink v1.6.0 // indirect
	github.com/mdlayher/socket v0.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	go4.org/intern v0.0.0-20211027215823-ae77deb06f29 // indirect
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20220617031537-928513b29760 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	golang.org/x/sys v0.0.0-20220408201424-a24fb2fb8a0f // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
