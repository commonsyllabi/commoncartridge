// Code generated by zek; DO NOT EDIT.

package types

import "encoding/xml"

// WebLink was generated 2022-05-10 19:56:34 by pierre on archpierre.
type WebLink struct {
	XMLName        xml.Name `xml:"webLink"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Title          string   `xml:"title"`
	URL            struct {
		Text           string `xml:",chardata"`
		Href           string `xml:"href,attr"`
		Target         string `xml:"target,attr"`
		WindowFeatures string `xml:"windowFeatures,attr"`
	} `xml:"url"`
}
