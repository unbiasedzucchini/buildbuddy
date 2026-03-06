package distributed_client

import (
	"context"
	"testing"
	"time"

	dcpb "github.com/buildbuddy-io/buildbuddy/proto/distributed_cache"
	repb "github.com/buildbuddy-io/buildbuddy/proto/remote_execution"
	rspb "github.com/buildbuddy-io/buildbuddy/proto/resource"
	"github.com/buildbuddy-io/buildbuddy/server/util/rpcutil"
	"github.com/buildbuddy-io/buildbuddy/server/util/status"
	"github.com/buildbuddy-io/buildbuddy/server/util/testing/flags"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

type fakeWriteClient struct {
	ctx      context.Context
	sendErr  chan error
	closeRsp chan struct {
		rsp *dcpb.WriteResponse
		err error
	}
}

func (c *fakeWriteClient) Send(*dcpb.WriteRequest) error {
	select {
	case err := <-c.sendErr:
		return err
	case <-c.ctx.Done():
		return c.ctx.Err()
	}
}

func (c *fakeWriteClient) CloseAndRecv() (*dcpb.WriteResponse, error) {
	select {
	case msg := <-c.closeRsp:
		return msg.rsp, msg.err
	case <-c.ctx.Done():
		return nil, c.ctx.Err()
	}
}

func (c *fakeWriteClient) Header() (metadata.MD, error) { return nil, nil }
func (c *fakeWriteClient) Trailer() metadata.MD         { return nil }
func (c *fakeWriteClient) CloseSend() error             { return nil }
func (c *fakeWriteClient) Context() context.Context     { return c.ctx }
func (c *fakeWriteClient) SendMsg(any) error            { return nil }
func (c *fakeWriteClient) RecvMsg(any) error            { return nil }

func testResourceName() *rspb.ResourceName {
	return &rspb.ResourceName{
		Digest: &repb.Digest{
			Hash:      "abcd",
			SizeBytes: 4,
		},
		CacheType:    rspb.CacheType_CAS,
		InstanceName: "test",
	}
}

func TestStreamWriteCloserWriteTimeout(t *testing.T) {
	flags.Set(t, "cache.distributed_cache.peer_write_timeout", time.Millisecond)

	ctx, cancel := context.WithCancel(context.Background())
	stream := &fakeWriteClient{
		ctx:     ctx,
		sendErr: make(chan error),
		closeRsp: make(chan struct {
			rsp *dcpb.WriteResponse
			err error
		}),
	}
	wc := &streamWriteCloser{
		cancelFunc: cancel,
		sender:     rpcutil.NewSender[*dcpb.WriteRequest, *dcpb.WriteResponse](ctx, stream),
		stream:     stream,
		peer:       "peer-1",
		r:          testResourceName(),
	}

	_, err := wc.Write([]byte("data"))
	require.Error(t, err)
	require.True(t, status.IsDeadlineExceededError(err), "expected deadline exceeded, got %s", err)
	require.ErrorIs(t, ctx.Err(), context.Canceled)
}

func TestStreamWriteCloserCommitCloseAndRecvTimeout(t *testing.T) {
	flags.Set(t, "cache.distributed_cache.peer_write_timeout", time.Millisecond)

	ctx, cancel := context.WithCancel(context.Background())
	stream := &fakeWriteClient{
		ctx:     ctx,
		sendErr: make(chan error, 1),
		closeRsp: make(chan struct {
			rsp *dcpb.WriteResponse
			err error
		}),
	}
	stream.sendErr <- nil
	wc := &streamWriteCloser{
		cancelFunc: cancel,
		sender:     rpcutil.NewSender[*dcpb.WriteRequest, *dcpb.WriteResponse](ctx, stream),
		stream:     stream,
		peer:       "peer-1",
		r:          testResourceName(),
	}

	err := wc.Commit()
	require.Error(t, err)
	require.True(t, status.IsDeadlineExceededError(err), "expected deadline exceeded, got %s", err)
	require.ErrorIs(t, ctx.Err(), context.Canceled)
}
