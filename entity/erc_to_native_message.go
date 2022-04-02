package entity

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type ErcToNativeMessage struct {
	ID        uint           `db:"id"`
	BridgeID  string         `db:"bridge_id"`
	MsgHash   common.Hash    `db:"msg_hash"`
	Direction Direction      `db:"direction"`
	Receiver  common.Address `db:"receiver"`
	Value     string         `db:"value"`
	CreatedAt *time.Time     `db:"created_at"`
	UpdatedAt *time.Time     `db:"updated_at"`
}

type ErcToNativeMessagesRepo interface {
	Ensure(ctx context.Context, msg *ErcToNativeMessage) error
	FindByMsgHash(ctx context.Context, bridgeID string, msgHash common.Hash) (*ErcToNativeMessage, error)
}