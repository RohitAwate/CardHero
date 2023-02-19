package db

import (
	"CardHero/models"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
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
		err = conn.Preload(clause.Associations).Find(&folder, condition, owner.ID, name).Error
	} else {
		condition += "= ?"
		err = conn.Preload(clause.Associations).Find(&folder, condition, owner.ID, name, parent.ID).Error
	}

	return folder, err
}

func ResolveFolder(path string, user models.User) (*models.Folder, error) {
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

	return parent, nil
}

const (
	GetFolderHierarchyCTE = `
		with recursive folder_cte as (
			select f.id, f.name, f.parent_id, f.owner_id 
			from folders f 
			where parent_id = ?
			and owner_id = ?
			
			union all 	
			
			select f.id, f.name, f.parent_id, f.owner_id 
			from folders f
			inner join folder_cte fcte on f.parent_id = fcte.id
		)
		select *
		from folder_cte;
	`
)

func GetFolderHierarchy(parent models.Folder) (*models.FolderHierarchy, error) {
	conn := getConn()

	var folders []models.Folder
	err := conn.Raw(GetFolderHierarchyCTE, parent.ID, parent.OwnerID).Scan(&folders).Error
	if err != nil {
		return nil, err
	}

	fh := models.BuildHierarchy(parent, folders)
	return &fh, nil
}

func GetCardsInFolder(folderID uuid.UUID, owner models.User) ([]models.Card, error) {
	var cards []models.Card

	conn := getConn()

	err := conn.Order("timestamp desc").Find(&cards, "folder_id = ? and owner_id = ?", folderID, owner.ID).Error
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
