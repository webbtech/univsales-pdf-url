module github.com/pulpfree/univsales-pdf-url

go 1.13

require (
	github.com/aws/aws-lambda-go v1.14.0
	github.com/aws/aws-sdk-go v1.29.6
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/pulpfree/lambda-utils v0.0.0-20200219231926-889196177675
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v2 v2.2.8
)

// replace github.com/pulpfree/lambda-utils v0.0.0 => ./github.com/pulpfree/lambda-utils
