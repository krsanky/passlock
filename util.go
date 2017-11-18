package passlock

import "github.com/dchest/uniuri"

func PasslockPassword() string {
	return uniuri.New()
}

/* view.go
package passlock

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gowww/router"
	"github.com/justinas/nosurf"
	"github.com/krsanky/lg"

	"oldcode.org/webplay/account"
	"oldcode.org/webplay/view"
)

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := view.NewCtx()
	ctx["token"] = nosurf.Token(r)

	now := time.Now()
	ctx["year"] = strconv.Itoa(now.Year())
	ctx["month"] = strconv.Itoa(int(now.Month()))
	ctx["day"] = strconv.Itoa(now.Day())
	ctx["hour"] = strconv.Itoa(now.Hour())
	ctx["minute"] = 0

	if "POST" == r.Method {
		lg.Log.Printf("POST to passlock")
		_, errs := validateAndCreate(r)
		for _, e := range errs {
			lg.Log.Printf("err:%s", e.Error())
		}
	}

	u, ok := account.UserFromContext(r.Context())
	if ok && (u != nil) {
		pls, err := GetAll(u)
		if err == nil {
			ctx["passlocks"] = pls
		}
	}

	view.Render(w, r, "passlock/index.html", ctx)
}

func checkErr(err error, errs []error) []error {
	if err != nil {
		errs = append(errs, err)
	}
	return errs
}
func validateAndCreate(r *http.Request) (*Passlock, []error) {
	var errs []error

	u, ok := account.UserFromContext(r.Context())
	if !ok || (u == nil) {
		errs = append(errs, errors.New("hey!"))
		return nil, errs
	}

	name := r.FormValue("name")
	lg.Log.Printf("name:%s", name)

	year, err := strconv.Atoi(r.FormValue("year"))
	errs = checkErr(err, errs)
	month_, err := strconv.Atoi(r.FormValue("month"))
	errs = checkErr(err, errs)
	month := time.Month(month_)

	day, err := strconv.Atoi(r.FormValue("day"))
	errs = checkErr(err, errs)
	hour, err := strconv.Atoi(r.FormValue("hour"))
	errs = checkErr(err, errs)
	minute, err := strconv.Atoi(r.FormValue("minute"))
	errs = checkErr(err, errs)

	ts := time.Date(year, month, day, hour, minute, 0, 0, time.UTC)
	pl := Create(u.Id, name, PasslockPassword(), ts)
	pl.Save()

	return pl, errs
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := router.Parameter(r, "id")
	lg.Log.Printf("delete passlock %s ...", id)

	// how can i factor out all these nested if/else ?
	u, ok := account.UserFromContext(r.Context())
	if !ok {
		lg.Log.Printf("user problem cant delete passlock:%s", id)
	} else {
		id_, err := strconv.Atoi(id)
		if err != nil {
			lg.Log.Printf("bad id:%s", id)
		} else {
			pl, err := Get(id_)
			if err != nil {
				lg.Log.Printf("error getting passlock")
			} else {
				if pl.AccountId == u.Id {
					pl.Delete()
				} else {
					lg.Log.Printf("logged in user does not own passlock to delete")
				}
			}
		}
	}

	http.Redirect(w, r, "/passlock", 303)
	return
}
*/
