include .env

default: build awspackage awsdeploy

clean:
	@rm -rf dist
	@mkdir -p dist

build: clean
	@for dir in `ls handler`; do \
		GOOS=linux go build -o dist/$$dir github.com/pulpfree/univsales-pdf-url/handler/$$dir; \
	done
	@cp ./config/defaults.yml dist/
	@echo "build successful"

run: build
	sam local start-api -n env.json

validate:
	sam validate

awspackage:
	@aws cloudformation package \
	--template-file ${FILE_TEMPLATE} \
	--output-template-file ${FILE_PACKAGE} \
	--s3-bucket $(AWS_BUCKET_NAME) \
	--s3-prefix $(AWS_BUCKET_PREFIX) \
	--profile $(AWS_PROFILE) \
	--region $(AWS_REGION)

awsdeploy:
	@aws cloudformation deploy \
	--template-file ${FILE_PACKAGE} \
	--region $(AWS_REGION) \
	--stack-name $(PROJECT_NAME) \
	--capabilities CAPABILITY_IAM \
	--profile $(AWS_PROFILE) \
	--force-upload \
	--parameter-overrides \
	 	ParamKMSKeyID=$(KMS_KEY_ID) \
		ParamBucketName=$(AWS_BUCKET_NAME) \
		ParamThundraKey=$(THUNDRA_API_KEY) \
		ParamAuthLambda=$(LAMBDA_AUTHORIZER_ARN)

describe:
	@aws cloudformation describe-stacks \
		--region $(AWS_REGION) \
		--stack-name $(PROJECT_NAME)

outputs:
	@ make describe \
		| jq -r '.Stacks[0].Outputs'