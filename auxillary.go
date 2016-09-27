package ipfixmessage

import "fmt"

/*

Auxillary functions

*/

// EncodeVariableLength returns the bytes for encoding a variable length as specified in RFC 7011, section 7
// In the Template Set, the Information Element Field Length is recorded as 65535.
// This reserved length value notifies the Collecting Process that the length value of the Information Element will be carried in the Information Element content itself.
// In most cases, the length of the Information Element will be less than 255 octets. In this case 1 byte is sufficient to encode the length.
// The length may also be encoded into 3 octets before the Information Element, allowing the length of the Information Element to be greater than or equal to 255 octets.
// In this case, the first octet of the Length field MUST be 255, and the length is carried in the second and third octets.
// The octets carrying the length (either the first or the first three octets) MUST NOT be included in the length of the Information Element.
func EncodeVariableLength(content []byte) ([]byte, error) {
	retval := []byte{}
	if len(content) < 255 {
		retval = []byte{uint8(len(content))}
	} else {
		if len(content) > 65535 {
			return []byte{}, fmt.Errorf("Content too large, maximum of 65535 octets, but it is %d", len(content))
		}
		lengthBytes := []byte{255}
		lengthContentBytes, err := marshalBinarySingleValue(uint16(len(content)))
		if err != nil {
			return []byte{}, err
		}
		retval = append(lengthBytes, lengthContentBytes...)
	}
	return retval, nil
}

// DecodeVariableLength returns the length for decoding a variable length as specified in RFC 7011, section 7
func DecodeVariableLength(content []byte) (uint16, error) {
	retval := uint16(0)
	if content[0] == 0 {
		return 0, fmt.Errorf("Content can not be 0 in length.")
	}
	if content[0] < 255 {
		retval = uint16(content[0])
	} else {
		retval = uint16(256*uint16(content[1])) + uint16(content[2])
	}
	return retval, nil
}
