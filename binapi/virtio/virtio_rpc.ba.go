// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package virtio

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/edwarnicke/govpp/binapi/memclnt"
)

// RPCService defines RPC service virtio.
type RPCService interface {
	SwInterfaceVirtioPciDump(ctx context.Context, in *SwInterfaceVirtioPciDump) (RPCService_SwInterfaceVirtioPciDumpClient, error)
	VirtioPciCreate(ctx context.Context, in *VirtioPciCreate) (*VirtioPciCreateReply, error)
	VirtioPciCreateV2(ctx context.Context, in *VirtioPciCreateV2) (*VirtioPciCreateV2Reply, error)
	VirtioPciDelete(ctx context.Context, in *VirtioPciDelete) (*VirtioPciDeleteReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) SwInterfaceVirtioPciDump(ctx context.Context, in *SwInterfaceVirtioPciDump) (RPCService_SwInterfaceVirtioPciDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwInterfaceVirtioPciDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwInterfaceVirtioPciDumpClient interface {
	Recv() (*SwInterfaceVirtioPciDetails, error)
	api.Stream
}

type serviceClient_SwInterfaceVirtioPciDumpClient struct {
	api.Stream
}

func (c *serviceClient_SwInterfaceVirtioPciDumpClient) Recv() (*SwInterfaceVirtioPciDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwInterfaceVirtioPciDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) VirtioPciCreate(ctx context.Context, in *VirtioPciCreate) (*VirtioPciCreateReply, error) {
	out := new(VirtioPciCreateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) VirtioPciCreateV2(ctx context.Context, in *VirtioPciCreateV2) (*VirtioPciCreateV2Reply, error) {
	out := new(VirtioPciCreateV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) VirtioPciDelete(ctx context.Context, in *VirtioPciDelete) (*VirtioPciDeleteReply, error) {
	out := new(VirtioPciDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
