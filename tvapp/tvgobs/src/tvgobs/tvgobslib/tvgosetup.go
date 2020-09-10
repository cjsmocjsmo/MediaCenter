package tvgobslib

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/globalsign/mgo"
)

//DBcon is exported for all our db connection objects
func DBcon() *mgo.Session {
	fmt.Println("Starting Update db session")
	s, err := mgo.Dial(os.Getenv("TVGOBS_MONGODB_ADDRESS"))
	if err != nil {
		fmt.Println("this is dial err")
		panic(err)
	}
	return s
}

func isDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func processTVShowInfo(pAth string) {
	log.Println("Process_Movs has started")
	offset := len(pAth) - 4
	noextPath := pAth[:offset]
	jpgextPath := noextPath + ".jpg"
	var tvpicPath string
	if _, err := os.Stat(jpgextPath); err == nil {
		tvpicPath = jpgextPath
	} else {
		tvpicPath = os.Getenv("TVGOBS_NO_ART_PIC_PATH")
	}

	var TvI TvShowInfoS
	TvI = getTvShowInfo(pAth, tvpicPath)
	ses := DBcon()
	defer ses.Close()
	MTc := ses.DB("tvgobs").C("tvgobs")
	err := MTc.Insert(&TvI)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func myDirVisit(pAth string, f os.FileInfo, err error) error {
	log.Printf("this is path: %s", pAth)
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	if f.IsDir() {
		return nil // not a file.  ignore.
	}

	// ext := filepath.Ext(pAth)
	// if ext == "" {
	// 	return nil //not a pic or movie
	// } else if ext == ".nfo" {
	// 	return nil //not a file we are interested in
	// } else if ext == ".txt" {
	// 	return nil // not a file we are interested in
	// }

	// switch ext {
	// 	case ".mp4":
	// 		processTVShowInfo(pAth)
	// 	case ".mkv":
	// 		processTVShowInfo(pAth)
	// 	case ".avi":
	// 		processTVShowInfo(pAth)
	// 	case ".m4v":
	// 		processTVShowInfo(pAth)
	// }

	ext := filepath.Ext(pAth)
	switch ext {
	case "":
		return nil
	case ".nfo":
		return nil
	case ".txt":
		return nil
	case ".mp4":
		processTVShowInfo(pAth)
	case ".mkv":
		processTVShowInfo(pAth)
	case ".avi":
		processTVShowInfo(pAth)
	case ".m4v":
		processTVShowInfo(pAth)
	}
	return nil
}

//SetUp is exported to main
func SetUp() (ExStat int) {
	//Start the timer
	startTime := time.Now().Unix()
	fmt.Printf("setup function has started at: %T", startTime)
	//Connect to the DB
	sess := DBcon()
	err := sess.DB("tvgobs").DropDatabase()
	if err != nil {
		fmt.Println(err)
	}

	err = filepath.Walk(os.Getenv("TVGOBS_TVSHOWS_PATH"), myDirVisit)
	if err != nil {
		fmt.Println(err)
	}

	os.Setenv("TVGOBS_SETUP", "0")
	fmt.Printf("this is noartlist :: %s", NoArtList)
	fmt.Println(startTime)
	stopTime := time.Now().Unix()
	fmt.Println(stopTime)
	etime := stopTime - startTime
	fmt.Println(etime)
	fmt.Println("SETUP IS COMPLETE")
	ExStat = 0
	return
}
