---
layout: "ibm"
page_title: "IBM : cloud_shell_account_settings"
description: |-
  Manages cloud_shell_account_settings.
subcategory: "IBM Cloud Shell"
---

# ibm\_cloud_shell_account_settings

Provides a resource for cloud_shell_account_settings. This allows cloud_shell_account_settings to be created, updated and deleted.

## Example Usage

```hcl
resource "cloud_shell_account_settings" "cloud_shell_account_settings" {
  account_id = "account_id"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required, Forces new resource, string) The account ID in which the account settings belong to.
* `rev` - (Optional, string) Unique revision number for the settings object.
* `default_enable_new_features` - (Optional, bool) You can choose which Cloud Shell features are available in the account and whether any new features are enabled as they become available. The feature settings apply only to the enabled Cloud Shell locations.
* `default_enable_new_regions` - (Optional, bool) Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and session data are stored. By default, users are routed to the nearest available location.
* `enabled` - (Optional, bool) When enabled, Cloud Shell is available to all users in the account.
* `features` - (Optional, List) List of Cloud Shell features.
  * `enabled` - (Optional, bool) State of the feature.
  * `key` - (Optional, string) Name of the feature.
* `regions` - (Optional, List) List of Cloud Shell region settings.
  * `enabled` - (Optional, bool) State of the region.
  * `key` - (Optional, string) Name of the region.
* `tags` - (Optional, array of strings) Tags associated with the cloud_shell_account_settings.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloud_shell_account_settings.
* `id` - Unique id of the settings object.
* `created_at` - Creation timestamp in Unix epoch time.
* `created_by` - IAM ID of creator.
* `type` - Type of api response object.
* `updated_at` - Timestamp of last update in Unix epoch time.
* `updated_by` - IAM ID of last updater.

## Import

You can import the `cloud_shell_account_settings` resource by using `account_id`.
The `account_id` property can be formed from `account_id`, and `account_id` in the following format:

```
<account_id>/<account_id>
```
* `account_id`: A string. The account ID in which the account settings belong to.
* `account_id`: A string. The account ID in which the account settings belong to.

```
$ terraform import cloud_shell_account_settings.cloud_shell_account_settings <account_id>/<account_id>
```
