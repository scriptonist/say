package main

import (
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"

	pb "github.com/scriptonist/say/api"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address to listen to")
	output := flag.String("o", "output.wav", "file to store output")
	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	checkError(err, "Cannot reach GRPC Server")
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	text := &pb.Text{
		Text: "Hello there",
	}

	res, err := client.Say(context.Background(), text)
	checkError(err, "flite error,backend error")

	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("Could not write %s:%v", *output, err)
	}
}

func checkError(err error, message string) {
	if err != nil {
		log.Fatalf(message, err)
	}
}
