package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceCSNATPool() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCSNATPoolRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_link": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceCSNATPoolRead(d *schema.ResourceData, m interface{}) error {
	filteredCSNATPools := vspk.CSNATPoolsList{}
	err := &bambou.Error{}
	fetchFilter := &bambou.FetchingInfo{}

	filters, filtersOk := d.GetOk("filter")
	if filtersOk {
		fetchFilter = bambou.NewFetchingInfo()
		for _, v := range filters.(*schema.Set).List() {
			m := v.(map[string]interface{})
			if fetchFilter.Filter != "" {
				fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string), m["operator"].(string), m["value"].(string))
			} else {
				fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
			}

		}
	}
	parent := &vspk.Link{ID: d.Get("parent_link").(string)}
	filteredCSNATPools, err = parent.CSNATPools(fetchFilter)
	if err != nil {
		return err
	}

	CSNATPool := &vspk.CSNATPool{}

	if len(filteredCSNATPools) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredCSNATPools) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	CSNATPool = filteredCSNATPools[0]

	d.Set("end_address", CSNATPool.EndAddress)
	d.Set("start_address", CSNATPool.StartAddress)

	d.Set("id", CSNATPool.Identifier())
	d.Set("parent_id", CSNATPool.ParentID)
	d.Set("parent_type", CSNATPool.ParentType)
	d.Set("owner", CSNATPool.Owner)

	d.SetId(CSNATPool.Identifier())

	return nil
}
