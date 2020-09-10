package tvgolib

import (
	"fmt"
	// its blank because I want it
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
	"github.com/disintegration/imaging"
	"github.com/globalsign/mgo/bson"
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

//ThumbInFo struct exported to setup
type ThumbInFo struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	MovName   string        `bson:"movname"`
	BasePath  string        `bson:"baspath"`
	DirPATH   string        `bson:"dirpath"`
	ThumbPath string        `bson:"thumbpath"`
	ThumbID   string        `bson:"thumbid"`
}

//CreateMoviesThumbnail exported to setup
func CreateMoviesThumbnail(p string) (thumbINFO ThumbInFo) {
	dirpath, basepath, movname, ext := myPathSplit(p)
	if ext == ".txt" {
		fmt.Print("what the fuck a text file")
	} else {

		thumbpath := os.Getenv("MOVIEGOBS_THUMBNAIL_PIC_PATH") + "/" + basepath

		pic, err := imaging.Open(p)
		if err != nil {
			log.Printf("\n this is file Open error jpgthumb %v \n", p)
		}
		thumbImage := imaging.Resize(pic, 0, 400, imaging.Lanczos)
		err = imaging.Save(thumbImage, thumbpath)
		if err != nil {
			fmt.Println(err)
		}

		tid := UUID()
		ThumbINFO := ThumbInFo{ID: bson.NewObjectId(),
			MovName:   movname,   //ex mythumbnail
			BasePath:  basepath,  //ex mythumbnail.jpg
			DirPATH:   dirpath,   //ex /usr/share/moviegobs
			ThumbPath: thumbpath, //our static folder path
			ThumbID:   tid,
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
		log.Println(&thumbINFO)
	}
	return thumbINFO
}

//NoArtList exported to setup
var NoArtList []string

// FindPicPaths is exported because I want it
func FindPicPaths(mpath string, noartpicpath string) (result string) {
	_, _, movename, _ := myPathSplit(mpath)
	fmt.Printf("this is movename from findpicpaths %s", movename)
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
	fmt.Printf("this is result %s", result)
	return
}

/*
func CreateTVShowsThumbnail(p string) string {
	_, fnn := path.Split(p)
	nfn1 := "./static/images/thumbnails/" + fnn
	nfn1 := "./static/test/" + fnn
	pic, err := imaging.Open(p)
	if err != nil {
		fmt.Printf("\n this is file Open error jpgthumb %v \n", p)
	}

   if strings.Contains(p, "TVShows") {
        pic, err := imaging.Open(p)
        if err != nil {
            fmt.Printf("\n this is file Open error jpgthumb %v \n", p)
        }

	    cjImage := imaging.Thumbnail(pic, 100, 100, imaging.Lanczos)
        err = imaging.Save(cjImage, nfn1)
        if err != nil {
		  fmt.Println(err)
	   }
    }
    return nfn1[1:]
}

*/
