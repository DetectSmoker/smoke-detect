package models

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"
)

type GoogleCloud struct {}

func (model *Model) UploadFile(file multipart.File, object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := model.Config.Client.Bucket(model.Config.BucketName).Object(model.Config.UploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}