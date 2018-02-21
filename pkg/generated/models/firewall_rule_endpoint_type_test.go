
package models

import "testing"


import "fmt"
func TestFirewallRuleEndpointType(t *testing.T) {
    model := MakeFirewallRuleEndpointType()
    fmt.Println(model)
}


func TestFirewallRuleEndpointTypeFQName(t *testing.T) {
    obj := MakeFirewallRuleEndpointType()
    if fqname := obj.GetFQName(); len(fqname != 0) {
        t.Errorf("Initial FQName is not empty for FirewallRuleEndpointType")
    }
    parent := "fake-parent"
    firstName := "first name"
    fakeFQName := []string{"some", "fake", firstName}
    newname := "fake-me"
    obj.SetFQName(parent, fakeFQName)
    fqname = obj.GetFQName()
    if len(fqname) != len(fakeFQName) {
        t.Errorf("Wrong FQName length, should be %v instead of %v", len(fakeFQName), len(fqname))
    }
    if fqname[len(fqname)-1] != newname {
        t.Errorf("Name should be %v after SetFQName", newname)
    }
}

