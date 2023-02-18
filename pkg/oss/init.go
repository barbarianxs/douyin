package oss

import (
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/consts"
	"os"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Bucket *oss.Bucket
var Path string

func Init() {
	//算出绝对路径，防止service层测试时路径错误
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	Path = strings.Split(dir, "/cmd")[0]
	//打开oss的Bucket
	endpoint := consts.OssEndPoint
	accesskeyid := consts.OssAccessKeyId
	accessKeySecret := consts.OssAccessKeySecret
	bucket := consts.OssBucket
	client, err := oss.New(endpoint, accesskeyid, accessKeySecret)
	if err != nil {
		panic(err)
	}
	Bucket, err = client.Bucket(bucket)
	if err != nil {
		panic(err)
	}
}
