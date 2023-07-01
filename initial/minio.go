package initial

import (
	"github.com/minio/minio-go/v6"
)

const (
	url               = "101.43.39.61"
	port              = "9000"
	access_key_id     = "username"
	secret_access_key = "password"
)

var MinioClient *minio.Client

func InitMinio() error {
	var err error
	MinioClient, err = minio.New(url+":"+port, access_key_id, secret_access_key, false)
	return err
}
