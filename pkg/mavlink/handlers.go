package mavlink

import (
	"github.com/bluenviron/gomavlib/v2"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/common"
	"github.com/fairytale5571/mav-cli/pkg/simulator"
)

type Handler struct {
	node       *gomavlib.Node
	simulators simulator.SimulatorsInterface
	done       chan bool
}

func NewHandler(node *gomavlib.Node, simulators simulator.SimulatorsInterface) *Handler {
	return &Handler{
		node:       node,
		simulators: simulators,
		done:       make(chan bool),
	}
}

func (h *Handler) HandleMessages(id int) {
	for evt := range h.node.Events() {
		select {
		case <-h.done:
			return
		default:
			if frm, ok := evt.(*gomavlib.EventFrame); ok {
				sim, exists := h.simulators.Get(id)
				if !exists {
					h.simulators.Add(id)
				}
				h.handleFrame(frm, sim)
			}
		}
	}
}

func (h *Handler) handleFrame(frm *gomavlib.EventFrame, sim *simulator.Simulator) {
	switch msg := frm.Message().(type) {
	case *common.MessageGpsRawInt:
		sim.Messages = sim.Messages.Next()
		sim.Messages.Value = msg
	}
}

func (h *Handler) Stop() {
	h.done <- true
}

func (h *Handler) DoneCh() <-chan bool {
	return h.done
}
