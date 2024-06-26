package panos

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/fpluchorg/pango"
	"github.com/fpluchorg/pango/netw/routing/protocol/bgp/exp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePanoramaBgpExportRuleGroup() *schema.Resource {
	return &schema.Resource{
		Create: createUpdatePanoramaBgpExportRuleGroup,
		Read:   readPanoramaBgpExportRuleGroup,
		Update: createUpdatePanoramaBgpExportRuleGroup,
		Delete: deletePanoramaBgpExportRuleGroup,

		Schema: map[string]*schema.Schema{
			"template":       templateSchema(true),
			"template_stack": templateStackSchema(),
			"virtual_router": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"position_keyword": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "",
				ValidateFunc: validateStringIn(movementKeywords()...),
				ForceNew:     true,
			},
			"position_reference": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"rule": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"enable": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"used_by": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"match_as_path_regex": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_community_regex": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_extended_community_regex": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_med": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"match_route_table": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateStringIn("", exp.MatchRouteTableUnicast, exp.MatchRouteTableMulticast, exp.MatchRouteTableBoth),
						},
						"match_address_prefix": {
							Type:     schema.TypeSet,
							Optional: true,
							Set:      resourceMatchAddressPrefixHash,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:     schema.TypeString,
										Required: true,
									},
									"exact": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"match_next_hops": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"match_from_peers": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"action": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      exp.ActionAllow,
							ValidateFunc: validateStringIn(exp.ActionAllow, exp.ActionDeny),
						},
						"local_preference": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"med": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"next_hop": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"origin": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateStringIn("", exp.OriginIgp, exp.OriginEgp, exp.OriginIncomplete),
						},
						"as_path_limit": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"as_path_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateStringIn("", exp.AsPathTypeNone, exp.AsPathTypeRemove, exp.AsPathTypePrepend, exp.AsPathTypeRemoveAndPrepend),
						},
						"as_path_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"community_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateStringIn("", exp.CommunityTypeNone, exp.CommunityTypeRemoveAll, exp.CommunityTypeRemoveRegex, exp.CommunityTypeAppend, exp.CommunityTypeOverwrite),
						},
						"community_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"extended_community_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateStringIn("", exp.CommunityTypeNone, exp.CommunityTypeRemoveAll, exp.CommunityTypeRemoveRegex, exp.CommunityTypeAppend, exp.CommunityTypeOverwrite),
						},
						"extended_community_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func parsePanoramaBgpExportRuleGroup(d *schema.ResourceData) (string, string, string, string, int, []exp.Entry) {
	tmpl := d.Get("template").(string)
	ts := d.Get("template_stack").(string)
	vr := d.Get("virtual_router").(string)
	oRule := d.Get("position_reference").(string)
	move := movementAtoi(d.Get("position_keyword").(string))

	rlist := d.Get("rule").([]interface{})
	ans := make([]exp.Entry, 0, len(rlist))
	for i := range rlist {
		elm := rlist[i].(map[string]interface{})
		o := exp.Entry{
			Name:                        elm["name"].(string),
			Enable:                      elm["enable"].(bool),
			UsedBy:                      asStringList(elm["used_by"].([]interface{})),
			MatchAsPathRegex:            elm["match_as_path_regex"].(string),
			MatchCommunityRegex:         elm["match_community_regex"].(string),
			MatchExtendedCommunityRegex: elm["match_extended_community_regex"].(string),
			MatchMed:                    elm["match_med"].(string),
			MatchRouteTable:             elm["match_route_table"].(string),
			MatchNextHop:                asStringList(elm["match_next_hops"].([]interface{})),
			MatchFromPeer:               asStringList(elm["match_from_peers"].([]interface{})),
			Action:                      elm["action"].(string),
			LocalPreference:             elm["local_preference"].(string),
			Med:                         elm["med"].(string),
			NextHop:                     elm["next_hop"].(string),
			Origin:                      elm["origin"].(string),
			AsPathLimit:                 elm["as_path_limit"].(int),
			AsPathType:                  elm["as_path_type"].(string),
			AsPathValue:                 elm["as_path_value"].(string),
			CommunityType:               elm["community_type"].(string),
			CommunityValue:              elm["community_value"].(string),
			ExtendedCommunityType:       elm["extended_community_type"].(string),
			ExtendedCommunityValue:      elm["extended_community_value"].(string),
		}

		sl := elm["match_address_prefix"].(*schema.Set).List()
		if len(sl) != 0 {
			o.MatchAddressPrefix = make(map[string]bool)
			for i := range sl {
				sli := sl[i].(map[string]interface{})
				o.MatchAddressPrefix[sli["prefix"].(string)] = sli["exact"].(bool)
			}
		}

		ans = append(ans, o)
	}

	return tmpl, ts, vr, oRule, move, ans
}

func parsePanoramaBgpExportRuleGroupId(v string) (string, string, string, int, string, []string) {
	t := strings.Split(v, IdSeparator)
	move, _ := strconv.Atoi(t[3])
	joined, _ := base64.StdEncoding.DecodeString(t[5])
	names := strings.Split(string(joined), "\n")
	return t[0], t[1], t[2], move, t[4], names
}

