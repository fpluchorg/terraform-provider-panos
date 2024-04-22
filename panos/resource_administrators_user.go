package panos

import (
	"github.com/fpluchorg/pango"
	"github.com/fpluchorg/pango/mgtconfig/user"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
)

// Local constants init
const (
	Name      = "name"
	Password  = "password"
	PublicKey = "public_key"
	RoleBased = "role_based"
	Template  = "template"
	Type      = "type"
)

// resourceAdministratorsUser create administrators user throw panorama or firewall
func resourceAdministratorsUser() *schema.Resource {
	return &schema.Resource{
		Create: createAdministratorsUser,
		Read:   readAdministratorsUser,
		Update: updateAdministratorsUser,
		Delete: deleteAdministratorsUser,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: administratorsUserSchema(),
	}
}

// administratorsUserSchema initialize the entry params
func administratorsUserSchema() map[string]*schema.Schema {
	ans := map[string]*schema.Schema{
		Name: &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		Template: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		PublicKey: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		RoleBased: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		Password: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		Type: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	}

	return ans
}

// parseUser parse the entry params to the user struct and get the template value
func parseUser(d *schema.ResourceData) (user.Entry, string) {
	tmpl := d.Get(Template).(string)

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

	return o, tmpl
}

// createAdministratorsUser this func will create the administrators user
func createAdministratorsUser(d *schema.ResourceData, meta interface{}) error {
	o, tmpl := parseUser(d)

	if tmpl != EmptyString {
		pano := meta.(*pango.Panorama)
		if err := pano.MGTConfig.User.Set(tmpl, o); err != nil {
			return err
		}
		d.SetId(buildPanoramaUserId(tmpl, o.Name))
	} else {
		fw := meta.(*pango.Firewall)
		if err := fw.MGTConfig.User.Set(o); err != nil {
			return err
		}
		d.SetId(o.Name)
	}

	return readAdministratorsUser(d, meta)
}

// buildPanoramaUserId this function will build the administrators user id in case of a template value in other word through panorama
func buildPanoramaUserId(a, c string) string {
	return strings.Join([]string{a, c}, IdSeparator)
}

// createAdministratorsUser this func will read the administrators users
func readAdministratorsUser(d *schema.ResourceData, meta interface{}) error {

	o, tmpl := parseUser(d)

	if tmpl != EmptyString {
		pano := meta.(*pango.Panorama)
		if _, err := pano.MGTConfig.User.Get(tmpl, o.Name); err != nil {
			if isObjectNotFound(err) {
				d.SetId(EmptyString)
				return nil
			}
			return err
		}
	} else {
		fw := meta.(*pango.Firewall)
		if _, err := fw.MGTConfig.User.Get(o.Name); err != nil {
			if isObjectNotFound(err) {
				d.SetId(EmptyString)
				return nil
			}
			return err
		}
	}

	return nil
}

// updateAdministratorsUser this func will update the administrators user
func updateAdministratorsUser(d *schema.ResourceData, meta interface{}) error {

	o, tmpl := parseUser(d)

	if tmpl != EmptyString {
		pano := meta.(*pango.Panorama)
		lo, err := pano.MGTConfig.User.Get(tmpl, o.Name)
		if err != nil {
			return err
		}
		lo.Copy(o)
		if err = pano.MGTConfig.User.Edit(tmpl, o); err != nil {
			return err
		}
	} else {
		fw := meta.(*pango.Firewall)
		lo, err := fw.MGTConfig.User.Get(o.Name)
		if err != nil {
			return err
		}
		lo.Copy(o)
		if err = fw.MGTConfig.User.Edit(o); err != nil {
			return err
		}
	}

	return readAdministratorsUser(d, meta)
}

// deleteAdministratorsUser this func will delete the administrators user
func deleteAdministratorsUser(d *schema.ResourceData, meta interface{}) error {
	o, tmpl := parseUser(d)

	if tmpl != EmptyString {
		pano := meta.(*pango.Panorama)
		if err := pano.MGTConfig.User.Delete(tmpl, o.Name); err != nil {
			if !isObjectNotFound(err) {
				return err
			}
		}
	} else {
		fw := meta.(*pango.Firewall)
		if err := fw.MGTConfig.User.Delete(o.Name); err != nil {
			if !isObjectNotFound(err) {
				return err
			}
		}
	}

	d.SetId(EmptyString)
	return nil
}
