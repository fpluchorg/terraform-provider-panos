package panos

import (
	"github.com/fpluchorg/pango/dev/general"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	Panorama = "panorama"
	Device   = "device"
)

func resourceGeneralSettings() *schema.Resource {
	return &schema.Resource{
		Create: createUpdateGeneralSettings,
		Read:   readGeneralSettings,
		Update: createUpdateGeneralSettings,
		Delete: deleteGeneralSettings,

		Schema: map[string]*schema.Schema{
			Template: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Template in case of panorama device",
			},
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The hostname",
			},
			"timezone": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Timezone",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Domain",
			},
			"login_banner": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Login Banner",
			},
			"update_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "updates.paloaltonetworks.com",
				Description: "PANOS update server",
			},
			"verify_update_server": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Verify update server identity",
			},
			"proxy_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxy_user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"proxy_password_enc": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"panorama_primary": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Primary Panorama server address",
			},
			"panorama_secondary": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Secondary Panorama server address",
			},
			"dns_primary": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Primary DNS IP address",
			},
			"dns_secondary": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Secondary DNS IP address",
			},
			"ntp_primary_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Primary NTP server",
			},
			"ntp_primary_auth_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "NTP auth type (none, autokey, symmetric-key)",
				ValidateFunc: validateStringIn("none", "autokey", "symmetric-key"),
			},
			"ntp_primary_key_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "NTP symmetric-key key ID",
			},
			"ntp_primary_algorithm": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "NTP symmetric-key algorithm (sha1 or md5)",
				ValidateFunc: validateStringIn("sha1", "md5"),
			},
			"ntp_primary_auth_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "NTP symmetric-key auth key",
			},
			"ntp_secondary_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Secondary NTP server",
			},
			"ntp_secondary_auth_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "NTP auth type (none, autokey, symmetric-key)",
				ValidateFunc: validateStringIn("none", "autokey", "symmetric-key"),
			},
			"ntp_secondary_key_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "NTP symmetric-key key ID",
			},
			"ntp_secondary_algorithm": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "NTP symmetric-key algorithm (sha1 or md5)",
				ValidateFunc: validateStringIn("sha1", "md5"),
			},
			"ntp_secondary_auth_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "NTP symmetric-key auth key",
			},
		},
	}
}

func parseGeneralSettings(d *schema.ResourceData) general.Config {
	return general.Config{
		Template:              d.Get("template").(string),
		Hostname:              d.Get("hostname").(string),
		Timezone:              d.Get("timezone").(string),
		Domain:                d.Get("domain").(string),
		UpdateServer:          d.Get("update_server").(string),
		VerifyUpdateServer:    d.Get("verify_update_server").(bool),
		ProxyServer:           d.Get("proxy_server").(string),
		ProxyPort:             d.Get("proxy_port").(int),
		ProxyUser:             d.Get("proxy_user").(string),
		ProxyPassword:         d.Get("proxy_password").(string),
		PanoramaPrimary:       d.Get("panorama_primary").(string),
		PanoramaSecondary:     d.Get("panorama_secondary").(string),
		DnsPrimary:            d.Get("dns_primary").(string),
		DnsSecondary:          d.Get("dns_secondary").(string),
		NtpPrimaryAddress:     d.Get("ntp_primary_address").(string),
		NtpPrimaryAuthType:    d.Get("ntp_primary_auth_type").(string),
		NtpPrimaryKeyId:       d.Get("ntp_primary_key_id").(int),
		NtpPrimaryAlgorithm:   d.Get("ntp_primary_algorithm").(string),
		NtpPrimaryAuthKey:     d.Get("ntp_primary_auth_key").(string),
		NtpSecondaryAddress:   d.Get("ntp_secondary_address").(string),
		NtpSecondaryAuthType:  d.Get("ntp_secondary_auth_type").(string),
		NtpSecondaryKeyId:     d.Get("ntp_secondary_key_id").(int),
		NtpSecondaryAlgorithm: d.Get("ntp_secondary_algorithm").(string),
		NtpSecondaryAuthKey:   d.Get("ntp_secondary_auth_key").(string),
		LoginBanner:           d.Get("login_banner").(string),
	}
}

func createUpdateGeneralSettings(d *schema.ResourceData, meta interface{}) error {
	template := d.Get("template").(string)
	if template == EmptyString {
		fw, err := firewall(meta, EmptyString)
		if err != nil {
			return err
		}

		o, err := fw.Device.GeneralSettings.Get()
		if err != nil {
			return err
		}

		o.Merge(parseGeneralSettings(d))
		if err = fw.Device.GeneralSettings.Edit(o); err != nil {
			return err
		}

		lo, err := fw.Device.GeneralSettings.Get()
		if err != nil {
			o = parseGeneralSettings(d)
		} else {
			o.Merge(parseGeneralSettings(d))
		}

		d.SetId(Device)
		err = d.Set("proxy_password_enc", lo.ProxyPassword)
		if err != nil {
			return err
		}
	} else {
		pano, err := panorama(meta, EmptyString)
		if err != nil {
			return err
		}

		o, err := pano.Device.GeneralSettings.Get(template, EmptyString, EmptyString)
		if err != nil {
			o = parseGeneralSettings(d)
		} else {
			o.Merge(parseGeneralSettings(d))
		}

		if err = pano.Device.GeneralSettings.Edit(template, EmptyString, EmptyString, o); err != nil {
			return err
		}

		lo, err := pano.Device.GeneralSettings.Get(template, EmptyString, EmptyString)
		if err != nil {
			return err
		}

		d.SetId(template)
		err = d.Set("proxy_password_enc", lo.ProxyPassword)
		if err != nil {
			return err
		}
	}

	return readGeneralSettings(d, meta)
}

