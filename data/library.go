package data

import "os"


func ListFolder(folder string) ([]string, error) {
	file, err := os.Open(folder)
	if err != nil { return []string{}, err }
	defer file.Close()
	list, err := file.Readdirnames(0)
	if err != nil { return []string{}, err }
	return list, nil
}


func Contains(list []string, filename string) bool {
	for _, file := range list {
		if file == filename { return true }
	}
	return false
}