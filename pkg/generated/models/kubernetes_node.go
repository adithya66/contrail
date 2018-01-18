package models

// KubernetesNode

import "encoding/json"

// KubernetesNode
type KubernetesNode struct {
	UUID                      string         `json:"uuid,omitempty"`
	ParentUUID                string         `json:"parent_uuid,omitempty"`
	ParentType                string         `json:"parent_type,omitempty"`
	ProvisioningLog           string         `json:"provisioning_log,omitempty"`
	ProvisioningStartTime     string         `json:"provisioning_start_time,omitempty"`
	Perms2                    *PermType2     `json:"perms2,omitempty"`
	FQName                    []string       `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName               string         `json:"display_name,omitempty"`
	ProvisioningProgress      int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState         string         `json:"provisioning_state,omitempty"`
	Annotations               *KeyValuePairs `json:"annotations,omitempty"`
}

// String returns json representation of the object
func (model *KubernetesNode) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeKubernetesNode makes KubernetesNode
func MakeKubernetesNode() *KubernetesNode {
	return &KubernetesNode{
		//TODO(nati): Apply default
		Perms2:                    MakePermType2(),
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		ProvisioningLog:           "",
		ProvisioningStartTime:     "",
		Annotations:               MakeKeyValuePairs(),
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		ProvisioningState:         "",
	}
}

// MakeKubernetesNodeSlice() makes a slice of KubernetesNode
func MakeKubernetesNodeSlice() []*KubernetesNode {
	return []*KubernetesNode{}
}
