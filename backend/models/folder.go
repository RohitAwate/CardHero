package models

import (
	uuid "github.com/satori/go.uuid"
)

type Folder struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key"`

	Name string `json:"name,omitempty" gorm:"uniqueIndex:unique_folder"`

	ParentID *uuid.UUID `json:"-" gorm:"uniqueIndex:unique_folder"`
	Parent   *Folder    `json:"parent,omitempty"`

	OwnerID uuid.UUID `json:"-" gorm:"uniqueIndex:unique_folder"`
	Owner   User      `json:"owner"`
}

func NewFolder(name string, parent *Folder, owner User) Folder {
	folder := Folder{
		ID:      uuid.NewV4(),
		Name:    name,
		OwnerID: owner.ID,
		Owner:   owner,
	}

	if parent != nil {
		parentCopy := *parent
		folder.Parent = &parentCopy
		folder.ParentID = &parentCopy.ID
	}

	return folder
}

const (
	RootFolderName = "Root"
)

func NewRoot(owner User) Folder {
	return NewFolder(RootFolderName, nil, owner)
}
