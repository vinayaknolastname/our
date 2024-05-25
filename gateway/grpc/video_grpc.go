package grpcHandlers

import (
	context "context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vinayaknolastname/our/gateway/utils"
	"github.com/vinayaknolastname/our/protobuf/video"

	"google.golang.org/grpc"
)

var (
// tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
// caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
// serverAddrVideo = flag.String("addr", "localhost:9001", "The server address in the format of host:port")

// serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

type VideoGrpcService struct {
	conn               *grpc.ClientConn
	videoServiceClient video.VideoServiceClient
}

var videoGrpcClient VideoGrpcService

func ConnectWithVideoGrpcMiddleWare(c *gin.Context) {
	utils.LogSomething("Calling ConnectVideoServiceGrpc --- Middle", videoGrpcClient.conn, 1)

	ConnectWithVideoGrpc()
	c.Next()
}

func ConnectWithVideoGrpc() {

	if videoGrpcClient.conn == nil {

		connection, err := grpc.NewClient("localhost:9001", grpc.WithInsecure())

		if err != nil {
			utils.LogSomething("err in dialing video grpc", err, 0)
		}

		client := video.NewVideoServiceClient(connection)

		videoGrpcClient = VideoGrpcService{
			conn:               connection,
			videoServiceClient: client,
		}
	}
}

func StartVideoStream(c *gin.Context) {
	inputVideo := "source_video.mp4"
	// outputPrefix := "output_chunk"

	// dur, _ := getVideoDuration(inputVideo)

	// var wg sync.WaitGroup

	// numWorkers := 4
	// chunks := make(chan int, numWorkers)

	// for i := 0; i < numWorkers; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for chunk := range chunks {
	// 			log.Printf("Error splitting video chunk %d: %v", chunk)

	// 			err := splitVideo(inputVideo, outputPrefix, chunk)
	// 			if err != nil {
	// 				log.Printf("Error splitting video chunk %d: %v", chunk, err)
	// 			}
	// 		}
	// 	}()

	// }

	// for i := 0; i < dur; i++ {
	// 	chunks <- i
	// }
	// close(chunks)

	// // Wait for all worker goroutines to finish
	// wg.Wait()
	// splitVideo(inputVideo, outputPrefix, dur)

	// ConvertBigVideoInSmallDurationChunks()
	// videoData := &video.Video{
	// 	Content: []byte{111},
	// }

	client := videoGrpcClient.videoServiceClient

	utils.LogSomething("videoProces", videoGrpcClient.videoServiceClient, 0)

	stream, err := client.VideoProccess(context.Background())

	if err != nil {
		utils.LogSomething("videoProces", err, 0)
	}

	buffer := make([]byte, 1024)
	// go func() {
	// 	for {
	// 		// res, err := stream.Recv()
	// 		// if err == io.EOF {
	// 		// 	close(waitc)
	// 		// 	return
	// 		// }
	// 		// if err != nil {
	// 		// 	log.Fatalf("failed to receive a message : %v", err)
	// 		// }
	// 		// log.Printf("Received message from %s: %s", res.StatusCode, res.Message)
	// 	}
	// }()

	// users := []string{"Alice", "Bob", "Charlie"}

	file, err := os.Open(inputVideo)

	if err != nil {
		log.Fatalf("error reading file d: %v", err)
	}
	// defer file.Close()

	go func() {
		for {
			bytesRead, err := file.Read(buffer)
			if err != nil {
				log.Println("error reading file ch: %v", err)
			}

			err = stream.Send(&video.Video{Content: buffer[:bytesRead]})
			if err != nil {
				log.Println("error sending chunk: %v", err)
			}
			if bytesRead == 0 {
				file.Close()
				break
			}
		}
	}()
	// go func() {
	// 	for {
	// 		if err := stream.Send(videoData); err != nil {
	// 			log.Fatalf("failed to send a message: %v", err)
	// 		}
	// 		time.Sleep(1 * time.Second)
	// 	}

	// }()
	// stream.CloseSend()

}

func PlayVideo(c *gin.Context) {
	// playlistPath := filepath.Join(hlsDir, "master.m3u8")
	// if _, err := os.Stat(playlistPath); os.IsNotExist(err) {
	// http.Error(w, "Master playlist not found", http.StatusNotFound)
	// return
	// }
	c.File(mas)
}

func splitVideo(inputVideo, outputPrefix string, duration int) error {
	// Calculate the number of chunks
	// numChunks := duration

	// Split the video into chunks
	// for i := 0; i < numChunks; i++ {
	outputFile := fmt.Sprintf("%s_%d.mp4", outputPrefix, duration)
	cmd := exec.Command("ffmpeg", "-ss", strconv.Itoa(duration), "-i", inputVideo, "-t", "1", "-c", "copy", outputFile)
	if err := cmd.Run(); err != nil {
		return err
		// }
	}
	return nil
}
func ConvertBigVideoInSmallDurationChunks() {
	inputVideo := "source_video.mp4"
	outputPrefix := "output_chunk"

	cmd := exec.Command("ffmpeg", "-i", inputVideo, "-c", "copy", "-segment_time", "1", "-f", "segment", outputPrefix+"%03d.mp4")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error splitting video: %v", err)
	}

	fmt.Println("Video split into 1-second chunks successfully.")
}

func getVideoDuration(inputVideo string) (int, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", inputVideo)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}
	durationStr := strings.TrimSpace(string(output))
	durationFloat, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, err
	}
	duration := int(durationFloat)
	return duration, nil
}
