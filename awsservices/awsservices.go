package awsservices

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pulpfree/univsales-pdf-url/config"

	log "github.com/sirupsen/logrus"
)

// S3Service struct
type S3Service struct {
	cfg     *config.Config
	session *session.Session
}

// NewS3 function
func NewS3(cfg *config.Config) (svc *S3Service, err error) {

	svc = &S3Service{
		cfg: cfg,
	}

	svc.session, err = session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return svc, err
}

// GetSignedURL method
func (s *S3Service) GetSignedURL(prefix string) (signedURL string, err error) {

	svc := s3.New(s.session)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.cfg.S3Bucket),
		Key:    aws.String(prefix),
	})

	urlStr, err := req.Presign(30 * time.Minute)
	if err != nil {
		log.Errorf("Failed to sign request: %s", err.Error())
		return "", err
	}

	return urlStr, err
}
