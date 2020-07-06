package moviegolib

import (
	"fmt"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func getmovName(movname string) string {
	var movName string
	_, fname := path.Split(movname)
	if strings.Contains(fname, "(") {
		fi := strings.Index(fname, "(")
		fdex := fi - 1
		movName = fname[:fdex]
	} else {
		ddex := len(fname) - 11
		movName = fname[ddex:]
	}
	return movName
}

func getMovieYear(apath string) (movyr string) {
	_, filename := path.Split(apath)
	fi := strings.Index(filename, "(")
	fdex := fi + 1
	ldex := strings.LastIndex(filename, ")")
	movyr = filename[fdex:ldex]
	return
}

//MOVI is exported because I want it so
type MOVI struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	DirPath        string        `bson:"dirpath"`
	Filepath       string        `bson:"filepath"`
	MediaID        string        `bson:"mediaid"`
	Movname        string        `bson:"movname"`
	Genre          string        `bson:"genre"`
	Catagory       string        `bson:"catagory"`
	MovFSPath      string        `bson:"movfspath"`
	ThumbPath      string        `bson:"thumbpath"`
	CarosThumbPath string        `bson:"carosthumbpath"`
	MovYear        string        `bson:"movyear"`
}

// GetMovieInfo comment
func GetMovieInfo(apath string, movpicInfo string) (MovInfo MOVI) {
	var filesystempath = "/media/pi/PiTB/media/" + apath[13:len(apath)]
	fmt.Printf("THIS IS FILESYSTEMPATH  %s", filesystempath)
	switch {
	case strings.Contains(apath, "SciFi"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "SciFi",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "Cartoons"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "Cartoons",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "Godzilla"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "Godzilla",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "Kingsman"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "Kingsman",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "StarTrek") && !strings.Contains(apath, " STTV "):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "StarTrek",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "StarWars"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "StarWars",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "SuperHeros"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "SuperHeros",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "IndianaJones"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "IndianaJones",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "Action"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "Action",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "Comedy"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "Comedy",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "Drama"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "Drama",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "JurassicPark"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "JurassicPark",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "JohnWayne"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "JohnWayne",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "JohnWick"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "JohnWick",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "MenInBlack"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "MenInBlack",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "HarryPotter"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "HarryPotter",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "Tremors"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "Tremors",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "Misc"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "Misc",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	case strings.Contains(apath, "Pirates"):
		dirp, _ := path.Split(apath)
		MovInfo = MOVI{ID: bson.NewObjectId(),
			DirPath:   dirp,
			Filepath:  apath,
			MediaID:   moviesUUID(),
			Movname:   getmovName(apath),
			Genre:     "movies",
			Catagory:  "Pirates",
			MovFSPath: filesystempath,
			ThumbPath: movpicInfo,
			MovYear: getMovieYear(apath),
		}
	}
	return
}

func moviesUUID() (UUID string) {
	aseed := time.Now()
	aSeed := aseed.UnixNano()
	rand.Seed(aSeed)
	u := rand.Int63n(aSeed)
	p := strconv.FormatInt(u, 10)
	return p
}