func readGeneralSettings(d *schema.ResourceData, meta interface{}) error {
	template := d.Get("template").(string)
	if template == EmptyString {
		fw, err := firewall(meta, EmptyString)
		if err != nil {
			return err
		}
		o, err := fw.Device.GeneralSettings.Get()
		if err != nil {
			// I don't think you can delete the general settings from a firewall,
			// so any error is a real error.
			return err
		}

		err = d.Set("hostname", o.Hostname)
		err = d.Set("timezone", o.Timezone)
		err = d.Set("domain", o.Domain)
		err = d.Set("update_server", o.UpdateServer)
		err = d.Set("verify_update_server", o.VerifyUpdateServer)
		err = d.Set("proxy_server", o.ProxyServer)
		err = d.Set("proxy_port", o.ProxyPort)
		err = d.Set("proxy_user", o.ProxyUser)
		if d.Get("proxy_password_enc").(string) != o.ProxyPassword {
			err = d.Set("proxy_password", "(incorrect proxy password)")
		}
		err = d.Set("panorama_primary", o.PanoramaPrimary)
		err = d.Set("panorama_secondary", o.PanoramaSecondary)
		err = d.Set("dns_primary", o.DnsPrimary)
		err = d.Set("dns_secondary", o.DnsSecondary)
		err = d.Set("ntp_primary_address", o.NtpPrimaryAddress)
		err = d.Set("ntp_primary_auth_type", o.NtpPrimaryAuthType)
		err = d.Set("ntp_primary_key_id", o.NtpPrimaryKeyId)
		err = d.Set("ntp_primary_algorithm", o.NtpPrimaryAlgorithm)
		err = d.Set("ntp_primary_auth_key", o.NtpPrimaryAuthKey)
		err = d.Set("ntp_secondary_address", o.NtpSecondaryAddress)
		err = d.Set("ntp_secondary_auth_type", o.NtpSecondaryAuthType)
		err = d.Set("ntp_secondary_key_id", o.NtpSecondaryKeyId)
		err = d.Set("ntp_secondary_algorithm", o.NtpSecondaryAlgorithm)
		err = d.Set("ntp_secondary_auth_key", o.NtpSecondaryAuthKey)
		err = d.Set("login_banner", o.LoginBanner)
		if err != nil {
			return err
		}
	} else {
		pano, err := panorama(meta, EmptyString)
		if err != nil {
			return err
		}
		o, err := pano.Device.GeneralSettings.Get(template, EmptyString, EmptyString)
		if err != nil {
			// I don't think you can delete the general settings from a firewall,
			// so any error is a real error.
			return err
		}

		err = d.Set("hostname", o.Hostname)
		err = d.Set("timezone", o.Timezone)
		err = d.Set("domain", o.Domain)
		err = d.Set("update_server", o.UpdateServer)
		err = d.Set("verify_update_server", o.VerifyUpdateServer)
		err = d.Set("proxy_server", o.ProxyServer)
		err = d.Set("proxy_port", o.ProxyPort)
		err = d.Set("proxy_user", o.ProxyUser)
		if d.Get("proxy_password_enc").(string) != o.ProxyPassword {
			err = d.Set("proxy_password", "(incorrect proxy password)")
		}
		err = d.Set("panorama_primary", o.PanoramaPrimary)
		err = d.Set("panorama_secondary", o.PanoramaSecondary)
		err = d.Set("dns_primary", o.DnsPrimary)
		err = d.Set("dns_secondary", o.DnsSecondary)
		err = d.Set("ntp_primary_address", o.NtpPrimaryAddress)
		err = d.Set("ntp_primary_auth_type", o.NtpPrimaryAuthType)
		err = d.Set("ntp_primary_key_id", o.NtpPrimaryKeyId)
		err = d.Set("ntp_primary_algorithm", o.NtpPrimaryAlgorithm)
		err = d.Set("ntp_primary_auth_key", o.NtpPrimaryAuthKey)
		err = d.Set("ntp_secondary_address", o.NtpSecondaryAddress)
		err = d.Set("ntp_secondary_auth_type", o.NtpSecondaryAuthType)
		err = d.Set("ntp_secondary_key_id", o.NtpSecondaryKeyId)
		err = d.Set("ntp_secondary_algorithm", o.NtpSecondaryAlgorithm)
		err = d.Set("ntp_secondary_auth_key", o.NtpSecondaryAuthKey)
		err = d.Set("login_banner", o.LoginBanner)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteGeneralSettings(d *schema.ResourceData, meta interface{}) error {
	d.SetId(EmptyString)
	return nil
}
