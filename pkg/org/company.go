package org

// A Company is a collection of names and mission statements.
type Company struct {
	ID      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Mission string `db:"mission" json:"mission"`
}
