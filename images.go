package buki
import (
	"net/http"
	"os"
	"io"
	"path"
	"errors"
	"io/ioutil"
	"os/exec"
)

const baseImagePath = "/home/mustafa/buki/images"
const baseVMPath = "/home/mustafa/buki/vms"

func InitStorage(){

}

func GetAvailableImages() []string{
	files, _ := ioutil.ReadDir(baseImagePath)

	s := make([]string, 0, 5);
	for _, f := range files {
		s = append(s, f.Name())
	}

	return s
}

func GetImagePath(name string) string{
	return path.Join(baseImagePath, string(name + ".img"))
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

func GetVMPrimaryDiskName(vmName string) string{
	vmFolder := path.Join(baseVMPath, vmName)
	vmFileName := path.Join(vmFolder, "disk0.img")
	return vmFileName
}

func CopyImage(imgName, vmName, size string) error{
	// First Copy the image to place
	// TODO: Add error checks
	imageFileName := GetImagePath(imgName)
	vmFolder := path.Join(baseVMPath, vmName)
	vmFileName := GetVMPrimaryDiskName(vmName)
	os.MkdirAll(vmFolder, 0777)

	CopyFile(imageFileName, vmFileName)
	cmd  := exec.Command("qemu-img", "resize", vmFileName, size)
	err := cmd.Run()
	return err
}

func CreateCloudConfig(vmName, userData string) (string, error) {
	// Writes files to
	vmFolder := path.Join(baseVMPath, vmName)
	userDataFilePath := path.Join(vmFolder,  "user-data")
	cloudInitImgPath := path.Join(vmFolder, "cloud-init.img")

	ioutil.WriteFile(userDataFilePath , []byte(userData), 0777 )

	cmd  := exec.Command("cloud-localds", cloudInitImgPath, userDataFilePath)
	err := cmd.Run()
	return cloudInitImgPath, err
}