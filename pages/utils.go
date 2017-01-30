package pages

import (
	"crypto/sha512"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"time"
	"github.com/gchaincl/dotsql"
	"database/sql"
)

const image_file_path = "static/img/"

func Init() {
	body, err := ioutil.ReadFile(about_path)
	if err != nil {
		log.Fatal(err)
		return
	}
	about = string(body)
	body, err = ioutil.ReadFile(contact_path)
	if err != nil {
		log.Fatal(err)
		return
	}
	contact = string(body)
}

func addStandardData(model map[string]interface{}, db *sql.DB, dot *dotsql.DotSql, activeTab ...string) {
	model["current_year"] = time.Now().Year()
	var err error
	model["social_accounts"], err = loadSocialAccounts(db, dot)
	if err != nil {
		log.Fatalln(err)
	}
	if len(activeTab) != 0 {
		model["active_tab"] = activeTab[0]
	}
}

func sha512Hash(input string) ([]byte, error) {
	h512 := sha512.New()
	_, err := io.WriteString(h512, input)
	if err != nil {
		return nil, err
	}
	body := h512.Sum(nil)
	return body, nil
}

func arrayContains(array []string, item string) bool {
	for _, i := range array {
		if i == item {
			return true
		}
	}
	return false
}

func saveImage(header *multipart.FileHeader) (*string, error) {
	name := image_file_path + header.Filename
	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(name, body, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return &name, nil
}
