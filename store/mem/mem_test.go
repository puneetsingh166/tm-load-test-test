package mem_test

import (
	"testing"

	"github.com/puneetsingh166/tm-load-test/store/cachekv"

	"github.com/stretchr/testify/require"

	"github.com/puneetsingh166/tm-load-test/store/mem"
	"github.com/puneetsingh166/tm-load-test/store/types"
)

func TestStore(t *testing.T) {
	db := mem.NewStore()
	key, value := []byte("key"), []byte("value")

	require.Equal(t, types.StoreTypeMemory, db.GetStoreType())

	require.Nil(t, db.Get(key))
	db.Set(key, value)
	require.Equal(t, value, db.Get(key))

	newValue := []byte("newValue")
	db.Set(key, newValue)
	require.Equal(t, newValue, db.Get(key))

	db.Delete(key)
	require.Nil(t, db.Get(key))

	cacheWrapper := db.CacheWrap()
	require.IsType(t, &cachekv.Store{}, cacheWrapper)

	cacheWrappedWithTrace := db.CacheWrapWithTrace(nil, nil)
	require.IsType(t, &cachekv.Store{}, cacheWrappedWithTrace)

	cacheWrappedWithListeners := db.CacheWrapWithListeners(nil, nil)
	require.IsType(t, &cachekv.Store{}, cacheWrappedWithListeners)
}

func TestCommit(t *testing.T) {
	db := mem.NewStore()
	key, value := []byte("key"), []byte("value")

	db.Set(key, value)
	id := db.Commit()
	require.True(t, id.IsZero())
	require.True(t, db.LastCommitID().IsZero())
	require.Equal(t, value, db.Get(key))
}
