package main

import (
	"context"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// endpoint := "http://127.0.0.1:9000/"
	endpoint := "minio-server"

	accessKeyId := "chenhn"
	secretAccessKey := "chn123456789"
	useSSL := false
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v\n", minioClient)

	presignedURL, err := minioClient.PresignedPutObject(context.Background(), "my-minio", "images/微信图片_20220928154250.jpg", time.Duration(1000)*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(presignedURL)
}
