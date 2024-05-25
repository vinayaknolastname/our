package gApi

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/vinayaknolastname/our/protobuf/video"
)

func (server *GrpcServer) VideoProccess(stream video.VideoService_VideoProccessServer) error {
	inputVideo := "received_videos.mp4"
	outputPrefix := "output_chunk"

	file, err := os.Create(inputVideo)
	if err != nil {
		log.Printf("err in creating itemp img %s", err)

		return err
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Received message from %s", err)

			// return nil
		}
		if err != nil {
			log.Printf("Received message from 12 %s", err)

			// return err

		}

		// log.Printf("Received message from  22: %s", in.Content)

		_, errf := file.Write(in.GetContent())

		if err != nil {
			log.Printf("Erro in writing file: %s", errf)
			// return nil, err
		}
		log.Printf("Erro in writing file: %s", len(in.Content))

		if len(in.Content) == 0 {
			break
		}

		// if err := stream.Send(&video.CommonResponseVideo{StatusCode: 500, Message: "Hello ", Success: false}); err != nil {
		// 	return err
		// }
	}

	log.Printf("Erro in writing fildde: %s")

	dur, err := getVideoDuration(inputVideo)

	if err != nil {
		log.Printf("Erro in writing file: %s", err)

		return err
	}
	log.Printf("Erro in writing file: %s", dur)

	numberOfWorker := 4

	var wg sync.WaitGroup

	chunks := make(chan int, numberOfWorker)

	for i := 0; i < numberOfWorker; i++ {
		log.Printf("l1: %s", i)

		wg.Add(1)
		go func() {
			defer wg.Done()

			for chunk := range chunks {
				log.Printf("l1: %s", chunks)
				splitVideo(inputVideo, outputPrefix, chunk)
			}
		}()
	}

	for i := 0; i < dur; i++ {
		chunks <- i
	}
	println("weg")

	wg.Wait()
	println("wef in video resolution.")

	changeVideoResolutionOperation()

	return nil
}

func changeVideoResolutionOperation() {
	println("err in video resolution.")
	dir := "output_videos"
	outputPrefix := "output_chunk"

	numOfVideo, _ := countNumberOfVideos(dir)

	numOfWorkers := 4

	resolutions := []string{"640x360", "1280x720", "1920x1080"}

	var wg sync.WaitGroup

	videos := make(chan int, numOfWorkers)
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		defer wg.Done()
		go func() {
			for video := range videos {
				file := fmt.Sprintf("%s_%d.mp4", outputPrefix, video)
				input := filepath.Join(dir, file)
				changeVideoResolution(input, resolutions[0], resolutions[0])
			}
		}()
	}

	for i := 0; i < numOfVideo; i++ {
		videos <- i
	}
	wg.Wait()
}

func countNumberOfVideos(sourceDirectory string) (int, error) {
	fileCount := 0

	err := filepath.Walk(sourceDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileCount++
		}
		return nil
	})

	return fileCount, err
}

func changeVideoResolution(inputVideo string, resolution string, output string) {
	// outputFilePath := filepath.Join(os.TempDir(), "converted_video_"+resolution+".mp4")
	cmd := exec.Command("ffmpeg", "-i", inputVideo, "-vf", "scale="+resolution, output)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("err in video resolution.")

	}

}

func splitVideo(inputVideo, outputPrefix string, duration int) error {
	// Calculate the number of chunks
	// numChunks := duration

	outputDir := "output_videos"

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}
	// Split the video into chunks
	// for i := 0; i < numChunks; i++ {
	outputFile := fmt.Sprintf("%s_%d.mp4", outputPrefix, duration)
	cmd := exec.Command("ffmpeg", "-ss", strconv.Itoa(duration), "-i", inputVideo, "-t", "1", "-c", "copy", filepath.Join(outputDir, outputFile))
	if err := cmd.Run(); err != nil {
		fmt.Println("Video split into 1-second chunks successfully.")

		return err
		// }
	}
	fmt.Println("Video split into 1-second chunks successfully.")

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
