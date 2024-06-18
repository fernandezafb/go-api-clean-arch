package domain

type ItemRequest struct {
	Name     string `json:"name"`
	Quantity uint64 `json:"quantity"`
	SellerId string `json:"sellerId"`
	Price    int64  `json:"price"`
}

type ItemResponse struct {
	Id int64 `json:"id"`
}

type ShowItemPropertiesResponse struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type Item struct {
	Id       int64
	Name     string
	Quantity uint64
	SellerId string
	Price    int64
}

type ItemRepository interface {
	Create(item *Item) (int64, error)
	FetchByItemId(id int64) (*Item, error)
	Update(item *Item) (int64, error)
	DeleteByItemId(id int64) (int64, error)
}

type ItemUsecase interface {
	SellItem(r ItemRequest) (ItemResponse, error)
	ShowItemProperties(id int64) (ShowItemPropertiesResponse, error)
	UpdateItemProperties(id int64, r ItemRequest) (ItemResponse, error)
	DeleteItem(id int64) (ItemResponse, error)
}
