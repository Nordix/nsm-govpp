// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package arp

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/edwarnicke/govpp/binapi/memclnt"
)

// RPCService defines RPC service arp.
type RPCService interface {
	ProxyArpAddDel(ctx context.Context, in *ProxyArpAddDel) (*ProxyArpAddDelReply, error)
	ProxyArpDump(ctx context.Context, in *ProxyArpDump) (RPCService_ProxyArpDumpClient, error)
	ProxyArpIntfcDump(ctx context.Context, in *ProxyArpIntfcDump) (RPCService_ProxyArpIntfcDumpClient, error)
	ProxyArpIntfcEnableDisable(ctx context.Context, in *ProxyArpIntfcEnableDisable) (*ProxyArpIntfcEnableDisableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) ProxyArpAddDel(ctx context.Context, in *ProxyArpAddDel) (*ProxyArpAddDelReply, error) {
	out := new(ProxyArpAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ProxyArpDump(ctx context.Context, in *ProxyArpDump) (RPCService_ProxyArpDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_ProxyArpDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_ProxyArpDumpClient interface {
	Recv() (*ProxyArpDetails, error)
	api.Stream
}

type serviceClient_ProxyArpDumpClient struct {
	api.Stream
}

func (c *serviceClient_ProxyArpDumpClient) Recv() (*ProxyArpDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *ProxyArpDetails:
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

func (c *serviceClient) ProxyArpIntfcDump(ctx context.Context, in *ProxyArpIntfcDump) (RPCService_ProxyArpIntfcDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_ProxyArpIntfcDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_ProxyArpIntfcDumpClient interface {
	Recv() (*ProxyArpIntfcDetails, error)
	api.Stream
}

type serviceClient_ProxyArpIntfcDumpClient struct {
	api.Stream
}

func (c *serviceClient_ProxyArpIntfcDumpClient) Recv() (*ProxyArpIntfcDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *ProxyArpIntfcDetails:
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

func (c *serviceClient) ProxyArpIntfcEnableDisable(ctx context.Context, in *ProxyArpIntfcEnableDisable) (*ProxyArpIntfcEnableDisableReply, error) {
	out := new(ProxyArpIntfcEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
