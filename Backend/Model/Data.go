package Model

type Data struct {
	ID      string `json:"id" bson:"_id,omitempty"`
	Percent string `json:"percent"`
}
