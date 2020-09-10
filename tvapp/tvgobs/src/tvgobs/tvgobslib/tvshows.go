package tvgolib

import (
	"fmt"
	"time"
	"math/rand"
	"path"
	"strings"
	"strconv"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func tvshowsUUID() (UUID string) {
	aseed := time.Now()
	aSeed := aseed.UnixNano()
	rand.Seed(aSeed)
	u := rand.Int63n(aSeed)
	UUID = strconv.FormatInt(u, 10)
	return
}

// TVShowInfoS is needed because I want it
type TVShowInfoS struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
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
	TvFSPath string `bson:"tvfspath"`
}

func getTvShowInfo(apath string, tvshowpicPath string) (TvSI TVShowInfoS) {
	_, filename := path.Split(apath)
	// /root/fsData/TVShows/Enterprise/S1/filename.mp4
	fspath := apath[19:]
	boo := len(filename) - 4
	TvSI.ID = bson.NewObjectId()
	TvSI.FilePath = apath
	TvSI.MediaID = tvshowsUUID()
	TvSI.Genre = "TVShows"
	TvSI.TVShowPicPath = tvshowpicPath
	TvSI.TvFSPath = fspath
	switch {
		case strings.Contains(apath, "TVShows/TNG"):
			TvSI.Catagory = "TNG"
			TvSI.Season = filename[15:17]
			TvSI.Episode = filename[18:20]
			TvSI.Title = filename[21:boo]
			TvSI.Series = filename[21:boo]
		case strings.Contains(apath, " STTV "):
			TvSI.Catagory = "STTV"
			TvSI.Season = filename[16:18]
			TvSI.Episode = filename[19:21]
			TvSI.Title = filename[21:boo]
			TvSI.Series = "Star Trek"
		case strings.Contains(apath, "The Orville"):
			TvSI.Catagory = "The Orville"
			TvSI.Season = filename[13:15]
			TvSI.Episode = filename[16:18]
			TvSI.Title = filename[19:boo]
			TvSI.Series = "The Orville"
		case strings.Contains(apath, "Voyager"):
			TvSI.Catagory = "Voyager"
			TvSI.Season = filename[19:21]
			TvSI.Episode = filename[22:24]
			TvSI.Title = filename[24:boo]
			TvSI.Series = "Voyager"
		case strings.Contains(apath, "Discovery"):
			TvSI.Catagory = "Discovery"
			TvSI.Season = filename[21:23]
			TvSI.Episode = filename[24:26]
			TvSI.Title = filename[27:boo]
			TvSI.Series = "Discovery"
		case strings.Contains(apath, "ENT"):
			TvSI.Catagory = "Enterprise"
			TvSI.Season = filename[15:17]
			TvSI.Episode = filename[18:20]
			TvSI.Title = filename[20:boo]
			TvSI.Series = "Enterprise"
		case strings.Contains(apath, "The Last Ship"):
			TvSI.Catagory = "The Last Ship"
			TvSI.Season = filename[15:17]
			TvSI.Episode = filename[18:20]
			TvSI.Title = filename[21:boo]
			TvSI.Series = "The Last Ship"
	}
	fmt.Printf("\n THIS IS TVI FROM TVSHOWS \n %s \n", TvSI)
	return
}

