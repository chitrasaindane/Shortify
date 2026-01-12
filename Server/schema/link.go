package schema

type Link struct {
	Title          *string `json:"title" bson:"title"`
	Description    *string `json:"description" bson:"description"`
	DestinationURL *string `json:"destination_url" bson:"destination_url"`
	Slug           *string `json:"slug" bson:"slug"`
}

type LinkRequest struct {
	Data *Link `json:"data"`
}
