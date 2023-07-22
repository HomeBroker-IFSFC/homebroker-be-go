package transformer

import (
	"github.com/JoaoRafa19/homebroker/go/internal/market/dto"
	"github.com/JoaoRafa19/homebroker/go/internal/market/entity"
)

func TransformInput(input dto.TradeInput) *entity.Order {
	asset:= entity.NewAsset(input.AssetID, input.AssetID, 1000)
	investor:=entity.NewInvestor(input.InvestorID)
	order:=entity.NewOrder(input.OrderID, investor, asset, input.Shares, input.Price, entity.OrderType(input.OrderType))

	if input.CurrentShares > 0 {
		assetPosition := entity.NewInvestorAssetPosition(input.AssetID, input.CurrentShares)
		investor.AddAssetPosition(assetPosition)
	}
	return order
}


func TransformOutput (order *entity.Order) *dto.OrderOutput {
	output:= &dto.OrderOutput{
		OrderID: order.ID,
		InvestorID: order.Investor.ID,
		AssetID: order.Asset.ID,
		OrderType: order.OrderType.String(),
		Status: order.Status.String(),
		Shares: order.Shares,
		Partial: order.PendingShares,
		
	}

	var transactionsOutput []*dto.TransactionOutput
	for _,t := range order.Transactions {
		transactionOutput := &dto.TransactionOutput{
			TransactionID: t.ID,
			BuyerID: t.PurchaseOrder.ID,
			SellerID: t.SalesOrder.ID,
			AssetID: t.SalesOrder.Asset.ID,
			Price: t.Price,
			Shares: t.SalesOrder.Shares - t.SalesOrder.PendingShares,

		}
		transactionsOutput = append(transactionsOutput, transactionOutput)

	}
	output.TransactionOutput = transactionsOutput
	return output
}