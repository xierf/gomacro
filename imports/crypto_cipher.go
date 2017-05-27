// this file was generated by gomacro command: import _b "crypto/cipher"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/cipher"
)

// reflection: allow interpreted code to import "crypto/cipher"
func init() {
	Packages["crypto/cipher"] = Package{
	Binds: map[string]Value{
		"NewCBCDecrypter":	ValueOf(cipher.NewCBCDecrypter),
		"NewCBCEncrypter":	ValueOf(cipher.NewCBCEncrypter),
		"NewCFBDecrypter":	ValueOf(cipher.NewCFBDecrypter),
		"NewCFBEncrypter":	ValueOf(cipher.NewCFBEncrypter),
		"NewCTR":	ValueOf(cipher.NewCTR),
		"NewGCM":	ValueOf(cipher.NewGCM),
		"NewGCMWithNonceSize":	ValueOf(cipher.NewGCMWithNonceSize),
		"NewOFB":	ValueOf(cipher.NewOFB),
	},
	Types: map[string]Type{
		"AEAD":	TypeOf((*cipher.AEAD)(nil)).Elem(),
		"Block":	TypeOf((*cipher.Block)(nil)).Elem(),
		"BlockMode":	TypeOf((*cipher.BlockMode)(nil)).Elem(),
		"Stream":	TypeOf((*cipher.Stream)(nil)).Elem(),
		"StreamReader":	TypeOf((*cipher.StreamReader)(nil)).Elem(),
		"StreamWriter":	TypeOf((*cipher.StreamWriter)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"AEAD":	TypeOf((*AEAD_crypto_cipher)(nil)).Elem(),
		"Block":	TypeOf((*Block_crypto_cipher)(nil)).Elem(),
		"BlockMode":	TypeOf((*BlockMode_crypto_cipher)(nil)).Elem(),
		"Stream":	TypeOf((*Stream_crypto_cipher)(nil)).Elem(),
	},
	Wrappers: map[string][]string{
	} }
}

// --------------- proxy for crypto/cipher.AEAD ---------------
type AEAD_crypto_cipher struct {
	Object	interface{}
	NonceSize_	func() int
	Open_	func(dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error)
	Overhead_	func() int
	Seal_	func(dst []byte, nonce []byte, plaintext []byte, additionalData []byte) []byte
}
func (Proxy *AEAD_crypto_cipher) NonceSize() int {
	return Proxy.NonceSize_()
}
func (Proxy *AEAD_crypto_cipher) Open(dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error) {
	return Proxy.Open_(dst, nonce, ciphertext, additionalData)
}
func (Proxy *AEAD_crypto_cipher) Overhead() int {
	return Proxy.Overhead_()
}
func (Proxy *AEAD_crypto_cipher) Seal(dst []byte, nonce []byte, plaintext []byte, additionalData []byte) []byte {
	return Proxy.Seal_(dst, nonce, plaintext, additionalData)
}

// --------------- proxy for crypto/cipher.Block ---------------
type Block_crypto_cipher struct {
	Object	interface{}
	BlockSize_	func() int
	Decrypt_	func(dst []byte, src []byte) 
	Encrypt_	func(dst []byte, src []byte) 
}
func (Proxy *Block_crypto_cipher) BlockSize() int {
	return Proxy.BlockSize_()
}
func (Proxy *Block_crypto_cipher) Decrypt(dst []byte, src []byte)  {
	Proxy.Decrypt_(dst, src)
}
func (Proxy *Block_crypto_cipher) Encrypt(dst []byte, src []byte)  {
	Proxy.Encrypt_(dst, src)
}

// --------------- proxy for crypto/cipher.BlockMode ---------------
type BlockMode_crypto_cipher struct {
	Object	interface{}
	BlockSize_	func() int
	CryptBlocks_	func(dst []byte, src []byte) 
}
func (Proxy *BlockMode_crypto_cipher) BlockSize() int {
	return Proxy.BlockSize_()
}
func (Proxy *BlockMode_crypto_cipher) CryptBlocks(dst []byte, src []byte)  {
	Proxy.CryptBlocks_(dst, src)
}

// --------------- proxy for crypto/cipher.Stream ---------------
type Stream_crypto_cipher struct {
	Object	interface{}
	XORKeyStream_	func(dst []byte, src []byte) 
}
func (Proxy *Stream_crypto_cipher) XORKeyStream(dst []byte, src []byte)  {
	Proxy.XORKeyStream_(dst, src)
}
