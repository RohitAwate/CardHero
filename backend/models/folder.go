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
	RootFolderName    = "Root"
	DefaultFolderName = "Default"
)

func NewRoot(owner User) Folder {
	return NewFolder(RootFolderName, nil, owner)
}

const (
	FolderStartCommand = "//"
	FolderDelimiter    = "/"
)

type FolderHierarchy struct {
	ID         uuid.UUID         `json:"id"`
	FolderName string            `json:"name"`
	Children   []FolderHierarchy `json:"children"`
}

func BuildHierarchy(parent Folder, folders []Folder) FolderHierarchy {
	fh := FolderHierarchy{
		ID: parent.ID, FolderName: parent.Name, Children: []FolderHierarchy{},
	}

	for _, folder := range folders {
		if folder.ID != parent.ID && *folder.ParentID == parent.ID {
			childFH := BuildHierarchy(folder, folders)
			fh.Children = append(fh.Children, childFH)
		}
	}

	return fh
}
