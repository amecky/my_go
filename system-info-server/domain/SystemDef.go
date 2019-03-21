package domain

type SystemDef struct {
	Name string
	Url string
	BasicAuth string
	MediaType string
	SupportMetrics bool
}

type InfoData struct {
    ServiceName string `xml:"serviceName"`
	ServiceHealth string `xml:"serviceHealth"`
	ServiceDescription string `xml:"serviceDescription" json:"healthDescription"`
	WarVersion string `xml:"warVersion"`
	CalledServices map[string]InfoData `json:"calledServices"`
	Url string
	SupportMetrics bool
	Environment string
}

type CustomLogData struct {
	Path string
	Elapsed int
	RequestTimestamp string
	RequestParams map[string]string
	PathParams map[string]string
	HeaderParams map[string]string
}

type Metric struct {
	Path string
	TotalCalls int
	Min float32
	Max float32
	Average float32
	Samples int
}
