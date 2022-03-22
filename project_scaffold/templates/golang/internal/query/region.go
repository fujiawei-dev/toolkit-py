{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

type RegionLocation struct {
	Region   string    `json:"region" example:"区域名称"`
	Location [2]string `json:"location" example:"经纬度"`
}
