package db

import (
	"database/sql"

	"github.com/Neil-Uli/Restful-go/items-api/models"
)

// File responsible for interacting with the items table. Such interactions
// include fetching all list items, creating an item, fetching an item using
// its ID as well as updating and deleting them:

func (db Database) GetAllItems() (*models.ItemList, error) {
	list := &models.ItemList{}
	rows, err := db.Conn.Query("SELECT * FROM items ORDER BY ID DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}
func (db Database) AddItem(item *models.Item) error {
	var id int
	var createdAt string
	query := `INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, item.Name, item.Description).Scan(&id, &createdAt)
	if err != nil {
		return err
	}
	item.ID = id
	item.CreatedAt = createdAt
	return nil
}
func (db Database) GetItemById(itemId int) (models.Item, error) {
	item := models.Item{}
	query := `SELECT * FROM items WHERE id = $1;`
	row := db.Conn.QueryRow(query, itemId)
	switch err := row.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt); err {
	case sql.ErrNoRows:
		return item, ErrNoMatch
	default:
		return item, err
	}
}
