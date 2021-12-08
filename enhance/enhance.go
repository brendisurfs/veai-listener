package enhance

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

// RunVideoEnhance - runs the veai.exe cmd.
func execEnhance(ipath, format, opath string) (string, error) {

	var file string

	// read the directory of the files (always dropbox veai)
	fp, err := os.ReadDir(ipath)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// read dir for files, if the index is greater than zero, process index[1]
	for _, v := range fp {
		if len(fp) > 0 {
			file = v.Name()
		}
	}

	// var cmdString string = fmt.Sprintf("veai.exe -i %s -f %s -o %s", file, format, opath)
	var standinCmd string = fmt.Sprintf("-i %s -f %s -o %s", file, format, opath+file)

	return standinCmd, nil

}

func RunCmd(ipath, format, opath string, event fsnotify.Event) {

	command, err := execEnhance(ipath, format, opath)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("VideoEnhanceAI", command)

	if err = cmd.Start(); err != nil {
		log.Fatal("could not run command\n", err.Error())
	}
	cmd.Process.Wait()

	log.Println("done running video enhance.")

}
