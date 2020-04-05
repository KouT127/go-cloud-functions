package image

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"io"
	"io/ioutil"
	"mime/multipart"
)

const publicURL = "https://storage.googleapis.com/%s/%s"
const bucketName = "images"

type StorageClient interface {
	Upload(file multipart.File, header *multipart.FileHeader) (string, error)
}

type storageClient struct {
	BucketHandle *storage.BucketHandle
}

func NewStorageClient(opt option.ClientOption) (StorageClient, error) {
	ctx := context.Background()
	bucket, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	client := storageClient{
		BucketHandle: bucket.Bucket(bucketName),
	}
	return &client, nil
}

func (s *storageClient) Upload(file multipart.File, header *multipart.FileHeader) (string, error) {
	ctx := context.Background()

	uuidObject, err := uuid.NewUUID()
	if err != nil {
		return "", nil
	}

	bytes, err := ioutil.ReadAll(file)
	name := uuid.NewSHA1(uuidObject, bytes)

	writer := s.BucketHandle.Object(name.String() + "/" + header.Filename).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	writer.ContentType = header.Header.Get("Content-Type")
	writer.CacheControl = "public, max-age=86400"

	if _, err := io.Copy(writer, file); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}

	return fmt.Sprintf(publicURL, bucketName, name), nil
}
