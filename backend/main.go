package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"net"

	pb "github.com/scriptonist/say/api"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	log.Info("Listening on port ", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	checkError(err, fmt.Sprintf("Cannot Listen on port %d", *port))

	s := grpc.NewServer()

	pb.RegisterTextToSpeechServer(s, server{})
	err = s.Serve(lis)
	checkError(err, "cannot start server")

}

type server struct {
}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	return nil, fmt.Errorf("Not Implemented")
}
func checkError(err error, message string) {
	if err != nil {
		log.Fatalf(message, err)
	}
}
