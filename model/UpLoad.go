package model

import (
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	putPolic := storage.PutPolicy{
		Scope: utils.Bucket,
	}
	mac := qbox.NewMac(utils.AccessKey, utils.SecretKey)
	upToken := putPolic.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		fmt.Println("上传失败")
		return "", errmsg.ERROR
	}
	url := utils.Url + ret.Key
	return url, errmsg.SUCCESS
}
