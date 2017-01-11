package pages

import (
	"gopkg.in/macaron.v1"
	"fmt"
	"log"
	"github.com/go-macaron/session"
)

type (
	LoginForm struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}

	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Salt     string `json:"salt"`
	}
)

func Login(ctx *macaron.Context) {
	addStandardData(ctx.Data)
	ctx.HTMLSet(200, "base", "login")
}

func isValidPassword(password string, u *User) (bool, error) {
	body, err := sha512Hash(u.Salt + password)
	if err != nil {
		return false, err
	}
	return fmt.Sprintf("%x", body) == u.Password, nil
}

func LoginSubmit(ctx *macaron.Context, sess session.Store, form LoginForm, users []User) {
	authenticated := sess.Get("authenticated")
	if authenticated != nil && authenticated.(bool) {
		ctx.Redirect("/?alert=You+are+already+logged+in!")
		return
	}
	var u *User
	for _, usr := range users {
		if usr.Username == usr.Username {
			u = &usr
			break
		}
	}
	if u == nil {
		ctx.Redirect("/login?alert=Invalid username or password!")
		return
	}
	var valid bool
	var err error
	if valid, err = isValidPassword(form.Password, u); err != nil {
		log.Fatal(err)
		return
	}
	if !valid {
		ctx.Redirect("/login?alert=Invalid username or password!")
		return
	}
	sess.Set("authenticated", true)
	ctx.Redirect("/admin?alert=Logged in!")
}

func Logout(ctx *macaron.Context, sess session.Store) {
	authenticated := sess.Get("authenticated")
	if authenticated == nil || !authenticated.(bool) {
		ctx.Redirect("/?alert=You+are+not+logged+in!")
		return
	}
	sess.Delete("authenticated")
	ctx.Redirect("/?alert=Logged out!")
}
