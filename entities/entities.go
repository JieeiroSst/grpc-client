package entities

type Item struct {
	ID string `bson:"_id,omitempty"` 
	Quantity int 
}