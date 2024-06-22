package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type JazidaApiStackProps struct {
	awscdk.StackProps
}

func NewJazidaApiStack(scope constructs.Construct, id string, props *JazidaApiStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	loadsTable := awsdynamodb.NewTable(stack, jsii.String("LoadsTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("loads"),
	})

	newLoadFunction := awslambda.NewFunction(stack, jsii.String("load"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.AssetCode_FromAsset(jsii.String("load/function.zip"), nil),
		Handler: jsii.String("main"),
	})

	loadsTable.GrantReadWriteData(newLoadFunction)

	api := awsapigateway.NewRestApi(stack, jsii.String("apiGateway"), &awsapigateway.RestApiProps{
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowHeaders: jsii.Strings("Content-Type", "Authrization"),
			AllowMethods: jsii.Strings("GET", "POST", "OPTIONS"),
			AllowOrigins: jsii.Strings("*"),
		},
		DeployOptions: &awsapigateway.StageOptions{
			LoggingLevel: awsapigateway.MethodLoggingLevel_INFO,
		},
	})

	loadIntegration := awsapigateway.NewLambdaIntegration(newLoadFunction, nil)

	registerResource := api.Root().AddResource(jsii.String("loads"), nil)
	registerResource.AddMethod(jsii.String("POST"), loadIntegration, nil)
	registerResource.AddMethod(jsii.String("GET"), loadIntegration, nil)

	signatureBucket := awss3.NewBucket(stack, jsii.String("SIGNATURES"), &awss3.BucketProps{
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	signatureUploadFunction := awslambda.NewFunction(stack, jsii.String("upload"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.AssetCode_FromAsset(jsii.String("upload/function.zip"), nil),
		Handler: jsii.String("main"),
		Environment: &map[string]*string{
			"BUCKET_NAME": signatureBucket.BucketName(),
		},
	})

	signatureBucket.GrantReadWrite(signatureUploadFunction, nil)

	signatureIntegration := awsapigateway.NewLambdaIntegration(signatureUploadFunction, nil)

	signatureUploadResource := api.Root().AddResource(jsii.String("signature"), nil)
	keySignatureResource := signatureUploadResource.AddResource(jsii.String("{key}"), nil)

	signatureUploadResource.AddMethod(jsii.String("POST"), signatureIntegration, nil)
	keySignatureResource.AddMethod(jsii.String("GET"), signatureIntegration, nil)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewJazidaApiStack(app, "JazidaApiStack", &JazidaApiStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
