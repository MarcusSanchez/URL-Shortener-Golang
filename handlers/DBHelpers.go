package handlers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"urlShortener/database"
)

type URL struct {
	ObjectID primitive.ObjectID `bson:"_id,omitempty"`
	LongURL  string             `bson:"LongURL"`
	ShortURL string             `bson:"ShortURL"`
}

func renderResults() ([]URL, error) {
	collection := database.GetCollection("main")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return []URL{}, err
	}
	defer func() {
		if err = cursor.Close(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	var results []URL
	for cursor.Next(context.Background()) {
		var url URL
		if err = cursor.Decode(&url); err != nil {
			return []URL{}, err
		}
		results = append(results, url)
	}

	if err = cursor.Err(); err != nil {
		return []URL{}, err
	}

	return results, nil
}

func saveURL(url string) (string, error) {
	if err := verifyUrl(url); err != nil {
		return "", err
	}

	collection := database.GetCollection("main")

	count, err := collection.CountDocuments(context.Background(), bson.M{"LongURL": url})
	if err != nil {
		return "", err
	}

	if count != 0 {
		var result URL
		err = collection.FindOne(context.Background(), bson.M{"LongURL": url}, options.FindOne().SetProjection(bson.M{"ShortURL": 1, "_id": 0})).Decode(&result)
		if err != nil {
			return "", err
		}

		shortUrl := result.ShortURL
		return shortUrl, nil
	}

	var shortUrl string
	for {
		shortUrl = generateShortenedUrl()

		count, err = collection.CountDocuments(context.Background(), bson.M{"ShortURL": shortUrl})
		if err != nil {
			return "", err
		}

		if count == 0 {
			break
		}
	}

	document := URL{
		LongURL:  url,
		ShortURL: shortUrl,
	}

	_, err = collection.InsertOne(context.Background(), document)
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func retrieveUrl(base62Key string) (string, error) {
	key := database.Route + "/s/" + base62Key

	// Query the MongoDB collection for a matching document
	collection := database.GetCollection("main")

	var result URL
	err := collection.FindOne(context.Background(), bson.M{"ShortURL": key}, options.FindOne().SetProjection(bson.M{"LongURL": 1, "_id": 0})).Decode(&result)
	if err != nil {
		return "", err
	}

	// Access the original URL from the result
	fmt.Println(result.ObjectID.IsZero())
	url := result.LongURL

	return url, nil
}

func clearDB() error {
	collection := database.GetCollection("main")

	// Delete all documents from the MongoDB collection
	_, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	return nil
}

func deleteUrl(shortUrl string) error {
	collection := database.GetCollection("main")

	_, err := collection.DeleteOne(context.Background(), bson.M{"ShortURL": shortUrl})
	if err != nil {
		return err
	}

	return nil
}
