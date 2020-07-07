package moviegolib

import (
	"fmt"
	// "io"
	"log"
	// "os"
	// "path/filepath"
	// "time"
	// "strings"
	// "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func nameInDbCheck(fn string) (result bool) {
	sess := DBcon()
	defer sess.Close()
	MTc := sess.DB("moviegobs").C("moviegobs")
	b1 := bson.M{"filepath": fn}
	b2 := bson.M{"_Id": 0}
	var PMedia []map[string]string
	err := MTc.Find(b1).Select(b2).All(&PMedia)
	if err != nil {
		log.Println(err)
	}
	num := len(PMedia)
	if num != 0 {
		result = true
	} else {
		result = false
	}
	return
}

//UpdateMain comment
func UpdateMain(filename string) (finished bool) {
	namecheck := nameInDbCheck(filename)
	if namecheck {
		fmt.Println("Movie already in DB")
	} else {
		picinfo := CreateMoviesThumbnail(filename)
		GetMovieInfo(filename, picinfo.ThumbPath)
	}
	return true
}