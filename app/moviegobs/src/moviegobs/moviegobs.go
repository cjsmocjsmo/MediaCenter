///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// LICENSE: GNU General Public License, version 2 (GPLv2)
// Copyright 2016, Charlie J. Smotherman
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License v2
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	moviegolib "moviegobs/moviegobslib"
	"net/http"
	"net/url"
	"os"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

func dBcon() *mgo.Session {
	s, err := mgo.Dial(os.Getenv("MOVIEGOBS_MONGODB_ADDRESS"))
	if err != nil {
		log.Println("Session creation dial error")
		log.Println(err)
	}
	log.Println("Session Connection to db established")
	return s
}

func showMovieGo(w http.ResponseWriter, r *http.Request) {
	tmppath := os.Getenv("MOVIEGOBS_TEMPLATE_ADDRESS")
	tmpl := template.Must(template.ParseFiles(tmppath))
	tmpl.Execute(w, tmpl)
}

func intActionHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var ActionMedia []map[string]string
	b1 := bson.M{"catagory": "Action"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&ActionMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(ActionMedia)
}

func intCartoonsHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	b1 := bson.M{"catagory": "Cartoons"}
	b2 := bson.M{"_id": 0}
	var CartoonMedia []map[string]string
	err := MTc.Find(b1).Select(b2).All(&CartoonMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(CartoonMedia)
}

func intComedyHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var ComedyMedia []map[string]string
	b1 := bson.M{"catagory": "Comedy"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&ComedyMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(ComedyMedia)
}

func intDramaHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var DramaMedia []map[string]string
	b1 := bson.M{"catagory": "Drama"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&DramaMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(DramaMedia)
}

func intGodzillaHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var GodzillaMedia []map[string]string
	b1 := bson.M{"catagory": "Godzilla"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&GodzillaMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(GodzillaMedia)
}

func intIndianaJonesHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTrc := ses.DB("moviegobs").C("moviegobs")
	var IndianaJonesMedia []map[string]string
	b1 := bson.M{"catagory": "IndianaJones"}
	b2 := bson.M{"_id": 0}
	err := MTrc.Find(b1).Select(b2).All(&IndianaJonesMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(IndianaJonesMedia)
}

func intJohnWayneHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var JohnWayneMedia []map[string]string
	b1 := bson.M{"catagory": "JohnWayne"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JohnWayneMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(JohnWayneMedia)
}

func intJurassicParkHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var JurassicParkMedia []map[string]string
	b1 := bson.M{"catagory": "JurassicPark"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&JurassicParkMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(JurassicParkMedia)
}

func intKingsManHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var KingsmanMedia []map[string]string
	b1 := bson.M{"catagory": "Kingsman"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&KingsmanMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(KingsmanMedia)
}

func intHarryPotterHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var HarryPotterMedia []map[string]string
	b1 := bson.M{"catagory": "HarryPotter"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&HarryPotterMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(HarryPotterMedia)
}

func intMenInBlackHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var MenInBlackMedia []map[string]string
	b1 := bson.M{"catagory": "MenInBlack"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&MenInBlackMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(MenInBlackMedia)
}

func intMiscHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var MiscMedia []map[string]string
	b1 := bson.M{"catagory": "Misc"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&MiscMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(MiscMedia)
}

func intSciFiHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var SciFiMedia []map[string]string
	b1 := bson.M{"catagory": "SciFi"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&SciFiMedia)
	if err != nil {
		log.Println("this is a database lookup error for the InitSciFiHandler function")
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(SciFiMedia)
}

func intStarTrekHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var StarTrekMedia []map[string]string
	b1 := bson.M{"catagory": "StarTrek"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&StarTrekMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(StarTrekMedia)
}

func intStarWarsHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var StarWarsMedia []map[string]string
	b1 := bson.M{"catagory": "StarWars"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&StarWarsMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(StarWarsMedia)
}

func intSuperHerosHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var SuperHerosMedia []map[string]string
	b1 := bson.M{"catagory": "SuperHeros"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&SuperHerosMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(SuperHerosMedia)
}

func intTremorsHandler(w http.ResponseWriter, r *http.Request) {
	ses := dBcon()
	defer ses.Close()
	MTc := ses.DB("moviegobs").C("moviegobs")
	var TremorsMedia []map[string]string
	b1 := bson.M{"catagory": "Tremors"}
	b2 := bson.M{"_id": 0}
	err := MTc.Find(b1).Select(b2).All(&TremorsMedia)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(TremorsMedia)
}

func playMediaHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	mf := m["movie"][0]
	ses := dBcon()
	defer ses.Close()
	var MediaInfo map[string]string
	MTc := ses.DB("moviegobs").C("moviegobs")
	b1 := bson.M{"movfspath": mf}
	b2 := bson.M{"_id": 0}
	err = MTc.Find(b1).Select(b2).One(&MediaInfo)
	if err != nil {
		fmt.Println(err)
	}
	omxAddr := os.Getenv("MOVIEGOBS_OMXPLAYER_ADDRESS")
	u, _ = url.Parse(omxAddr)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("medPath", omxAddr)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	Abody := string(body)
	fmt.Println(Abody)
	fmt.Printf("this is mediainfo sent to browser: %s", MediaInfo)
	log.Printf("this is mediainfo filepath: %s", MediaInfo)
	log.Println("Sending info to omxplayer server")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&MediaInfo)
}

//UpdateHandler updates the db with newly added music
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	j, err := os.OpenFile(os.Getenv("MOVIEGOBS_LOG"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer j.Close()
	wrtupdate := io.MultiWriter(os.Stdout, j)
	log.SetOutput(wrtupdate)

	log.Println("UpdateHandler is starting")
	val, _ := os.LookupEnv("MOVIEGOBS_SETUP")
	fmt.Println("this is moviegosetup val")
	fmt.Println(val)
	var exitstatus int
	if val == "0" {
		fmt.Println("Ok is Ok")
		log.Println("MOVIEGOBS_SETUP environment variable is set, starting MOVIEGO")
		exitstatus = 0

	} else {
		fmt.Println("not OK")
		log.Println("MOVIEGOBS_SETUP environment variable is not set, starting SETUP")
		exitstatus = moviegolib.SetUp()
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exitstatus)
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/static").Subrouter()
	r.HandleFunc("/moviego", showMovieGo)
	r.HandleFunc("/intAction", intActionHandler)
	r.HandleFunc("/intCartoons", intCartoonsHandler)
	r.HandleFunc("/intComedy", intComedyHandler)
	r.HandleFunc("/intDrama", intDramaHandler)
	r.HandleFunc("/intGodzilla", intGodzillaHandler)
	r.HandleFunc("/intHarryPotter", intHarryPotterHandler)
	r.HandleFunc("/intIndianaJones", intIndianaJonesHandler)
	r.HandleFunc("/intMenInBlack", intMenInBlackHandler)
	r.HandleFunc("/intMisc", intMiscHandler)
	r.HandleFunc("/intJohnWayne", intJohnWayneHandler)
	r.HandleFunc("/intJurassicPark", intJurassicParkHandler)
	r.HandleFunc("/intKingsMan", intKingsManHandler)
	r.HandleFunc("/intSciFi", intSciFiHandler)
	r.HandleFunc("/intStarTrek", intStarTrekHandler)
	r.HandleFunc("/intStarWars", intStarWarsHandler)
	r.HandleFunc("/intSuperHeros", intSuperHerosHandler)
	r.HandleFunc("/intTremors", intTremorsHandler)
	r.HandleFunc("/playMedia", playMediaHandler)
	r.HandleFunc("/Update", UpdateHandler)
	s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/media/"))))
	http.ListenAndServe(":8888", (r))
}
