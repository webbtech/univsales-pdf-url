# Instructions On How To Create and Upload A Resource Policy

This is taken from [aws-samples/aws-sam-movies-api-resource-policy](https://github.com/aws-samples/aws-sam-movies-api-resource-policy)

```bash
policy=`cat policy.json`
API_ID=0d9cxfnacl
aws apigateway update-rest-api --rest-api-id $API_ID --patch-operations op=replace,path=/policy,value="$policy"
```
