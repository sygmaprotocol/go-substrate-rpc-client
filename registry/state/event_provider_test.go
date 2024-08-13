package state

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sygmaprotocol/go-substrate-rpc-client/v4/rpc/state/mocks"
	"github.com/sygmaprotocol/go-substrate-rpc-client/v4/types"
	"github.com/sygmaprotocol/go-substrate-rpc-client/v4/types/codec"
)

func TestProvider_GetStorageEvents(t *testing.T) {
	stateRPCMock := mocks.NewState(t)

	provider := NewEventProvider(stateRPCMock)

	testHash := types.Hash{}

	var testMeta types.Metadata

	// Empty metadata should cause an error when creating the storage key.
	res, err := provider.GetStorageEvents(&testMeta, testHash)
	assert.ErrorIs(t, err, ErrEventStorageKeyCreation)
	assert.Nil(t, res)

	err = codec.DecodeFromHex(types.MetadataV14Data, &testMeta)
	assert.NoError(t, err)

	storageKey, err := types.CreateStorageKey(&testMeta, storagePrefix, storageMethod, nil)
	assert.NoError(t, err)

	storageData := &types.StorageDataRaw{}

	stateRPCMock.On("GetStorageRaw", storageKey, testHash).
		Return(storageData, nil).
		Once()

	res, err = provider.GetStorageEvents(&testMeta, testHash)
	assert.NoError(t, err)
	assert.Equal(t, storageData, res)

	stateRPCError := errors.New("error")

	stateRPCMock.On("GetStorageRaw", storageKey, testHash).
		Return(nil, stateRPCError).
		Once()

	res, err = provider.GetStorageEvents(&testMeta, testHash)
	assert.ErrorIs(t, err, ErrEventStorageRetrieval)
	assert.Nil(t, res)
}
