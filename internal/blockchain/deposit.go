package blockchain

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DepositData struct {
	Pubkey                string `json:"pubkey"`
	WithdrawalCredentials string `json:"withdrawal_credentials"`
	Amount                uint64 `json:"amount"`
	Signature             string `json:"signature"`
}

func ExecuteDepositTransaction(clientURL, privateKeyHex, depositDataFile string) (string, error) {
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		return "", fmt.Errorf("connection failed: %w", err)
	}
	defer client.Close()

	data, err := loadDepositData(depositDataFile)
	if err != nil {
		return "", fmt.Errorf("deposit data error: %w", err)
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	if len(privateKeyHex) != 64 {
		return "", fmt.Errorf("invalid key length: %d chars", len(privateKeyHex))
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %w", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return "", fmt.Errorf("balance check failed: %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("gas price check failed: %w", err)
	}

	gasLimit := uint64(300000)
	gasCost := new(big.Int).Mul(big.NewInt(int64(gasLimit)), gasPrice)
	requiredFunds := new(big.Int).Add(big.NewInt(int64(data.Amount)), gasCost)

	if balance.Cmp(requiredFunds) < 0 {
		return "", fmt.Errorf("insufficient funds: balance %s wei < required %s wei",
			balance.String(), requiredFunds.String())
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("nonce error: %w", err)
	}

	tx := types.NewTransaction(
		nonce,
		common.HexToAddress("0x4242424242424242424242424242424242424242"),
		big.NewInt(int64(data.Amount)),
		gasLimit,
		gasPrice,
		nil,
	)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(17000)), privateKey)
	if err != nil {
		return "", fmt.Errorf("signing failed: %w", err)
	}

	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		return "", fmt.Errorf("transaction failed: %w", err)
	}

	return signedTx.Hash().Hex(), nil
}

func loadDepositData(filename string) (*DepositData, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("file read error: %w", err)
	}

	var depositData []DepositData
	if err := json.Unmarshal(data, &depositData); err != nil {
		return nil, fmt.Errorf("json parse error: %w", err)
	}

	if len(depositData) == 0 {
		return nil, fmt.Errorf("empty deposit data")
	}

	return &depositData[0], nil
}
