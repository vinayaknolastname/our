package mediaservice

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/vinayaknolastname/our/gateway/utils"
)

type MediaDB struct {
	MediaDB *cloudinary.Cloudinary
}

var MDB *MediaDB

func NewMediaDB() {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================

	if err != nil {
		utils.LogSomething("cld", err, 0)
	}

	utils.LogSomething("cld", cld, 0)
	// cld.Config.URL.Secure = true
	// ctx := context.Background()
	MDB = &MediaDB{
		MediaDB: cld,
	}
	// return cld, ctx
}
