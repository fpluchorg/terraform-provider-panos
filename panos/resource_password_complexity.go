package panos

import (
	"github.com/fpluchorg/pango"
	"github.com/fpluchorg/pango/mgtconfig/passwordcomplexity"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Local constants init
const (
	Template                       = "template"
	minimumLength                  = "minimum_length"
	enabled                        = "enabled"
	minimumUppercaseLetters        = "minimum_uppercase_letters"
	minimumLowercaseLetters        = "minimum_lowercase_letters"
	minimumNumericLetters          = "minimum_numeric_letters"
	minimumSpecialCharacters       = "minimum_special_characters"
	blockRepeatedCharacters        = "block_repeated_characters"
	blockUsernameInclusion         = "block_username_inclusion"
	newPasswordDiffersByCharacters = "new_password_differs_by_characters"
	passwordChangeOnFirstLogin     = "password_change_on_first_login"
	passwordHistoryCount           = "password_history_count"
	passwordChangePeriodBlock      = "password_change_period_block"
	expirationPeriod               = "expiration_period"
	expirationWarningPeriod        = "expiration_warning_period"
	postExpirationAdminLoginCount  = "post_expiration_admin_login_count"
	postExpirationGracePeriod      = "post_expiration_grace_period"
)

// resourcePasswordComplexity create Password Complexity throw panorama or firewall
func resourcePasswordComplexity() *schema.Resource {
	return &schema.Resource{
		Create: createPasswordComplexity,
		Read:   readPasswordComplexity,
		Update: updatePasswordComplexity,
		Delete: deletePasswordComplexity,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: PasswordComplexitySchema(),
	}
}

// PasswordComplexitySchema initialize the entry params
func PasswordComplexitySchema() map[string]*schema.Schema {
	ans := map[string]*schema.Schema{
		Template: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		minimumLength: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		enabled: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		minimumUppercaseLetters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		minimumLowercaseLetters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		minimumNumericLetters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		minimumSpecialCharacters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		blockRepeatedCharacters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		blockUsernameInclusion: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		newPasswordDiffersByCharacters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		passwordChangeOnFirstLogin: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		passwordHistoryCount: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		passwordChangePeriodBlock: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		expirationPeriod: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		expirationWarningPeriod: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		postExpirationAdminLoginCount: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		postExpirationGracePeriod: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
	}

	return ans
}

// parsePasswordComplexity parse the entry params to the user struct and get the template value
func parsePasswordComplexity(d *schema.ResourceData) (passwordcomplexity.Entry, string) {
	tmpl := d.Get(Template).(string)

	o := passwordcomplexity.Entry{}

	if minimumLength := d.Get(minimumLength); minimumLength != nil {
		o.MinimumLength = minimumLength.(int)
	}

	if enabled := d.Get(enabled); enabled != nil {
		o.Enabled = enabled.(string)
	}

	if minimumUppercaseLetters := d.Get(minimumUppercaseLetters); minimumUppercaseLetters != nil {
		o.MinimumUppercaseLetters = minimumUppercaseLetters.(int)
	}

	if minimumLowercaseLetters := d.Get(minimumLowercaseLetters); minimumLowercaseLetters != nil {
		o.MinimumLowercaseLetters = minimumLowercaseLetters.(int)
	}

	if minimumNumericLetters := d.Get(minimumNumericLetters); minimumNumericLetters != nil {
		o.MinimumNumericLetters = minimumNumericLetters.(int)
	}

	if minimumSpecialCharacters := d.Get(minimumSpecialCharacters); minimumSpecialCharacters != nil {
		o.MinimumSpecialCharacters = minimumSpecialCharacters.(int)
	}

	if blockRepeatedCharacters := d.Get(blockRepeatedCharacters); blockRepeatedCharacters != nil {
		o.BlockRepeatedCharacters = blockRepeatedCharacters.(int)
	}

	if blockUsernameInclusion := d.Get(blockUsernameInclusion); blockUsernameInclusion != nil {
		o.BlockUsernameInclusion = blockUsernameInclusion.(string)
	}

	if newPasswordDiffersByCharacters := d.Get(newPasswordDiffersByCharacters); newPasswordDiffersByCharacters != nil {
		o.NewPasswordDiffersByCharacters = newPasswordDiffersByCharacters.(int)
	}

	if passwordChangeOnFirstLogin := d.Get(passwordChangeOnFirstLogin); passwordChangeOnFirstLogin != nil {
		o.PasswordChangeOnFirstLogin = passwordChangeOnFirstLogin.(string)
	}

	if passwordHistoryCount := d.Get(passwordHistoryCount); passwordHistoryCount != nil {
		o.PasswordHistoryCount = passwordHistoryCount.(int)
	}

	if passwordChangePeriodBlock := d.Get(passwordChangePeriodBlock); passwordChangePeriodBlock != nil {
		o.PasswordChangePeriodBlock = passwordChangePeriodBlock.(int)
	}

	if expirationPeriod := d.Get(expirationPeriod); expirationPeriod != nil {
		o.ExpirationPeriod = expirationPeriod.(int)
	}

	if expirationWarningPeriod := d.Get(expirationWarningPeriod); expirationWarningPeriod != nil {
		o.ExpirationWarningPeriod = expirationWarningPeriod.(int)
	}

	if postExpirationAdminLoginCount := d.Get(postExpirationAdminLoginCount); postExpirationAdminLoginCount != nil {
		o.PostExpirationAdminLoginCount = postExpirationAdminLoginCount.(int)
	}

	if postExpirationGracePeriod := d.Get(postExpirationGracePeriod); postExpirationGracePeriod != nil {
		o.PostExpirationGracePeriod = postExpirationGracePeriod.(int)
	}

	return o, tmpl
}

// createPasswordComplexity this func will create the Password Complexity
func createPasswordComplexity(d *schema.ResourceData, meta interface{}) error {
	o, tmpl := parsePasswordComplexity(d)

	if tmpl != EmptyString {
		pano := meta.(*pango.Panorama)
		if err := pano.MGTConfig.PasswordComplexity.Set(tmpl, o); err != nil {
			return err
		}
		d.SetId(tmpl)
	} else {
		fw := meta.(*pango.Firewall)
		if err := fw.MGTConfig.PasswordComplexity.Set(o); err != nil {
			return err
		}
		d.SetId(uuid.New().String())
	}

	return readPasswordComplexity(d, meta)
}

// readPasswordComplexity this func will read the Password Complexity
func readPasswordComplexity(d *schema.ResourceData, meta interface{}) error {

	_, tmpl := parsePasswordComplexity(d)

	if tmpl != EmptyString {
		pano := meta.(*pango.Panorama)
		if _, err := pano.MGTConfig.PasswordComplexity.Get(tmpl); err != nil {
			if isObjectNotFound(err) {
				d.SetId(tmpl)
				return nil
			}
			return err
		}
	} else {
		fw := meta.(*pango.Firewall)
		if _, err := fw.MGTConfig.PasswordComplexity.Get(); err != nil {
			if isObjectNotFound(err) {
				d.SetId(EmptyString)
				return nil
			}
			return err
		}
	}

	return nil
}

// updatePasswordComplexity this func will update the Password Complexity
func updatePasswordComplexity(d *schema.ResourceData, meta interface{}) error {

	o, tmpl := parsePasswordComplexity(d)

	if tmpl != EmptyString {
		pano := meta.(*pango.Panorama)
		lo, err := pano.MGTConfig.PasswordComplexity.Get(tmpl)
		if err != nil {
			return err
		}
		lo.Copy(o)
		if err = pano.MGTConfig.PasswordComplexity.Edit(tmpl, o); err != nil {
			return err
		}
	} else {
		fw := meta.(*pango.Firewall)
		lo, err := fw.MGTConfig.PasswordComplexity.Get()
		if err != nil {
			return err
		}
		lo.Copy(o)
		if err = fw.MGTConfig.PasswordComplexity.Edit(o); err != nil {
			return err
		}
	}

	return readPasswordComplexity(d, meta)
}

// deletePasswordComplexity this func will delete the Password Complexity
func deletePasswordComplexity(d *schema.ResourceData, meta interface{}) error {
	_, tmpl := parsePasswordComplexity(d)

	if tmpl != EmptyString {
		pano := meta.(*pango.Panorama)
		if err := pano.MGTConfig.PasswordComplexity.Delete(tmpl); err != nil {
			if !isObjectNotFound(err) {
				return err
			}
		}
	} else {
		fw := meta.(*pango.Firewall)
		if err := fw.MGTConfig.PasswordComplexity.Delete(); err != nil {
			if !isObjectNotFound(err) {
				return err
			}
		}
	}

	d.SetId(EmptyString)
	return nil
}
