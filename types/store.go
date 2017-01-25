package types

import (
	"github.com/noypi/kv"
)

type VertexDeserializer func([]byte) (Vertex, error)

type StoreConstructor func(kv.KVStore, VertexDeserializer) (Store, error)

type Store interface {
	NewAdjacency() Adjacency
	Store() kv.KVStore
}
