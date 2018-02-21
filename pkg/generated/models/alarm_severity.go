
package models
// AlarmSeverity


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type AlarmSeverity int

// MakeAlarmSeverity makes AlarmSeverity
func MakeAlarmSeverity() AlarmSeverity {
    var data AlarmSeverity
    return data
}



// MakeAlarmSeveritySlice makes a slice of AlarmSeverity
func MakeAlarmSeveritySlice() []AlarmSeverity {
    return []AlarmSeverity{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AlarmSeverity) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AlarmSeverity) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmSeverity) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AlarmSeverity) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmSeverity) GetFQName() []string {
    return model.FQName
}

func (model *AlarmSeverity) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AlarmSeverity) GetParentType() string {
    return model.ParentType
}

func (model *AlarmSeverity) GetUuid() string {
    return model.UUID
}

func (model *AlarmSeverity) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AlarmSeverity) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AlarmSeverity) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AlarmSeverity) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AlarmSeverity) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *AlarmSeverity) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AlarmSeverity) UpdateReferences() error {
    return nil
}