func buildPanoramaBgpExportRuleGroupId(a, b, c string, d int, e string, f []exp.Entry) string {
	var buf bytes.Buffer
	for i := range f {
		if i != 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(f[i].Name)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())

	return strings.Join([]string{a, b, c, strconv.Itoa(d), e, enc}, IdSeparator)
}

func createUpdatePanoramaBgpExportRuleGroup(d *schema.ResourceData, meta interface{}) error {
	var err error

	pano := meta.(*pango.Panorama)
	tmpl, ts, vr, oRule, move, list := parsePanoramaBgpExportRuleGroup(d)

	if !movementIsRelative(move) && oRule != "" {
		return fmt.Errorf("'position_reference' must be empty for non-relative movement")
	}

	if err = pano.Network.BgpExport.Edit(tmpl, ts, vr, list[0]); err != nil {
		return err
	}
	dl := make([]interface{}, len(list)-1)
	for i := 1; i < len(list); i++ {
		dl = append(dl, list[i])
	}
	_ = pano.Network.BgpExport.Delete(tmpl, ts, vr, dl...)
	if err = pano.Network.BgpExport.Set(tmpl, ts, vr, list[1:len(list)]...); err != nil {
		return err
	}
	if err = pano.Network.BgpExport.MoveGroup(tmpl, ts, vr, move, oRule, list...); err != nil {
		return err
	}

	d.SetId(buildPanoramaBgpExportRuleGroupId(tmpl, ts, vr, move, oRule, list))
	return readPanoramaBgpExportRuleGroup(d, meta)
}

func readPanoramaBgpExportRuleGroup(d *schema.ResourceData, meta interface{}) error {
	var err error

	pano := meta.(*pango.Panorama)
	tmpl, ts, vr, move, oRule, rules := parsePanoramaBgpExportRuleGroupId(d.Id())

	list, err := pano.Network.BgpExport.GetList(tmpl, ts, vr)
	if err != nil {
		return err
	}

	fIdx, oIdx := -1, -1
	for i := range list {
		if list[i] == rules[0] {
			fIdx = i
		} else if list[i] == oRule {
			oIdx = i
		}
		if fIdx != -1 && oIdx != -1 {
			break
		}
	}

	if fIdx == -1 {
		// First policy is MIA, but others may be present, so report an
		// empty ruleset to force rules to be recreated.
		d.Set("rule", nil)
		return nil
	} else if oIdx == -1 && movementIsRelative(move) {
		return fmt.Errorf("Can't position group %s %q: rule is not present", movementItoa(move), oRule)
	}

	d.Set("template", tmpl)
	d.Set("template_stack", ts)
	d.Set("virtual_router", vr)
	d.Set("position_keyword", movementItoa(move))
	if groupPositionIsOk(move, fIdx, oIdx, list, rules) {
		d.Set("position_reference", oRule)
	} else {
		d.Set("position_reference", "(incorrect group positioning)")
	}

	ilist := make([]interface{}, 0, len(rules))
	for i := 0; i+fIdx < len(list) && i < len(rules); i++ {
		if list[i+fIdx] != rules[i] {
			// Rules must be contiguous.
			break
		}
		o, err := pano.Network.BgpExport.Get(tmpl, ts, vr, rules[i])
		if err != nil {
			return err
		}
		aps := &schema.Set{
			F: resourceMatchAddressPrefixHash,
		}
		for k, v := range o.MatchAddressPrefix {
			aps.Add(map[string]interface{}{
				"prefix": k,
				"exact":  v,
			})
		}
		m := map[string]interface{}{
			"name":                           o.Name,
			"enable":                         o.Enable,
			"used_by":                        o.UsedBy,
			"match_as_path_regex":            o.MatchAsPathRegex,
			"match_community_regex":          o.MatchCommunityRegex,
			"match_extended_community_regex": o.MatchExtendedCommunityRegex,
			"match_med":                      o.MatchMed,
			"match_route_table":              o.MatchRouteTable,
			"match_address_prefix":           aps,
			"match_next_hops":                o.MatchNextHop,
			"match_from_peers":               o.MatchFromPeer,
			"action":                         o.Action,
			"local_preference":               o.LocalPreference,
			"med":                            o.Med,
			"next_hop":                       o.NextHop,
			"origin":                         o.Origin,
			"as_path_limit":                  o.AsPathLimit,
			"as_path_type":                   o.AsPathType,
			"as_path_value":                  o.AsPathValue,
			"community_type":                 o.CommunityType,
			"community_value":                o.CommunityValue,
			"extended_community_type":        o.ExtendedCommunityType,
			"extended_community_value":       o.ExtendedCommunityValue,
		}

		ilist = append(ilist, m)
	}

	if err = d.Set("rule", ilist); err != nil {
		log.Printf("[WARN] Error setting 'rule' param for %q: %s", d.Id(), err)
	}

	return nil
}

func deletePanoramaBgpExportRuleGroup(d *schema.ResourceData, meta interface{}) error {
	pano := meta.(*pango.Panorama)
	tmpl, ts, vr, _, _, list := parsePanoramaBgpExportRuleGroupId(d.Id())

	ilist := make([]interface{}, len(list))
	for i := range list {
		ilist[i] = list[i]
	}

	if err := pano.Network.BgpExport.Delete(tmpl, ts, vr, ilist...); err != nil {
		if isObjectNotFound(err) {
			return err
		}
	}

	d.SetId("")
	return nil
}
