package constants

// ContextKey Use for context
type ContextKey int

//Define key send through context
const (
	UserIDContextKey ContextKey = iota
	UsernameContextKey
	UserRoleContextKey
)
