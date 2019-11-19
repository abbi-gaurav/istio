module istio.io/istio/mixer/adapter/mygrpcadapter

go 1.13

require (
	github.com/gogo/protobuf v1.3.1
	golang.org/x/net v0.0.0-20191112182307-2180aed22343 // indirect
	golang.org/x/sys v0.0.0-20191113165036-4c7a9d0fe056 // indirect
	google.golang.org/genproto v0.0.0-20191114150713-6bbd007550de // indirect
	google.golang.org/grpc v1.25.1
	istio.io/api v0.0.0-20191111210003-35e06ef8d838
	istio.io/istio v0.0.0-20191114021746-c4def934e4f8
)

replace k8s.io/api => k8s.io/api v0.0.0-20191003000013-35e20aa79eb8

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20191003000419-f68efa97b39e

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191003002041-49e3d608220c

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20191003002408-6e42c232ac7d

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191003002707-f6b7b0f55cc0
