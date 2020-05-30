package models

import (
	"coinproject/config"
	"hash/fnv"
	"log"
)

type Url struct {
	ID       int    `json:"id"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
	UrlId    string `json:"url_id"`
}

func InsertUrl(url string, surl string, urlId uint32) {
	longurl := url
	urlid := urlId
	shorturl := surl
	var err error
	db := config.GetDB()
	sqlStatement := `INSERT INTO tinyurl (long_url, short_url, url_id) VALUES (?,?,?)`
	_, err = db.Exec(sqlStatement, longurl, shorturl, urlid)

	if err != nil {
		log.Println("Query failed.....")
		log.Println(err.Error())
	}
}

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func GetOriginalUrl(urlID string) string {
	var longurl string
	db := config.GetDB()
	rows, err := db.Query("SELECT long_url FROM tinyurl where url_id = " + urlID + ";")
	if err != nil {
		log.Println("Query failed.....")
		log.Println(err.Error())
	}
	log.Println("Query pass.....")

	for rows.Next() {
		err := rows.Scan(&longurl)
		if err != nil {
			log.Fatal(err)
		}
	}
	return longurl
}

func SingleUrlDetail(sid string) []Url {
	db := config.GetDB()
	rows, err := db.Query("SELECT id, long_url, short_url, url_id FROM tinyurl where id = " + sid + ";")
	if err != nil {
		log.Println("Query failed.....")
		log.Println(err.Error())
		return nil
	}
	log.Println("Query pass.....")
	urldata := []Url{}

	var id int
	var longurl string
	var shorturl string
	var urlid string
	for rows.Next() {

		rows.Scan(&id, &longurl, &shorturl, &urlid)

		urldata = append(urldata, Url{ID: id, LongUrl: longurl, ShortUrl: shorturl, UrlId: urlid})
	}
	return urldata

}

func AllUrlDetails() []Url {
	db := config.GetDB()
	rows, err := db.Query(`SELECT id, long_url, short_url, url_id FROM tinyurl`)
	if err != nil {
		log.Println("Query failed.....")
		log.Println(err.Error())
		return nil
	}
	log.Println("Query pass.....")
	urldatas := []Url{}

	var id int
	var longurl string
	var shorturl string
	var urlid string
	for rows.Next() {

		rows.Scan(&id, &longurl, &shorturl, &urlid)

		urldatas = append(urldatas, Url{ID: id, LongUrl: longurl, ShortUrl: shorturl, UrlId: urlid})
	}
	return urldatas
}
