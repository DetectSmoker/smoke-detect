package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
)

type Config struct {
	Client         *storage.Client
	FirebaseClient *firestore.Client
	ProjectID      string
	BucketName     string
	UploadPath     string
	ImageURL       string
}

const CREDENTIAL_PATH = "../credential/detect-smoker-1.json"

func InitConfig() *Config {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", CREDENTIAL_PATH)

	ctx := context.Background()
	//initial gcs
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	firebaseClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	projectId := "detect-smoker"
	bucketName := "detect_smoker-1"
	imgUrl := fmt.Sprintf("https://storage.googleapis.com/%s/", bucketName)
	return &Config{
		Client:         client,
		FirebaseClient: firebaseClient,
		ProjectID:      projectId,
		BucketName:     bucketName,
		UploadPath:     "",
		ImageURL:       imgUrl,
	}
}
