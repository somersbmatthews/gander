package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	alert "alert/proto/alert"
)

type Alert struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Alert) Call(ctx context.Context, req *alert.Request, rsp *alert.Response) error {
	log.Info("Received Alert.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Alert) Stream(ctx context.Context, req *alert.StreamingRequest, stream alert.Alert_StreamStream) error {
	log.Infof("Received Alert.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&alert.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Alert) PingPong(ctx context.Context, stream alert.Alert_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&alert.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
