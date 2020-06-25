package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	provision "provision/proto/provision"
)

type Provision struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Provision) Call(ctx context.Context, req *provision.Request, rsp *provision.Response) error {
	log.Info("Received Provision.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Provision) Stream(ctx context.Context, req *provision.StreamingRequest, stream provision.Provision_StreamStream) error {
	log.Infof("Received Provision.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&provision.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Provision) PingPong(ctx context.Context, stream provision.Provision_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&provision.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
