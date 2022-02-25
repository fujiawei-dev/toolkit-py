{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/go-co-op/gocron"
	"github.com/guonaihong/gout"

	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
	"{{GOLANG_MODULE}}/internal/query"
)

func httpUserLogin() (string, error) {
	body := struct {
		query.Response
		Result string `json:"result"`
	}{}

	err := gout.POST(conf.ExternalHttpHostPort() + conf.BasePath() + "/user/login").
		SetJSON(form.UserLogin{
			Username: entity.Admin.Username,
			Password: entity.Admin.Password,
		}).
		Debug(true).
		BindJSON(&body).
		Do()

	if body.Code != 0 && body.Code != 200 {
		log.Printf("corn: http user login failed %s", body.Msg)
	}

	return body.Result, err
}

func httpGetExamples(page int) (entity.Examples, error) {
	body := struct {
		query.Response
		Result struct {
			Pager form.Pager      `json:"pager"`
			List  entity.Examples `json:"list"`
		} `json:"result"`
	}{}

	token, err := httpUserLogin()
	if err != nil {
		return nil, err
	}

	err = gout.GET(conf.ExternalHttpHostPort() + conf.BasePath() + "/examples").
		SetHeader(gout.H{"Authorization": conf.JWTScheme() + " " + token}).
		SetQuery(form.Pager{Page: page, PageSize: 5}).
		Debug(true).
		BindJSON(&body).
		Do()

	if body.Code != 0 {
		log.Printf("corn: get examples message %s", body.Msg)
	}

	return body.Result.List, err
}

func cornHttpGetExamples(cron *gocron.Scheduler) {
	job, err := cron.Every(3).Seconds().Do(func() {

		var (
			page     = 1
			examples entity.Examples
			err      error
		)

		for {
			if examples, err = httpGetExamples(page); err != nil {
				log.Printf("corn: get examples error %v", err)
				return
			}

			if len(examples) == 0 {
				log.Printf("corn: get len(examples)==0, done")
				break
			}

			page++
		}

	})

	if err != nil {
		panic(err)
	}

	job.SingletonMode()
}
