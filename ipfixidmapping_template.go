// This is the template for ipfixidmapping, all lines until the marker will be removed before output
// +build ignore

//***GENERATEMARKER***
//Generated by generateipfixidmapping.go using ipfixidmapping_template
//Generation timestamp: {{.TimeStamp}} 
//Do not edit manually, run go generate
package ipfixmessage

import (
	"fmt"
	"sync"
)

var(
	customIPFIXIDMap = map[uint32]map[uint16]customFieldType{}
	customMapLock sync.Mutex
)

type customFieldType struct{
	FieldLength uint16
	Description string
	FieldVal FieldValue
}

//RegisterCustomField allows runtime addition of new elements
func RegisterCustomField(enterpriseid uint32,elementid uint16,fieldlen uint16,desc string,val FieldValue)error{
	if enterpriseid==0{
		return NewError("Can not use IANA specified enterprise id.",ErrCritical)
	}
	customMapLock.Lock()
	defer customMapLock.Unlock()
	if _,found:=customIPFIXIDMap[enterpriseid];!found{
		customIPFIXIDMap[enterpriseid]=make(map[uint16]customFieldType)
	}
	customIPFIXIDMap[enterpriseid][elementid]=customFieldType{
		FieldLength:fieldlen,
		Description:desc,
		FieldVal:val,
	}

 return nil
}

//UnregisterCustomField removes a custom element
func UnregisterCustomField(enterpriseid uint32,elementid uint16)error{
	customMapLock.Lock()
	defer customMapLock.Unlock()
	if _,found:=customIPFIXIDMap[enterpriseid];!found{
		return NewError(fmt.Sprintf("Did not find enterprise id %d",enterpriseid),ErrCritical)
 	}
	if _,found:=customIPFIXIDMap[enterpriseid][elementid];!found{
		return NewError(fmt.Sprintf("Did not find enterprise id %d, element id",enterpriseid,elementid),ErrCritical)
 	}
	delete(customIPFIXIDMap[enterpriseid],elementid)
	if len(customIPFIXIDMap[enterpriseid])==0{
		delete(customIPFIXIDMap,enterpriseid)
	}
	
 return nil
}

//GetCustomField returns a custom element
func GetCustomField(enterpriseid uint32,elementid uint16)(customFieldType,error){
	customMapLock.Lock()
	defer customMapLock.Unlock()
	if _,found:=customIPFIXIDMap[enterpriseid];!found{
		return customFieldType{}, NewError(fmt.Sprintf("No such element: E%did%d", enterpriseid, elementid),ErrCritical)
 	}
	if field,found:=customIPFIXIDMap[enterpriseid][elementid];!found{
		return customFieldType{}, NewError(fmt.Sprintf("No such element: E%did%d", enterpriseid, elementid),ErrCritical)
 	}else{
		return field, nil
	 }
}

// NewFieldValueByID returns an empty FieldValue that matches the enterprise id and element id
func NewFieldValueByID(enterpriseid uint32, elementid uint16) (FieldValue, error) {
	switch enterpriseid {
{{range $_, $enterpriseid := .EnterpriseOrder}}
case {{$enterpriseid}}: // {{index $.Sources $enterpriseid}}
{{$elements := (index $.Elements $enterpriseid)}}
    switch elementid { {{range $_, $elementid:= (index $.ElementsOrder $enterpriseid)}}
        case {{$elementid}}:
        return &{{(index $elements $elementid).GoFieldValue}}{}, nil // {{(index $elements $elementid).Name}}{{end}}
        default:
           return nil,NewError(fmt.Sprintf("No such element: E%did%d",enterpriseid,elementid),ErrCritical)
    }
{{end}}
	default://Checking if we registered any custom elements
	    custfield,err:=GetCustomField(enterpriseid,elementid)
		if err!=nil{
			return nil,err
		}
	 	retval,err:=getNewFieldValue(custfield)
	 	if err!=nil{
		 	return nil,err
	 	}
	 	return retval,nil
	}
}

