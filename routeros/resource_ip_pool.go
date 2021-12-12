package routeros

import (
	"fmt"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPPoolCreate,
		Read:   resourceIPPoolRead,
		Update: resourceIPPoolUpdate,
		Delete: resourceIPPoolDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ranges": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIPPoolCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)

	ip_pool := new(roscl.IPPool)
	ip_pool.Name = d.Get("name").(string)
	ip_pool.Ranges = d.Get("ranges").(string)

	res, err := c.CreateIPPool(ip_pool)
	if err != nil {
		return fmt.Errorf("error creating ip pool: %s", err.Error())
	}

	d.SetId(res.ID)
	return nil
}

func resourceIPPoolRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ip_pool, err := c.ReadIPPool(d.Id())

	if err != nil {
		return fmt.Errorf("error fetching ip pool: %s", err.Error())
	}

	d.SetId(ip_pool.ID)
	d.Set("name", ip_pool.Name)
	d.Set("ranges", ip_pool.Ranges)

	return nil

}

func resourceIPPoolUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ip_pool := new(roscl.IPPool)
	ip_pool.Name = d.Get("name").(string)
	ip_pool.Ranges = d.Get("ranges").(string)

	res, err := c.UpdateIPPool(d.Id(), ip_pool)

	if err != nil {
		return fmt.Errorf("error updating ip address: %s", err.Error())
	}

	d.SetId(res.ID)

	return nil
}

func resourceIPPoolDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ip_pool, _ := c.ReadIPPool(d.Id())
	err := c.DeleteIPPool(ip_pool)
	if err != nil {
		return fmt.Errorf("error deleting ip address: %s", err.Error())
	}
	d.SetId("")
	return nil
}
