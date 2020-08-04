package model

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

type Video struct {
	gorm.Model
	Title  string
	Url    string
	Info   string
	Avatar string
}

// AvatarURL 封面地址
func (video *Video) AvatarURL() string {
	client, _ := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI4FyBXLb7ikTZoUJE6zy4", "pUb5pjMwJH2C64xbjE8YUcCutk5Hip")
	bucket, _ := client.Bucket("csyishero")
	signedGetURL, _ := bucket.SignURL(video.Avatar, oss.HTTPGet, 600)
	fmt.Println(signedGetURL)
	return signedGetURL
}
