package postgres

import (
	"amb-monitor/db"
	"amb-monitor/entity"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type signedMessagesRepo struct {
	table string
	db    *db.DB
}

func NewSignedMessagesRepo(table string, db *db.DB) entity.SignedMessagesRepo {
	return &signedMessagesRepo{
		table: table,
		db:    db,
	}
}

func (r *signedMessagesRepo) Ensure(ctx context.Context, msg *entity.SignedMessage) error {
	q, args, err := sq.Insert(r.table).
		Columns("log_id", "bridge_id", "msg_hash", "signer").
		Values(msg.LogID, msg.BridgeID, msg.MsgHash, msg.Signer).
		Suffix("ON CONFLICT (log_id) DO UPDATE SET updated_at = NOW()").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't build query: %w", err)
	}
	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("can't insert signed message: %w", err)
	}
	return nil
}