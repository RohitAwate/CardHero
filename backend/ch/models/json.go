package models

type JSONTransformer interface {
	JSON() []byte
}
