package parser

type Items struct {
	Items []Item `xml:"item"`
}

type Item struct {
	URL      string   `xml:"url"`
	Host     Host     `xml:"host"`
	Port     int      `xml:"port"`
	Method   string   `xml:"method"`
	Path     string   `xml:"path"`
	Request  Request  `xml:"request"`
	Response Response `xml:"response"`
}

type Host struct {
	IP   string `xml:"ip,attr"`
	Name string `xml:",chardata"`
}

type Request struct {
	Base64 bool   `xml:"base64,attr"`
	Data   string `xml:",chardata"`
}

type Response struct {
	Base64 bool   `xml:"base64,attr"`
	Data   string `xml:",chardata"`
}
