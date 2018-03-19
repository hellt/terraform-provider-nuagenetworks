package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceHostInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceHostInterfaceCreate,
		Read:   resourceHostInterfaceRead,
		Update: resourceHostInterfaceUpdate,
		Delete: resourceHostInterfaceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vport_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vport_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tier_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_decision_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_floating_ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attached_network_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attached_network_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_vport": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceHostInterfaceCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize HostInterface object
	o := &vspk.HostInterface{}
	if attr, ok := d.GetOk("mac"); ok {
		o.MAC = attr.(string)
	}
	if attr, ok := d.GetOk("ip_address"); ok {
		o.IPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("vport_id"); ok {
		o.VPortID = attr.(string)
	}
	if attr, ok := d.GetOk("vport_name"); ok {
		o.VPortName = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("network_name"); ok {
		o.NetworkName = attr.(string)
	}
	if attr, ok := d.GetOk("tier_id"); ok {
		o.TierID = attr.(string)
	}
	if attr, ok := d.GetOk("policy_decision_id"); ok {
		o.PolicyDecisionID = attr.(string)
	}
	if attr, ok := d.GetOk("domain_id"); ok {
		o.DomainID = attr.(string)
	}
	if attr, ok := d.GetOk("domain_name"); ok {
		o.DomainName = attr.(string)
	}
	if attr, ok := d.GetOk("zone_id"); ok {
		o.ZoneID = attr.(string)
	}
	if attr, ok := d.GetOk("zone_name"); ok {
		o.ZoneName = attr.(string)
	}
	if attr, ok := d.GetOk("associated_floating_ip_address"); ok {
		o.AssociatedFloatingIPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("attached_network_id"); ok {
		o.AttachedNetworkID = attr.(string)
	}
	if attr, ok := d.GetOk("attached_network_type"); ok {
		o.AttachedNetworkType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.VPort{ID: d.Get("parent_vport").(string)}
	err := parent.CreateHostInterface(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceHostInterfaceRead(d, m)
}

func resourceHostInterfaceRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.HostInterface{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("mac", o.MAC)
	d.Set("ip_address", o.IPAddress)
	d.Set("vport_id", o.VPortID)
	d.Set("vport_name", o.VPortName)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("gateway", o.Gateway)
	d.Set("netmask", o.Netmask)
	d.Set("network_name", o.NetworkName)
	d.Set("tier_id", o.TierID)
	d.Set("entity_scope", o.EntityScope)
	d.Set("policy_decision_id", o.PolicyDecisionID)
	d.Set("domain_id", o.DomainID)
	d.Set("domain_name", o.DomainName)
	d.Set("zone_id", o.ZoneID)
	d.Set("zone_name", o.ZoneName)
	d.Set("associated_floating_ip_address", o.AssociatedFloatingIPAddress)
	d.Set("attached_network_id", o.AttachedNetworkID)
	d.Set("attached_network_type", o.AttachedNetworkType)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceHostInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.HostInterface{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("mac"); ok {
		o.MAC = attr.(string)
	}
	if attr, ok := d.GetOk("ip_address"); ok {
		o.IPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("vport_id"); ok {
		o.VPortID = attr.(string)
	}
	if attr, ok := d.GetOk("vport_name"); ok {
		o.VPortName = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("network_name"); ok {
		o.NetworkName = attr.(string)
	}
	if attr, ok := d.GetOk("tier_id"); ok {
		o.TierID = attr.(string)
	}
	if attr, ok := d.GetOk("policy_decision_id"); ok {
		o.PolicyDecisionID = attr.(string)
	}
	if attr, ok := d.GetOk("domain_id"); ok {
		o.DomainID = attr.(string)
	}
	if attr, ok := d.GetOk("domain_name"); ok {
		o.DomainName = attr.(string)
	}
	if attr, ok := d.GetOk("zone_id"); ok {
		o.ZoneID = attr.(string)
	}
	if attr, ok := d.GetOk("zone_name"); ok {
		o.ZoneName = attr.(string)
	}
	if attr, ok := d.GetOk("associated_floating_ip_address"); ok {
		o.AssociatedFloatingIPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("attached_network_id"); ok {
		o.AttachedNetworkID = attr.(string)
	}
	if attr, ok := d.GetOk("attached_network_type"); ok {
		o.AttachedNetworkType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceHostInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.HostInterface{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
