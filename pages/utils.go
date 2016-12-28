package pages

import "time"

func addStandardData(model map[string]interface{}, activeTab ...string) {
	model["current_year"] = time.Now().Year()
	if len(activeTab) != 0 {
		model["active_tab"] = activeTab[0]
	}
}
