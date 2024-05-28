package panos

import (
	"github.com/fpluchorg/pango/dev/settingmanagement"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePanoramaSettingManagement() *schema.Resource {
	return &schema.Resource{
		Create: createUpdatePanoramaSettingManagement,
		Read:   readPanoramaSettingManagement,
		Update: createUpdatePanoramaSettingManagement,
		Delete: deletePanoramaSettingManagement,

		Schema: map[string]*schema.Schema{
			"hostname_type_in_syslog": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Send HOSTNAME in Syslog",
			},
			"failed_attempts": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Failed Attempts",
			},
			"lockout_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Lockout Time (min)",
			},
			"max_session_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Max Session Count (number)",
			},
			"max_session_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Max Session Time (min)",
			},
			"idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Idle Timeout (min)",
			},
		},
	}
}

func parsePanoramaSettingManagement(d *schema.ResourceData) settingmanagement.Config {
	response := settingmanagement.Config{
		HostnameTypeInSyslog: d.Get("hostname_type_in_syslog").(string),
	}
	failedAttempts := d.Get("failed_attempts").(int)
	response.FailedAttempts = &failedAttempts
	lockoutTime := d.Get("lockout_time").(int)
	response.LockoutTime = &lockoutTime
	maxSessionCount := d.Get("max_session_count").(int)
	response.MaxSessionCount = &maxSessionCount
	maxSessionTime := d.Get("max_session_time").(int)
	response.MaxSessionTime = &maxSessionTime
	idleTimeout := d.Get("idle_timeout").(int)
	response.IdleTimeout = &idleTimeout
	return response
}

func createUpdatePanoramaSettingManagement(d *schema.ResourceData, meta interface{}) error {

	pano, err := panorama(meta, EmptyString)
	if err != nil {
		return err
	}

	o, err := pano.Device.SettingManagement.Get(EmptyString, EmptyString, EmptyString)
	if err != nil {
		return err
	}

	o.Merge(parsePanoramaSettingManagement(d))
	if err = pano.Device.SettingManagement.Edit(EmptyString, EmptyString, EmptyString, o); err != nil {
		return err
	}

	if _, err := pano.Device.SettingManagement.Get(EmptyString, EmptyString, EmptyString); err != nil {
		return err
	}
	d.SetId(Panorama)

	return readPanoramaSettingManagement(d, meta)
}

func readPanoramaSettingManagement(d *schema.ResourceData, meta interface{}) error {
	pano, err := panorama(meta, EmptyString)
	if err != nil {
		return err
	}
	o, err := pano.Device.SettingManagement.Get(EmptyString, EmptyString, EmptyString)
	if err != nil {
		return err
	}
	err = d.Set("hostname_type_in_syslog", o.HostnameTypeInSyslog)
	err = d.Set("failed_attempts", o.FailedAttempts)
	err = d.Set("lockout_time", o.LockoutTime)
	err = d.Set("max_session_count", o.MaxSessionCount)
	err = d.Set("max_session_time", o.MaxSessionTime)
	err = d.Set("idle_timeout", o.IdleTimeout)
	if err != nil {
		return err
	}
	return nil
}

func deletePanoramaSettingManagement(d *schema.ResourceData, meta interface{}) error {
	d.SetId(EmptyString)
	return nil
}
