package pages

import (
	"io/ioutil"
	"log"
	"time"
	"io"
	"crypto/sha512"
)

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

func addStandardData(model map[string]interface{}, activeTab ...string) {
	model["current_year"] = time.Now().Year()
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