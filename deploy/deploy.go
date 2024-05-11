package deploy

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

// TerratestGenOpts generates the corresponding configuration needed and returns *terraform.Options and the the path of the module.
func TerratestGenOpts(t *testing.T, awsRegion, bucketName, modName, modPath, awsBucketKey string, feedVars map[string]interface{}) (*terraform.Options, string) {
	modulePath := modPath
	s3Key := ""
	if awsBucketKey != "" {
		s3Key = awsBucketKey
	} else {
		s3Key = fmt.Sprintf("%s/terratest/%s/terraform.tfstate", awsRegion, modName)
	}

	moduleOptions := &terraform.Options{
		TerraformDir: test_structure.CopyTerraformFolderToTemp(t, "../", modulePath),
		Reconfigure:  true,
		Vars:         feedVars,
		BackendConfig: map[string]interface{}{
			"bucket":         bucketName,
			"region":         awsRegion,
			"key":            s3Key,
			"dynamodb_table": bucketName,
			"encrypt":        true,
		},
	}
	return moduleOptions, "../" + modulePath
}

// PlanModule runs terraform init and plan with the given options and returns stdout/stderr from the plan command.
func PlanModule(t *testing.T, tfOpts *terraform.Options) {
	terraform.InitAndPlan(t, tfOpts)
}

// DeployModule deploys a given module and save the terraform options object on disk for further use
func DeployModule(t *testing.T, tfOpts *terraform.Options, modpath string) {
	test_structure.SaveTerraformOptions(t, modpath, tfOpts)
	terraform.InitAndApply(t, tfOpts)
}

// UnDeployModule undeploys a given module and save the terraform options object on disk for further use
func UnDeployModule(t *testing.T, modpath string) {
	tfOpts := test_structure.LoadTerraformOptions(t, modpath)
	defer terraform.Destroy(t, tfOpts)
}
