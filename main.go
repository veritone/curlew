package main

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Scheme = "s3"
)

func main() {
	objPath := getObjPath()

	obj, err := getObj(objPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed getting object:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, obj)
}

func getObjPath() (objPath string) {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error: wrong number of arguments provided")
		fmt.Fprintln(os.Stderr, "")
		help()
	}

	switch os.Args[1] {
	case "h", "-h", "--h", "help", "-help", "--help":
		help()
	}

	objPath = os.Args[1]
	return
}

func getObj(objPath string) (objReader io.ReadCloser, err error) {
	s3URL, err := url.Parse(objPath)
	if err != nil {
		return
	}

	if strings.ToLower(s3URL.Scheme) != s3Scheme {
		err = fmt.Errorf("Error: scheme is not s3: %s", s3URL.Scheme)
		return
	}

	bucket := s3URL.Host
	key := s3URL.Path

	sess := session.New()
	svcS3 := s3.New(sess)

	in := s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	out, err := svcS3.GetObject(&in)
	if err != nil {
		return
	}

	objReader = out.Body
	return
}

func help() {
	fmt.Fprintln(os.Stderr, `Get a single object from an S3 bucket

Usage:
	curlew s3://bucket-name-here/path/to/object.txt
	curlew s3://bucket-name-here/path/to/object.txt > object.txt`)
	fmt.Fprintln(os.Stderr, "")

	os.Exit(1)
}
