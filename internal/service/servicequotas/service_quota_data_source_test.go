package servicequotas_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/service/servicequotas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccServiceQuotasServiceQuotaDataSource_quotaCode(t *testing.T) {
	const dataSourceName = "data.aws_servicequotas_service_quota.test"

	const serviceCode = "vpc"
	const quotaCode = "L-F678F1CE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(servicequotas.EndpointsID, t)
			preCheckServiceQuotaSet(serviceCode, quotaCode, t)
		},
		ErrorCheck: acctest.ErrorCheck(t, servicequotas.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceQuotaQuotaCodeDataSourceConfig(serviceCode, quotaCode),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "adjustable", "true"),
					acctest.CheckResourceAttrRegionalARN(dataSourceName, "arn", "servicequotas", fmt.Sprintf("%s/%s", serviceCode, quotaCode)),
					resource.TestCheckResourceAttr(dataSourceName, "default_value", "5"),
					resource.TestCheckResourceAttr(dataSourceName, "global_quota", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "quota_code", quotaCode),
					resource.TestCheckResourceAttr(dataSourceName, "quota_name", "VPCs per Region"),
					resource.TestCheckResourceAttr(dataSourceName, "service_code", serviceCode),
					resource.TestCheckResourceAttr(dataSourceName, "service_name", "Amazon Virtual Private Cloud (Amazon VPC)"),
					resource.TestMatchResourceAttr(dataSourceName, "value", regexp.MustCompile(`^\d+$`)),
				),
			},
		},
	})
}

func TestAccServiceQuotasServiceQuotaDataSource_quotaCode_Unset(t *testing.T) {
	const dataSourceName = "data.aws_servicequotas_service_quota.test"

	const serviceCode = "s3"
	const quotaCode = "L-DC2B2D3D"
	const quotaName = "Buckets"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(servicequotas.EndpointsID, t)
			preCheckServiceQuotaUnset(serviceCode, quotaCode, t)
		},
		ErrorCheck: acctest.ErrorCheck(t, servicequotas.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceQuotaQuotaCodeDataSourceConfig(serviceCode, quotaCode),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceAttrRegionalARNNoAccount(dataSourceName, "arn", "servicequotas", fmt.Sprintf("%s/%s", serviceCode, quotaCode)),
					resource.TestCheckResourceAttr(dataSourceName, "adjustable", "true"),
					resource.TestMatchResourceAttr(dataSourceName, "default_value", regexp.MustCompile(`^\d+$`)),
					resource.TestCheckResourceAttr(dataSourceName, "global_quota", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "quota_code", quotaCode),
					resource.TestCheckResourceAttr(dataSourceName, "quota_name", quotaName),
					resource.TestCheckResourceAttr(dataSourceName, "service_code", serviceCode),
					resource.TestCheckResourceAttr(dataSourceName, "service_name", "Amazon Simple Storage Service (Amazon S3)"),
					resource.TestMatchResourceAttr(dataSourceName, "value", regexp.MustCompile(`^\d+$`)),
					resource.TestCheckResourceAttrPair(dataSourceName, "value", dataSourceName, "default_value"),
				),
			},
		},
	})
}

func TestAccServiceQuotasServiceQuotaDataSource_PermissionError_quotaCode(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			testAccPreCheck(t)
			acctest.PreCheckAssumeRoleARN(t)
		},
		ErrorCheck:   acctest.ErrorCheck(t, servicequotas.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config:      testAccServiceQuotaDataSourceConfig_PermissionError_QuotaCode("elasticloadbalancing", "L-53DA6B97"),
				ExpectError: regexp.MustCompile(`DEPENDENCY_ACCESS_DENIED_ERROR`),
			},
		},
	})
}

func TestAccServiceQuotasServiceQuotaDataSource_quotaName(t *testing.T) {
	dataSourceName := "data.aws_servicequotas_service_quota.test"

	const serviceCode = "vpc"
	const quotaCode = "L-F678F1CE"
	const quotaName = "VPCs per Region"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(servicequotas.EndpointsID, t)
			preCheckServiceQuotaSet(serviceCode, quotaCode, t)
		},
		ErrorCheck: acctest.ErrorCheck(t, servicequotas.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceQuotaQuotaNameDataSourceConfig("vpc", quotaName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "adjustable", "true"),
					acctest.CheckResourceAttrRegionalARN(dataSourceName, "arn", "servicequotas", fmt.Sprintf("%s/%s", serviceCode, quotaCode)),
					resource.TestCheckResourceAttr(dataSourceName, "default_value", "5"),
					resource.TestCheckResourceAttr(dataSourceName, "global_quota", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "quota_code", quotaCode),
					resource.TestCheckResourceAttr(dataSourceName, "quota_name", quotaName),
					resource.TestCheckResourceAttr(dataSourceName, "service_code", serviceCode),
					resource.TestCheckResourceAttr(dataSourceName, "service_name", "Amazon Virtual Private Cloud (Amazon VPC)"),
					resource.TestMatchResourceAttr(dataSourceName, "value", regexp.MustCompile(`^\d+$`)),
				),
			},
		},
	})
}

