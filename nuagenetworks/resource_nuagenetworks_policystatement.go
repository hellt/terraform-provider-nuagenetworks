package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourcePolicyStatement() *schema.Resource {
	return &schema.Resource{
		Create: resourcePolicyStatementCreate,
		Read:   resourcePolicyStatementRead,
		Update: resourcePolicyStatementUpdate,
		Delete: resourcePolicyStatementDelete,
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
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_link": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePolicyStatementCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize PolicyStatement object
	o := &vspk.PolicyStatement{}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	parent := &vspk.Link{ID: d.Get("parent_link").(string)}
	err := parent.CreatePolicyStatement(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourcePolicyStatementRead(d, m)
}

func resourcePolicyStatementRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PolicyStatement{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("description", o.Description)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourcePolicyStatementUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PolicyStatement{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}

	o.Save()

	return nil
}

func resourcePolicyStatementDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PolicyStatement{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
