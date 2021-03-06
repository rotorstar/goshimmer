package protocol

import (
	"github.com/iotaledger/goshimmer/plugins/autopeering/instances/chosenneighbors"
	"github.com/iotaledger/goshimmer/plugins/autopeering/instances/knownpeers"
	"github.com/iotaledger/goshimmer/plugins/autopeering/types/response"
	"github.com/iotaledger/hive.go/events"
	"github.com/iotaledger/hive.go/node"
)

func createIncomingResponseProcessor(plugin *node.Plugin) *events.Closure {
	return events.NewClosure(func(peeringResponse *response.Response) {
		go processIncomingResponse(plugin, peeringResponse)
	})
}

func processIncomingResponse(plugin *node.Plugin, peeringResponse *response.Response) {
	log.Debugf("received peering response from %s", peeringResponse.Issuer.String())

	if conn := peeringResponse.Issuer.GetConn(); conn != nil {
		_ = conn.Close()
	}

	knownpeers.INSTANCE.AddOrUpdate(peeringResponse.Issuer)
	for _, peer := range peeringResponse.Peers {
		knownpeers.INSTANCE.AddOrUpdate(peer)
	}

	if peeringResponse.Type == response.TYPE_ACCEPT {
		defer chosenneighbors.INSTANCE.Lock()()

		chosenneighbors.INSTANCE.AddOrUpdate(peeringResponse.Issuer)
	}
}
