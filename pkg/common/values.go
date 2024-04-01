package common

import (
	"errors"
	"os"
)

const (
	// DefaultConfigAccountAliasesDirName from accounts
	DefaultConfigAccountAliasesDirName = "keys"
	// DefaultPassphrase for accounts
	DefaultPassphrase = ""
	// Secp256k1PrivateKeyBytesLength privete key
	Secp256k1PrivateKeyBytesLength = 32
	// AmountDecimalPoint TRX decimal point
	AmountDecimalPoint = 6
)

var (
	// DefaultConfigDirName for wallets
	DefaultConfigDirName = "storage/crypto_coins/.tron"
	DebugGRPC            = false
	DebugTransaction     = false
	ErrNotAbsPath        = errors.New("keypath is not absolute path")
	ErrBadKeyLength      = errors.New("invalid private key (wrong length)")
	ErrFoundNoPass       = errors.New("found no passphrase file")
)

func init() {
	if _, enabled := os.LookupEnv("TRONCTL_GRPC_DEBUG"); enabled != false {
		DebugGRPC = true
	}
	if _, enabled := os.LookupEnv("TRONCTL_TX_DEBUG"); enabled != false {
		DebugTransaction = true
	}
	if _, enabled := os.LookupEnv("TRONCTL_ALL_DEBUG"); enabled != false {
		EnableAllVerbose()
	}
}

// EnableAllVerbose sets debug vars to true
func EnableAllVerbose() {
	DebugGRPC = true
	DebugTransaction = true
}
