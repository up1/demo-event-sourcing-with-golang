package event

import (
	"demo/bank"
	"errors"
)

func (e CreateEvent) Process() (*bank.BankAccount, error) {
	return bank.UpdateAccount(e.AccId, map[string]interface{}{
		"Id":      e.AccId,
		"Name":    e.AccName,
		"Balance": "0",
	})
}

func (e DepositEvent) Process() (*bank.BankAccount, error) {
	if acc, err := bank.FetchAccount(e.AccId); err != nil {
		return nil, err
	} else {
		if acc == nil {
			return nil, errors.New("Account not found")
		}
		newBalance := acc.Balance + e.Amount
		return bank.UpdateAccount(e.AccId, map[string]interface{}{
			"Balance": newBalance,
		})
	}
}

func (e WithdrawEvent) Process() (*bank.BankAccount, error) {
	if acc, err := bank.FetchAccount(e.AccId); err != nil {
		return nil, err
	} else {
		if acc == nil {
			return nil, errors.New("Account not found")
		}
		if acc.Balance >= e.Amount {
			newBalance := acc.Balance - e.Amount
			return bank.UpdateAccount(e.AccId, map[string]interface{}{
				"Balance": newBalance,
			})
		} else {
			return nil, errors.New("Insufficient amount")
		}
	}
}

func (e TransferEvent) Process() (*bank.BankAccount, error) {
	if acc, err := bank.FetchAccount(e.AccId); err != nil {
		return nil, err
	} else {
		if acc == nil {
			return nil, errors.New("Account not found")
		}
		if destAcc, err := bank.FetchAccount(e.TargetId); err != nil {
			return nil, err
		} else {
			if destAcc == nil {
				return nil, errors.New("Destination bank account not found")
			}
			if acc.Balance >= e.Amount {
				acc.Balance -= e.Amount
				destAcc.Balance += e.Amount
				if _, err := bank.UpdateAccount(destAcc.Id, map[string]interface{}{
					"Balance": destAcc.Balance,
				}); err != nil {
					return nil, err
				} else {
					return bank.UpdateAccount(acc.Id, map[string]interface{}{
						"Balance": acc.Balance,
					})
				}
			} else {
				return nil, errors.New("Insufficient amount")
			}
		}
	}
}
