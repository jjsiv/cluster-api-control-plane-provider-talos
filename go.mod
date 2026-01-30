module github.com/siderolabs/cluster-api-control-plane-provider-talos

go 1.25.3

require (
	github.com/coreos/go-semver v0.3.1
	github.com/go-logr/logr v1.4.3
	github.com/gobuffalo/flect v1.0.3
	github.com/google/uuid v1.6.0
	github.com/onsi/gomega v1.38.2
	github.com/pkg/errors v0.9.1
	github.com/siderolabs/capi-utils v0.0.0-20251124160722-4ee8a1b7d4d0
	github.com/siderolabs/cluster-api-bootstrap-provider-talos v0.6.11
	github.com/siderolabs/crypto v0.6.4
	github.com/siderolabs/gen v0.8.6
	github.com/siderolabs/go-retry v0.3.3
	github.com/siderolabs/talos/pkg/machinery v1.12.0
	github.com/spf13/pflag v1.0.10
	github.com/stretchr/testify v1.11.1
	golang.org/x/sync v0.18.0
	google.golang.org/grpc v1.76.0
	google.golang.org/protobuf v1.36.10
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.34.3
	k8s.io/apiextensions-apiserver v0.34.3
	k8s.io/apimachinery v0.34.3
	k8s.io/apiserver v0.34.3
	k8s.io/client-go v0.34.3
	k8s.io/component-base v0.34.3
	k8s.io/klog/v2 v2.130.1
	k8s.io/utils v0.0.0-20251002143259-bc988d571ff4
	sigs.k8s.io/cluster-api v1.12.2
	sigs.k8s.io/controller-runtime v0.22.5
)
