package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gmalheirog/go-course/util"
)

func createRandomTransfer(t *testing.T, account1 Account, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.Amount, arg.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t, createRandomAccount(t), createRandomAccount(t))
	transferInDb, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transferInDb)
	require.NotZero(t, transferInDb.ID)

	require.Equal(t, transferInDb.ID, transfer.ID)
	require.Equal(t, transferInDb.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transferInDb.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transferInDb.Amount, transfer.Amount)
}

func TestUpdateTransfer(t *testing.T) {
	transfer := createRandomTransfer(t, createRandomAccount(t), createRandomAccount(t))

	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: util.RandomMoney(),
	}

	updatedTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotZero(t, updatedTransfer)

	require.Equal(t, updatedTransfer.ID, transfer.ID)
	require.Equal(t, updatedTransfer.Amount, arg.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t, createRandomAccount(t), createRandomAccount(t))

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)

	transferInDb, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err, sql.ErrNoRows.Error())

	require.Empty(t, transferInDb)
}
