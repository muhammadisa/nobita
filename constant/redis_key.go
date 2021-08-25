package constant

const (
	AutoDept    ProcType = `auto_debt`
	Transaction ProcType = `transaction`
)

const (
	// RedisHaltKey key format lock_for:<type>:<account_id>
	RedisHaltKey = `halt:%s:%d`
)
