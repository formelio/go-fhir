package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// DigitalMediaType is documented here http://hl7.org/fhir/ValueSet/digital-media-type
type DigitalMediaType int

const (
	DigitalMediaTypePhoto DigitalMediaType = iota
	DigitalMediaTypeVideo
	DigitalMediaTypeAudio
)

func (code DigitalMediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DigitalMediaType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "photo":
		*code = DigitalMediaTypePhoto
	case "video":
		*code = DigitalMediaTypeVideo
	case "audio":
		*code = DigitalMediaTypeAudio
	default:
		return fmt.Errorf("unknown DigitalMediaType code `%s`", s)
	}
	return nil
}
func (code DigitalMediaType) String() string {
	return code.Code()
}
func (code DigitalMediaType) Code() string {
	switch code {
	case DigitalMediaTypePhoto:
		return "photo"
	case DigitalMediaTypeVideo:
		return "video"
	case DigitalMediaTypeAudio:
		return "audio"
	}
	return "<unknown>"
}
func (code DigitalMediaType) Display() string {
	switch code {
	case DigitalMediaTypePhoto:
		return "Photo"
	case DigitalMediaTypeVideo:
		return "Video"
	case DigitalMediaTypeAudio:
		return "Audio"
	}
	return "<unknown>"
}
func (code DigitalMediaType) Definition() string {
	switch code {
	case DigitalMediaTypePhoto:
		return "The media consists of one or more unmoving images, including photographs, computer-generated graphs and charts, and scanned documents"
	case DigitalMediaTypeVideo:
		return "The media consists of a series of frames that capture a moving image"
	case DigitalMediaTypeAudio:
		return "The media consists of a sound recording"
	}
	return "<unknown>"
}
