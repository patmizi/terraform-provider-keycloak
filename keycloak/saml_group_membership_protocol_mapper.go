package keycloak

type SamlGroupMembershipProtocolMapper struct {
	Id            string
	Name          string
	RealmId       string
	ClientId      string
	ClientScopeId string

	GroupAttributeName      string
	SamlAttributeNameFormat string
	FriendlyName            string
	FullGroupPath           string
	SingleGroupAtrribute    string
}

func (mapper *SamlGroupMembershipProtocolMapper) convertToGenericProtocolMapper() *protocolMapper {
	return &protocolMapper{
		Id:             mapper.Id,
		Name:           mapper.Name,
		Protocol:       "saml",
		ProtocolMapper: "saml-group-membership-mapper",
		Config: map[string]string{
			attributeNameField:       mapper.GroupAttributeName,
			attributeNameFormatField: mapper.SamlAttributeNameFormat,
			friendlyNameField:        mapper.FriendlyName,
			fullPathField:            mapper.FullGroupPath,
			singleField:              mapper.SingleGroupAtrribute,
		},
	}
}
