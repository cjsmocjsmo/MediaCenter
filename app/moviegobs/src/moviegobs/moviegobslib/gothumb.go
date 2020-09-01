package moviegolib

import (
	"fmt"
	//because I want it
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
	"net/url"
	"github.com/disintegration/imaging"
	"github.com/globalsign/mgo/bson"
)

//ThumbInFo struct exported to setup
type ThumbInFo struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	MovName   string        `bson:"movname"`
	BasePath  string        `bson:"baspath"`
	DirPATH   string        `bson:"dirpath"`
	ThumbPath string        `bson:"thumbpath"`
	ThumbID   string        `bson:"thumbid"`
	ThumbPathTwo string `bson:"thumbpathtwo"`
}

//UUID holds the unique identifier for the file
func UUID() (UUID string) {
	aSeed := time.Now()
	aseed := aSeed.UnixNano()
	rand.Seed(aseed)
	u := rand.Int63n(aseed)
	UUID = strconv.FormatInt(u, 10)
	return
}

func myPathSplit(myPath string) (DirPath string, BaseNAme string, MOvName string, Ext string) {
	DirPath, BaseNAme = path.Split(myPath)
	Ext = BaseNAme[3:]
	factor := len(BaseNAme) - 4
	MOvName = BaseNAme[:factor]
	return
}

func getServerAddr() (addr string) {
	addr = os.Getenv("MOVIEGOBS_SERVER_ADDRESS")
	return
}

func getServerPort() (port string) {
	port = os.Getenv("MOVIEGOBS_SERVER_PORT")
	return
}

func getThumbPath() (tpath string) {
	tpath = os.Getenv("MOVIEGOBS_THUMBNAIL_PIC_PATH")
	return
}
//CreateMoviesThumbnail exported to setup
func CreateMoviesThumbnail(p string) (ThumbINFO ThumbInFo) {
// func CreateMoviesThumbnail(p string) (thumbINFO ThumbInFo) {
	dirpath, basepath, movname, ext := myPathSplit(p)
	ThumbINFO.ID = bson.NewObjectId()
	ThumbINFO.MovName = movname
	ThumbINFO.BasePath = basepath
	ThumbINFO.DirPATH = dirpath

	if ext == ".txt" {
		fmt.Print("what the fuck a text file")
	} else {
		// MSA := os.Getenv("MOVIEGOBS_SERVER_ADDRESS")
		MSA := getServerAddr()

		// MSP := os.Getenv("MOVIEGOBS_SERVER_PORT")
		MSP := getServerPort()

		// MTPP := os.Getenv("MOVIEGOBS_THUMBNAIL_PIC_PATH")
		MTPP := getThumbPath()

		BP := "/" + url.QueryEscape(basepath)
		
		ThumbINFO.ThumbPathTwo = MSA + ":" + MSP + MTPP[1:] + BP

		thumbpath := MTPP + "/" + basepath
		// log.Printf("\n\n THIS IS THUMBPATH %v \n\n", thumbpath)
		// log.Printf("\n\n THIS IS THUMBPATH2 %v \n\n", thumbpathtwo)
		ThumbINFO.ThumbPath = thumbpath
		ThumbINFO.ThumbID = UUID()
		_, err := os.Stat(thumbpath)
		if err == nil {
			log.Printf("file %s exists", thumbpath)
		} else if os.IsNotExist(err) {
			pic, err := imaging.Open(p)
			if err != nil {
				log.Printf("\n this is file Open error jpgthumb %v \n", p)
			}
			thumbImage := imaging.Resize(pic, 0, 250, imaging.Lanczos)
			err = imaging.Save(thumbImage, thumbpath)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			log.Printf("file %s stat error: %v", thumbpath, err)
		}

		
		
		fmt.Println(&ThumbINFO)
		cmtses := DBcon()
		defer cmtses.Close()
		CMTc := cmtses.DB("movbsthumb").C("movbsthumb")
		err = CMTc.Insert(ThumbINFO)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("THIS IS THUMBINFO")
		log.Println(&ThumbINFO)
	}
	return
	// return ThumbINFO
	// return thumbINFO
}

//NoArtList exported to setup
var NoArtList []string

//FindPicPaths exported to setup
func FindPicPaths(mpath string, noartpicpath string) (result string) {
	_, _, movename, _ := myPathSplit(mpath)
	Tses := DBcon()
	defer Tses.Close()
	MTc := Tses.DB("movbsthumb").C("movbsthumb")
	b1 := bson.M{"movname": movename}
	b2 := bson.M{"_Id": 0}
	var ThumbI []map[string]string
	err := MTc.Find(b1).Select(b2).All(&ThumbI)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ThumbI)
	LenI := len(ThumbI)
	if LenI == 0 {
		NoArtList = append(NoArtList, mpath)
		result = noartpicpath
	} else {
		result = ThumbI[0]["thumbpath"]
		fmt.Printf("THIS IS THUMBI.THUMBPATH:  %s", ThumbI[0]["thumbpath"])
	}
	return
}

