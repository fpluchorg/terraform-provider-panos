package panos

import (
	"github.com/fpluchorg/pango/dev/settingmanagement"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSettingManagement() *schema.Resource {
	return &schema.Resource{
		Create: createUpdateSettingManagement,
		Read:   readSettingManagement,
		Update: createUpdateSettingManagement,
		Delete: deleteSettingManagement,

		Schema: map[string]*schema.Schema{
			"template": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Template name if exist",
			},
			"enable_log_high_dp_load": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Enable Log on High DP Load",
			},
			"enable_high_speed_log_forwarding": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Enable High Speed Log Forwarding",
			},
			"support_utf8_for_log_output": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Support UTF-8 For Log Output",
			},
			"traffic_stop_on_logdb_full": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Stop Traffic when LogDb Ful",
			},
			"idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Idle Timeout (min)",
			},
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
			"threat_vault_access": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Threat Vault Access",
			},
		},
	}
}

func parsePanosSettingManagement(d *schema.ResourceData) settingmanagement.Config {
	config := settingmanagement.Config{
		Template:             d.Get("template").(string),
		HostnameTypeInSyslog: d.Get("hostname_type_in_syslog").(string),
	}
	failedAttempts := d.Get("failed_attempts").(int)
	config.FailedAttempts = &failedAttempts
	lockoutTime := d.Get("lockout_time").(int)
	config.LockoutTime = &lockoutTime
	maxSessionCount := d.Get("max_session_count").(int)
	config.MaxSessionCount = &maxSessionCount
	maxSessionTime := d.Get("max_session_time").(int)
	config.MaxSessionTime = &maxSessionTime
	idleTimeout := d.Get("idle_timeout").(int)
	config.IdleTimeout = &idleTimeout
	enableLogHighDpLoad := d.Get("enable_log_high_dp_load").(bool)
	config.EnableLogHighDpLoad = &enableLogHighDpLoad
	enableHighSpeedLogForwarding := d.Get("enable_high_speed_log_forwarding").(bool)
	config.EnableHighSpeedLogForwarding = &enableHighSpeedLogForwarding
	supportUtf8ForLogOutput := d.Get("support_utf8_for_log_output").(bool)
	config.SupportUtf8ForLogOutput = &supportUtf8ForLogOutput
	trafficStopOnLogdbFull := d.Get("traffic_stop_on_logdb_full").(bool)
	config.TrafficStopOnLogdbFull = &trafficStopOnLogdbFull
	threatVaultAccess := d.Get("threat_vault_access").(bool)
	config.ThreatVaultAccess = &threatVaultAccess
	return config
}

func parseFwSettingManagement(d *schema.ResourceData) settingmanagement.Config {
	config := settingmanagement.Config{
		HostnameTypeInSyslog: d.Get("hostname_type_in_syslog").(string),
	}
	failedAttempts := d.Get("failed_attempts").(int)
	config.FailedAttempts = &failedAttempts
	lockoutTime := d.Get("lockout_time").(int)
	config.LockoutTime = &lockoutTime
	maxSessionCount := d.Get("max_session_count").(int)
	config.MaxSessionCount = &maxSessionCount
	maxSessionTime := d.Get("max_session_time").(int)
	config.MaxSessionTime = &maxSessionTime
	idleTimeout := d.Get("idle_timeout").(int)
	config.IdleTimeout = &idleTimeout
	enableLogHighDpLoad := d.Get("enable_log_high_dp_load").(bool)
	config.EnableLogHighDpLoad = &enableLogHighDpLoad
	supportUtf8ForLogOutput := d.Get("support_utf8_for_log_output").(bool)
	config.SupportUtf8ForLogOutput = &supportUtf8ForLogOutput
	trafficStopOnLogdbFull := d.Get("traffic_stop_on_logdb_full").(bool)
	config.TrafficStopOnLogdbFull = &trafficStopOnLogdbFull
	threatVaultAccess := d.Get("threat_vault_access").(bool)
	config.ThreatVaultAccess = &threatVaultAccess
	return config
}

