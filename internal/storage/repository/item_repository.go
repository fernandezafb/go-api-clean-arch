package repository

import (
	"fmt"

	"github.com/fernandezafb/go-api-clean-arch/domain"
	"github.com/fernandezafb/go-api-clean-arch/internal/storage/inmemory"
)

type itemRepository struct {
	db inmemory.DbClient
}

func NewItemRepository(db inmemory.DbClient) domain.ItemRepository {
	return &itemRepository{
		db: db,
	}
}

func (i *itemRepository) Create(item *domain.Item) (int64, error) {
	return i.db.InsertOne(item)
}

func (i *itemRepository) FetchByItemId(id int64) (*domain.Item, error) {
	item, err := i.db.FindOne(id)
	if err != nil {
		return nil, fmt.Errorf("error fetching item %d: %w", id, err)
	}
	return item, nil
}

func (i *itemRepository) Update(item *domain.Item) (int64, error) {
	id, err := i.db.UpdateOne(item)
	if err != nil {
		return 0, fmt.Errorf("error updating item %d: %w", item.Id, err)
	}
	return id, nil
}

func (i *itemRepository) DeleteByItemId(id int64) (int64, error) {
	rid, err := i.db.DeleteOne(id)
	if err != nil {
		return 0, fmt.Errorf("error deleting item %d: %w", id, err)
	}
	return rid, nil
}
