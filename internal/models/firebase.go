package models

import (
	"context"
	"log"

	"google.golang.org/api/iterator"
)

func (model *Model) CreateImageData(data map[string]interface{}) error {
	ref, res, err := model.Config.FirebaseClient.Collection("detect_smoker").Add(context.Background(), data)
	if err != nil {
		return err
	}
	log.Println(ref)
	log.Println(res)
	return nil
}

func (model *Model) GetAllData() []map[string]interface{} {
	iter := model.Config.FirebaseClient.Collection("detect_smoker").Documents(context.Background())
	var data []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		data = append(data, doc.Data())
		log.Println(doc.Data())
	}
	return data
}
