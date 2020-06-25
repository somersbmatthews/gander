package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	job "job/proto/job"
)

type Job struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Job) Call(ctx context.Context, req *job.Request, rsp *job.Response) error {
	log.Info("Received Job.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Job) Stream(ctx context.Context, req *job.StreamingRequest, stream job.Job_StreamStream) error {
	log.Infof("Received Job.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&job.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Job) PingPong(ctx context.Context, stream job.Job_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&job.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
