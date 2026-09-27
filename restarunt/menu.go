package restaurant

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type MenuItem struct {
	ItemID       int     `json:"item_id"`
	Name         string  `json:"name"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	IsVegetarian bool    `json:"is_vegetarian"`
}

type Menu struct {
	Items []MenuItem `json:"items"`
}

// AddMenuItem adds a new item
func (m *Menu) AddMenuItem(item MenuItem) error {

	if strings.TrimSpace(item.Name) == "" {
		return errors.New("item name cannot be empty")
	}

	if item.Price <= 0 {
		return errors.New("price must be positive")
	}

	m.Items = append(m.Items, item)

	return nil
}

// FindItemByName searches item by name
func (m *Menu) FindItemByName(name string) (*MenuItem, error) {

	for i := range m.Items {

		if strings.EqualFold(m.Items[i].Name, name) {
			return &m.Items[i], nil
		}
	}

	return nil, errors.New("item not found")
}

// UpdatePrice updates price
func (m *Menu) UpdatePrice(id int, newPrice float64) error {

	if newPrice <= 0 {
		return errors.New("price must be positive")
	}

	for i := range m.Items {

		if m.Items[i].ItemID == id {
			m.Items[i].Price = newPrice
			return nil
		}
	}

	return errors.New("item not found")
}

// RemoveMenuItem removes item
func (m *Menu) RemoveMenuItem(id int) error {

	for i := range m.Items {

		if m.Items[i].ItemID == id {

			m.Items = append(m.Items[:i], m.Items[i+1:]...)

			return nil
		}
	}

	return errors.New("item not found")
}

// SaveToFile saves menu data into JSON file
func (m *Menu) SaveToFile(filename string) error {

	data, err := json.MarshalIndent(m, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// LoadFromFile loads menu data from JSON file
func (m *Menu) LoadFromFile(filename string) error {

	data, err := os.ReadFile(filename)

	if err != nil {

		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return json.Unmarshal(data, m)
}