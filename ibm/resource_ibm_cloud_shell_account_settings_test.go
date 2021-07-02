// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/ibmcloudshellv1"
)

func TestAccIBMCloudShellAccountSettingsBasic(t *testing.T) {
	var conf ibmcloudshellv1.AccountSettings
	accountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCloudShellAccountSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsConfigBasic(accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCloudShellAccountSettingsExists("ibm_cloud_shell_account_settings.cloud_shell_account_settings", conf),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "account_id", accountID),
				),
			},
		},
	})
}

func TestAccIBMCloudShellAccountSettingsAllArgs(t *testing.T) {
	var conf ibmcloudshellv1.AccountSettings
	accountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	rev := fmt.Sprintf("tf_rev_%d", acctest.RandIntRange(10, 100))
	defaultEnableNewFeatures := "false"
	defaultEnableNewRegions := "true"
	enabled := "false"
	tags := "[ \"tag1\", \"tag2\" ]"
	revUpdate := fmt.Sprintf("tf_rev_%d", acctest.RandIntRange(10, 100))
	defaultEnableNewFeaturesUpdate := "true"
	defaultEnableNewRegionsUpdate := "false"
	enabledUpdate := "true"
	tagsUpdate := "[ \"tag3\" ]"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCloudShellAccountSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsConfig(accountID, rev, defaultEnableNewFeatures, defaultEnableNewRegions, enabled, tags),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCloudShellAccountSettingsExists("ibm_cloud_shell_account_settings.cloud_shell_account_settings", conf),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "rev", rev),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "default_enable_new_features", defaultEnableNewFeatures),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "default_enable_new_regions", defaultEnableNewRegions),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "enabled", enabled),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "tags.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCloudShellAccountSettingsConfig(accountID, revUpdate, defaultEnableNewFeaturesUpdate, defaultEnableNewRegionsUpdate, enabledUpdate, tagsUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "rev", revUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "default_enable_new_features", defaultEnableNewFeaturesUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "default_enable_new_regions", defaultEnableNewRegionsUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "enabled", enabledUpdate),
					resource.TestCheckResourceAttr("ibm_cloud_shell_account_settings.cloud_shell_account_settings", "tags.#", "1"),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_cloud_shell_account_settings.cloud_shell_account_settings",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"tags"},
			},
		},
	})
}

func testAccCheckIBMCloudShellAccountSettingsConfigBasic(accountID string) string {
	return fmt.Sprintf(`

		resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings" {
			account_id = "%s"
		}
	`, accountID)
}

func testAccCheckIBMCloudShellAccountSettingsConfig(accountID string, rev string, defaultEnableNewFeatures string, defaultEnableNewRegions string, enabled string, tags string) string {
	return fmt.Sprintf(`

		resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings" {
			account_id = "%s"
			rev = "%s"
			default_enable_new_features = %s
			default_enable_new_regions = %s
			enabled = %s
			features = { example: "object" }
			regions = { example: "object" }
			tags = %s
		}
	`, accountID, rev, defaultEnableNewFeatures, defaultEnableNewRegions, enabled, tags)
}

func testAccCheckIBMCloudShellAccountSettingsExists(n string, obj ibmcloudshellv1.AccountSettings) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		ibmCloudShellClient, err := testAccProvider.Meta().(ClientSession).IBMCloudShellV1()
		if err != nil {
			return err
		}

		getAccountSettingsOptions := &ibmcloudshellv1.GetAccountSettingsOptions{}

		parts, err := sepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAccountSettingsOptions.SetAccountID(parts[0])
		getAccountSettingsOptions.SetAccountID(parts[1])

		accountSettings, _, err := ibmCloudShellClient.GetAccountSettings(getAccountSettingsOptions)
		if err != nil {
			return err
		}

		obj = *accountSettings
		return nil
	}
}

func testAccCheckIBMCloudShellAccountSettingsDestroy(s *terraform.State) error {
	ibmCloudShellClient, err := testAccProvider.Meta().(ClientSession).IBMCloudShellV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cloud_shell_account_settings" {
			continue
		}

		getAccountSettingsOptions := &ibmcloudshellv1.GetAccountSettingsOptions{}

		parts, err := sepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAccountSettingsOptions.SetAccountID(parts[0])
		getAccountSettingsOptions.SetAccountID(parts[1])

		// Try to find the key
		_, response, err := ibmCloudShellClient.GetAccountSettings(getAccountSettingsOptions)

		if err == nil {
			return fmt.Errorf("cloud_shell_account_settings still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cloud_shell_account_settings (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
