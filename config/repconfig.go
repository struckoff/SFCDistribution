package config

type GeoJsonReportConfig struct {
	NodesNumber  int
	NodesCap     float64
	NoiseRate    float64
	OutputFmt    string
	PointsNumber int
	Clustermode  string
	GroupBy      string
	TestMode     string
	CurveName    string
	CurveBits    int64
	CurveDims    int64
	ConsMode     string
	ConsLoad     float64
}
