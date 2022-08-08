package mono

import (
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/types"
)

func GetAccountId(request *types.MonoAccountIdRequest) (*types.MonoAccountIdResponse, error) {
	var resp *types.MonoAccountIdResponse
	v, err := Post("/v1/accounts/auth", request, resp)
	if err != nil {
		return nil, err
	}
	res := v.(*types.MonoAccountIdResponse)
	return res, nil
}

func GetAccountDetails(id string) (*types.MonoAccountResponse, error) {
	var resp *types.MonoAccountResponse
	v, err := Get(fmt.Sprintf("/v1/accounts/%s", id), "", resp)
	if err != nil {
		return nil, err
	}
	res := v.(*types.MonoAccountResponse)
	return res, nil
}

func GetAccountTransactions(id string) (*types.MonoTransactionResponse, error) {
	var resp *types.MonoTransactionResponse
	v, err := Get(fmt.Sprintf("/v1/accounts/%s/transactions", id), fmt.Sprintf("?limit=%d&paginate=%t", 30, false), resp)
	if err != nil {
		return nil, err
	}
	res := v.(*types.MonoTransactionResponse)
	return res, nil
}
