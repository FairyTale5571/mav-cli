package simulator

import (
	"container/ring"

	"github.com/bluenviron/gomavlib/v2/pkg/dialects/ardupilotmega"
)

type Simulator struct {
	id        int
	Messages  *ring.Ring
	messageCh chan *ardupilotmega.MessageGpsRawInt
}

type SimulatorsInterface interface {
	Add(id int)
	Get(id int) (*Simulator, bool)
	GetAll() map[int]*Simulator
}

type Simulators struct {
	mapping map[int]*Simulator
}

func NewSimulators() SimulatorsInterface {
	return &Simulators{
		mapping: make(map[int]*Simulator),
	}
}

func (s *Simulators) Add(id int) {
	s.mapping[id] = &Simulator{
		id:        id,
		Messages:  ring.New(100),
		messageCh: make(chan *ardupilotmega.MessageGpsRawInt, 1),
	}
}

func (s *Simulators) Get(id int) (*Simulator, bool) {
	sim, exists := s.mapping[id]
	return sim, exists
}

func (s *Simulators) GetAll() map[int]*Simulator {
	return s.mapping
}
