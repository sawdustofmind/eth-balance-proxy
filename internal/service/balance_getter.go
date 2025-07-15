package service

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"

	"github.com/sawdustofmind/eth-balance-proxy/internal/log"
	"github.com/sawdustofmind/eth-balance-proxy/internal/monitoring"
)

type BalanceGetterConfig struct {
	DataSources []string `mapstructure:"data_sources"`
}

type BalanceGetter struct {
	clients []balanceGetterClient
}

type balanceGetterClient struct {
	url    string
	client *ethclient.Client
}

func NewBalanceGetter(cfg *BalanceGetterConfig) (*BalanceGetter, error) {
	clients := make([]balanceGetterClient, 0, len(cfg.DataSources))

	for _, ds := range cfg.DataSources {
		client, err := ethclient.Dial(ds)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
		}
		clients = append(clients, balanceGetterClient{
			url:    ds,
			client: client,
		})
	}

	return &BalanceGetter{
		clients: clients,
	}, nil
}

func (bg *BalanceGetter) GetBalance(ctx context.Context, address string) (string, error) {
	for _, client := range bg.clients {
		balance, err := bg.getBalance(ctx, client, address)
		if err == nil {
			return balance, nil
		}
		log.Error("getting balance from client", zap.String("address", address), zap.Error(err))
	}
	return "", fmt.Errorf("could not get balance for %s", address)
}

func (bg *BalanceGetter) getBalance(ctx context.Context, client balanceGetterClient, address string) (string, error) {
	account := common.HexToAddress(address)
	balance, err := client.client.BalanceAt(ctx, account, nil)
	if err != nil {
		monitoring.ProviderRequestFailures.WithLabelValues(client.url).Inc()
		return "", fmt.Errorf("failed to get balance: %w", err)
	}

	return balance.String(), nil
}
