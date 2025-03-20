package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// EntropyxMainnetPrivate is the version that is used for
// entropyx mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var EntropyxMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// EntropyxMainnetPublic is the version that is used for
// entropyx mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var EntropyxMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// EntropyxTestnetPrivate is the version that is used for
// entropyx testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var EntropyxTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// EntropyxTestnetPublic is the version that is used for
// entropyx testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var EntropyxTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// EntropyxDevnetPrivate is the version that is used for
// entropyx devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var EntropyxDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// EntropyxDevnetPublic is the version that is used for
// entropyx devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var EntropyxDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// EntropyxSimnetPrivate is the version that is used for
// entropyx simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var EntropyxSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// EntropyxSimnetPublic is the version that is used for
// entropyx simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var EntropyxSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case EntropyxMainnetPrivate:
		return EntropyxMainnetPublic, nil
	case EntropyxTestnetPrivate:
		return EntropyxTestnetPublic, nil
	case EntropyxDevnetPrivate:
		return EntropyxDevnetPublic, nil
	case EntropyxSimnetPrivate:
		return EntropyxSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case EntropyxMainnetPrivate:
		return true
	case EntropyxTestnetPrivate:
		return true
	case EntropyxDevnetPrivate:
		return true
	case EntropyxSimnetPrivate:
		return true
	}

	return false
}
