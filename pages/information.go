package pages

import (
	"database/sql"
	"github.com/gchaincl/dotsql"
	"gopkg.in/macaron.v1"
	"io/ioutil"
	"log"
	"os"
)

const about_path, contact_path = "about.txt", "contact.txt"

type InformationForm struct {
	About   string `form:"about"`
	Contact string `form:"contact"`
}

var about, contact string

func Information(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	addStandardData(ctx.Data, "about")
	ctx.Data["about_text"] = about
	ctx.Data["contact_text"] = contact
	var err error
	ctx.Data["social_accounts"], err = loadSocialAccounts(db, dot)
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx.HTMLSet(200, "base", "about")
}

func AdminInformation(ctx *macaron.Context, form InformationForm) {
	err := ioutil.WriteFile(about_path, []byte(form.About), os.ModeAppend)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = ioutil.WriteFile(contact_path, []byte(form.Contact), os.ModeAppend)
	if err != nil {
		log.Fatal(err)
		return
	}
	about, contact = form.About, form.Contact
	ctx.Redirect("/admin?alert=Edited information!#admin-information")
}
