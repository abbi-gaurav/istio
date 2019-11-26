module istio.io/istio/mixer/adapter/mygrpcadapter

go 1.13

replace k8s.io/api => k8s.io/api v0.0.0-20191003000013-35e20aa79eb8

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20191003000419-f68efa97b39e

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191003002041-49e3d608220c

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20191003002408-6e42c232ac7d

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191003002707-f6b7b0f55cc0

require (
	github.com/go-redis/redis v6.10.2+incompatible
	github.com/gogo/protobuf v1.3.1
	google.golang.org/grpc v1.25.1
	istio.io/api v0.0.0-20191126023742-2a7248f229ee
	istio.io/istio v0.0.0-20191126034942-9b19bb2eec42
	istio.io/pkg v0.0.0-20191125150539-18bbd272ee5c
)
