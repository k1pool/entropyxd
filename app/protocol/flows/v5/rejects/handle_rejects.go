package rejects

import (
	"github.com/k1pool/entropyxd/app/appmessage"
	"github.com/k1pool/entropyxd/app/protocol/protocolerrors"
	"github.com/k1pool/entropyxd/infrastructure/network/netadapter/router"
)

// HandleRejectsContext is the interface for the context needed for the HandleRejects flow.
type HandleRejectsContext interface {
}

type handleRejectsFlow struct {
	HandleRejectsContext
	incomingRoute, outgoingRoute *router.Route
}

// HandleRejects handles all reject messages coming through incomingRoute.
// This function assumes that incomingRoute will only return MsgReject.
func HandleRejects(context HandleRejectsContext, incomingRoute *router.Route, outgoingRoute *router.Route) error {
	flow := &handleRejectsFlow{
		HandleRejectsContext: context,
		incomingRoute:        incomingRoute,
		outgoingRoute:        outgoingRoute,
	}
	return flow.start()
}

func (flow *handleRejectsFlow) start() error {
	message, err := flow.incomingRoute.Dequeue()
	if err != nil {
		return err
	}
	rejectMessage := message.(*appmessage.MsgReject)

	return protocolerrors.Errorf(false, "got reject message: `%s`", rejectMessage.Reason)
}
