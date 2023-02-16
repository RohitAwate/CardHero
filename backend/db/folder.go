package db

import (
	"CardHero/models"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"regexp"
	"strings"
)

func IngestCard(card models.Card, user models.User) {
	BuildFoldersFromCard(&card, user)
	SaveCard(card)
}

func BuildFoldersFromCard(card *models.Card, owner models.User) {
	whitespaceRegex := regexp.MustCompile("\\s+")
	tokens := whitespaceRegex.Split(card.Contents, -1)

	root := models.NewRoot(owner)
	SaveFolder(&root)

	ch := make(chan models.Folder)
	go CreateDefaultFolder(root, ch)

	parent := root
	for _, token := range tokens {
		if strings.HasPrefix(token, models.FolderStartCommand) {
			// Stripping away the double forward slash
			token = token[2:]

			folders := strings.Split(token, models.FolderDelimiter)

			for _, folderName := range folders {
				// This could happen if there are redundant slashes
				//like so
				// - //dev/jetbrains/goland/ -- the trailing slash here
				// - ///recipes/pasta -- the 3 slashes at the start
				if folderName == "" {
					continue
				}

				newFolder := models.NewFolder(folderName, &parent, owner)
				SaveFolder(&newFolder)

				parent = newFolder
			}
		}
	}

	if parent != root {
		card.AssignFolder(parent)
	} else {
		defaultFolder := <-ch
		card.AssignFolder(defaultFolder)
	}
}

func CreateDefaultFolder(root models.Folder, ch chan models.Folder) {
	defaultFolder := models.NewFolder("Default", &root, root.Owner)
	SaveFolder(&defaultFolder)
	ch <- defaultFolder
}

func GetFolder(name string, parent *models.Folder, owner models.User) (models.Folder, error) {
	conn := getConn()

	condition := "owner_id = ? and name = ? and parent_id "

	var folder models.Folder
	var err error

	if parent == nil {
		condition += "is null"
		err = conn.Find(&folder, condition, owner.ID, name).Error
	} else {
		condition += "= ?"
		err = conn.Find(&folder, condition, owner.ID, name, parent.ID).Error
	}

	return folder, err
}

func GetFolderStructure(path string, user models.User) (*models.FolderStructure, error) {
	// Traverse to that folder first
	folders := strings.Split(path, models.FolderDelimiter)
	var parent *models.Folder = nil
	for _, folderName := range folders {
		if folderName == "" {
			continue
		}

		folder, err := GetFolder(folderName, parent, user)
		if err != nil {
			return nil, err
		}

		parent = &folder
	}

	fs := models.FolderStructure{ID: parent.ID, FolderName: parent.Name}
	children, err := GetChildFolders(parent)
	if err != nil {
		return nil, err
	}

	for _, child := range children {
		childPath := path + "/" + child.Name
		childFS, err := GetFolderStructure(childPath, user)
		if err != nil {
			return nil, err
		}

		fs.Children = append(fs.Children, *childFS)
	}

	// If no children, don't leave the array nil
	// Make it an empty array
	if fs.Children == nil {
		fs.Children = []models.FolderStructure{}
	}

	return &fs, nil
}

func GetChildFolders(parent *models.Folder) ([]models.Folder, error) {
	conn := getConn()

	condition := "owner_id = ? and parent_id "

	var children []models.Folder
	var err error

	if parent == nil {
		condition += "is null"
		err = conn.Find(&children, condition, parent.OwnerID).Error
	} else {
		condition += "= ?"
		err = conn.Find(&children, condition, parent.OwnerID, parent.ID).Error
	}

	return children, err
}

func GetCardsInFolder(folderID uuid.UUID, owner models.User) ([]models.Card, error) {
	var cards []models.Card

	fmt.Println(folderID, owner.ID)

	conn := getConn()

	err := conn.Find(&cards, "folder_id = ? and owner_id = ?", folderID, owner.ID).Error
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func SaveFolder(folder *models.Folder) {
	conn := getConn()

	existingFolder, err := GetFolder(folder.Name, folder.Parent, folder.Owner)
	if err != nil {
		fmt.Println(err)
		return
	}

	if existingFolder.ID == uuid.Nil {
		conn.Create(folder)
	} else {
		*folder = existingFolder
	}
}
