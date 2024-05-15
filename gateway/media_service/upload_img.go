package mediaservice

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/vinayaknolastname/our/gateway/utils"
)

func HandleImgMessage(base64Data string) {

	imageData, err := base64.StdEncoding.DecodeString(string(base64Data))
	if err != nil {
		log.Println("Error decoding base64 data:", err)
		return
	}

	// Write image data to file
	var imgPath = "received_image.jpg"
	err = os.WriteFile(imgPath, imageData, 0644)
	if err != nil {
		log.Println("Error writing image data to file:", err)
		return
	}

	fmt.Println("Image received and saved successfully")
	// 	var ctx = context.Background()
	// resp, err := cld.Upload.Upload(ctx, "my_picture.jpg", uploader.UploadParams{PublicID: "my_image"});
	StoreImgInCloud(imgPath)

}

func StoreImgInCloud(path string) {
	ctx := context.Background()

	file, err := os.ReadFile(path)

	if err != nil {
		utils.LogSomething("error in read file", err, 0)
	}
	utils.LogSomething("error in read file", file, 0)
	resp, err := MDB.MediaDB.Upload.Upload(ctx, path, uploader.UploadParams{
		PublicID:       "quickstart_butterfly",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	if err != nil {
		fmt.Println("error", err)
	}

	// Log the delivery URL
	fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL, "\n")
}
