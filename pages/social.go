package pages

import (
	"database/sql"
	"github.com/gchaincl/dotsql"
	"gopkg.in/macaron.v1"
	"log"
)

type (
	socialAccount struct {
		id   string
		link string
	}
	SocialAccountsForm struct {
		Instagram string `form:"instagram"`
		Twitter   string `form:"twitter"`
		Facebook  string `form:"facebook"`
		Youtube   string `form:"youtube"`
		Behance   string `form:"behance"`
		Linkedin  string `form:"linkedin"`
	}
)

var allowedSocialAccounts = [...]string{"instagram", "twitter", "facebook", "youtube", "behance", "linkedin"}

func loadSocialAccounts(db *sql.DB, dot *dotsql.DotSql) ([]socialAccount, error) {
	rows, err := dot.Query(db, "select-social-accounts")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	accounts := make([]socialAccount, 0)
	for rows.Next() {
		var account socialAccount
		err = rows.Scan(&account.id, &account.link)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		accounts = append(accounts, account)
	}
	err = rows.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return accounts, nil
}

func getSocialAccounts(db *sql.DB, dot *dotsql.DotSql) ([]socialAccount, error) {
	accounts, err := loadSocialAccounts(db, dot)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(accounts))
	for i, account := range accounts {
		names[i] = account.id
	}
	for _, name := range allowedSocialAccounts {
		if !arrayContains(names, name) {
			account := new(socialAccount)
			account.id = name
			accounts = append(accounts, *account)
		}
	}
	return accounts, nil
}

func insertSocialAccount(db *sql.DB, dot *dotsql.DotSql, accounts []socialAccount, id, link string) error {
	update := false
	for _, account := range accounts {
		if account.id == id {
			update = true
			break
		}
	}
	var err error
	if update {
		_, err = dot.Exec(db, "update-social-account", link, id)
	} else {
		_, err = dot.Exec(db, "insert-social-account", id, link)
	}
	return err
}

// TODO one sql statement
func AdminSocialAccounts(ctx *macaron.Context, form SocialAccountsForm, db *sql.DB, dot *dotsql.DotSql) {
	accounts, err := loadSocialAccounts(db, dot)
	err = insertSocialAccount(db, dot, accounts, "instagram", form.Instagram)
	err = insertSocialAccount(db, dot, accounts, "twitter", form.Twitter)
	err = insertSocialAccount(db, dot, accounts, "facebook", form.Facebook)
	err = insertSocialAccount(db, dot, accounts, "youtube", form.Youtube)
	err = insertSocialAccount(db, dot, accounts, "behance", form.Behance)
	err = insertSocialAccount(db, dot, accounts, "linkedin", form.Linkedin)
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx.Redirect("/admin?alert=Update+social+accounts!#admin-social-accounts")
}
