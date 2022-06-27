package caps

import (
	"encoding/json"
	"log"

	v1 "gitlab.com/dyrector_io/dyrector.io/go/pkg/api/v1"
)

type Port struct {
	Listening int64 `json:"listening"`
	Exposed   bool  `json:"exposed"`
}

type NetworkLabel struct {
	Ports []Port `json:"ports"`
}

func ParseLabelsIntoContainerConfig(labels map[string]string, config *v1.ContainerConfig) {
	for key, value := range labels {
		if key != "io.dyrector.cap.network.v1" {
			continue
		} else {
			network := NetworkLabel{}

			err := json.Unmarshal([]byte(value), &network)
			if err != nil {
				log.Println(err.Error())
			}

			ports := []v1.PortBinding{}
			if config.Ports != nil && len(config.Ports) > 0 {
				ports = config.Ports
			}
			for i := range network.Ports {
				ports = append(ports, v1.PortBinding{ExposedPort: uint16(network.Ports[i].Listening), PortBinding: 0})
			}

			config.Ports = ports
		}
	}
}
