package models

type Txn struct {
	ID string
	Value float64
	Card  Card
	Timestamp
}
