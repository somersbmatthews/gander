package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	policy "policy/proto/policy"
)

type Policy struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Policy) Call(ctx context.Context, req *policy.Request, rsp *policy.Response) error {
	log.Info("Received Policy.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Policy) Stream(ctx context.Context, req *policy.StreamingRequest, stream policy.Policy_StreamStream) error {
	log.Infof("Received Policy.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&policy.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Policy) PingPong(ctx context.Context, stream policy.Policy_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&policy.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
