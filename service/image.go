package service

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime/multipart"
	"sfts/initial"
	"strings"

	"github.com/minio/minio-go/v6"
	log "github.com/sirupsen/logrus"
)

const (
	url  = "101.43.39.61"
	port = "9000"
)

func HashSHAFile(file multipart.File) (string, error) {
	var hashValue string
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return hashValue, err
	}
	hashInBytes := hash.Sum(nil)
	hashValue = hex.EncodeToString(hashInBytes)
	return hashValue, nil
}

func CreateImage(imageName string, file multipart.File, fileSize int64) (imageURL string, err error) {
	bucketName := "image"
	objectName := imageName
	if err != nil {
		return
	}
	imageURL = "http://" + url + ":" + port + "/" + bucketName + "/" + objectName

	_, err = initial.MinioClient.PutObject(bucketName, objectName, file, fileSize, minio.PutObjectOptions{})
	if err != nil {
		return
	}
	return imageURL, nil
}

func UploadImage(FileHeader *multipart.FileHeader) (string, error) {
	var imageURL string
	file, err := FileHeader.Open()
	if err != nil {
		log.Errorf("Open MultipartFile error: %v", err)
		return imageURL, err
	}

	hashValue, err := HashSHAFile(file)
	if err != nil {
		return "", err
	}

	fileArrIndex := strings.LastIndex(FileHeader.Filename, ".")
	fileSuf := FileHeader.Filename[fileArrIndex:]
	fileName := hashValue + fileSuf
	fileSize := FileHeader.Size

	file, err = FileHeader.Open()
	if err != nil {
		log.Errorf("Open MultipartFile error: %v", err)
		return "", err
	}

	imageURL, err = CreateImage(fileName, file, fileSize)
	if err != nil {
		log.Errorf("保存图片错误: %v", err)
		return imageURL, err
	}

	if err = file.Close(); err != nil {
		log.Errorf("Close file error: %v", err)
		return imageURL, err
	}
	return imageURL, nil
}
