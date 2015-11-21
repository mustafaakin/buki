package buki
import (
	"net/http"
	"os"
	"io"
	"path"
	"errors"
	"io/ioutil"

)

var basePath = "/home/mustafa/buki/images"

func InitStorage(){

}

func GetAvailableImages() []string{
	files, _ := ioutil.ReadDir(basePath)

	s := make([]string, 0, 5);
	for _, f := range files {
		s = append(s, f.Name())
	}

	return s
}

func GetImagePath(name string) string{
	return path.Join(basePath, string(name + ".img"))
}

func DownloadImage(path, name string) (error){
	// TODO: If zipped, unzip automatically
	savePath := GetImagePath(name)

	// Check if already downloaded
	if 	_, err := os.Stat(savePath); os.IsExist(err) {
		return errors.New("Image already exists")
	}

	// Else download image
	out, err := os.Create(savePath)
	if err != nil {
		return errors.New("Could not create the file in path")
	}

	defer out.Close()
	resp, err := http.Get(path)
	if err != nil {
		return errors.New("Could not access the given image path")
	}

	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return errors.New("Could not finish downloading the image")
	}

	return nil
}