package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	endpoint        = "203.194.113.6:9000"
	accessKeyID     = "MidFtK0wfiZ6AUjDfZbz"
	secretAccessKey = "KxkgFNq196ok2AKq9U5h2naOUq0Akpi8HyjA4RO3"
	bucketName      = "smk-telkom"
)

func main() {
	// Initialize MinIO client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully connected to MinIO")

	// Example usage
	listObjects(minioClient)
	uploadFile(minioClient, "example.txt")
	downloadFile(minioClient, "example.txt", "downloaded_example.txt")
	generatePresignedURL(minioClient, "example.txt")
	generatePublicURL("example.txt")
	deleteFile(minioClient, "example.txt")
}

func listObjects(client *minio.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectCh := client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{})
	fmt.Println("Files in bucket:")
	for object := range objectCh {
		if object.Err != nil {
			log.Fatalln(object.Err)
		}
		fmt.Println(object.Key)
	}
}

func uploadFile(client *minio.Client, filePath string) {
	ctx := context.Background()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	fileStat, _ := file.Stat()
	_, err = client.PutObject(ctx, bucketName, filePath, file, fileStat.Size(), minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("File uploaded successfully")
}

func downloadFile(client *minio.Client, objectName, filePath string) {
	ctx := context.Background()
	err := client.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("File downloaded successfully")
}

func deleteFile(client *minio.Client, objectName string) {
	ctx := context.Background()
	err := client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("File deleted successfully")
}

func generatePresignedURL(client *minio.Client, objectName string) {
	ctx := context.Background()
	expiry := time.Second * 60
	url, err := client.PresignedGetObject(ctx, bucketName, objectName, expiry, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Presigned URL:", url)
}

func generatePublicURL(objectName string) {
	fmt.Printf("Public URL: http://%s/%s/%s\n", endpoint, bucketName, objectName)
}
