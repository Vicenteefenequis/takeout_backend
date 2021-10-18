package domain

type Post struct {
	NameUser    string `bson:"name_user,omitempty" json:"name_user,omitempty" validate:"required,maxlength:2,label:Nome de Usuário"`
	TypePost    string `bson:"type_post,omitempty" json:"type_post,omitempty" validate:"required,maxlength:10,label:Tipo de Post"`
	Visibility  string `bson:"visibility,omitempty" json:"visibility,omitempty" validate:"required,maxlength:10,label:Visibilidade"`
	CaseStatus  string `bson:"case_status,omitempty" json:"case_status,omitempty" validate:"required,maxlength:10,label:Status do Caso"`
	Image       string `bson:"image,omitempty" json:"image,omitempty" validate:"required,maxlength:255,label:Imagem"`
	Description string `bson:"description,omitempty" json:"description,omitempty" validate:"required,label:Description"`
}
