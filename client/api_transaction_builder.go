package client

import (
	"context"

	"github.com/howjmay/go-sui-sdk/sui_types"
	"github.com/howjmay/go-sui-sdk/types"
)

// TODO: execution_mode : <SuiTransactionBlockBuilderMode>
func (c *Client) BatchTransaction(
	ctx context.Context,
	signer *sui_types.SuiAddress,
	txnParams []map[string]interface{},
	gas *sui_types.ObjectID,
	gasBudget uint64,
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, batchTransaction, signer, txnParams, gas, gasBudget)
}

// MergeCoins Create an unsigned transaction to merge multiple coins into one coin.
func (c *Client) MergeCoins(
	ctx context.Context,
	signer *sui_types.SuiAddress,
	primaryCoin, coinToMerge *sui_types.ObjectID,
	gas *sui_types.ObjectID,
	gasBudget types.SafeSuiBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, mergeCoins, signer, primaryCoin, coinToMerge, gas, gasBudget)
}

// MoveCall Create an unsigned transaction to execute a Move call on the network, by calling the specified function in the module of a given package.
// TODO: not support param `typeArguments` yet.
// So now only methods with `typeArguments` are supported
// TODO: execution_mode : <SuiTransactionBlockBuilderMode>
func (c *Client) MoveCall(
	ctx context.Context,
	signer *sui_types.SuiAddress,
	packageId *sui_types.ObjectID,
	module, function string,
	typeArgs []string,
	arguments []any,
	gas *sui_types.ObjectID,
	gasBudget types.SafeSuiBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(
		ctx,
		&resp,
		moveCall,
		signer,
		packageId,
		module,
		function,
		typeArgs,
		arguments,
		gas,
		gasBudget,
	)
}

func (c *Client) Pay(
	ctx context.Context,
	signer *sui_types.SuiAddress,
	inputCoins []sui_types.ObjectID,
	recipients []*sui_types.SuiAddress,
	amount []types.SafeSuiBigInt[uint64],
	gas *sui_types.ObjectID,
	gasBudget types.SafeSuiBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, pay, signer, inputCoins, recipients, amount, gas, gasBudget)
}

// PayAllSui Create an unsigned transaction to send all SUI coins to one recipient.
func (c *Client) PayAllSui(
	ctx context.Context,
	signer, recipient *sui_types.SuiAddress,
	inputCoins []sui_types.ObjectID,
	gasBudget types.SafeSuiBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, payAllSui, signer, inputCoins, recipient, gasBudget)
}

func (c *Client) PaySui(
	ctx context.Context,
	signer *sui_types.SuiAddress,
	inputCoins []sui_types.ObjectID,
	recipients []*sui_types.SuiAddress,
	amount []types.SafeSuiBigInt[uint64],
	gasBudget types.SafeSuiBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, paySui, signer, inputCoins, recipients, amount, gasBudget)
}

func (c *Client) Publish(
	ctx context.Context,
	sender *sui_types.SuiAddress,
	compiledModules []*suiBase64Data,
	dependencies []sui_types.ObjectID,
	gas *sui_types.ObjectID,
	gasBudget uint,
) (*types.TransactionBytes, error) {
	var resp types.TransactionBytes
	return &resp, c.CallContext(ctx, &resp, publish, sender, compiledModules, dependencies, gas, gasBudget)
}

func (c *Client) RequestAddStake(
	ctx context.Context,
	signer *sui_types.SuiAddress,
	coins []sui_types.ObjectID,
	amount types.SuiBigInt,
	validator *sui_types.SuiAddress,
	gas *sui_types.ObjectID,
	gasBudget types.SuiBigInt,
) (*types.TransactionBytes, error) {
	var resp types.TransactionBytes
	return &resp, c.CallContext(ctx, &resp, requestAddStake, signer, coins, amount, validator, gas, gasBudget)
}

func (c *Client) RequestWithdrawStake(
	ctx context.Context,
	signer *sui_types.SuiAddress,
	stakedSuiId sui_types.ObjectID,
	gas *sui_types.ObjectID,
	gasBudget types.SuiBigInt,
) (*types.TransactionBytes, error) {
	var resp types.TransactionBytes
	return &resp, c.CallContext(ctx, &resp, requestWithdrawStake, signer, stakedSuiId, gas, gasBudget)
}

// SplitCoin Create an unsigned transaction to split a coin object into multiple coins.
func (c *Client) SplitCoin(
	ctx context.Context,
	signer *sui_types.SuiAddress,
	Coin *sui_types.ObjectID,
	splitAmounts []types.SafeSuiBigInt[uint64],
	gas *sui_types.ObjectID,
	gasBudget types.SafeSuiBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, splitCoin, signer, Coin, splitAmounts, gas, gasBudget)
}

// SplitCoinEqual Create an unsigned transaction to split a coin object into multiple equal-size coins.
func (c *Client) SplitCoinEqual(
	ctx context.Context,
	signer *sui_types.SuiAddress,
	Coin *sui_types.ObjectID,
	splitCount types.SafeSuiBigInt[uint64],
	gas *sui_types.ObjectID,
	gasBudget types.SafeSuiBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, splitCoinEqual, signer, Coin, splitCount, gas, gasBudget)
}

// TransferObject Create an unsigned transaction to transfer an object from one address to another. The object's type must allow public transfers
func (c *Client) TransferObject(
	ctx context.Context,
	signer, recipient *sui_types.SuiAddress,
	objID *sui_types.ObjectID,
	gas *sui_types.ObjectID,
	gasBudget types.SafeSuiBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, transferObject, signer, objID, gas, gasBudget, recipient)
}

// TransferSui Create an unsigned transaction to send SUI coin object to a Sui address. The SUI object is also used as the gas object.
func (c *Client) TransferSui(
	ctx context.Context, signer, recipient *sui_types.SuiAddress, suiObjID *sui_types.ObjectID, amount,
	gasBudget types.SafeSuiBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, transferSui, signer, suiObjID, gasBudget, recipient, amount)
}
