package domain

type Post struct {
	NameUser    string `bson:"name_user,omitempty" json:"name_user,omitempty"`
	TypePost    string `bson:"type_post,omitempty" json:"type_post,omitempty"`
	Visibility  string `bson:"visibility,omitempty" json:"visibility,omitempty"`
	CaseStatus  string `bson:"case_status,omitempty" json:"case_status,omitempty"`
	Image       string `bson:"image,omitempty" json:"image,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}
