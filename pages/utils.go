package pages

import "time"

func addStandardData(model map[string]interface{}) {
	model["current_year"] = time.Now().Year()
}
