// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure alarm.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAlarm() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAlarmUpdate,
		Read:   resourceSystemAlarmRead,
		Update: resourceSystemAlarmUpdate,
		Delete: resourceSystemAlarmDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"audible": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"admin_auth_failure_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),
							Optional:     true,
						},
						"admin_auth_lockout_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),
							Optional:     true,
						},
						"user_auth_failure_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),
							Optional:     true,
						},
						"user_auth_lockout_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),
							Optional:     true,
						},
						"replay_attempt_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),
							Optional:     true,
						},
						"self_test_failure_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
						},
						"log_full_warning_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),
							Optional:     true,
						},
						"encryption_failure_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),
							Optional:     true,
						},
						"decryption_failure_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),
							Optional:     true,
						},
						"fw_policy_violations": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1024),
										Optional:     true,
									},
									"src_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dst_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"src_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"dst_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
								},
							},
						},
						"fw_policy_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fw_policy_id_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),
							Optional:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemAlarmUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSystemAlarm(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlarm resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAlarm(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlarm resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAlarm")
	}

	return resourceSystemAlarmRead(d, m)
}

func resourceSystemAlarmDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemAlarm(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemAlarm resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAlarm(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemAlarm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAlarmRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadSystemAlarm(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlarm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAlarm(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlarm resource from API: %v", err)
	}
	return nil
}

func flattenSystemAlarmStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAlarmAudible(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAlarmGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenSystemAlarmGroupsId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "period"
		if cur_v, ok := i["period"]; ok {
			tmp["period"] = flattenSystemAlarmGroupsPeriod(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "admin_auth_failure_threshold"
		if cur_v, ok := i["admin-auth-failure-threshold"]; ok {
			tmp["admin_auth_failure_threshold"] = flattenSystemAlarmGroupsAdminAuthFailureThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "admin_auth_lockout_threshold"
		if cur_v, ok := i["admin-auth-lockout-threshold"]; ok {
			tmp["admin_auth_lockout_threshold"] = flattenSystemAlarmGroupsAdminAuthLockoutThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_auth_failure_threshold"
		if cur_v, ok := i["user-auth-failure-threshold"]; ok {
			tmp["user_auth_failure_threshold"] = flattenSystemAlarmGroupsUserAuthFailureThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_auth_lockout_threshold"
		if cur_v, ok := i["user-auth-lockout-threshold"]; ok {
			tmp["user_auth_lockout_threshold"] = flattenSystemAlarmGroupsUserAuthLockoutThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "replay_attempt_threshold"
		if cur_v, ok := i["replay-attempt-threshold"]; ok {
			tmp["replay_attempt_threshold"] = flattenSystemAlarmGroupsReplayAttemptThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "self_test_failure_threshold"
		if cur_v, ok := i["self-test-failure-threshold"]; ok {
			tmp["self_test_failure_threshold"] = flattenSystemAlarmGroupsSelfTestFailureThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_full_warning_threshold"
		if cur_v, ok := i["log-full-warning-threshold"]; ok {
			tmp["log_full_warning_threshold"] = flattenSystemAlarmGroupsLogFullWarningThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encryption_failure_threshold"
		if cur_v, ok := i["encryption-failure-threshold"]; ok {
			tmp["encryption_failure_threshold"] = flattenSystemAlarmGroupsEncryptionFailureThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "decryption_failure_threshold"
		if cur_v, ok := i["decryption-failure-threshold"]; ok {
			tmp["decryption_failure_threshold"] = flattenSystemAlarmGroupsDecryptionFailureThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fw_policy_violations"
		if cur_v, ok := i["fw-policy-violations"]; ok {
			tmp["fw_policy_violations"] = flattenSystemAlarmGroupsFwPolicyViolations(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fw_policy_id"
		if cur_v, ok := i["fw-policy-id"]; ok {
			tmp["fw_policy_id"] = flattenSystemAlarmGroupsFwPolicyId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fw_policy_id_threshold"
		if cur_v, ok := i["fw-policy-id-threshold"]; ok {
			tmp["fw_policy_id_threshold"] = flattenSystemAlarmGroupsFwPolicyIdThreshold(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemAlarmGroupsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsAdminAuthFailureThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsAdminAuthLockoutThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsUserAuthFailureThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsUserAuthLockoutThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsReplayAttemptThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsSelfTestFailureThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsLogFullWarningThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsEncryptionFailureThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsDecryptionFailureThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsFwPolicyViolations(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenSystemAlarmGroupsFwPolicyViolationsId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if cur_v, ok := i["threshold"]; ok {
			tmp["threshold"] = flattenSystemAlarmGroupsFwPolicyViolationsThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if cur_v, ok := i["src-ip"]; ok {
			tmp["src_ip"] = flattenSystemAlarmGroupsFwPolicyViolationsSrcIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if cur_v, ok := i["dst-ip"]; ok {
			tmp["dst_ip"] = flattenSystemAlarmGroupsFwPolicyViolationsDstIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port"
		if cur_v, ok := i["src-port"]; ok {
			tmp["src_port"] = flattenSystemAlarmGroupsFwPolicyViolationsSrcPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port"
		if cur_v, ok := i["dst-port"]; ok {
			tmp["dst_port"] = flattenSystemAlarmGroupsFwPolicyViolationsDstPort(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemAlarmGroupsFwPolicyViolationsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsFwPolicyViolationsThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsFwPolicyViolationsSrcIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAlarmGroupsFwPolicyViolationsDstIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAlarmGroupsFwPolicyViolationsSrcPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsFwPolicyViolationsDstPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsFwPolicyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAlarmGroupsFwPolicyIdThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSystemAlarm(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("status", flattenSystemAlarmStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("audible", flattenSystemAlarmAudible(o["audible"], d, "audible", sv)); err != nil {
		if !fortiAPIPatch(o["audible"]) {
			return fmt.Errorf("Error reading audible: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("groups", flattenSystemAlarmGroups(o["groups"], d, "groups", sv)); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenSystemAlarmGroups(o["groups"], d, "groups", sv)); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemAlarmFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemAlarmStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmAudible(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemAlarmGroupsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "period"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["period"], _ = expandSystemAlarmGroupsPeriod(d, i["period"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["period"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "admin_auth_failure_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["admin-auth-failure-threshold"], _ = expandSystemAlarmGroupsAdminAuthFailureThreshold(d, i["admin_auth_failure_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["admin-auth-failure-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "admin_auth_lockout_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["admin-auth-lockout-threshold"], _ = expandSystemAlarmGroupsAdminAuthLockoutThreshold(d, i["admin_auth_lockout_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["admin-auth-lockout-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_auth_failure_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["user-auth-failure-threshold"], _ = expandSystemAlarmGroupsUserAuthFailureThreshold(d, i["user_auth_failure_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["user-auth-failure-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_auth_lockout_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["user-auth-lockout-threshold"], _ = expandSystemAlarmGroupsUserAuthLockoutThreshold(d, i["user_auth_lockout_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["user-auth-lockout-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "replay_attempt_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["replay-attempt-threshold"], _ = expandSystemAlarmGroupsReplayAttemptThreshold(d, i["replay_attempt_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["replay-attempt-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "self_test_failure_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["self-test-failure-threshold"], _ = expandSystemAlarmGroupsSelfTestFailureThreshold(d, i["self_test_failure_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["self-test-failure-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_full_warning_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log-full-warning-threshold"], _ = expandSystemAlarmGroupsLogFullWarningThreshold(d, i["log_full_warning_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["log-full-warning-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encryption_failure_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["encryption-failure-threshold"], _ = expandSystemAlarmGroupsEncryptionFailureThreshold(d, i["encryption_failure_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["encryption-failure-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "decryption_failure_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["decryption-failure-threshold"], _ = expandSystemAlarmGroupsDecryptionFailureThreshold(d, i["decryption_failure_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["decryption-failure-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fw_policy_violations"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fw-policy-violations"], _ = expandSystemAlarmGroupsFwPolicyViolations(d, i["fw_policy_violations"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["fw-policy-violations"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fw_policy_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fw-policy-id"], _ = expandSystemAlarmGroupsFwPolicyId(d, i["fw_policy_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["fw-policy-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fw_policy_id_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fw-policy-id-threshold"], _ = expandSystemAlarmGroupsFwPolicyIdThreshold(d, i["fw_policy_id_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["fw-policy-id-threshold"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAlarmGroupsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsAdminAuthFailureThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsAdminAuthLockoutThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsUserAuthFailureThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsUserAuthLockoutThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsReplayAttemptThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsSelfTestFailureThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsLogFullWarningThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsEncryptionFailureThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsDecryptionFailureThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsFwPolicyViolations(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemAlarmGroupsFwPolicyViolationsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold"], _ = expandSystemAlarmGroupsFwPolicyViolationsThreshold(d, i["threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src-ip"], _ = expandSystemAlarmGroupsFwPolicyViolationsSrcIp(d, i["src_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst-ip"], _ = expandSystemAlarmGroupsFwPolicyViolationsDstIp(d, i["dst_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src-port"], _ = expandSystemAlarmGroupsFwPolicyViolationsSrcPort(d, i["src_port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["src-port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst-port"], _ = expandSystemAlarmGroupsFwPolicyViolationsDstPort(d, i["dst_port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dst-port"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAlarmGroupsFwPolicyViolationsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsFwPolicyViolationsThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsFwPolicyViolationsSrcIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsFwPolicyViolationsDstIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsFwPolicyViolationsSrcPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsFwPolicyViolationsDstPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsFwPolicyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlarmGroupsFwPolicyIdThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAlarm(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemAlarmStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("audible"); ok {
		if setArgNil {
			obj["audible"] = nil
		} else {
			t, err := expandSystemAlarmAudible(d, v, "audible", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["audible"] = t
			}
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		if setArgNil {
			obj["groups"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemAlarmGroups(d, v, "groups", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["groups"] = t
			}
		}
	}

	return &obj, nil
}
