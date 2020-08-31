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
	// "strings"
)

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

// Star Trek STTV S01E29 The City On The Edge Of Forever.jpg'
// func myPathSplitTV(myPath string) (DirPath string, BaseName string, Ext string, Season string, Episode string, EpisodeName string) {
// 	DirPath, BaseName = path.Split(myPath)
// 	Ext = BaseName[3:]
// 	Season = BaseName[17:19]
// 	Episode = BaseName[20:22]
// 	factor := len(BaseName) - 4
// 	EpisodeName = BaseName[24:factor]
// 	return
// }

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

//CreateMoviesThumbnail exported to setup
func CreateMoviesThumbnail(p string) (thumbINFO ThumbInFo) {
	dirpath, basepath, movname, ext := myPathSplit(p)
	if ext == ".txt" {
		fmt.Print("what the fuck a text file")
	} else {
		MSA := os.Getenv("MOVIEGOBS_SERVER_ADDRESS")
		MSP := os.Getenv("MOVIEGOBS_SERVER_PORT")
		MTPP := os.Getenv("MOVIEGOBS_THUMBNAIL_PIC_PATH")
		BP := "/" + url.QueryEscape(basepath)
		thumbpathtwo := MSA + ":" + MSP + MTPP + BP
		
		thumbpath := os.Getenv("MOVIEGOBS_THUMBNAIL_PIC_PATH") + "/" + basepath
		log.Printf("\n this is thumbpath %v \n", thumbpath)


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
			// log.Printf("file %s not exists", file)
		} else {
			log.Printf("file %s stat error: %v", thumbpath, err)
		}


		tid := UUID()
		ThumbINFO := ThumbInFo{ID: bson.NewObjectId(),
			MovName:   movname,   //ex mythumbnail
			BasePath:  basepath,  //ex mythumbnail.jpg
			DirPATH:   dirpath,   //ex /usr/share/moviegobs
			ThumbPath: thumbpath, //our static folder path
			ThumbID:   tid,
			ThumbPathTwo: thumbpathtwo,
		}
		fmt.Println(ThumbINFO)
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
	return thumbINFO
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

//FindPicPathsTV comment
// func FindPicPathsTV(mpath string, noartpicpath string) (result string) {
// 	// DirPath: dirpath,
// 	// BaseName: basename,
// 	// Ext: ext,
// 	// Season: season,
// 	// Episode: episode,
// 	// EpisodeName: episodename,
// 	_, _, _, Season, Episode, EpisodeName := myPathSplitTV(mpath)
// 	Tses := DBcon()
// 	defer Tses.Close()
// 	MTc := Tses.DB("moviegobs").C("tvbsthumb")
// 	b1 := bson.M{"season": Season, "episode": Episode, "episodename": EpisodeName}
// 	b2 := bson.M{"_Id": 0}
// 	var ThumbII []map[string]string
// 	err := MTc.Find(b1).Select(b2).All(&ThumbII)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Println(ThumbII)
// 	LenI := len(ThumbII)
// 	if LenI == 0 {
// 		NoArtList = append(NoArtList, mpath)
// 		result = noartpicpath
// 	} else {
// 		result = ThumbII[0]["thumbpath"]
// 		fmt.Printf("THIS IS THUMBI.THUMBPATH:  %s", ThumbII[0]["thumbpath"])
// 	}
// 	return
// }




//CreateTVShowsThumbnail comment
// func CreateTVShowsThumbnail(p string) string {
// 	_, fnn := path.Split(p)
// 	nfn1 := "./static/images/thumbnails/" + fnn
// 	//nfn1 := "./static/test/" + fnn
// 	// pic, err := imaging.Open(p)
// 	// if err != nil {
// 	// 	fmt.Printf("\n this is file Open error jpgthumb %v \n", p)
// 	// }

   
//     pic, err := imaging.Open(p)
//     if err != nil {
//         fmt.Printf("\n this is file Open error jpgthumb %v \n", p)
//     }

// 	cjImage := imaging.Thumbnail(pic, 100, 100, imaging.Lanczos)
//     err = imaging.Save(cjImage, nfn1)
//     if err != nil {
// 		fmt.Println(err)
// 	}
//     return nfn1[1:]
// }

