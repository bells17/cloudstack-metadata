package metadata

import (
	"io/ioutil"
	"net/http"
	"sync"
)

// type Metadata interface{
// 	Request(string) (string, error)
// }

type metadata struct {
	URL    string
	client *http.Client
}

func NewMetadata(url string) *metadata {
	return &metadata{
		URL:    url,
		client: new(http.Client),
	}
}

func (m *metadata) Request(metadataType string) (string, error) {
	req, err := http.NewRequest("GET", m.URL+"/latest/"+metadataType, nil)
	if err != nil {
		return "", err
	}

	res, err := m.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(byteArray), nil
}

func (m *metadata) ServiceOffering() (string, error) {
	return m.Request("service-offering")
}

func (m *metadata) AvailabilityZone() (string, error) {
	return m.Request("availability-zone")
}

func (m *metadata) LocalIpv4() (string, error) {
	return m.Request("local-ipv4")
}

func (m *metadata) LocalHostname() (string, error) {
	return m.Request("local-hostname")
}

func (m *metadata) PublicIpv4() (string, error) {
	return m.Request("public-ipv4")
}

func (m *metadata) PublicHostname() (string, error) {
	return m.Request("public-hostname")
}

func (m *metadata) InstanceID() (string, error) {
	return m.Request("instance-id")
}

func (m *metadata) UserData() (string, error) {
	return m.Request("user-data")
}

func (m *metadata) FetchAll() (map[string]string, error) {
	return m.FetchData([]string{
		"service-offering",
		"availability-zone",
		"local-ipv4",
		"local-hostname",
		"public-ipv4",
		"public-hostname",
		"instance-id",
		"user-data",
	})
}

func (m *metadata) FetchData(types []string) (map[string]string, error) {
	var result map[string]string
	var err error
	wg := sync.WaitGroup{}

	for _, t := range types {
		wg.Add(1)

		go func(t string) {
			switch t {
			case "service-offering":
				res, e := m.ServiceOffering()
				if e != nil {
					err = e
				}
				result["ServiceOffering"] = res

			case "availability-zone":
				res, e := m.AvailabilityZone()
				if e != nil {
					err = e
				}
				result["AvailabilityZone"] = res

			case "local-ipv4":
				res, e := m.LocalIpv4()
				if e != nil {
					err = e
				}
				result["LocalIpv4"] = res

			case "local-hostname":
				res, e := m.LocalHostname()
				if e != nil {
					err = e
				}
				result["LocalHostname"] = res

			case "public-ipv4":
				res, e := m.PublicIpv4()
				if e != nil {
					err = e
				}
				result["PublicIpv4"] = res

			case "public-hostname":
				res, e := m.PublicHostname()
				if e != nil {
					err = e
				}
				result["PublicHostname"] = res

			case "instance-id":
				res, e := m.InstanceID()
				if e != nil {
					err = e
				}
				result["InstanceID"] = res

			case "user-data":
				res, e := m.UserData()
				if e != nil {
					err = e
				}
				result["UserData"] = res
			}

			wg.Done()
		}(t)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}
