package models

// VirtualIP

// VirtualIP
//proteus:generate
type VirtualIP struct {
	UUID                string         `json:"uuid,omitempty"`
	ParentUUID          string         `json:"parent_uuid,omitempty"`
	ParentType          string         `json:"parent_type,omitempty"`
	FQName              []string       `json:"fq_name,omitempty"`
	IDPerms             *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName         string         `json:"display_name,omitempty"`
	Annotations         *KeyValuePairs `json:"annotations,omitempty"`
	Perms2              *PermType2     `json:"perms2,omitempty"`
	VirtualIPProperties *VirtualIpType `json:"virtual_ip_properties,omitempty"`

	LoadbalancerPoolRefs        []*VirtualIPLoadbalancerPoolRef        `json:"loadbalancer_pool_refs,omitempty"`
	VirtualMachineInterfaceRefs []*VirtualIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
}

// VirtualIPLoadbalancerPoolRef references each other
type VirtualIPLoadbalancerPoolRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// VirtualIPVirtualMachineInterfaceRef references each other
type VirtualIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// MakeVirtualIP makes VirtualIP
func MakeVirtualIP() *VirtualIP {
	return &VirtualIP{
		//TODO(nati): Apply default
		UUID:                "",
		ParentUUID:          "",
		ParentType:          "",
		FQName:              []string{},
		IDPerms:             MakeIdPermsType(),
		DisplayName:         "",
		Annotations:         MakeKeyValuePairs(),
		Perms2:              MakePermType2(),
		VirtualIPProperties: MakeVirtualIpType(),
	}
}

// MakeVirtualIPSlice() makes a slice of VirtualIP
func MakeVirtualIPSlice() []*VirtualIP {
	return []*VirtualIP{}
}