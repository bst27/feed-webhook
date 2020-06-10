package atom

import "encoding/xml"

type feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []entry  `xml:"entry"`
}

type entry struct {
	ID      string `xml:"id"`
	Title   string `xml:"title"`
	Link    link   `xml:"link"`
	Summary string `xml:"summary"`
	Updated string `xml:"updated"`
}

type link struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}
