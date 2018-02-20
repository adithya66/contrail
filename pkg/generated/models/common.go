package models

//ListSpec is configuration option for select query.
type ListSpec struct {
	Filter       Filter
	Limit        int
	Offset       int
	Detail       bool
	Count        bool
	Shared       bool
	ExcludeHrefs bool
	ParentFQName []string
	ParentType   string
	ParentUUIDs  []string
	BackRefUUIDs []string
	ObjectUUIDs  []string
	Fields       []string
}

//Filter is used to filter API response.
type Filter map[string][]string

//AppendValues appends filter values for key.
func (filter Filter) AppendValues(key string, values []string) {
	if filter == nil {
		return
	}
	if values == nil {
		return
	}
	f, ok := filter[key]
	if !ok {
		f = []string{}
	}
	filter[key] = append(f, values...)
}