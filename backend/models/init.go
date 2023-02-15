package models

func GetAll() []interface{} {
	return []interface{}{
		&User{}, &Card{}, &Folder{},
	}
}
