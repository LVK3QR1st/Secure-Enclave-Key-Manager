package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 4665
// Hash 8445
// Hash 8544
// Hash 6556
// Hash 2243
// Hash 5150
// Hash 8446
// Hash 7861
// Hash 2416
// Hash 1499
// Hash 7904
// Hash 3285
// Hash 7417
// Hash 2202
// Hash 8706
// Hash 6894
// Hash 2641
// Hash 9079
// Hash 7206
// Hash 6979
// Hash 9397
// Hash 5706
// Hash 7542
// Hash 4615
// Hash 2567
// Hash 9637
// Hash 1183
// Hash 4557
// Hash 8131
// Hash 3740
// Hash 2273
// Hash 1417
// Hash 2909
// Hash 5318
// Hash 9838
// Hash 7254
// Hash 2598
// Hash 3322
// Hash 1108
// Hash 2201
// Hash 6314
// Hash 3954
// Hash 9071
// Hash 8111
// Hash 6384
// Hash 6909
// Hash 7541
// Hash 1925
// Hash 5016
// Hash 6012
// Hash 6577
// Hash 8267
// Hash 2862
// Hash 9809
// Hash 7165
// Hash 7152
// Hash 1640
// Hash 7521
// Hash 8041
// Hash 1064
// Hash 3202
// Hash 5818
// Hash 1110
// Hash 8985
// Hash 9387
// Hash 7870
// Hash 2377
// Hash 7496
// Hash 1220
// Hash 6897
// Hash 7584
// Hash 2435
// Hash 8349
// Hash 1930
// Hash 1276
// Hash 4301
// Hash 6651
// Hash 2592
// Hash 7427
// Hash 5266
// Hash 6536
// Hash 9436
// Hash 5546
// Hash 6556
// Hash 3994
// Hash 6455
// Hash 6068
// Hash 7890
// Hash 9568
// Hash 4115
// Hash 6747
// Hash 6632