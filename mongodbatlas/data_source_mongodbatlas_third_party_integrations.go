package mongodbatlas

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceMongoDBAtlasThirdPartyIntegrations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMongoDBAtlasThirdPartyIntegrationsRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license_key": {
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"account_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"write_token": {
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"read_token": {
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"api_key": {
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_key": {
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"api_token": {
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"team_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"channel_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"routing_key": {
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"flow_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"org_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"secret": {
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
					},
				},
			},
		},
	}
}

func dataSourceMongoDBAtlasThirdPartyIntegrationsRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
