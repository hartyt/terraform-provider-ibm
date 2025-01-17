package ibm

import (
	"context"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMAppIDThemeColor() *schema.Resource {
	return &schema.Resource{
		Description: "Colors of the App ID login widget",
		ReadContext: dataSourceIBMAppIDThemeColorRead,
		Schema: map[string]*schema.Schema{
			"tenant_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The AppID instance GUID",
			},
			"header_color": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMAppIDThemeColorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	appIDClient, err := meta.(ClientSession).AppIDAPI()

	if err != nil {
		return diag.FromErr(err)
	}

	tenantID := d.Get("tenant_id").(string)

	colors, _, err := appIDClient.GetThemeColorWithContext(ctx, &appid.GetThemeColorOptions{
		TenantID: &tenantID,
	})

	if err != nil {
		return diag.Errorf("Error getting AppID theme colors: %s", err)
	}

	if colors.HeaderColor != nil {
		d.Set("header_color", *colors.HeaderColor)
	}

	d.SetId(tenantID)

	return nil
}
