package suisigner

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/pattonkan/sui-go/sui"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/blake2b"
)

const (
	SignatureFlagEd25519   = 0x0
	SignatureFlagSecp256k1 = 0x1

	// IOTA_DIFF 4218 is for iota
	DerivationPathEd25519   = `m/44'/784'/0'/0'/0'`
	DerivationPathSecp256k1 = `m/54'/784'/0'/0/0`
)

var (
	TEST_MNEMONIC = "ordinary cry margin host traffic bulb start zone mimic wage fossil eight diagram clay say remove add atom"
	TEST_SEED     = []byte{4, 66, 186, 181, 112, 134, 111, 192, 149, 13, 68, 115, 67, 195, 58, 59, 33, 20, 200, 10, 150, 185, 145, 3, 106, 160, 105, 37, 4, 153, 172, 103, 69, 228, 114, 210, 176, 182, 208, 21, 252, 59, 50, 82, 135, 160, 1, 131, 156, 104, 159, 240, 183, 20, 250, 216, 26, 228, 91, 220, 15, 222, 75, 91}
	TEST_ADDRESS  = sui.MustAddressFromHex("0x1a02d61c6434b4d0ff252a880c04050b5f27c8b574026c98dd72268865c0ede5")
)

// FIXME support more than ed25519
type Signer struct {
	Ed25519Keypair *KeypairEd25519
	// secp256k1Keypair *KeypairSecp256k1
	Address *sui.Address
}

func NewSigner(seed []byte, flag KeySchemeFlag) *Signer {
	prikey := ed25519.NewKeyFromSeed(seed[:ed25519.SeedSize])
	pubkey := prikey.Public().(ed25519.PublicKey)

	// IOTA_DIFF iota ignore flag when signature scheme is ed25519
	var buf []byte
	switch flag {
	case KeySchemeFlagEd25519:
		buf = []byte{KeySchemeFlagEd25519.Byte()}
	case KeySchemeFlagSecp256k1:
		buf = []byte{KeySchemeFlagEd25519.Byte()}
	case KeySchemeFlagIotaEd25519:
		buf = []byte{}
	default:
		panic("unrecognizable key scheme flag")
	}
	buf = append(buf, pubkey...)
	addrBytes := blake2b.Sum256(buf)
	addr := "0x" + hex.EncodeToString(addrBytes[:])

	return &Signer{
		Ed25519Keypair: &KeypairEd25519{
			PriKey: prikey,
			PubKey: pubkey,
		},
		Address: sui.MustAddressFromHex(addr),
	}
}

// there are only 256 different signers can be generated
func NewSignerByIndex(seed []byte, flag KeySchemeFlag, index int) *Signer {
	seed[0] = seed[0] + byte(index)
	return NewSigner(seed, flag)
}

// generate keypair (signer) with mnemonic which is referring the Sui monorepo in the following code
//
// let phrase = "asset pink record dawn hundred sure various crime client enforce carbon blossom";
// let mut keystore = Keystore::from(InMemKeystore::new_insecure_for_tests(0));
// let generated_address = keystore.import_from_mnemonic(&phrase, SignatureScheme::ED25519, None, None).unwrap();
func NewSignerWithMnemonic(mnemonic string, flag KeySchemeFlag) (*Signer, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}
	key, err := DeriveForPath(DerivationPathEd25519, seed)
	if err != nil {
		return nil, err
	}
	return NewSigner(key.Key, flag), nil
}

func (s *Signer) PrivateKey() []byte {
	switch {
	case s.Ed25519Keypair != nil:
		return s.Ed25519Keypair.PriKey
	default:
		return nil
	}
}

func (s *Signer) PublicKey() []byte {
	switch {
	case s.Ed25519Keypair != nil:
		return s.Ed25519Keypair.PubKey
	default:
		return nil
	}
}

func (s *Signer) Sign(data []byte) Signature {
	// FIXME support more than ed25519
	return Signature{
		Ed25519SuiSignature: NewEd25519SuiSignature(s, data),
	}
}

// FIXME support more than ed25519
func (a *Signer) SignTransactionBlock(txnBytes []byte, intent Intent) (Signature, error) {
	data := MessageWithIntent(intent, bcsBytes(txnBytes))
	hash := blake2b.Sum256(data)
	return a.Sign(hash[:]), nil
}

type bcsBytes []byte

func (b bcsBytes) MarshalBCS() ([]byte, error) {
	return b, nil
}
