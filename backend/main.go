package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io/ioutil"
	"os/exec"

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
	f, err := ioutil.TempFile("", "")
	checkError(err, "could'nt create file")
	err = f.Close()
	checkError(err, "Could not close file")

	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())
	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed with :  %s ", data)
	}
	data, err := ioutil.ReadFile(f.Name())
	checkError(err, "could not read temp file")
	return &pb.Speech{Audio: data}, nil

}
func checkError(err error, message string) {
	if err != nil {
		log.Fatalf(message, err)
	}
}
