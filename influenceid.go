package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

type loginParam struct {
	urlLogin         string
	user             string
	pass             string
	selectorUser     string
	selectorPassword string
	selectorButton   string
}

type result struct {
	h1Text  string
	request []byte
}

type navigateInfluencerListParam struct {
	selectorButtons   string
	urlListInfluencer string
	h1Text            *string
	request           []byte
}

func scrapeInfluenceId(db *sql.DB) {
	/**
	TOTAL DATA = 3800
	*/

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.WindowSize(1920, 1080),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	result := result{
		h1Text:  "",
		request: []byte{},
	}

	urlLogin := "https://advertiser.influence.id/login"
	user := "qkhddhifgflneq@cutradition.com"
	pass := "Vagrant!234"
	selectorUser := `//div[@class="login-empty-input"]/div/div/div/span/input`
	selectorPass := `//div[@class="login-empty-password"]/div/div/div/span/input`
	selectorSubmit := `//button[@type="submit"]`
	paramSubmitLogin := loginParam{
		urlLogin:         urlLogin,
		user:             user,
		pass:             pass,
		selectorUser:     selectorUser,
		selectorPassword: selectorPass,
		selectorButton:   selectorSubmit,
	}
	submitLoginTasks := submitLogin(paramSubmitLogin)

	urlListInfluencer := "https://advertiser.influence.id/influencer/discovery"
	selectorButtons := `(//button[@class="identity-id-button box-shadow ant-btn"])`
	paramNavigateListInfluencer := navigateInfluencerListParam{
		selectorButtons:   selectorButtons,
		urlListInfluencer: urlListInfluencer,
		h1Text:            &result.h1Text,
		request:           result.request,
	}
	navigateToInfluencerListTasks := navigateToInfluencerList(paramNavigateListInfluencer)

	chromedp.ListenTarget(
		ctx,
		func(ev interface{}) {
			if ev, ok := ev.(*network.EventResponseReceived); ok {
				if ev.Type == "XHR" && strings.Contains(ev.Response.URL, "api.influence.id/adv/v1/influencer/participant") {
					// json, _ := json.Marshal(ev.Request.Headers)
					// fmt.Println(string(json))
					go func() {
						fmt.Println(ev.Response)
						c := chromedp.FromContext(ctx)
						rbp := network.GetResponseBody(ev.RequestID)
						body, err := rbp.Do(cdp.WithExecutor(ctx, c.Target))
						if err != nil {
							fmt.Println(err)
						}

						_, err = db.Exec("INSERT INTO influenceid VALUES ($1, $2)", ev.RequestID, string(body))
						if err != nil {
							fmt.Println("Error: %v", err)
						}
						fmt.Printf("INSERT %s", ev.RequestID)
					}()
				}
			}
		},
	)

	err := chromedp.Run(
		ctx,
		buildTasks(
			submitLoginTasks,
			navigateToInfluencerListTasks,
		),
	)
	if err != nil {
		fmt.Print("Error: %v", err)
	}
}

func submitLogin(param loginParam) chromedp.Tasks {
	return chromedp.Tasks{
		emulation.SetDeviceMetricsOverride(int64(1920), int64(1080), 1.0, false),
		chromedp.Navigate(param.urlLogin),
		chromedp.WaitVisible(param.selectorUser),
		chromedp.SetValue(param.selectorUser, param.user, chromedp.BySearch),
		chromedp.SetValue(param.selectorPassword, param.pass, chromedp.BySearch),
		chromedp.Click(param.selectorButton),
		chromedp.Sleep(3 * time.Second),
	}
}

func navigateToInfluencerList(param navigateInfluencerListParam) chromedp.Tasks {
	totalInfluencers := 3800
	totalPerPage := 20
	counterPage := 1
	counterInfluencers := 1

	tasks := chromedp.Tasks{
		network.Enable(),
		chromedp.Navigate("https://advertiser.influence.id/influencer/discovery"),
	}

	for counterInfluencers <= totalInfluencers/totalPerPage {
		for counterPage <= totalPerPage {
			tasks = append(
				tasks,
				chromedp.Click(fmt.Sprintf("%s[%d]", param.selectorButtons, counterPage), chromedp.BySearch),
				chromedp.Sleep(5*time.Second),
				chromedp.Click(`//button[@class="ant-drawer-close"]`, chromedp.BySearch),
				chromedp.Sleep(2*time.Second),
			)
			counterPage += 1
		}
		tasks = append(
			tasks,
			chromedp.Click(`(//button[@class="pagination-border ant-btn"])[2]`, chromedp.BySearch),
			chromedp.Sleep(7*time.Second),
		)
		counterPage = 1
		counterInfluencers += 1
	}

	return tasks
}

func buildTasks(tasks ...chromedp.Action) chromedp.Tasks {
	result := chromedp.Tasks{}
	for _, v := range tasks {
		result = append(result, v)
	}
	return result
}
