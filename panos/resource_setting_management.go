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
		},
	}
}

func parsePanosSettingManagement(d *schema.ResourceData) settingmanagement.Config {
	response := settingmanagement.Config{
		Template: d.Get("template").(string),
	}
	enableLogHighDpLoad := d.Get("enable_log_high_dp_load").(bool)
	response.EnableLogHighDpLoad = &enableLogHighDpLoad
	enableHighSpeedLogForwarding := d.Get("enable_high_speed_log_forwarding").(bool)
	response.EnableHighSpeedLogForwarding = &enableHighSpeedLogForwarding
	supportUtf8ForLogOutput := d.Get("support_utf8_for_log_output").(bool)
	response.SupportUtf8ForLogOutput = &supportUtf8ForLogOutput
	trafficStopOnLogdbFull := d.Get("traffic_stop_on_logdb_full").(bool)
	response.TrafficStopOnLogdbFull = &trafficStopOnLogdbFull
	return response
}

func parseFwSettingManagement(d *schema.ResourceData) settingmanagement.Config {
	config := settingmanagement.Config{}
	enableLogHighDpLoad := d.Get("enable_log_high_dp_load").(bool)
	config.EnableLogHighDpLoad = &enableLogHighDpLoad
	supportUtf8ForLogOutput := d.Get("support_utf8_for_log_output").(bool)
	config.SupportUtf8ForLogOutput = &supportUtf8ForLogOutput
	trafficStopOnLogdbFull := d.Get("traffic_stop_on_logdb_full").(bool)
	config.TrafficStopOnLogdbFull = &trafficStopOnLogdbFull
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

		err = d.Set("enable_log_high_dp_load", o.EnableLogHighDpLoad)
		err = d.Set("enable_high_speed_log_forwarding", o.EnableHighSpeedLogForwarding)
		err = d.Set("support_utf8_for_log_output", o.SupportUtf8ForLogOutput)
		err = d.Set("traffic_stop_on_logdb_full", o.TrafficStopOnLogdbFull)
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
