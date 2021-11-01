// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package pbl

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/edwarnicke/govpp/binapi/memclnt"
)

// RPCService defines RPC service pbl.
type RPCService interface {
	PblClientDel(ctx context.Context, in *PblClientDel) (*PblClientDelReply, error)
	PblClientDump(ctx context.Context, in *PblClientDump) (RPCService_PblClientDumpClient, error)
	PblClientUpdate(ctx context.Context, in *PblClientUpdate) (*PblClientUpdateReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) PblClientDel(ctx context.Context, in *PblClientDel) (*PblClientDelReply, error) {
	out := new(PblClientDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PblClientDump(ctx context.Context, in *PblClientDump) (RPCService_PblClientDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_PblClientDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_PblClientDumpClient interface {
	Recv() (*PblClientDetails, error)
	api.Stream
}

type serviceClient_PblClientDumpClient struct {
	api.Stream
}

func (c *serviceClient_PblClientDumpClient) Recv() (*PblClientDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *PblClientDetails:
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

func (c *serviceClient) PblClientUpdate(ctx context.Context, in *PblClientUpdate) (*PblClientUpdateReply, error) {
	out := new(PblClientUpdateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
