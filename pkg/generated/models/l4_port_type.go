
package models
// L4PortType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type L4PortType int

// MakeL4PortType makes L4PortType
func MakeL4PortType() L4PortType {
    var data L4PortType
    return data
}



// MakeL4PortTypeSlice makes a slice of L4PortType
func MakeL4PortTypeSlice() []L4PortType {
    return []L4PortType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *L4PortType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *L4PortType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *L4PortType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *L4PortType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *L4PortType) GetFQName() []string {
    return model.FQName
}

func (model *L4PortType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *L4PortType) GetParentType() string {
    return model.ParentType
}

func (model *L4PortType) GetUuid() string {
    return model.UUID
}

func (model *L4PortType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *L4PortType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *L4PortType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *L4PortType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *L4PortType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *L4PortType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *L4PortType) UpdateReferences() error {
    return nil
}


