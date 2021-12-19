package sqs

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/labstack/echo/v4"
)

type GetQueuesResponse struct {
	QueueURLs []string
}

func GetQueues(c echo.Context) error {
	sess, err := loadSession()
	if err != nil {
		return err
	}

	svc := sqs.New(sess)
	queues, err := svc.ListQueues(nil)
	if err != nil {
		log.Println(fmt.Sprintf("ListQueues failed: %s\n", err))
		return err
	}

	queueURLs := make([]string, 0, len(queues.QueueUrls))
	for _, url := range queues.QueueUrls {
		queueURLs = append(queueURLs, *url)
	}

	res := GetQueuesResponse{
		QueueURLs: queueURLs,
	}
	return c.JSON(http.StatusOK, res)
}

type awsConfig struct {
	ID       string
	SECRET   string
	Region   string
	Endpoint string
}

func newAWSConfig() (*awsConfig, error) {
	awsID := os.Getenv("AWS_ID")
	if awsID == "" {
		log.Println(fmt.Sprintf("Env is empty: %s\n", "AWS_ID"))
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "unknown error")
	}

	awsSecret := os.Getenv("AWS_SECRET")
	if awsSecret == "" {
		log.Println(fmt.Sprintf("Env is empty: %s\n", "AWS_SECRET"))
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "unknown error")
	}

	region := os.Getenv("AWS_REGION")
	if region == "" {
		log.Println(fmt.Sprintf("Env is empty: %s\n", " AWS_REGION"))
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "unknown error")
	}

	endpoint := os.Getenv("AWS_ENDPOINT")
	if endpoint == "" {
		log.Println(fmt.Sprintf("Env is empty: %s\n", "AWS_ENDPOINT"))
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "unknown error")
	}

	return &awsConfig{
		ID:       awsID,
		SECRET:   awsSecret,
		Region:   region,
		Endpoint: endpoint,
	}, nil
}

func loadSession() (*session.Session, error) {
	cfg, err := newAWSConfig()
	if err != nil {
		return nil, err
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(cfg.Region),
			Endpoint:    aws.String(cfg.Endpoint),
			Credentials: credentials.NewStaticCredentials(cfg.ID, cfg.SECRET, ""),
		},
	}))
	return sess, nil
}
