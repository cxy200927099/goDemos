package n63

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"syscall"
)

type JobItem struct {
	DirKey string
	DirName string
	PicUrl	string
}


var downloadChannel chan *JobItem

func InitDownloadManager(){
	downloadChannel = make(chan *JobItem, 20)
	go handleDownloadJob(downloadChannel)
}

func DownloadJob(job *JobItem){
	go func() {
		downloadChannel<-job
	}()
}

func handleDownloadJob(dlChan chan *JobItem){

	for{
		job := <- dlChan
		doDownloadJob(job)
	}
}

func doDownloadJob(item *JobItem){
	log.Printf("download %s to dir:%s\n", item.PicUrl, item.DirName)
	err := downloadFile(item.PicUrl, item.DirName)
	if err != nil{
		log.Printf("Error download file failed:%s errMsg:%s\n", item.PicUrl, err.Error())
	}
}

func createDirIfNotExist(dir string) error{
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}
	return nil
}

// DownloadFile will download a url and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
func downloadFile(url string, filepath string) error {
	// Create the file
	dir := path.Dir(filepath)
	createDirIfNotExist(dir)
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}