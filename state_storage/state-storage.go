package state_storage

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// TerratestTempBucket creates a temporal S3 bucket to hold the state files for example modules that will be tested.
func TerratestTempBucketAndTable(t *testing.T, awsAccountID, awsRegion, tempBucket string) *terraform.Options {
	currentDateTime := strings.ToLower(time.Now().Local().Format("2006-01-02T15-04-05.005"))
	bucketOpts := &terraform.Options{
		TerraformDir: tempBucket,
		Reconfigure:  true,
		Vars: map[string]interface{}{
			"aws_account_id":      awsAccountID,
			"aws_region":          awsRegion,
			"bucket_name":         fmt.Sprintf("%s-%s-terratest-%s", awsAccountID, awsRegion, currentDateTime),
			"dynamodb_table_name": fmt.Sprintf("%s-%s-terratest-%s", awsAccountID, awsRegion, currentDateTime),
			"custom_tags": map[string]interface{}{
				"Environment":    "Automation Environment",
				"DeployStrategy": "Terraform started by Terratest",
			},
		},
	}
	return bucketOpts
}
