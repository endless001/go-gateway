package conf

import (
	"fmt"
	"testing"
)

type SiteConfig struct {
	HttpPort  int
	HttpsOn   bool
	Domain    string
	HttpsPort int
}

type NginxConfig struct {
	Port    int
	LogPath string
	Path    string
}

func TestConfigYaml(t *testing.T) {
	c2 := ConfigEngine{}
	c2.Load("test.yaml")

	siteConf := SiteConfig{}
	c2.GetStruct("Site", &siteConf)
	fmt.Println(siteConf)

	nginxConfig := NginxConfig{}
	c2.GetStruct("Nginx", &nginxConfig)
	fmt.Println(nginxConfig)

	siteName := c2.GetString("SiteName")
	siteAddr := c2.GetString("SiteAddr")
	fmt.Println(siteName, siteAddr)
}
