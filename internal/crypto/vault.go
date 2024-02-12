package vaultcrypto

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/binary"
	"errors"
	"hash"
	"io"

	"github.com/Chystik/pass-man/internal/vault/usecases"
)

const (
	BUFFER_SIZE int  = 4 * 1024
	IV_SIZE     int  = 16
	V1          byte = 0x1
	hmacSize         = sha512.Size
)

type vaultCryptor struct {
	key usecases.VaultKeyStore
}

func NewVaultCryptor(keyStore usecases.VaultKeyStore) *vaultCryptor {
	return &vaultCryptor{
		key: keyStore,
	}
}

type streamReader struct {
	stream cipher.Stream
	out    io.Writer
}

func newStreamReader(stream cipher.Stream, out io.Writer) *streamReader {
	return &streamReader{
		stream: stream,
		out:    out,
	}
}

func (sr *streamReader) decrypt(in io.Reader, h hash.Hash, iv []byte) (int, error) {
	var err error
	var sum int

	h.Write(iv)
	mac := make([]byte, hmacSize)

	buf := bufio.NewReaderSize(in, BUFFER_SIZE)
	var limit int
	var b []byte

	for {
		b, err = buf.Peek(BUFFER_SIZE)
		if err != nil && err != io.EOF {
			return sum, err
		}

		limit = len(b) - hmacSize

		if err == io.EOF {
			left := buf.Buffered()
			if left < hmacSize {
				return sum, errors.New("not enough left")
			}

			copy(mac, b[left-hmacSize:left])

			if left == hmacSize {
				break
			}
		}

		h.Write(b[:limit])

		// We always leave at least hmacSize bytes left in the buffer
		// That way, our next Peek() might be EOF, but we will still have enough
		outBuf := make([]byte, int64(limit))
		_, err = buf.Read(b[:limit])
		if err != nil {
			return sum, err
		}
		sr.stream.XORKeyStream(outBuf, b[:limit])
		b, err := sr.out.Write(outBuf)
		if err != nil {
			return sum, err
		}
		sum += b

		if err == io.EOF {
			break
		}
	}

	if !hmac.Equal(mac, h.Sum(nil)) {
		return sum, errors.New("invalid hmac")
	}

	return sum, nil
}

type streamWriter struct {
	buf    []byte
	stream cipher.Stream
	out    io.Writer
}

func newStreamWriter(stream cipher.Stream, ow io.Writer) *streamWriter {
	return &streamWriter{
		buf:    make([]byte, BUFFER_SIZE),
		stream: stream,
		out:    ow,
	}
}

func (sw *streamWriter) encrypt(in io.Reader) (int, error) {
	var err error
	var sum int

	for {
		n, err := in.Read(sw.buf)
		if err != nil && err != io.EOF {
			return sum, err
		}

		if n != 0 {
			outBuf := make([]byte, n)
			sw.stream.XORKeyStream(outBuf, sw.buf[:n])
			b, err := sw.out.Write(outBuf)
			if err != nil {
				return sum, err
			}
			sum += b
		}

		if err == io.EOF {
			break
		}
	}

	return sum, err
}

func (v *vaultCryptor) Encrypt(in io.Reader, out io.Writer, userID string) (int, error) {
	var n int

	key, err := v.key.GetKey(userID)
	if err != nil {
		return 0, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	iv := make([]byte, IV_SIZE)
	_, err = rand.Read(iv)
	if err != nil {
		return 0, err
	}

	keyHmac := key // TODO:make hmac key

	h := hmac.New(sha512.New, keyHmac)

	streamMode := cipher.NewCTR(block, iv)

	cw := newStreamWriter(streamMode, out)

	// Version
	_, err = cw.out.Write([]byte{V1})
	if err != nil {
		return 0, err
	}

	cw.out = io.MultiWriter(out, h)

	// Write version
	nv, err := cw.out.Write(iv)
	if err != nil {
		return n + nv, err
	}

	en, err := cw.encrypt(in)
	if err != nil {
		return n + en, err
	}

	nh, err := cw.out.Write(h.Sum(nil))
	if err != nil {
		return n + nh, err
	}

	return n, nil
}

func (v *vaultCryptor) Decrypt(in io.Reader, out io.Writer, userID string) (int, error) {
	var n int

	key, err := v.key.GetKey(userID)
	if err != nil {
		return 0, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	// Read version (up to 0-255)
	var version int8
	err = binary.Read(in, binary.LittleEndian, &version)
	if err != nil {
		return n, err
	}

	iv := make([]byte, IV_SIZE)
	_, err = io.ReadFull(in, iv)
	if err != nil {
		return n, err
	}
	keyHmac := key // TODO:make hmac key

	h := hmac.New(sha512.New, keyHmac)

	streamMode := cipher.NewCTR(block, iv)

	cr := newStreamReader(streamMode, out)

	dn, err := cr.decrypt(in, h, iv)
	if err != nil {
		return n + dn, err
	}

	return n, nil //decrypt(in, out, key, key)
}
