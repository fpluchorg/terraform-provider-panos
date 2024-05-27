package panos

import (
	"github.com/fpluchorg/pango"
	"github.com/fpluchorg/pango/mgtconfig/passwordcomplexity"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourcePanoramaPasswordComplexity create Password Complexity throw panorama or firewall
func resourcePanoramaPasswordComplexity() *schema.Resource {
	return &schema.Resource{
		Create: createPanoramaPasswordComplexity,
		Read:   readPanoramaPasswordComplexity,
		Update: updatePanoramaPasswordComplexity,
		Delete: deletePanoramaPasswordComplexity,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: PanoramaPasswordComplexitySchema(),
	}
}

// PanoramaPasswordComplexitySchema initialize the entry params
func PanoramaPasswordComplexitySchema() map[string]*schema.Schema {
	ans := map[string]*schema.Schema{
		Template: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		minimumLength: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		enabled: &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		minimumUppercaseLetters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		minimumLowercaseLetters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		minimumNumericLetters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		minimumSpecialCharacters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		blockRepeatedCharacters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		blockUsernameInclusion: &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		newPasswordDiffersByCharacters: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		passwordChangeOnFirstLogin: &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		passwordHistoryCount: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		passwordChangePeriodBlock: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		expirationPeriod: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		expirationWarningPeriod: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		postExpirationAdminLoginCount: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		postExpirationGracePeriod: &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
	}

	return ans
}

// parsePanoramaPasswordComplexity parse the entry params to the user struct and get the template value
func parsePanoramaPasswordComplexity(d *schema.ResourceData) passwordcomplexity.Entry {
	o := passwordcomplexity.Entry{}

	if minimumLength := d.Get(minimumLength); minimumLength != nil {
		o.MinimumLength = minimumLength.(int)
	}

	if enabled := d.Get(enabled); enabled != nil {
		o.Enabled = enabled.(bool)
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
		o.BlockUsernameInclusion = blockUsernameInclusion.(bool)
	}

	if newPasswordDiffersByCharacters := d.Get(newPasswordDiffersByCharacters); newPasswordDiffersByCharacters != nil {
		o.NewPasswordDiffersByCharacters = newPasswordDiffersByCharacters.(int)
	}

	if passwordChangeOnFirstLogin := d.Get(passwordChangeOnFirstLogin); passwordChangeOnFirstLogin != nil {
		o.PasswordChangeOnFirstLogin = passwordChangeOnFirstLogin.(bool)
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

	return o
}

// createPanoramaPasswordComplexity this func will create the Panorama Password Complexity
func createPanoramaPasswordComplexity(d *schema.ResourceData, meta interface{}) error {
	o := parsePanoramaPasswordComplexity(d)

	pano := meta.(*pango.Panorama)
	if err := pano.MGTConfig.PasswordComplexity.Set(EmptyString, o); err != nil {
		return err
	}
	d.SetId(uuid.New().String())

	return nil
}

// readPanoramaPasswordComplexity this func will read the Panorama Password Complexity
func readPanoramaPasswordComplexity(d *schema.ResourceData, meta interface{}) error {

	parsePanoramaPasswordComplexity(d)

	pano := meta.(*pango.Panorama)
	if _, err := pano.MGTConfig.PasswordComplexity.Get(EmptyString); err != nil {
		if isObjectNotFound(err) {
			d.SetId(EmptyString)
			return nil
		}
		return err
	}

	return nil
}

// updatePanoramaPasswordComplexity this func will update the Panorama Password Complexity
func updatePanoramaPasswordComplexity(d *schema.ResourceData, meta interface{}) error {

	o := parsePanoramaPasswordComplexity(d)

	pano := meta.(*pango.Panorama)
	lo, err := pano.MGTConfig.PasswordComplexity.Get(EmptyString)
	if err != nil {
		return err
	}
	lo.Copy(o)
	if err = pano.MGTConfig.PasswordComplexity.Edit(EmptyString, o); err != nil {
		return err
	}

	return nil
}

// deletePanoramaPasswordComplexity this func will delete the Panorama Password Complexity
func deletePanoramaPasswordComplexity(d *schema.ResourceData, meta interface{}) error {
	parsePanoramaPasswordComplexity(d)

	pano := meta.(*pango.Panorama)
	if err := pano.MGTConfig.PasswordComplexity.Delete(EmptyString); err != nil {
		if !isObjectNotFound(err) {
			return err
		}
	}

	d.SetId(EmptyString)
	return nil
}
