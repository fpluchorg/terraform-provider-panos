package panos

import (
	"github.com/fpluchorg/pango"
	"github.com/fpluchorg/pango/mgtconfig/user"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourcePanoramaAdministratorsUser create panorama administrators user
func resourcePanoramaAdministratorsUser() *schema.Resource {
	return &schema.Resource{
		Create: createPanoramaAdministratorsUser,
		Read:   readPanoramaAdministratorsUser,
		Update: updatePanoramaAdministratorsUser,
		Delete: deletePanoramaAdministratorsUser,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: panoramaAdministratorsUserSchema(),
	}
}

// panoramaAdministratorsUserSchema initialize the entry params
func panoramaAdministratorsUserSchema() map[string]*schema.Schema {
	ans := map[string]*schema.Schema{
		Name: &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The name of the user to create the panorama administrators for.",
		},
		PublicKey: &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The public key of the user",
		},
		RoleBased: &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Role of the user",
			ValidateFunc: validateStringIn(
				"superuser",
				"superreader",
				"panorama-admin",
				"installeradmin",
				"gpio-esi-ro",
				"api-rules-automation",
				"api-object-automation",
				"api-vmauthkey-automation",
			),
		},

		Password: &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The password of the user",
		},
		Type: &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Type Role of the user",
			ValidateFunc: validateStringIn(
				"dynamic",
				"custom",
			),
		},
	}

	return ans
}

// parsePanoramaAdministratorsUser parse the entry params to the user struct
func parsePanoramaAdministratorsUser(d *schema.ResourceData) user.Entry {

	o := user.Entry{
		Name: d.Get(Name).(string),
	}

	if roleBased := d.Get(RoleBased); roleBased != nil {
		o.Role = roleBased.(string)
	}

	if publicKey := d.Get(PublicKey); publicKey != nil {
		o.PublicKey = publicKey.(string)
	}

	if roleType := d.Get(Type); roleType != nil {
		o.Type = roleType.(string)
	}

	if password := d.Get(Password); password != nil {
		o.PasswordHash = password.(string)
	}

	return o
}

// createPanoramaAdministratorsUser this func will create the panorama administrators user
func createPanoramaAdministratorsUser(d *schema.ResourceData, meta interface{}) error {
	o := parsePanoramaAdministratorsUser(d)

	pano := meta.(*pango.Panorama)
	if err := pano.MGTConfig.User.Set(EmptyString, o); err != nil {
		return err
	}
	d.SetId(buildPanoramaUserId(EmptyString, o.Name))

	return readPanoramaAdministratorsUser(d, meta)
}

// readPanoramaAdministratorsUser this func will read the panorama administrators users
func readPanoramaAdministratorsUser(d *schema.ResourceData, meta interface{}) error {

	o := parsePanoramaAdministratorsUser(d)

	pano := meta.(*pango.Panorama)
	if _, err := pano.MGTConfig.User.Get(EmptyString, o.Name); err != nil {
		if isObjectNotFound(err) {
			d.SetId(EmptyString)
			return nil
		}
		return err
	}

	return nil
}

// updatePanoramaAdministratorsUser this func will update the panorama administrators user
func updatePanoramaAdministratorsUser(d *schema.ResourceData, meta interface{}) error {

	o := parsePanoramaAdministratorsUser(d)
	pano := meta.(*pango.Panorama)
	lo, err := pano.MGTConfig.User.Get(EmptyString, o.Name)
	if err != nil {
		return err
	}
	lo.Copy(o)
	if err = pano.MGTConfig.User.Edit(EmptyString, o); err != nil {
		return err
	}

	return readPanoramaAdministratorsUser(d, meta)
}

// deletePanoramaAdministratorsUser this func will delete the panorama administrators user
func deletePanoramaAdministratorsUser(d *schema.ResourceData, meta interface{}) error {
	o := parsePanoramaAdministratorsUser(d)

	pano := meta.(*pango.Panorama)
	if err := pano.MGTConfig.User.Delete(EmptyString, o.Name); err != nil {
		if !isObjectNotFound(err) {
			return err
		}
	}

	d.SetId(EmptyString)
	return nil
}
