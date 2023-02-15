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
		if strings.HasPrefix(token, FolderStartCommand) {
			// Stripping away the double forward slash
			token = token[2:]

			folders := strings.Split(token, FolderDelimiter)

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
		*folder = existingFolder
	}
}
