// thresh-wallet
//
// Copyright 2019 by KeyFuse
//
// GPLv3 License

package library

import (
	"testing"

	"server"

	"github.com/stretchr/testify/assert"
)

func TestWalletBalance(t *testing.T) {
	var token string

	ts := server.MockServer()
	defer ts.Close()

	mobile := "10086"
	// Token.
	{
		body := APIGetToken(ts.URL, mobile, "vcode", mockMasterPubKey)
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	body := APIWalletBalance(ts.URL, token)
	rsp := &WalletBalanceResponse{}
	unmarshal(body, rsp)

	t.Logf("%+v", rsp)
	assert.Equal(t, 200, rsp.Code)
	assert.Equal(t, uint64(103266), rsp.AllBalance)
}

func TestAPIEcdsaNewAddress(t *testing.T) {
	var token string

	ts := server.MockServer()
	defer ts.Close()

	mobile := "10086"
	// Token.
	{
		body := APIGetToken(ts.URL, mobile, "vcode", mockMasterPubKey)
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	for i := 0; i < 3; i++ {
		body := APIEcdsaNewAddress(ts.URL, token)
		rsp := &EcdsaAddressResponse{}
		unmarshal(body, rsp)

		t.Logf("%+v", rsp)
		assert.Equal(t, 200, rsp.Code)
	}
}

func TestAPIWalletSend(t *testing.T) {
	var token string

	ts := server.MockServer()
	defer ts.Close()

	mobile := "10086"
	// Token.
	{
		body := APIGetToken(ts.URL, mobile, "vcode", mockMasterPubKey)
		rsp := &TokenResponse{}
		unmarshal(body, rsp)
		assert.Equal(t, 200, rsp.Code)
		token = rsp.Token
	}

	{
		body := APIWalletSend(ts.URL, token, "testnet", mockMasterPrvKey, "mmBRSnFG7o1BX5DaK8Da3xKxvjBh6fzNQq", 100000, 1000)
		rsp := &WalletSendResponse{}
		unmarshal(body, rsp)

		t.Logf("%+v", rsp)
		assert.Equal(t, 200, rsp.Code)

	}

	// Suffient value.
	{
		body := APIWalletSend(ts.URL, token, "testnet", mockMasterPrvKey, "mmBRSnFG7o1BX5DaK8Da3xKxvjBh6fzNQq", 1000000, 1000)
		rsp := &WalletSendResponse{}
		unmarshal(body, rsp)

		t.Logf("%+v", rsp)
		assert.Equal(t, 500, rsp.Code)
	}
}
