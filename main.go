package main

import (
        "encoding/base64"
        "github.com/aws/aws-lambda-go/lambda"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/kms"
		"os"
		"log"
)

var enClinetID string = os.Getenv("ClientID")
var enClinetSecret string = os.Getenv("ClientSecret")


var ClientID string
var ClientSecret string

func init() {
	ClientID = decrypt(enClinetID)
	ClientSecret = decrypt(enClinetSecret)
}

// Decrypting encrypted environment variables
func decrypt(value string) string {
	kmsClient := kms.New(session.New())
	decodedBytes, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
			panic(err)
	}
	input := &kms.DecryptInput{
			CiphertextBlob: decodedBytes,
	}
	response, err := kmsClient.Decrypt(input)
	if err != nil {
			panic(err)
	}
	// Plaintext is a byte array, so convert to string
	return string(response.Plaintext[:])
}

func HandleRequest() (string, error) {

	log.Println("INFO: ClientID: ", ClientID)
	log.Println("INFO: ClientSecret: ", ClientSecret)

    return "Hello, World!", nil
}

func main() {
    lambda.Start(HandleRequest)
}
