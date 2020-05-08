package constants

// TxStatus is transaction status
type TxStatus int

const (
	// PendingTxStatus is pending status
	PendingTxStatus TxStatus = iota
	// CommitedTxStatus is commited status
	CommitedTxStatus
	// RollbackedTxStatus is rollbacked status
	RollbackedTxStatus
)
