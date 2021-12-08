package enhance

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

// RunVideoEnhance - runs the veai.exe cmd.
func ExecEnhance(ipath, format, opath string) {

	var file string

	wg := new(sync.WaitGroup)
	wg.Add(2)
	// read the directory of the files (always dropbox veai)
	fp, err := os.ReadDir(ipath)
	if err != nil {
		log.Println(err)
		return
	}

	// read dir for files, if the index is greater than zero, process index[1]
	for _, v := range fp {
		if len(fp) > 0 {
			file = v.Name()
		}
	}

	// var cmdString string = fmt.Sprintf("veai.exe -i %s -f %s -o %s", file, format, opath)
	var standinCmd string = fmt.Sprintf("VideoEnhanceAI -i %s -f %s -o %s", file, format, opath)

	fmt.Println(standinCmd)

	// exec cmdString.
	cmd := exec.Command(standinCmd)
	if err = cmd.Run(); err != nil {
		log.Fatal("could not run veai command", err.Error())
	}

	log.Println("running veai.")
}
