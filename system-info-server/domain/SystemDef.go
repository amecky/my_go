package domain

type SystemDef struct {
	Name string
	Url string
	BasicAuth string
	MediaType string
}

type InfoData struct {
    ServiceName string `xml:"serviceName"`
	ServiceHealth string `xml:"serviceHealth"`
	ServiceDescription string `xml:"serviceDescription" json:"healthDescription"`
	WarVersion string `xml:"warVersion"`
	CalledServices map[string]InfoData `json:"calledServices"`
}