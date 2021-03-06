package z

import "github.com/go-serv/service/pkg/z/ancillary/crypto"

type (
	OnGCFn       func()
	SessionState int
)

type SessionInterface interface {
	Id() SessionId
	State() SessionState
	Nonce() []byte
	WithNonce([]byte)
	BlockCipher() crypto.AEAD_CipherInterface
	WithBlockCipher(crypto.AEAD_CipherInterface)
	Context() interface{}
	WithContext(interface{})
}
