package gApi

import (
	"io"
	"log"

	"github.com/vinayaknolastname/our/protobuf/video"
)

func (server *GrpcServer) VideoProccess(stream video.VideoService_VideoProccessServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received message from %s: %s", in.Content, in.Content)
		if err := stream.Send(&video.CommonResponse{StatusCode: 500, Message: "Hello ", Success: false}); err != nil {
			return err
		}
	}

}
