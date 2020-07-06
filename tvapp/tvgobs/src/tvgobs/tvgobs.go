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
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	tvgolib "tvgobs/tvgobslib"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

//DBcon is exported  to main
func DBcon() *mgo.Session {
	s, err := mgo.Dial(os.Getenv("TVGOBS_MONGODB_ADDRESS"))
	if err != nil {
		fmt.Println("Session creation dial error")
		fmt.Println(err)
	}
	fmt.Println("Session Connection to db established")
	return s
}

func intSTTVHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var STTVMedia []map[string]string
	b1 := bson.M{"catagory": "STTV", "season": s1}
	b2 := bson.M{"_id": 0}
	fmt.Printf("this is b1 %s", b1)
	errG := MTyc.Find(b1).Select(b2).All(&STTVMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(&STTVMedia)
}

func intTNGHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var nextGenerationMedia []map[string]string
	b1 := bson.M{"catagory": "TNG", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&nextGenerationMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(&nextGenerationMedia)
}

func intEnterpriseHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var enterpriseMedia []map[string]string
	b1 := bson.M{"catagory": "Enterprise", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&enterpriseMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(&enterpriseMedia)
}

func intDiscoveryHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var discoveryMedia []map[string]string
	b1 := bson.M{"catagory": "Discovery", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&discoveryMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(&discoveryMedia)
}

func intVoyagerHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var voyagerMedia []map[string]string
	b1 := bson.M{"catagory": "Voyager", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&voyagerMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(&voyagerMedia)

}

func intLastShipHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTyc := ses.DB("tvgobs").C("tvgobs")
	var lastshipMedia []map[string]string
	b1 := bson.M{"catagory": "LastShip", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTyc.Find(b1).Select(b2).All(&lastshipMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(&lastshipMedia)
}

func intOrvilleHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	m, eff := url.ParseQuery(u.RawQuery)
	if eff != nil {
		fmt.Println(eff)
	}
	s1 := m["season"][0]
	ses := DBcon()
	defer ses.Close()
	MTtc := ses.DB("tvgobs").C("tvgobs")
	var OrvilleMedia []map[string]string
	b1 := bson.M{"catagory": "Orville", "season": s1}
	b2 := bson.M{"_id": 0}
	errG := MTtc.Find(b1).Select(b2).All(&OrvilleMedia)
	if errG != nil {
		fmt.Println(errG)
	}
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=370739520, public")
	json.NewEncoder(w).Encode(&OrvilleMedia)
}

func playMediaHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	m, _ := url.ParseQuery(u.RawQuery)
	mf := m["tvshow"][0]
	ses := DBcon()
	defer ses.Close()
	var MediaInfo map[string]string
	MTc := ses.DB("tvgobs").C("tvgobs")
	b1 := bson.M{"movfspath": mf}
	b2 := bson.M{"_id": 0}
	err = MTc.Find(b1).Select(b2).One(&MediaInfo)
	if err != nil {
		fmt.Println(err)
	}
	omxAddr := os.Getenv("TVGOBS_OMXPLAYER_ADDRESS")
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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&MediaInfo)
}

func playMediaReactHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	m, _ := url.ParseQuery(u.RawQuery)
	mf := m["tvshow"][0]

	ses := DBcon()
	defer ses.Close()
	var MediaInfo map[string]string
	MTc := ses.DB("tvgobs").C("tvgobs")
	b1 := bson.M{"movfspath": mf}
	b2 := bson.M{"_id": 0}
	err = MTc.Find(b1).Select(b2).One(&MediaInfo)
	if err != nil {
		log.Printf("this is playmedia err %s", err)
	}


	omxAddr := os.Getenv("TVGOBS_OMXPLAYER_ADDRESS")

	u, _ = url.Parse(omxAddr)
	q, _ := url.ParseQuery(u.RawQuery)
	// q.Set("medPath", omxAddr)
	q.Set("medPath", mf)
	u.RawQuery = q.Encode()
	log.Printf("this is u.string %s", u.String())
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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&MediaInfo)
}

//UpdateHandler updates the db with newly added music
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	val, _ := os.LookupEnv("TVGOBS_SETUP")
	var exitstatus int
	if val == "0" {
		fmt.Println("Ok is Ok")
		fmt.Println("TVGOBS_SETUP environment variable is set, starting TVGO")
		exitstatus = 0
	} else {
		fmt.Println("not OK")
		fmt.Println("TVGOBS_SETUP environment variable is not set, starting SETUP")
		exitstatus = tvgolib.SetUp()
	}
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exitstatus)
}

//DropTVDataBaseHandler is crap
func DropTVDataBaseHandler(w http.ResponseWriter, r *http.Request) {
	sess := DBcon()
	err := sess.DB("tvgobs").DropDatabase()
	if err != nil {
		fmt.Println(err)
	}
}

//TVDBCountHandler bla bla
func TVDBCountHandler(w http.ResponseWriter, r *http.Request) {
	ses := DBcon()
	foo, err := ses.DB("tvgobs").C("tvgobs").Count()
	if err != nil {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(-0)
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foo)
}

//TVSetupStatusHandler bla bla
func TVSetupStatusHandler(w http.ResponseWriter, r *http.Request) {
	status := os.Getenv("TVGOBS_SETUP")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/static").Subrouter()
	r.HandleFunc("/intSTTV", intSTTVHandler)
	r.HandleFunc("/intTNG", intTNGHandler)
	r.HandleFunc("/intEnterprise", intEnterpriseHandler)
	r.HandleFunc("/intDiscovery", intDiscoveryHandler)
	r.HandleFunc("/intLastShip", intLastShipHandler)
	r.HandleFunc("/intOrville", intOrvilleHandler)
	r.HandleFunc("/intVoyager", intVoyagerHandler)
	r.HandleFunc("/playMedia", playMediaHandler)
	r.HandleFunc("/playMediaReact", playMediaReactHandler)
	r.HandleFunc("/Update", UpdateHandler)
	r.HandleFunc("/DropTVDataBase", DropTVDataBaseHandler)
	r.HandleFunc("/TVDBCount", TVDBCountHandler)

	r.HandleFunc("/TVSetupStatus", TVSetupStatusHandler)
	s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/media/"))))
	http.ListenAndServe(":9999", (r))
}
