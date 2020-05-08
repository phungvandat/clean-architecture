package transaction

import (
	"context"
	"testing"

	mongoConfig "github.com/phungvandat/clean-architecture/util/config/db/mongo"
)

func Test_tx_Begin(t *testing.T) {
	dbTest, cleanup := mongoConfig.CreateTestDatabase(t)
	defer cleanup()

	tx := &tx{
		connDB: dbTest,
	}

	_, err := tx.Begin(context.TODO())
	if err != nil {
		t.Fatalf("Test_tx_Begin failed by error: %v", err)
	}
}

func Test_tx_Commit(t *testing.T) {
	dbTest, cleanup := mongoConfig.CreateTestDatabase(t)
	defer cleanup()

	tx := &tx{
		connDB: dbTest,
	}

	pool, _ := tx.Begin(context.TODO())

	err := tx.Commit(context.TODO(), pool)
	if err != nil {
		t.Fatalf("Test_tx_Commit failed by error: %v", err)
	}
}

func Test_tx_RollBack(t *testing.T) {
	dbTest, cleanup := mongoConfig.CreateTestDatabase(t)
	defer cleanup()

	tx := &tx{
		connDB: dbTest,
	}

	pool, _ := tx.Begin(context.TODO())

	err := tx.RollBack(context.TODO(), pool)
	if err != nil {
		t.Fatalf("Test_tx_RollBack failed by error: %v", err)
	}
}
