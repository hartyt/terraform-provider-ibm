---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# IBM Cloud Shell resources
{: #ibm-cloud-shell-resources}

Create, update, or delete IBM Cloud Shell resources.
You can reference the output parameters for each resource in other resources or data sources by using Terraform interpolation syntax.

Before you start working with your resource, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters) 
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_cloud_shell_account_settings`
{: #cloud_shell_account_settings}

Create, update, or delete an cloud_shell_account_settings.
{: shortdesc}

### Sample Terraform code
{: #cloud_shell_account_settings-sample}

```
resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings" {
  account_id = "account_id"
  tags = "[ \"tag3\" ]"
}
```

### Input parameters
{: #cloud_shell_account_settings-input}

Review the input parameters that you can specify for your resource. {: shortdesc}

|Name|Data type|Required/optional|Description|Forces new resource|
|----|-----------|-------|----------|--------------------|
|`account_id`|String|Required|The account ID in which the account settings belong to.|Yes|
|`rev`|String|Optional|Unique revision number for the settings object.|No|
|`default_enable_new_features`|Boolean|Optional|You can choose which Cloud Shell features are available in the account and whether any new features are enabled as they become available. The feature settings apply only to the enabled Cloud Shell locations.|No|
|`default_enable_new_regions`|Boolean|Optional|Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and session data are stored. By default, users are routed to the nearest available location.|No|
|`enabled`|Boolean|Optional|When enabled, Cloud Shell is available to all users in the account.|No|
|`features`|List|Optional|List of Cloud Shell features.|No|
|`regions`|List|Optional|List of Cloud Shell region settings.|No|
|`tags`|List|Optional|The list of tags associated with the cloud_shell_account_settings. **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.|No|

### Output parameters
{: #cloud_shell_account_settings-output}

Review the output parameters that you can access after your resource is created. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`id`|String|The unique identifier of the cloud_shell_account_settings.|
|`id`|String|Unique id of the settings object.|
|`created_at`|Integer|Creation timestamp in Unix epoch time.|
|`created_by`|String|IAM ID of creator.|
|`type`|String|Type of api response object.|
|`updated_at`|Integer|Timestamp of last update in Unix epoch time.|
|`updated_by`|String|IAM ID of last updater.|

### Import
{: #cloud_shell_account_settings-import}

`ibm_cloud_shell_account_settings` can be imported by ID

```
$ terraform import ibm_cloud_shell_account_settings.example sample-id
```

