package ipfix

import "fmt"

// SubTemplateList represents a list of zero or more instances of a structured data type, where the data type of each list element is the same and corresponds with a single Template Record.
// Examples include a structured data type composed of multiple pairs of ("MPLS label stack entry position", "MPLS label stack value"), a structured data type composed of performance metrics, and a structured data type composed of multiple pairs of IP address, etc.
type SubTemplateList struct {
	Semantic            uint8            //one of: SemanticsNoneOf, ExactlyOneOf, OneOrMoreOf, AllOf, Ordered or Undefined
	TemplateID          uint16           //Each of the newly generated Template Records is given a unique Template ID.  This uniqueness is local to the Transport Session and Observation Domain that generated the Template ID. Template IDs 0-255 are reserved for Template Sets, Options Template Sets, and other reserved Sets yet to be created.  Template IDs of Data Sets are numbered from 256 to 65535.  There are no constraints regarding the order of the Template ID allocation.
	AssociatedTemplates *ActiveTemplates //We can only begin to marshal/unmarshal the records when we have the whole template belonging to the TemplateID
	Records             []Record         //The list of Records
}

// NewSubTemplateList returns a new SubTemplateList.
func NewSubTemplateList(semantic uint8, templateid uint16) (*SubTemplateList, error) {
	if templateid < 256 {
		return nil, NewError(fmt.Sprintf("Can not have a template id <256, but got %d", templateid), ErrCritical)
	}
	if semantic >= 0x05 && semantic <= 0xFE {
		return nil, NewError(fmt.Sprintf("Semantic undefined: %d", semantic), ErrCritical)
	}
	return &SubTemplateList{
		Semantic:   semantic,
		TemplateID: templateid,
		Records:    make([]Record, 0, 0),
	}, nil
}

// Len returns the length of the field specifier, in octets.
func (stl *SubTemplateList) Len() uint16 {
	stllen := uint16(3)
	for _, listitem := range stl.Records {
		stllen += listitem.Len()
	}
	return stllen
}

// AssociateTemplates sets the template to be used marshalling/unmarshalling this SubTemplateList
func (stl *SubTemplateList) AssociateTemplates(at *ActiveTemplates) error {
	if at == nil {
		return NewError(fmt.Sprintf("Can not use nil as Template List"), ErrCritical)
	}
	stl.AssociatedTemplates = at
	return nil
}
