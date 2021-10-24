package presenter

import (
	"amb-monitor/entity"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type MessageInfo struct {
	BridgeID  string
	MsgHash   common.Hash
	MessageID common.Hash
	Direction entity.Direction
	Sender    common.Address
	Executor  common.Address
	DataType  uint
}

type InformationRequestInfo struct {
	BridgeID  string
	MessageID common.Hash
	Direction entity.Direction
	Sender    common.Address
	Executor  common.Address
}

type TxInfo struct {
	BlockNumber uint
	Timestamp   time.Time
	Link        string
}

type EventInfo struct {
	Action         string
	LogID          uint            `json:"-"`
	Signer         *common.Address `json:",omitempty"`
	Data           hexutil.Bytes   `json:",omitempty"`
	Count          uint            `json:",omitempty"`
	Status         bool            `json:",omitempty"`
	CallbackStatus bool            `json:",omitempty"`
	*TxInfo
}

type SearchResult struct {
	Event         *EventInfo
	Message       interface{}
	RelatedEvents []*EventInfo
}