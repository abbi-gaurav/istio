// nolint:lll
// Generates the mygrpcadapter adapter's resource yaml. It contains the adapter's configuration, name,
// supported template names (metric in this case), and whether it is session or no-session based.
//go:generate $REPO_ROOT/bin/mixer_codegen.sh -a mixer/adapter/mygrpcadapter/config/config.proto -x "-s=false -n mygrpcadapter -t authorization"

package mygrpcadapter

import (
	"context"
	"fmt"
	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/passport_service"
	"net"

	"google.golang.org/grpc"

	"istio.io/api/mixer/adapter/model/v1beta1"
	"istio.io/istio/mixer/pkg/status"
	"istio.io/istio/mixer/template/authorization"
	"istio.io/pkg/log"
)

type (
	Server interface {
		Addr() string
		Close() error
		Run(shutdown chan error)
	}

	MyAuthAdapter struct {
		listener net.Listener
		server   *grpc.Server
	}
)

func (ma *MyAuthAdapter) HandleAuthorization(ctx context.Context, r *authorization.HandleAuthorizationRequest) (*v1beta1.CheckResult, error) {
	log.Infof("received request %v\n", *r)
	return &v1beta1.CheckResult{
		Status: status.OK,
	}, nil
}

func (ma *MyAuthAdapter) Addr() string {
	return ma.listener.Addr().String()
}

func (ma *MyAuthAdapter) Run(shutdown chan error) {
	shutdown <- ma.server.Serve(ma.listener)
}

func (ma *MyAuthAdapter) Close() error {
	if ma.server != nil {
		ma.server.GracefulStop()
	}

	if ma.listener != nil {
		_ = ma.listener.Close()
	}

	return nil
}

func NewMyAuthAdapter(addr string, ps *passport_service.PassportService) (Server, error) {
	if addr == "" {
		addr = "0"
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))

	if err != nil {
		return nil, fmt.Errorf("unable to listen on socket: %+v", err)
	}

	ma := &MyAuthAdapter{
		listener: listener,
	}

	log.Infof("listening on %+v", listener)

	ma.server = grpc.NewServer()
	authorization.RegisterHandleAuthorizationServiceServer(ma.server, ma)
	return ma, nil
}
