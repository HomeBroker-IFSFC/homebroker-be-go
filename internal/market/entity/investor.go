package entity

type Investor struct {
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) * Investor {
	return &Investor{
		ID: id,
		AssetPosition: []*InvestorAssetPosition{},
		
	}
}

func (i * Investor) AddAssetPosition(assetPosition *InvestorAssetPosition){
	i.AssetPosition = append(i.AssetPosition, assetPosition)
} 

func (i * Investor) UpdateAssetPosition(assetId string, qtdShares int){
	assetPosition := i.GetAssetPosition(assetId)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetId, qtdShares))
	}else {
		assetPosition.Shares += qtdShares
	}
}

func (i* Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

type InvestorAssetPosition struct {
	AssetID string
	Shares int 
}


func NewInvestorAssetPosition (assetID string, shares int) *InvestorAssetPosition{
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares: shares,
	}

}