// FieldLengthByID returns the default length that matches the enterprise id and element id
func FieldLengthByID(enterpriseid uint32, elementid uint16) (uint16, error) {
	switch enterpriseid {
{{range $_, $enterpriseid := .EnterpriseOrder}}
case {{$enterpriseid}}: // {{index $.Sources $enterpriseid}}
{{$elements := (index $.Elements $enterpriseid)}}
    switch elementid { {{range $_, $elementid:= (index $.ElementsOrder $enterpriseid)}}
        case {{$elementid}}:
        return {{(index $elements $elementid).GoFieldLength}}, nil // {{(index $elements $elementid).Name}}{{end}}
        default:
           return 0,NewError(fmt.Sprintf("No such element: E%did%d",enterpriseid,elementid),ErrCritical)
    }
{{end}}
	default://Checking if we registered any custom elements
	    custfield,err:=GetCustomField(enterpriseid,elementid)
		if err!=nil{
			return 0,err
		}
	 	return custfield.FieldLength,nil
	}
    return 0,nil
}

// FieldDescriptionByID returns the given semantic description that matches the enterprise id and element id
func FieldDescriptionByID(enterpriseid uint32, elementid uint16) (string, error) {
	switch enterpriseid {
{{range $_, $enterpriseid := .EnterpriseOrder}}
case {{$enterpriseid}}: // {{index $.Sources $enterpriseid}}
{{$elements := (index $.Elements $enterpriseid)}}
    switch elementid { {{range $_, $elementid:= (index $.ElementsOrder $enterpriseid)}}
        case {{$elementid}}:
        return "{{(index $elements $elementid).Name}}", nil{{end}}
        default:
           return "",NewError(fmt.Sprintf("No such element: E%did%d",enterpriseid,elementid),ErrCritical)
    }
{{end}}
	default://Checking if we registered any custom elements
	    custfield,err:=GetCustomField(enterpriseid,elementid)
		if err!=nil{
			return "",err
		}
	 	return custfield.Description,nil
	}
    return "",nil
}


//getNewFieldValue returns a new empty FieldValue based on the FieldValue provided.
//This is so the same pointer does not get re-used
func getNewFieldValue(val interface{})(FieldValue,error){
				switch val.(type) {
					case *FieldValueUnsigned8:
						return &FieldValueUnsigned8{},nil
					case *FieldValueUnsigned16:
						return &FieldValueUnsigned16{},nil
					case *FieldValueUnsigned32:
						return &FieldValueUnsigned32{},nil
					case *FieldValueUnsigned64:
						return &FieldValueUnsigned64{},nil

					case *FieldValueSigned8:
						return &FieldValueSigned8{},nil
					case *FieldValueSigned16:
						return &FieldValueSigned16{},nil
					case *FieldValueSigned32:
						return &FieldValueSigned32{},nil
					case *FieldValueSigned64:
						return &FieldValueSigned64{},nil

					case *FieldValueFloat32:
						return &FieldValueFloat32{},nil
					case *FieldValueFloat64:
						return &FieldValueFloat64{},nil

					case *FieldValueBoolean:
						return &FieldValueBoolean{},nil

					case *FieldValueMacAddress:
						return &FieldValueMacAddress{},nil

					case *FieldValueOctetArray:
						return &FieldValueOctetArray{},nil

					case *FieldValueString:
						return &FieldValueString{},nil

					case *FieldValueDateTimeSeconds:
						return &FieldValueDateTimeSeconds{},nil
					case *FieldValueDateTimeMilliseconds:
						return &FieldValueDateTimeMilliseconds{},nil
					case *FieldValueDateTimeMicroseconds:
						return &FieldValueDateTimeMicroseconds{},nil
					case *FieldValueDateTimeNanoseconds:
						return &FieldValueDateTimeNanoseconds{},nil

					case *FieldValueIPv4Address:
						return &FieldValueIPv4Address{},nil
					case *FieldValueIPv6Address:
						return &FieldValueIPv6Address{},nil

					case *FieldValueBasicList:
						return &FieldValueBasicList{},nil

					case *FieldValueSubTemplateList:
						return &FieldValueSubTemplateList{},nil
					case *FieldValueSubTemplateMultiList:
						return &FieldValueSubTemplateMultiList{},nil
					}
 		return nil,NewError(fmt.Sprintf("Unknown field value %#v.",val),ErrCritical)					
}