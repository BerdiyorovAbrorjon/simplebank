package db

import (
	"context"
	"fmt"
)

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfers `json:"transfer"`
	FromAccount Accounts  `json:"from_account"`
	ToAccount   Accounts  `json:"to_account"`
	FromEntry   Entries   `json:"from_entry"`
	ToEntry     Entries   `json:"to_entry"`
}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(context.Background(), func(q *Queries) error {
		var err error

		//create transfer
		result.Transfer, err = q.CreateTransfers(ctx, CreateTransfersParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			fmt.Println("ERROR: create transfer")
			return err
		}

		//create from entry
		result.FromEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			fmt.Println("ERROR: create from entry")
			return err
		}

		//create to entry
		result.ToEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			fmt.Println("ERROR: create from entry")
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		//update from account's balance
		// result.FromAccount, err = q.AddMoneyAccount(ctx, AddMoneyAccountParams{
		// 	Amount: -arg.Amount,
		// 	ID:     arg.FromAccountID,
		// })
		// if err != nil {
		// 	fmt.Println("ERROR: update from  balance")
		// 	return err
		// }

		// //update to account's balance
		// result.ToAccount, err = q.AddMoneyAccount(ctx, AddMoneyAccountParams{
		// 	Amount: arg.Amount,
		// 	ID:     arg.ToAccountID,
		// })
		// if err != nil {
		// 	fmt.Println("ERROR: update to  balance")
		// 	return err
		// }

		return nil
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Accounts, account2 Accounts, err error) {
	account1, err = q.AddMoneyAccount(ctx, AddMoneyAccountParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddMoneyAccount(ctx, AddMoneyAccountParams{
		ID:     accountID2,
		Amount: amount2,
	})
	return
}
