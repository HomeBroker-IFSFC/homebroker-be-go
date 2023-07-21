package entity

type Asset struct {
	ID           string
	Name         string
	MaeketVolume int
}

func NewAsset(id string, name string, marketVolume int) * Asset {
	return &Asset{
		Name: name,
		ID: id,
		MaeketVolume: marketVolume,
	}
}
