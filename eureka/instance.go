package eureka

import (
	"math/rand"
	"time"
)
type Apps struct {
	Applications Applications `json:"applications"`
}

type Instances struct {
	Instance []InstanceInfo `json:"instance"`
}
type Application struct {
	Instance InstanceInfo  `json:"instance"`
}
type InstanceInfo struct {
	InstanceID string `json:"instanceId"`
	HostName   string `json:"hostName"`
	App        string `json:"app"`
	IPAddr     string `json:"ipAddr"`
	VipAddress string `json:"vipAddress"`
	Status     string `json:"status" xml:"status"`
	Port       struct {
		NAMING_FAILED int  `json:"$"`
		Enabled       string `json:"@enabled"`
	} `json:"port"`
	SecurePort struct {
		NAMING_FAILED int  `json:"$"`
		Enabled       string `json:"@enabled"`
	} `json:"securePort"`
	HomePageURL    string `json:"homePageUrl"`
	StatusPageURL  string `json:"statusPageUrl"`
	HealthCheckURL string `json:"healthCheckUrl"`
	DataCenterInfo struct {
		Class string `json:"@class"`
		Name  string `json:"name"`
	} `json:"dataCenterInfo"`
	LeaseInfo struct {
		RenewalIntervalInSecs int `json:"renewalIntervalInSecs"`
		DurationInSecs        int `json:"durationInSecs"`
	} `json:"leaseInfo"`
}

type Applications struct {
	Application []struct {
		InstanceInfo []struct {
			InstanceID string `json:"instanceId"`
			HostName   string `json:"hostName"`
			App        string `json:"app"`
			IPAddr     string `json:"ipAddr"`
			VipAddress string `json:"vipAddress"`
			Status     string `json:"status" xml:"status"`
			Port       struct {
				NAMING_FAILED int  `json:"$"`
				Enabled       string `json:"@enabled"`
			} `json:"port"`
			SecurePort struct {
				NAMING_FAILED int  `json:"$"`
				Enabled       string `json:"@enabled"`
			} `json:"securePort"`
			HomePageURL    string `json:"homePageUrl"`
			StatusPageURL  string `json:"statusPageUrl"`
			HealthCheckURL string `json:"healthCheckUrl"`
			DataCenterInfo struct {
				Class string `json:"@class"`
				Name  string `json:"name"`
			} `json:"dataCenterInfo"`
			LeaseInfo struct {
				RenewalIntervalInSecs int `json:"renewalIntervalInSecs"`
				DurationInSecs        int `json:"durationInSecs"`
			} `json:"leaseInfo"`
		} `json:"instance"`
	} `json:"application"`
}


func (c *Client) NewInstance(appName , hostName, ipAddr string, port int, isSSL bool, homeURL, statusURL, healthCheckURL string) *Application {
	application := &Application{}
	application.Instance.InstanceID = genRandomNumber(13)
	application.Instance.HostName = hostName
	application.Instance.App = appName
	application.Instance.Status = "UP"
	application.Instance.IPAddr = ipAddr
	application.Instance.VipAddress = appName
	application.Instance.Port.NAMING_FAILED = port
	application.Instance.Port.Enabled = "true"
	if isSSL {
		application.Instance.SecurePort.NAMING_FAILED = port
		application.Instance.SecurePort.Enabled = "true"
	}
	application.Instance.HomePageURL = homeURL
	application.Instance.StatusPageURL = statusURL
	application.Instance.HealthCheckURL = healthCheckURL
	application.Instance.DataCenterInfo.Class = "com.netflix.appinfo.MyDataCenterInfo"
	application.Instance.DataCenterInfo.Name = "MyOwn"
	application.Instance.LeaseInfo.RenewalIntervalInSecs = 15
	application.Instance.LeaseInfo.DurationInSecs = 60
	return application
}


func genRandomNumber(length int) string {
	numberset := "0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = numberset[seededRand.Intn(len(numberset))]
	}
	return string(b)
}