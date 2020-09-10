package tvgobslib

import (
	"fmt"
	"math/rand"

	//"os"
	"path"
	"strconv"
	"strings"
	"time"

	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//TvShowInfoS is needed in main
type TvShowInfoS struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	FilePath      string        `bson:"filepath"`
	Catagory      string        `bson:"catagory"`
	MediaID       string        `bson:"mediaid"`
	Genre         string        `bson:"genre"`
	Season        string        `bson:"season"`
	Episode       string        `bson:"episode"`
	Title         string        `bson:"title"`
	Series        string        `bson:"series"`
	MovFSPath     string        `bson:"movfspath"`
	TVShowPicPath string        `bson:"tvshowpicpath"`
	// ThumbPath string `bson:"thumbpath"`
}

// func tvshowsDBcon() *mgo.Session {
// 	s, err := mgo.Dial(os.Getenv("TVGOBS_MONGODB_ADDRESS"))
// 	if err != nil {
// 		fmt.Println("this is dial err")
// 		panic(err)
// 	}
// 	return s
// }

func getTvShowInfo(apath string, tvshowpicPath string) (TvSI TvShowInfoS) {
	var filesystempath = "/media/pi/PiTB/media/" + apath[13:len(apath)]
	fmt.Println(filesystempath)
	switch {
	case strings.Contains(apath, "TVShows/TNG"):
		_, filename := path.Split(apath)
		var boo = len(filename) - 4
		TvSI = TvShowInfoS{ID: bson.NewObjectId(),
			FilePath:      apath,
			Catagory:      "TNG",
			MediaID:       tvshowsUUID(),
			Genre:         "TVShows",
			Season:        filename[15:17],
			Episode:       filename[18:20],
			Title:         filename[21:boo],
			Series:        filename[21:boo],
			MovFSPath:     filesystempath,
			TVShowPicPath: tvshowpicPath,
			// //ThumbPath: tvshowpicInfo,
		}
	case strings.Contains(apath, " STTV "):
		_, filename := path.Split(apath)
		var boo = len(filename) - 4
		TvSI = TvShowInfoS{ID: bson.NewObjectId(),
			FilePath:      apath,
			Catagory:      "STTV",
			MediaID:       tvshowsUUID(),
			Genre:         "TVShows",
			Season:        filename[16:18],
			Episode:       filename[19:21],
			Title:         filename[21:boo],
			Series:        "Star Trek",
			TVShowPicPath: tvshowpicPath,
			MovFSPath:     filesystempath,
			// //ThumbPath: tvshowpicInfo,
		}
	case strings.Contains(apath, "Orville"):
		_, filename := path.Split(apath)
		var boo = len(filename) - 4
		TvSI = TvShowInfoS{ID: bson.NewObjectId(),
			FilePath:      apath,
			Catagory:      "Orville",
			MediaID:       tvshowsUUID(),
			Genre:         "TVShows",
			Season:        filename[13:15],
			Episode:       filename[16:18],
			Title:         filename[19:boo],
			Series:        "Orville",
			TVShowPicPath: tvshowpicPath,
			MovFSPath:     filesystempath,
			//ThumbPath: tvshowpicInfo,
		}
	case strings.Contains(apath, "Voyager"):
		_, filename := path.Split(apath)
		var boo = len(filename) - 4
		TvSI = TvShowInfoS{ID: bson.NewObjectId(),
			FilePath:      apath,
			Catagory:      "Voyager",
			MediaID:       tvshowsUUID(),
			Genre:         "TVShows",
			Season:        filename[19:21],
			Episode:       filename[22:24],
			Title:         filename[24:boo],
			Series:        "Voyager",
			TVShowPicPath: tvshowpicPath,
			MovFSPath:     filesystempath,
			//ThumbPath: tvshowpicInfo,
		}
	case strings.Contains(apath, "Discovery"):
		_, filename := path.Split(apath)
		var boo = len(filename) - 4
		TvSI = TvShowInfoS{ID: bson.NewObjectId(),
			FilePath:      apath,
			Catagory:      "Discovery",
			MediaID:       tvshowsUUID(),
			Genre:         "TVShows",
			Season:        filename[21:23],
			Episode:       filename[24:26],
			Title:         filename[27:boo],
			Series:        "Discovery",
			TVShowPicPath: tvshowpicPath,
			MovFSPath:     filesystempath,
			//ThumbPath: tvshowpicInfo,
		}
	case strings.Contains(apath, "ENT"):
		_, filename := path.Split(apath)
		var boo = len(filename) - 4
		TvSI = TvShowInfoS{ID: bson.NewObjectId(),
			FilePath:      apath,
			Catagory:      "Enterprise",
			MediaID:       tvshowsUUID(),
			Genre:         "TVShows",
			Season:        filename[15:17],
			Episode:       filename[18:20],
			Title:         filename[20:boo],
			Series:        "Enterprise",
			TVShowPicPath: tvshowpicPath,
			MovFSPath:     filesystempath,
			//ThumbPath: tvshowpicInfo,
		}
	case strings.Contains(apath, "The Last Ship"):
		_, filename := path.Split(apath)
		var boo = len(filename) - 4
		TvSI = TvShowInfoS{ID: bson.NewObjectId(),
			FilePath:      apath,
			Catagory:      "LastShip",
			MediaID:       tvshowsUUID(),
			Genre:         "TVShows",
			Season:        filename[15:17],
			Episode:       filename[18:20],
			Title:         filename[21:boo],
			Series:        "LastShip",
			TVShowPicPath: tvshowpicPath,
			MovFSPath:     filesystempath,
			//ThumbPath: tvshowpicInfo,
		}
	case strings.Contains(apath, "Lost In Space"):
		_, filename := path.Split(apath)
		var boo = len(filename) - 4
		TvSI = TvShowInfoS{ID: bson.NewObjectId(),
			FilePath:      apath,
			Catagory:      "Lost In Space",
			MediaID:       tvshowsUUID(),
			Genre:         "TVShows",
			Season:        filename[15:17],
			Episode:       filename[18:20],
			Title:         filename[21:boo],
			Series:        "Lost In Space",
			TVShowPicPath: tvshowpicPath,
			MovFSPath:     filesystempath,
			//ThumbPath: tvshowpicInfo,
		}
	}
	return
}

func tvshowsUUID() (UUID string) {
	aseed := time.Now()
	aSeed := aseed.UnixNano()
	rand.Seed(aSeed)
	u := rand.Int63n(aSeed)
	p := strconv.FormatInt(u, 10)
	return p
}
