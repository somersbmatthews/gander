package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	catalog "catalog/proto/catalog"
)

type Catalog struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Catalog) Call(ctx context.Context, req *catalog.Request, rsp *catalog.Response) error {
	log.Info("Received Catalog.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Catalog) Stream(ctx context.Context, req *catalog.StreamingRequest, stream catalog.Catalog_StreamStream) error {
	log.Infof("Received Catalog.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&catalog.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Catalog) PingPong(ctx context.Context, stream catalog.Catalog_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&catalog.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
