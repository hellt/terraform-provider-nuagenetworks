package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceVirtualIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualIPCreate,
		Read:   resourceVirtualIPRead,
		Update: resourceVirtualIPUpdate,
		Delete: resourceVirtualIPDelete,
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
				Computed: true,
			},
			"ip_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "IPV4",
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_floating_ip_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_redirection_target": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport"},
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redirection_target"},
			},
		},
	}
}

func resourceVirtualIPCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize VirtualIP object
	o := &vspk.VirtualIP{
		VirtualIP: d.Get("virtual_ip").(string),
	}
	if attr, ok := d.GetOk("mac"); ok {
		o.MAC = attr.(string)
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_floating_ip_id"); ok {
		o.AssociatedFloatingIPID = attr.(string)
	}
	if attr, ok := d.GetOk("subnet_id"); ok {
		o.SubnetID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_redirection_target"); ok {
		parent := &vspk.RedirectionTarget{ID: attr.(string)}
		err := parent.CreateVirtualIP(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		err := parent.CreateVirtualIP(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceVirtualIPRead(d, m)
}

func resourceVirtualIPRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VirtualIP{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("mac", o.MAC)
	d.Set("ip_type", o.IPType)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("virtual_ip", o.VirtualIP)
	d.Set("entity_scope", o.EntityScope)
	d.Set("associated_floating_ip_id", o.AssociatedFloatingIPID)
	d.Set("subnet_id", o.SubnetID)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceVirtualIPUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VirtualIP{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.VirtualIP = d.Get("virtual_ip").(string)

	if attr, ok := d.GetOk("mac"); ok {
		o.MAC = attr.(string)
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_floating_ip_id"); ok {
		o.AssociatedFloatingIPID = attr.(string)
	}
	if attr, ok := d.GetOk("subnet_id"); ok {
		o.SubnetID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceVirtualIPDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VirtualIP{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
