package repository

import (
	"github.com/phungvandat/clean-architecture/util/transaction"
)

// RepoOptions struct
type RepoOptions struct {
	TX *transaction.Pool
}

// NewRepoOptions function return a RepoOptions
func NewRepoOptions() *RepoOptions {
	return new(RepoOptions)
}

// MergeRepoOptions function
func MergeRepoOptions(opts ...*RepoOptions) *RepoOptions {
	ro := NewRepoOptions()
	for _, o := range opts {
		if o.TX != nil {
			ro.TX = o.TX
		}
	}
	return ro
}
