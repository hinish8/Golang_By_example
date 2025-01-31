package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	// Base 64 encoding
	sample := "a@"
	encodedString_base64 := base64.StdEncoding.EncodeToString([]byte(sample))
	fmt.Println(encodedString_base64)

	originalStringBytes, err := base64.StdEncoding.DecodeString(encodedString_base64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(originalStringBytes))

	//URL encoding
	sample1 := "�"

	encodedStringURL := base64.URLEncoding.EncodeToString([]byte(sample1))
	fmt.Printf("URL Encoding: %s\n", encodedStringURL)
	encodedStringSTD := base64.StdEncoding.EncodeToString([]byte(sample1))
	fmt.Printf("STD Encoding: %s\n", encodedStringSTD)

	originalStringBytes, err = base64.URLEncoding.DecodeString(encodedStringURL)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))

	//Raw std encoding
	sample = "a@"
	encodedStringStdEncoding := base64.StdEncoding.EncodeToString([]byte(sample))
	fmt.Printf("STD Encoding: %s\n", encodedStringStdEncoding)

	encodedStringRawStdEncoding := base64.RawStdEncoding.EncodeToString([]byte(sample))
	fmt.Printf("Raw STD Encoding: %s\n", encodedStringRawStdEncoding)

	originalStringBytes, err = base64.RawStdEncoding.DecodeString(encodedStringRawStdEncoding)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))

	// Raw URL encoding
	sample = "a@"
	encodedStringURLEncoding := base64.URLEncoding.EncodeToString([]byte(sample))
	fmt.Printf("URL Encoding: %s\n", encodedStringURLEncoding)

	encodedStringRawURLEncoding := base64.RawURLEncoding.EncodeToString([]byte(sample))
	fmt.Printf("Raw URL Encoding: %s\n", encodedStringRawURLEncoding)

	originalStringBytes, err = base64.RawStdEncoding.DecodeString(encodedStringRawURLEncoding)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))

}
