package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gmalheirog/go-course/util"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)

	entryInDb, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entryInDb)
	require.NotZero(t, entryInDb.ID)

	require.Equal(t, entryInDb.Amount, entry.Amount)
	require.Equal(t, entryInDb.AccountID, entry.AccountID)
	require.Equal(t, entryInDb.ID, entry.ID)
}
