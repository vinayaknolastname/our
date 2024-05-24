package gApi

import "github.com/vinayaknolastname/our/protobuf/video"

func (server *GrpcServer) VideoProccess(stream video.VideoService_VideoProccessServer) error {

	print("dd")
	return nil
}
