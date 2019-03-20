package config

// defaults struct
type defaults struct {
	AWSAccessKeyID  string `yaml:"AWSAccessKeyID"`
	AWSRegion       string `yaml:"AWSRegion"`
	AWSSecretKey    string `yaml:"AWSSecretKey"`
	CognitoClientID string `yaml:"CognitoClientID"`
	CognitoPoolID   string `yaml:"CognitoPoolID"`
	CognitoRegion   string `yaml:"CognitoRegion"`
	S3Bucket        string `yaml:"S3Bucket"`
	SsmPath         string `yaml:"SsmPath"`
	Stage           string `yaml:"Stage"`
}

type config struct {
	AWSAccessKeyID  string
	AWSRegion       string
	AWSSecretKey    string
	CognitoClientID string
	CognitoPoolID   string
	CognitoRegion   string
	S3Bucket        string
	Stage           StageEnvironment
}
