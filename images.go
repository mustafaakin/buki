package buki

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
)

const baseImagePath = "/var/lib/libvirt/images/base-images"
const baseVMPath = "/var/lib/libvirt/images/vms"

func InitStorage() {

}

func GetAvailableImages() []string {
	files, _ := ioutil.ReadDir(baseImagePath)

	s := make([]string, 0, 5)
	for _, f := range files {
		s = append(s, f.Name())
	}

	return s
}

func GetImagePath(name string) string {
	return path.Join(baseImagePath, string(name+".img"))
}

func DownloadImage(path, name string) error {
	// TODO: If zipped, unzip automatically
	savePath := GetImagePath(name)

	// Check if already downloaded
	if _, err := os.Stat(savePath); os.IsExist(err) {
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

func GetVMPrimaryDiskName(vmName string) string {
	vmFolder := path.Join(baseVMPath, vmName)
	vmFileName := path.Join(vmFolder, "disk0.img")
	return vmFileName
}

// CopyImage copies the given image to output directory, with given resize option
func CopyImage(imgName, vmName, size string) error {
	// First Copy the image to place
	// TODO: Add error checks
	imageFileName := GetImagePath(imgName)
	vmFolder := path.Join(baseVMPath, vmName)
	vmFileName := GetVMPrimaryDiskName(vmName)
	os.MkdirAll(vmFolder, 0777)

	err2 := CopyFile(imageFileName, vmFileName)
	fmt.Println(err2)

	cmd := exec.Command("qemu-img", "resize", vmFileName, size)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))

	return err
}

// CreateCloudConfig generates valid cloud-init images for vms
func CreateCloudConfig(vmName, userData string) (string, error) {
	// Writes files to
	vmFolder := path.Join(baseVMPath, vmName)
	userDataFilePath := path.Join(vmFolder, "user-data")
	cloudInitImgPath := path.Join(vmFolder, "cloud-init.img")

	ioutil.WriteFile(userDataFilePath, []byte(userData), 0777)

	cmd := exec.Command("cloud-localds", cloudInitImgPath, userDataFilePath)
	err := cmd.Run()
	return cloudInitImgPath, err
}
