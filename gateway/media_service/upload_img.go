package mediaservice

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/vinayaknolastname/our/gateway/rtc/ws"
	"github.com/vinayaknolastname/our/gateway/utils"
)

var (
	_, b, _, _ = runtime.Caller(0)
	RootPath   = filepath.Join(filepath.Dir(b), "../..")
)

func fileNameWithoutExtension(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}

func FileUpload(c *gin.Context) {
	//1. Param input for multipart file upload
	/// Maximum of 200MB file allowed

	//2. Retrieve file from form-data

	userID, _ := strconv.Atoi(c.Param("userId"))
	chatID, _ := strconv.Atoi(c.Query("chatId"))

	//<Form-id> is the form key that we will read from. Client should use the same form key when uploading the file
	file, err := c.FormFile("form-id")
	if err != nil {
		errStr := fmt.Sprintf("Error in reading the file %s\n", err)
		fmt.Println(errStr)
		// fmt.Fprintf(c, errStr)
		return
	}

	localFilePath := saveFile(c, file)

	fileUrl := storeImgInCloud(localFilePath)

	go ws.DoThisOnImgMsg(ws.WsManagerIns, &ws.Message{
		ChatId:    int32(chatID),
		SenderId:  int32(userID),
		MediaLink: fileUrl,
	})

	// fmt.Fprintf(c, result)

}

func saveFile(c *gin.Context, file *multipart.FileHeader) (localImgPath string) {

	tempPath := fmt.Sprintf("assets/uploads/" + file.Filename)
	err := c.SaveUploadedFile(file, tempPath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar"})
		return
	}

	return tempPath

}

func storeImgInCloud(path string) string {
	ctx := context.Background()

	file, err := os.ReadFile(path)

	if err != nil {
		utils.LogSomething("error in read file", err, 0)
	}
	utils.LogSomething("error in read file", file, 0)
	resp, err := MDB.MediaDB.Upload.Upload(ctx, path, uploader.UploadParams{
		PublicID:       "path",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	if err != nil {
		fmt.Println("error", err)
	}

	// grpcHandlers.CreateMessage()
	// Log the delivery URL
	fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL, "\n")
	return resp.SecureURL
}
