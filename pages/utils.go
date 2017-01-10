package pages

import (
	"io/ioutil"
	"log"
	"time"
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
