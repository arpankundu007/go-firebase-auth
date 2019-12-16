package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

func GetFirebaseApp () (*firebase.App, error) {
	opt := option.WithCredentialsFile("/Users/arpankundu/Downloads/cafu-auth-test-firebase-adminsdk.json")
	config := &firebase.Config{ProjectID: "cafu-auth-test"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}
	return app, err
}
