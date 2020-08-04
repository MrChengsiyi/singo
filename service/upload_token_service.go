package service

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"mime"
	"path/filepath"
	"singo/serializer"
)

// UploadTokenService 获得上传oss token的服务
type UploadTokenService struct {
	Filename string `form:"filename" json:"filename"`
}

// Post 创建token
func (service *UploadTokenService) Post() serializer.Response {
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI4FyBXLb7ikTZoUJE6zy4", "pUb5pjMwJH2C64xbjE8YUcCutk5Hip")
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "oss配置失败",
			Error: err.Error(),
		}
	}

	// 获取存储空间。
	bucket, err := client.Bucket("csyishero")
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "oss配置失败",
			Error: err.Error(),
		}
	}
	// 获取扩展名
	ext := filepath.Ext(service.Filename)
	//fmt.Println(ext,mime.TypeByExtension(ext))
	// 带可选参数的签名直传
	option := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),

	}

	key := uuid.Must(uuid.NewRandom()).String() + ext

	// 签名直传。
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, option...)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "oss失败",
			Error: err.Error(),
		}
	}

	// 查看图片
	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "oss配置失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}
