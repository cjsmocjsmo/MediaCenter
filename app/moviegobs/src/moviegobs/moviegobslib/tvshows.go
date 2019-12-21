package moviegolib

import (
	"fmt"
	"time"
	"math/rand"
	"path"
	"strings"
	"strconv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type tvShowInfoS struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	FilePath string `bson:"filepath"`
	Catagory string `bson:"catagory"`
	MediaID string `bson:"MediaID"`
	Genre string `bson:"genre"`
	Season string `bson:"season"`
	Episode string `bson:"episode"`
	Title string `bson:"title"`
	Series string `bson:"series"`
	TVShowPicPath string `bson:"tvshowpicpath"`
	ThumbPath string `bson:"thumbpath"`
}

func tvshowsDBcon() *mgo.Session {
	s, err := mgo.Dial("127.0.0.1")
//	s, err := mgo.Dial("mongodb://192.168.1.101:27017")
	if err != nil {
		fmt.Println("this is dial err")
		panic(err)
	}
	return s
}

func getTvShowInfo(apath string, tvshowpicPath string, tvshowpicInfo string) (TvSI tvShowInfoS) {
	switch {
		case strings.Contains(apath, "TVShows/TNG"):
			_, filename := path.Split(apath)
			var boo = len(filename) - 4
			TvSI = tvShowInfoS{Id: bson.NewObjectId(),
				FilePath: apath,
				Catagory: "TNG",
				MediaID: tvshowsUUID(),
				Genre: "TVShows",
				Season: filename[15:17],
				Episode: filename[18:20],
				Title: filename[21:boo],
				Series: filename[21:boo],
				TVShowPicPath: tvshowpicPath,
				ThumbPath: tvshowpicInfo,
			}
		case strings.Contains(apath, " STTV "):
			_, filename := path.Split(apath)
			var boo = len(filename) - 4
			TvSI = tvShowInfoS{Id: bson.NewObjectId(),
				FilePath: apath,
				Catagory: "STTV",
				MediaID: tvshowsUUID(),
				Genre: "TVShows",
				Season: filename[16:18],
				Episode: filename[19:21],
				Title: filename[21:boo],
				Series: "Star Trek",
				TVShowPicPath: tvshowpicPath,
				ThumbPath: tvshowpicInfo,
			}
		case strings.Contains(apath, "The Orville"):
			_, filename := path.Split(apath)
			var boo = len(filename) - 4
			TvSI = tvShowInfoS{Id: bson.NewObjectId(),
				FilePath: apath,
				Catagory: "The Orville",
				MediaID: tvshowsUUID(),
				Genre: "TVShows",
				Season: filename[13:15],
				Episode: filename[16:18],
				Title: filename[19:boo],
				Series: "The Orville",
				TVShowPicPath: tvshowpicPath,
				ThumbPath: tvshowpicInfo,
			}
		case strings.Contains(apath, "Voyager"):
			_, filename := path.Split(apath)
			var boo = len(filename) - 4
			TvSI = tvShowInfoS{Id: bson.NewObjectId(),
				FilePath: apath,
				Catagory: "Voyager",
				MediaID: tvshowsUUID(),
				Genre: "TVShows",
				Season: filename[19:21],
				Episode: filename[22:24],
				Title: filename[24:boo],
				Series: "Voyager",
				TVShowPicPath: tvshowpicPath,
				ThumbPath: tvshowpicInfo,
			}
		case strings.Contains(apath, "Discovery"):
			_, filename := path.Split(apath)
			var boo = len(filename) - 4
			TvSI = tvShowInfoS{Id: bson.NewObjectId(),
				FilePath: apath,
				Catagory: "Discovery",
				MediaID: tvshowsUUID(),
				Genre: "TVShows",
				Season: filename[21:23],
				Episode: filename[24:26],
				Title: filename[27:boo],
				Series: "Discovery",
				TVShowPicPath: tvshowpicPath,
				ThumbPath: tvshowpicInfo,
			}
		case strings.Contains(apath, "ENT"):
			_, filename := path.Split(apath)
			var boo = len(filename) - 4
			TvSI = tvShowInfoS{Id: bson.NewObjectId(),
				FilePath: apath,
				Catagory: "Enterprise",
				MediaID: tvshowsUUID(),
				Genre: "TVShows",
				Season: filename[15:17],
				Episode: filename[18:20],
				Title: filename[20:boo],
				Series: "Enterprise",
				TVShowPicPath: tvshowpicPath,
				ThumbPath: tvshowpicInfo,
			}
		case strings.Contains(apath, "The Last Ship"):
			_, filename := path.Split(apath)
			var boo = len(filename) - 4
			TvSI = tvShowInfoS{Id: bson.NewObjectId(),
				FilePath: apath,
				Catagory: "The Last Ship",
				MediaID: tvshowsUUID(),
				Genre: "TVShows",
				Season: filename[15:17],
				Episode: filename[18:20],
				Title: filename[21:boo],
				Series: "The Last Ship",
				TVShowPicPath: tvshowpicPath,
				ThumbPath: tvshowpicInfo,
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