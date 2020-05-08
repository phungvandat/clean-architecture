package transaction

import (
	"context"

	"github.com/phungvandat/clean-architecture/util/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

// Txer interface
type Txer interface {
	Begin(ctx context.Context) (*Pool, error)
	Commit(ctx context.Context, p *Pool) error
	RollBack(ctx context.Context, p *Pool) error
}

// Pool contain session, context, status of a transaction
type Pool struct {
	session mongo.Session
	SCtx    mongo.SessionContext
	status  constants.TxStatus
}

type tx struct {
	connDB *mongo.Database
}

// NewTxer function return transactioner
func NewTxer(connDB *mongo.Database) Txer {
	return &tx{
		connDB: connDB,
	}
}

func (t *tx) Begin(ctx context.Context) (*Pool, error) {
	s, err := t.connDB.Client().StartSession()

	if err != nil {
		return nil, err
	}
	err = s.StartTransaction()
	if err != nil {
		return nil, err
	}
	return &Pool{
		SCtx:    mongo.NewSessionContext(ctx, s),
		status:  constants.PendingTxStatus,
		session: s,
	}, nil
}

func (t *tx) Commit(ctx context.Context, p *Pool) error {
	defer p.session.EndSession(ctx)
	if p == nil ||
		p.status == constants.CommitedTxStatus ||
		p.status == constants.RollbackedTxStatus {
		return nil
	}

	if p.session != nil && p.SCtx != nil {
		err := p.session.CommitTransaction(p.SCtx)
		if err != nil {
			return err
		}
	}
	p.status = constants.CommitedTxStatus
	return nil
}

func (t *tx) RollBack(ctx context.Context, p *Pool) error {
	defer p.session.EndSession(ctx)
	if p == nil ||
		p.status == constants.CommitedTxStatus ||
		p.status == constants.RollbackedTxStatus {
		return nil
	}

	if p.session != nil && p.SCtx != nil {
		err := p.session.AbortTransaction(p.SCtx)
		if err != nil {
			return err
		}
	}
	p.status = constants.RollbackedTxStatus
	return nil
}
