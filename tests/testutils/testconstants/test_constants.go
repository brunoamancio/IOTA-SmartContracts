package testconstants

const (
	// ContractName is defined in smartcontract/rust/Cargo.toml > package > name
	ContractName = "my_iota_sc"
	// Debug is used by Solo. 'true' for logging level 'debug', otherwise 'info'
	Debug = false
	// StackTrace is used by Solo. 'true' if stack trace must be printed in case of errors
	StackTrace = false

	// AccountsContractName sets the name of the Accounts contract, which is a root contract present in every chain
	AccountsContractName = "accounts"
)
