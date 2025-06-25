// Using json content to auth to invoke genai SDK.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"os"

	"cloud.google.com/go/auth/credentials"

	"google.golang.org/genai"
)

var (
	projectID = flag.String("project", "frankie1-422709", "the project id")
	location  = flag.String("location", "us-central1", "the location")
	modelName = flag.String("model", "gemini-2.0-flash", "the model name")
)

func chat(ctx context.Context) {

	//get json content.
	data, err := os.ReadFile("REPLACE_YOUR FILE PATH.")
	if err != nil {
		log.Fatal(err)
	}
	creds, err := credentials.DetectDefault(&credentials.DetectOptions{
		Scopes:          []string{"https://www.googleapis.com/auth/bigquery", "https://www.googleapis.com/auth/cloud-platform"},
		CredentialsJSON: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Project:  *projectID,
		Location: *location,
		Backend: genai.BackendVertexAI,
		Credentials: creds,
	})
	if err != nil {
		log.Fatal(err)
	}

	var config *genai.GenerateContentConfig = &genai.GenerateContentConfig{Temperature: genai.Ptr[float32](0.5)}

	// Create a new Chat.
	chat, err := client.Chats.Create(ctx, *modelName, config, nil)

	// Send first chat message.
	result, err := chat.SendMessage(ctx, genai.Part{Text: "What's the weather in San Francisco?"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())

	// Send second chat message.
	result, err = chat.SendMessage(ctx, genai.Part{Text: "How about New York?"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())
}

func main() {
	ctx := context.Background()
	flag.Parse()
	chat(ctx)
}
