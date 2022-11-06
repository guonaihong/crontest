package robfigcron

import (
	"fmt"

	//"github.com/jakecoffman/cron"
	"github.com/guonaihong/crontest/model"
	cron "github.com/robfig/cron/v3"
)

type Robfigcron struct {
	model.CoreOpt
}

func (r *Robfigcron) SubMain() {

	c := cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
	//c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	//c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	//c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	// Funcs are invoked in their own goroutine, asynchronously.
	// Funcs may also be added to a running Cron
	var err error
	for i := 0; i < r.Count; i++ {
		_, err = c.AddFunc(r.Crontab, func() { fmt.Println("Every Second") })
		if err != nil {
			panic(err.Error())
		}
	}

	// Inspect the cron job entries' next and previous run times.
	//inspect(c.Entries())
	//c.Stop() // Stop the s
	r.Sleep()
}
