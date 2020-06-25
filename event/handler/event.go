package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	event "event/proto/event"
)

type Event struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Event) Call(ctx context.Context, req *event.Request, rsp *event.Response) error {
	log.Info("Received Event.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Event) Stream(ctx context.Context, req *event.StreamingRequest, stream event.Event_StreamStream) error {
	log.Infof("Received Event.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&event.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Event) PingPong(ctx context.Context, stream event.Event_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&event.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
