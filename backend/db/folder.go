package db

import (
	"CardHero/models"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"regexp"
	"strings"
)

const (
	FolderStartCommand = "//"
	FolderDelimiter    = "/"
)

func BuildFoldersFromCard(card models.Card, owner models.User) {
	whitespaceRegex := regexp.MustCompile("\\s+")
	tokens := whitespaceRegex.Split(card.Contents, -1)

	parent := models.NewRoot(owner)
	SaveFolder(&parent)

	for _, token := range tokens {
		if strings.HasPrefix(token, FolderStartCommand) {
			folders := strings.Split(token, FolderDelimiter)

			for _, folder := range folders[2:] {
				// Skipping first 2 since they would be empty because of the "//"
				newFolder := models.NewFolder(folder, &parent, owner)
				SaveFolder(&newFolder)

				parent = newFolder
			}
		}
	}
}

func GetFolder(folder models.Folder) (models.Folder, error) {
	conn := getConn()

	condition := "owner_id = ? and name = ? and parent_id "
	args := []interface{}{folder.OwnerID, folder.Name}

	var existingFolder models.Folder
	var err error

	if folder.Parent == nil {
		condition += "is null"
		err = conn.Find(&existingFolder, condition, folder.OwnerID, folder.Name).Error
	} else {
		condition += "= ?"
		args = append(args, folder.ParentID)
		err = conn.Find(&existingFolder, condition, folder.OwnerID, folder.Name, folder.ParentID).Error
	}

	return existingFolder, err
}

func SaveFolder(folder *models.Folder) {
	conn := getConn()

	existingFolder, err := GetFolder(*folder)
	if err != nil {
		fmt.Println(err)
		return
	}

	if existingFolder.ID == uuid.Nil {
		conn.Create(folder)
	} else {
		fmt.Println("Found existing: ", existingFolder.Name)
		*folder = existingFolder
	}
}