func TestAccServiceQuotasServiceQuotaDataSource_quotaName_Unset(t *testing.T) {
	const dataSourceName = "data.aws_servicequotas_service_quota.test"

	const serviceCode = "s3"
	const quotaCode = "L-DC2B2D3D"
	const quotaName = "Buckets"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			acctest.PreCheckPartitionHasService(servicequotas.EndpointsID, t)
			preCheckServiceQuotaUnset(serviceCode, quotaCode, t)
		},
		ErrorCheck: acctest.ErrorCheck(t, servicequotas.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceQuotaQuotaNameDataSourceConfig(serviceCode, quotaName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceAttrRegionalARNNoAccount(dataSourceName, "arn", "servicequotas", fmt.Sprintf("%s/%s", serviceCode, quotaCode)),
					resource.TestCheckResourceAttr(dataSourceName, "adjustable", "true"),
					resource.TestMatchResourceAttr(dataSourceName, "default_value", regexp.MustCompile(`^\d+$`)),
					resource.TestCheckResourceAttr(dataSourceName, "global_quota", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "quota_code", quotaCode),
					resource.TestCheckResourceAttr(dataSourceName, "quota_name", quotaName),
					resource.TestCheckResourceAttr(dataSourceName, "service_code", serviceCode),
					resource.TestCheckResourceAttr(dataSourceName, "service_name", "Amazon Simple Storage Service (Amazon S3)"),
					resource.TestMatchResourceAttr(dataSourceName, "value", regexp.MustCompile(`^\d+$`)),
					resource.TestCheckResourceAttrPair(dataSourceName, "value", dataSourceName, "default_value"),
				),
			},
		},
	})
}

func TestAccServiceQuotasServiceQuotaDataSource_PermissionError_quotaName(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			testAccPreCheck(t)
			acctest.PreCheckAssumeRoleARN(t)
		},
		ErrorCheck:   acctest.ErrorCheck(t, servicequotas.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config:      testAccServiceQuotaDataSourceConfig_PermissionError_QuotaName("elasticloadbalancing", "Application Load Balancers per Region"),
				ExpectError: regexp.MustCompile(`DEPENDENCY_ACCESS_DENIED_ERROR`),
			},
		},
	})
}

func testAccServiceQuotaQuotaCodeDataSourceConfig(serviceCode, quotaCode string) string {
	return fmt.Sprintf(`
data "aws_servicequotas_service_quota" "test" {
  quota_code   = %[1]q
  service_code = %[2]q
}
`, quotaCode, serviceCode)
}

func testAccServiceQuotaDataSourceConfig_PermissionError_QuotaCode(serviceCode, quotaCode string) string {
	policy := `{
  "Version": "2012-10-17",
  "Statement": [
    {
  	  "Effect": "Allow",
  	  "Action": [
  	    "servicequotas:GetServiceQuota"
  	  ],
  	  "Resource": "*"
    },
    {
  	  "Effect": "Deny",
  	  "Action": [
  	    "elasticloadbalancing:*"
  	  ],
  	  "Resource": "*"
    }
  ]
}`

	return acctest.ConfigCompose(
		acctest.ConfigAssumeRolePolicy(policy),
		fmt.Sprintf(`
data "aws_servicequotas_service_quota" "test" {
  service_code = %[1]q
  quota_code   = %[2]q
}
`, serviceCode, quotaCode))
}

func testAccServiceQuotaQuotaNameDataSourceConfig(serviceCode, quotaName string) string {
	return fmt.Sprintf(`
data "aws_servicequotas_service_quota" "test" {
  quota_name   = %[1]q
  service_code = %[2]q
}
`, quotaName, serviceCode)
}

func testAccServiceQuotaDataSourceConfig_PermissionError_QuotaName(serviceCode, quotaName string) string {
	policy := `{
  "Version": "2012-10-17",
  "Statement": [
    {
  	  "Effect": "Allow",
  	  "Action": [
  	    "servicequotas:ListServiceQuotas"
  	  ],
  	  "Resource": "*"
    },
    {
  	  "Effect": "Deny",
  	  "Action": [
  	    "elasticloadbalancing:*"
  	  ],
  	  "Resource": "*"
    }
  ]
}`

	return acctest.ConfigCompose(
		acctest.ConfigAssumeRolePolicy(policy),
		fmt.Sprintf(`
data "aws_servicequotas_service_quota" "test" {
  service_code = %[1]q
  quota_name   = %[2]q
}
`, serviceCode, quotaName))
}
