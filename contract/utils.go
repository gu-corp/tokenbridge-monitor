package contract

import (
	"fmt"
	"tokenbridge-monitor/entity"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func Indexed(args abi.Arguments) abi.Arguments {
	var indexed abi.Arguments
	for _, arg := range args {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return indexed
}

func FindMatchingEventABI(contractABI abi.ABI, topics []common.Hash) *abi.Event {
	for _, e := range contractABI.Events {
		if e.ID == topics[0] {
			indexed := Indexed(e.Inputs)
			if len(indexed) == len(topics)-1 {
				return &e
			}
		}
	}
	return nil
}

func DecodeEventLog(event *abi.Event, topics []common.Hash, data []byte) (map[string]interface{}, error) {
	indexed := Indexed(event.Inputs)
	m := make(map[string]interface{})
	if len(indexed) < len(event.Inputs) {
		if err := event.Inputs.UnpackIntoMap(m, data); err != nil {
			return nil, fmt.Errorf("can't unpack data: %w", err)
		}
	}
	if err := abi.ParseTopicsIntoMap(m, indexed, topics[1:]); err != nil {
		return nil, fmt.Errorf("can't unpack topics: %w", err)
	}
	return m, nil
}

func ParseLog(contractABI abi.ABI, log *entity.Log) (string, map[string]interface{}, error) {
	topics := log.Topics()
	if len(topics) == 0 {
		return "", nil, fmt.Errorf("cannot process event without topics")
	}
	event := FindMatchingEventABI(contractABI, topics)
	if event == nil {
		return "", nil, nil
	}

	res, err := DecodeEventLog(event, topics, log.Data)
	if err != nil {
		return "", nil, fmt.Errorf("can't decode event log: %w", err)
	}
	return event.String(), res, nil
}