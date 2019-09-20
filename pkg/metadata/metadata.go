package metadata

import (
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	// ServiceOffering is a metadata name of service offering
	ServiceOffering = "service-offering"
	// AvailabilityZone is a metadata name of availability zone
	AvailabilityZone = "availability-zone"
	// LocalIpv4 is a metadata name of local ipv4
	LocalIpv4 = "local-ipv4"
	// LocalHostname is a metadata name of local hostname
	LocalHostname = "local-hostname"
	// PublicIpv4 is a metadata name of public ipv4
	PublicIpv4 = "public-ipv4"
	// PublicHostname is a metadata name of public hostname
	PublicHostname = "public-hostname"
	// InstanceID is a metadata name of instance id
	InstanceID = "instance-id"
	// UserData is a metadata name of userdata
	UserData = "user-data"
)

// ResponseGroup is a metadata data group for response
type ResponseGroup struct {
	ServiceOffering  string
	AvailabilityZone string
	LocalIpv4        string
	LocalHostname    string
	PublicIpv4       string
	PublicHostname   string
	InstanceID       string
	UserData         string
}

// Metadata is a interface of metadata struct
type Metadata interface {
	Request(string) (string, error)
	ServiceOffering() (string, error)
	AvailabilityZone() (string, error)
	LocalIpv4() (string, error)
	LocalHostname() (string, error)
	PublicIpv4() (string, error)
	PublicHostname() (string, error)
	InstanceID() (string, error)
	UserData() (string, error)
	FetchAll() (map[string]string, error)
	FetchData([]string) (*ResponseGroup, error)
}

type metadata struct {
	Domain string
	client *http.Client
}


// NewMetadata return *metadata
func NewMetadata(domain string) *metadata {
	return &metadata{
		Domain: domain,
		client: new(http.Client),
	}
}

func (m *metadata) request(metadataType string) (string, error) {
	req, err := http.NewRequest("GET", "http://"+m.Domain+"/latest/"+metadataType, nil)
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
	return m.request(ServiceOffering)
}

func (m *metadata) AvailabilityZone() (string, error) {
	return m.request(AvailabilityZone)
}

func (m *metadata) LocalIpv4() (string, error) {
	return m.request(LocalIpv4)
}

func (m *metadata) LocalHostname() (string, error) {
	return m.request(LocalHostname)
}

func (m *metadata) PublicIpv4() (string, error) {
	return m.request(PublicIpv4)
}

func (m *metadata) PublicHostname() (string, error) {
	return m.request(PublicHostname)
}

func (m *metadata) InstanceID() (string, error) {
	return m.request(InstanceID)
}

func (m *metadata) UserData() (string, error) {
	return m.request(UserData)
}

func (m *metadata) FetchAll() (*ResponseGroup, error) {
	return m.FetchData([]string{
		ServiceOffering,
		AvailabilityZone,
		LocalIpv4,
		LocalHostname,
		PublicIpv4,
		PublicHostname,
		InstanceID,
		UserData,
	})
}

func (m *metadata) FetchData(types []string) (*ResponseGroup, error) {
	group := &ResponseGroup{}
	var err error
	wg := sync.WaitGroup{}

	for _, t := range types {
		wg.Add(1)

		go func(t string) {
			switch t {
			case ServiceOffering:
				res, e := m.ServiceOffering()
				if e != nil {
					err = e
				}
				group.ServiceOffering = res

			case AvailabilityZone:
				res, e := m.AvailabilityZone()
				if e != nil {
					err = e
				}
				group.AvailabilityZone = res

			case LocalIpv4:
				res, e := m.LocalIpv4()
				if e != nil {
					err = e
				}
				group.LocalIpv4 = res

			case LocalHostname:
				res, e := m.LocalHostname()
				if e != nil {
					err = e
				}
				group.LocalHostname = res

			case PublicIpv4:
				res, e := m.PublicIpv4()
				if e != nil {
					err = e
				}
				group.PublicIpv4 = res

			case PublicHostname:
				res, e := m.PublicHostname()
				if e != nil {
					err = e
				}
				group.PublicHostname = res

			case InstanceID:
				res, e := m.InstanceID()
				if e != nil {
					err = e
				}
				group.InstanceID = res

			case UserData:
				res, e := m.UserData()
				if e != nil {
					err = e
				}
				group.UserData = res
			}

			wg.Done()
		}(t)
	}

	wg.Wait()

	if err != nil {
		return nil, err
	}

	return group, nil
}