func createUpdateSettingManagement(d *schema.ResourceData, meta interface{}) error {
	template := d.Get("template").(string)
	if template == EmptyString {
		fw, err := firewall(meta, EmptyString)
		if err != nil {
			return err
		}

		o, err := fw.Device.SettingManagement.Get()
		if err != nil {
			return err
		}

		o.Merge(parseFwSettingManagement(d))
		if err = fw.Device.SettingManagement.Edit(o); err != nil {
			return err
		}

		_, err = fw.Device.SettingManagement.Get()
		if err != nil {
			o = parseFwSettingManagement(d)
		} else {
			o.Merge(parseFwSettingManagement(d))
		}

		d.SetId(Device)

	} else {
		pano, err := panorama(meta, EmptyString)
		if err != nil {
			return err
		}

		o, err := pano.Device.SettingManagement.Get(template, EmptyString, EmptyString)
		if err != nil {
			o = parsePanosSettingManagement(d)
		} else {
			o.Merge(parsePanosSettingManagement(d))
		}

		if err = pano.Device.SettingManagement.Edit(template, EmptyString, EmptyString, o); err != nil {
			return err
		}

		if _, err := pano.Device.SettingManagement.Get(template, EmptyString, EmptyString); err != nil {
			return err
		}

		d.SetId(template)
	}

	return readSettingManagement(d, meta)
}

func readSettingManagement(d *schema.ResourceData, meta interface{}) error {
	template := d.Get("template").(string)
	if template == EmptyString {
		fw, err := firewall(meta, EmptyString)
		if err != nil {
			return err
		}
		o, err := fw.Device.SettingManagement.Get()
		if err != nil {
			return err
		}

		err = d.Set("enable_high_speed_log_forwarding", o.EnableHighSpeedLogForwarding)
		err = d.Set("enable_log_high_dp_load", o.EnableLogHighDpLoad)
		err = d.Set("support_utf8_for_log_output", o.SupportUtf8ForLogOutput)
		err = d.Set("traffic_stop_on_logdb_full", o.TrafficStopOnLogdbFull)
		err = d.Set("hostname_type_in_syslog", o.HostnameTypeInSyslog)
		err = d.Set("failed_attempts", o.FailedAttempts)
		err = d.Set("lockout_time", o.LockoutTime)
		err = d.Set("max_session_count", o.MaxSessionCount)
		err = d.Set("max_session_time", o.MaxSessionTime)
		err = d.Set("idle_timeout", o.IdleTimeout)
		err = d.Set("threat_vault_access", o.ThreatVaultAccess)
		if err != nil {
			return err
		}
	} else {
		pano, err := panorama(meta, EmptyString)
		if err != nil {
			return err
		}
		o, err := pano.Device.SettingManagement.Get(template, EmptyString, EmptyString)
		if err != nil {
			return err
		}

		err = d.Set("enable_log_high_dp_load", o.EnableLogHighDpLoad)
		err = d.Set("support_utf8_for_log_output", o.SupportUtf8ForLogOutput)
		err = d.Set("traffic_stop_on_logdb_full", o.TrafficStopOnLogdbFull)
		err = d.Set("hostname_type_in_syslog", o.HostnameTypeInSyslog)
		err = d.Set("failed_attempts", o.FailedAttempts)
		err = d.Set("lockout_time", o.LockoutTime)
		err = d.Set("max_session_count", o.MaxSessionCount)
		err = d.Set("max_session_time", o.MaxSessionTime)
		err = d.Set("idle_timeout", o.IdleTimeout)
		err = d.Set("threat_vault_access", o.ThreatVaultAccess)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteSettingManagement(d *schema.ResourceData, meta interface{}) error {
	d.SetId(EmptyString)
	return nil
}
