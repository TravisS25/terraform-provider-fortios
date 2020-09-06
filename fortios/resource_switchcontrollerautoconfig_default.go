// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure default auto-config QoS policy for FortiSwitch.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerAutoConfigDefault() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAutoConfigDefaultUpdate,
		Read:   resourceSwitchControllerAutoConfigDefaultRead,
		Update: resourceSwitchControllerAutoConfigDefaultUpdate,
		Delete: resourceSwitchControllerAutoConfigDefaultDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fgt_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"isl_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"icl_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchControllerAutoConfigDefaultUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerAutoConfigDefault(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigDefault resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerAutoConfigDefault(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigDefault resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerAutoConfigDefault")
	}

	return resourceSwitchControllerAutoConfigDefaultRead(d, m)
}

func resourceSwitchControllerAutoConfigDefaultDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerAutoConfigDefault(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerAutoConfigDefault resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAutoConfigDefaultRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerAutoConfigDefault(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigDefault resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAutoConfigDefault(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigDefault resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAutoConfigDefaultFgtPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigDefaultIslPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigDefaultIclPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerAutoConfigDefault(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fgt_policy", flattenSwitchControllerAutoConfigDefaultFgtPolicy(o["fgt-policy"], d, "fgt_policy")); err != nil {
		if !fortiAPIPatch(o["fgt-policy"]) {
			return fmt.Errorf("Error reading fgt_policy: %v", err)
		}
	}

	if err = d.Set("isl_policy", flattenSwitchControllerAutoConfigDefaultIslPolicy(o["isl-policy"], d, "isl_policy")); err != nil {
		if !fortiAPIPatch(o["isl-policy"]) {
			return fmt.Errorf("Error reading isl_policy: %v", err)
		}
	}

	if err = d.Set("icl_policy", flattenSwitchControllerAutoConfigDefaultIclPolicy(o["icl-policy"], d, "icl_policy")); err != nil {
		if !fortiAPIPatch(o["icl-policy"]) {
			return fmt.Errorf("Error reading icl_policy: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerAutoConfigDefaultFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerAutoConfigDefaultFgtPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigDefaultIslPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigDefaultIclPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerAutoConfigDefault(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fgt_policy"); ok {
		t, err := expandSwitchControllerAutoConfigDefaultFgtPolicy(d, v, "fgt_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-policy"] = t
		}
	}

	if v, ok := d.GetOk("isl_policy"); ok {
		t, err := expandSwitchControllerAutoConfigDefaultIslPolicy(d, v, "isl_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isl-policy"] = t
		}
	}

	if v, ok := d.GetOk("icl_policy"); ok {
		t, err := expandSwitchControllerAutoConfigDefaultIclPolicy(d, v, "icl_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icl-policy"] = t
		}
	}

	return &obj, nil
}
