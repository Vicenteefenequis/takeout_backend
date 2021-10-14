package domain

type Post struct {
	NameUser    string `bson:"name_user" json:"name_user"`
	TypePost    string `bson:"type_post" json:"type_post"`
	Visibility  string `bson:"visibility" json:"visibility"`
	CaseStatus  string `bson:"case_status" json:"case_status"`
	Image       string `bson:"image" json:"image"`
	Description string `bson:"description" json:"description"`
}
