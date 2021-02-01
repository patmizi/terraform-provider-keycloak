package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var keycloakSamlGroupAttributeProtocolMapperNameFormats = []string{"Basic", "URI Reference", "Unspecified"}

func resourceKeycloakSamlGroupMembershipProtocolMapper() *schema.Resource {
	return &schema.Resource{
		Create: nil,
		Read:   nil,
		Update: nil,
		Delete: nil,
		Importer: &schema.ResourceImporter{
			// import a mapper tied to a client:
			// {{realmId}}/client/{{clientId}}/{{protocolMapperId}}
			// or a client scope:
			// {{realmId}}/client-scope/{{clientScopeId}}/{{protocolMapperId}}
			State: genericProtocolMapperImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"realm_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_id": {
				Type:          schema.TypeString,
				Required:      true,
				ForceNew:      true,
				ConflictsWith: []string{"client_scope_id"},
			},
			"client_scope_id": {
				Type:          schema.TypeString,
				Required:      true,
				ForceNew:      true,
				ConflictsWith: []string{"client_id"},
			},
			"group_attribute_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml_attribute_name_format": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(keycloakSamlGroupAttributeProtocolMapperNameFormats, false),
			},
		},
	}
}

// func mapFromDataToSamlGroupMembershipProtocolMapper(data *schema.ResourceData) *

// func resourceKeycloakSamlGroupMembershipProtocolMapperCreate(data *schema.ResourceData, meta interface{}) error {
// 	keycloakClient := meta.(*keycloak.KeycloakClient)

// 	samlUserAttributeMapper := mapFromDataToSamlGroupMembershipProtocolMapper(data)
// }
