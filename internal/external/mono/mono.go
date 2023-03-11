package mono

import (
	"errors"
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/types"
)

func GetAccountId(request *types.MonoAccountIdRequest) (*types.MonoAccountIdResponse, error) {
	var resp *types.MonoAccountIdResponse
	v, err := Post("/v1/account/auth", request, resp)
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

func SyncUserData(id string) (bool, error) {
	var resp *types.MonoManualsyncResponse
	s, err := Get(fmt.Sprintf("/v1/accounts/%s/sync", id), "", resp)
	if err != nil {
		return false, err
	}
	res := s.(*types.MonoManualsyncResponse)
	if res.Status == "failed" {
		if res.Code == "REAUTHORISATION_REQUIRED" {
			return false, errors.New("REAUTHORISATION_REQUIRED")
		} else if res.Code == "SYNC_ERROR" {
			return false, errors.New("SYNC_ERROR")
		}
	}
	if res.Code == "INCOMPLETE_STATEMENT_ERROR" {
		s, err = Get(fmt.Sprintf("/v1/accounts/%s/sync", id), "?allow_incomplete_statement=true", resp)
		//if err != nil{
		//	return false, err
		//}
	}
	return true, nil
}

func ReauthoriseUser(id string) (*types.MonoReauthResponse, error) {
	var resp *types.MonoReauthResponse
	s, err := Get(fmt.Sprintf("/v1/accounts/%s/reauthorise", id), "", resp)
	if err != nil {
		return nil, err
	}
	res := s.(*types.MonoReauthResponse)
	return res, nil
}

func Unlink(id string) error {
	var resp string
	_, err := Get(fmt.Sprintf("/v1/accounts/%s/unlink", id), "", resp)
	if err != nil {
		return err
	}
	return nil
}
