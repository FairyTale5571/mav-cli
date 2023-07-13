package mavlink

import (
	"log"

	"github.com/bluenviron/gomavlib/v2"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/ardupilotmega"
)

func InitNode(address string, id int) *gomavlib.Node {
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointTCPClient{Address: address},
		},
		Dialect:     ardupilotmega.Dialect,
		OutVersion:  gomavlib.V2,
		OutSystemID: byte(id),
	})
	if err != nil {
		log.Fatalf("Error creating node: %v", err)
	}
	return node
}
