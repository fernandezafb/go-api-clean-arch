package inmemory

import (
	"github.com/fernandezafb/go-api-clean-arch/domain"
)

type db struct {
	items map[int64]*domain.Item
	seq   int64
}

func NewDb() DbClient {
	return &db{
		items: map[int64]*domain.Item{},
		seq:   1,
	}
}

type DbClient interface {
	FindAll() ([]*domain.Item, error)
	FindOne(id int64) (*domain.Item, error)
	InsertOne(item *domain.Item) (int64, error)
	UpdateOne(item *domain.Item) (int64, error)
	DeleteOne(id int64) (int64, error)
	CountAll() int64
}

func (d *db) FindAll() ([]*domain.Item, error) {
	list := []*domain.Item{}
	for _, i := range d.items {
		list = append(list, i)
	}
	return list, nil
}

func (d *db) FindOne(id int64) (*domain.Item, error) {
	item, ok := d.items[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return item, nil
}

func (d *db) InsertOne(item *domain.Item) (int64, error) {
	item.Id = d.seq
	d.items[item.Id] = item
	d.seq++

	return item.Id, nil
}

func (d *db) UpdateOne(i *domain.Item) (int64, error) {
	item, ok := d.items[i.Id]
	if !ok {
		return 0, domain.ErrNotFound
	}

	item.Name = i.Name
	item.Quantity = i.Quantity
	item.SellerId = i.SellerId
	item.Price = i.Price

	return i.Id, nil
}

func (d *db) DeleteOne(id int64) (int64, error) {
	_, ok := d.items[id]
	if !ok {
		return 0, domain.ErrNotFound
	}

	delete(d.items, id)
	d.seq--

	return id, nil
}

func (d *db) CountAll() int64 {
	return d.seq
}
