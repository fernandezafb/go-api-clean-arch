package usecase

import (
	"github.com/fernandezafb/go-api-clean-arch/domain"
)

type itemUsecase struct {
	itemRepository domain.ItemRepository
}

func NewItemUsecase(ir domain.ItemRepository) domain.ItemUsecase {
	return &itemUsecase{
		itemRepository: ir,
	}
}

func (i *itemUsecase) SellItem(r domain.ItemRequest) (domain.ItemResponse, error) {
	item := &domain.Item{
		Name:     r.Name,
		Quantity: r.Quantity,
		SellerId: r.SellerId,
		Price:    r.Price,
	}

	id, err := i.itemRepository.Create(item)
	if err != nil {
		return domain.ItemResponse{}, err
	}
	return domain.ItemResponse{Id: id}, nil
}

func (i *itemUsecase) ShowItemProperties(id int64) (domain.ShowItemPropertiesResponse, error) {
	item, err := i.itemRepository.FetchByItemId(id)
	if err != nil {
		return domain.ShowItemPropertiesResponse{}, err
	}
	return domain.ShowItemPropertiesResponse{Name: item.Name, Price: item.Price}, nil
}

func (i *itemUsecase) UpdateItemProperties(id int64, r domain.ItemRequest) (domain.ItemResponse, error) {
	item := &domain.Item{
		Id:       id,
		Name:     r.Name,
		Quantity: r.Quantity,
		SellerId: r.SellerId,
		Price:    r.Price,
	}

	id, err := i.itemRepository.Update(item)
	if err != nil {
		return domain.ItemResponse{}, err
	}
	return domain.ItemResponse{Id: id}, nil
}

func (i *itemUsecase) DeleteItem(id int64) (domain.ItemResponse, error) {
	id, err := i.itemRepository.DeleteByItemId(id)
	if err != nil {
		return domain.ItemResponse{}, err
	}
	return domain.ItemResponse{Id: id}, nil
}
