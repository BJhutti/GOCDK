package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GocdkStackProps struct {
	awscdk.StackProps
}

// scope is the app
func NewGocdkStack(scope constructs.Construct, id string, props *GocdkStackProps) awscdk.Stack {

	//construct our properties
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	//starting a new stsck
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	awslambda.NewFunction(stack, jsii.String("myLambdaFunction"), &awslambda.FunctionProps{
		//lambda specifc propertoes
		//Runtime: what is the backend writtien in?
		//Handler: where does this code exist?
		//Code: how do we execute our code?
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.AssetCode_FromAsset(jsii.String("lambda/function.zip"), nil), // nil means nothiing extra
		Handler: jsii.String("main"),
	})

	// example resource
	//queue := awssqs.NewQueue(stack, jsii.String("GocdkQueue"), &awssqs.QueueProps{
	//VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	//})

	return stack
}

func main() {
	defer jsii.Close() //jsii: another framework that lets us use cdk with not typescript. transpiles go into typescript
	//defer: run this code after everything else. deployed, sythed

	//cdk is like using lego pieces, we can attach things to thhings
	//app is what were going to bind everything to
	app := awscdk.NewApp(nil) //type construct

	//app will be passed into here
	//stacks: containds peices of infrastructure . tools in stack, which are in apps

	NewGocdkStack(app, "GocdkStack", &GocdkStackProps{ //returns a stack
		awscdk.StackProps{
			Env: env(), //pass in env
		},
	})

	app.Synth(nil) //app sythesises the stack
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil //default awss account credentials

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
