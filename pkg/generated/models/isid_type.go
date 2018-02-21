
package models
// IsidType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type IsidType int

// MakeIsidType makes IsidType
func MakeIsidType() IsidType {
    var data IsidType
    return data
}



// MakeIsidTypeSlice makes a slice of IsidType
func MakeIsidTypeSlice() []IsidType {
    return []IsidType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *IsidType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *IsidType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *IsidType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *IsidType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *IsidType) GetFQName() []string {
    return model.FQName
}

func (model *IsidType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *IsidType) GetParentType() string {
    return model.ParentType
}

func (model *IsidType) GetUuid() string {
    return model.UUID
}

func (model *IsidType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *IsidType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *IsidType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *IsidType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *IsidType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *IsidType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *IsidType) UpdateReferences() error {
    return nil
}


