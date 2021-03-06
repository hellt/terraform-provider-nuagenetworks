package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatewayCreate,
		Read:   resourceGatewayRead,
		Update: resourceGatewayUpdate,
		Delete: resourceGatewayDelete,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"redundancy_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"template_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pending": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"permitted_action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"personality": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_gateway_vlanvnid": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"vtep": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_disc_gateway_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceGatewayCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Gateway object
	o := &vspk.Gateway{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("redundancy_group_id"); ok {
		o.RedundancyGroupID = attr.(string)
	}
	if attr, ok := d.GetOk("peer"); ok {
		o.Peer = attr.(string)
	}
	if attr, ok := d.GetOk("template_id"); ok {
		o.TemplateID = attr.(string)
	}
	if attr, ok := d.GetOk("pending"); ok {
		Pending := attr.(bool)
		o.Pending = &Pending
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("personality"); ok {
		o.Personality = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("use_gateway_vlanvnid"); ok {
		UseGatewayVLANVNID := attr.(bool)
		o.UseGatewayVLANVNID = &UseGatewayVLANVNID
	}
	if attr, ok := d.GetOk("vtep"); ok {
		o.Vtep = attr.(string)
	}
	if attr, ok := d.GetOk("auto_disc_gateway_id"); ok {
		o.AutoDiscGatewayID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_me"); ok {
		parent := &vspk.Me{ID: attr.(string)}
		err := parent.CreateGateway(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		err := parent.CreateGateway(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceGatewayRead(d, m)
}

func resourceGatewayRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Gateway{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("redundancy_group_id", o.RedundancyGroupID)
	d.Set("peer", o.Peer)
	d.Set("template_id", o.TemplateID)
	d.Set("pending", o.Pending)
	d.Set("permitted_action", o.PermittedAction)
	d.Set("personality", o.Personality)
	d.Set("description", o.Description)
	d.Set("enterprise_id", o.EnterpriseID)
	d.Set("entity_scope", o.EntityScope)
	d.Set("use_gateway_vlanvnid", o.UseGatewayVLANVNID)
	d.Set("vtep", o.Vtep)
	d.Set("auto_disc_gateway_id", o.AutoDiscGatewayID)
	d.Set("external_id", o.ExternalID)
	d.Set("system_id", o.SystemID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceGatewayUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Gateway{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("redundancy_group_id"); ok {
		o.RedundancyGroupID = attr.(string)
	}
	if attr, ok := d.GetOk("peer"); ok {
		o.Peer = attr.(string)
	}
	if attr, ok := d.GetOk("template_id"); ok {
		o.TemplateID = attr.(string)
	}
	if attr, ok := d.GetOk("pending"); ok {
		Pending := attr.(bool)
		o.Pending = &Pending
	}
	if attr, ok := d.GetOk("permitted_action"); ok {
		o.PermittedAction = attr.(string)
	}
	if attr, ok := d.GetOk("personality"); ok {
		o.Personality = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("use_gateway_vlanvnid"); ok {
		UseGatewayVLANVNID := attr.(bool)
		o.UseGatewayVLANVNID = &UseGatewayVLANVNID
	}
	if attr, ok := d.GetOk("vtep"); ok {
		o.Vtep = attr.(string)
	}
	if attr, ok := d.GetOk("auto_disc_gateway_id"); ok {
		o.AutoDiscGatewayID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceGatewayDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Gateway{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
