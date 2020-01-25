package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type awsCloudformationDocJSON struct {
	Contents []awsCloudformationDocJSONMap
}

type awsCloudformationDocJSONMap struct {
	Title    string
	Href     string
	Contents []awsCloudformationDocJSONMapInner
}

type awsCloudformationDocJSONMapInner struct {
	Title    string
	Href     string
	Contents []awsCloudformationDocJSONMapEnd
}

type awsCloudformationDocJSONMapEnd struct {
	Title string
	Href  string
}

func main() {
	var cloudformationResourceServices = []string{
		"AWS_AccessAnalyzer",
		"Alexa_ASK",
		"AWS_AmazonMQ",
		"AWS_Amplify",
		"AWS_ApiGateway",
		"AWS_ApiGatewayV2",
		"AWS_ApplicationAutoScaling",
		"AWS_AppMesh",
		"AWS_AppStream",
		"AWS_AppSync",
		"AWS_Athena",
		"AWS_AutoScalingPlans",
		"AWS_AutoScaling",
		"AWS_Backup",
		"AWS_Batch",
		"AWS_Budgets",
		"AWS_CertificateManager",
		"AWS_Cloud9",
		"AWS_CloudFormation",
		"AWS_CloudFront",
		"AWS_ServiceDiscovery",
		"AWS_CloudTrail",
		"AWS_CloudWatch",
		"AWS_Logs",
		"AWS_Events",
		"AWS_CodeBuild",
		"AWS_CodeCommit",
		"AWS_CodeDeploy",
		"AWS_CodePipeline",
		"AWS_CodeStar",
		"AWS_CodeStarNotifications",
		"AWS_Cognito",
		"AWS_Config",
		"AWS_DataPipeline",
		"AWS_DAX",
		"AWS_DirectoryService",
		"AWS_DLM",
		"AWS_DMS",
		"AWS_DocDB",
		"AWS_DynamoDB",
		"AWS_EC2",
		"AWS_ECR",
		"AWS_ECS",
		"AWS_EFS",
		"AWS_EKS",
		"AWS_ElastiCache",
		"AWS_Elasticsearch",
		"AWS_ElasticBeanstalk",
		"AWS_ElasticLoadBalancing",
		"AWS_ElasticLoadBalancingV2",
		"AWS_EMR",
		"AWS_EventSchemas",
		"AWS_FSx",
		"AWS_GameLift",
		"AWS_Glue",
		"AWS_GuardDuty",
		"AWS_IAM",
		"AWS_Inspector",
		"AWS_IoT",
		"AWS_IoT1Click",
		"AWS_IoTAnalytics",
		"AWS_IoTEvents",
		"AWS_Greengrass",
		"AWS_IoTThingsGraph",
		"AWS_Kinesis",
		"AWS_KinesisAnalytics",
		"AWS_KinesisAnalyticsV2",
		"AWS_KinesisFirehose",
		"AWS_KMS",
		"AWS_LakeFormation",
		"AWS_Lambda",
		"AWS_ManagedBlockchain",
		"AWS_MediaConvert",
		"AWS_MediaLive",
		"AWS_MediaStore",
		"AWS_MSK",
		"AWS_Neptune",
		"AWS_OpsWorks",
		"AWS_OpsWorksCM",
		"AWS_Pinpoint",
		"AWS_PinpointEmail",
		"AWS_QLDB",
		"AWS_RAM",
		"AWS_RDS",
		"AWS_Redshift",
		"AWS_RoboMaker",
		"AWS_Route53",
		"AWS_Route53Resolver",
		"AWS_S3",
		"AWS_SageMaker",
		"AWS_SecretsManager",
		"AWS_ServiceCatalog",
		"AWS_SecurityHub",
		"AWS_SES",
		"AWS_SDB",
		"AWS_SNS",
		"AWS_SQS",
		"AWS_StepFunctions",
		"AWS_SSM",
		"AWS_Transfer",
		"AWS_WAF",
		"AWS_WAFv2",
		"AWS_WAFRegional",
		"AWS_WorkSpaces",
	}

	var resources = []string{}
	for _, service := range cloudformationResourceServices {
		url := fmt.Sprintf("https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/toc-%v.json", service)
		r, err := getResourceTypes(url)
		if err != nil {
			fmt.Println("Error getting resource tpyes for url: ", url)
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		resources = append(resources, r...)
	}
	for _, resource := range resources {
		println(resource)
	}
}

func getResourceTypes(url string) (resources []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return resources, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resources, err
	}
	var doc awsCloudformationDocJSON
	err = json.Unmarshal(body, &doc)
	if err != nil {
		return resources, err
	}
	for _, resource := range doc.Contents[0].Contents {
		resources = append(resources, resource.Title)
	}
	return
}
