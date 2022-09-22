package gym

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

const (
	IdDate   = "gym-date"
	IdPeriod = "gym-period"
)

func GetContext() (context.Context, context.CancelFunc) {
	opts := chromedp.DefaultExecAllocatorOptions[:]
	if viper.GetBool(ConfigKeyDebug) {
		log.Debug("debug mode")
		opts = append(opts, chromedp.Flag("headless", false))
	} else {
		opts = append(opts, chromedp.DisableGPU)
	}
	ctx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	return chromedp.NewContext(ctx, chromedp.WithLogf(log.WithField("from", "browser").Printf))
}

func PicName(date string, period string) string {
	return date + "-" + period + ".png"
}

type Action struct {
	A    chromedp.Action
	Desc string
}

// Runner is used for logging job steps and avoid bugs with sleep
type Runner struct {
	Actions []Action
	ctx     context.Context
}

func (r *Runner) Run() error {
	for _, action := range r.Actions {
		time.Sleep(10 * time.Millisecond)
		if err := chromedp.Run(r.ctx, action.A); err != nil {
			log.WithField("desc", action.Desc).Error("action failed")
			return err
		}
		log.WithField("desc", action.Desc).Info("action done")
	}
	return nil
}

func NewRunner(ctx context.Context, actions ...Action) *Runner {
	return &Runner{
		ctx:     ctx,
		Actions: actions,
	}
}

func GoToGym(ctx context.Context, timeFunc TimeFunc) (string, error) {
	date, period := timeFunc()
	log.WithField("date", date).WithField("period", period).Info("AhA, you wanna go to GYM!")
	abs, err := filepath.Abs(filepath.Join(viper.GetString(ConfigKeyStore), PicName(date, period)))
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(abs); err == nil {
		log.Info("WaHHH! Already hacked!")
		return abs, nil
	}
	log.Info("hacking...")
	var buf []byte
	if err := NewRunner(ctx,
		Action{chromedp.Navigate("https://ids.xmu.edu.cn/authserver/login?service=http://cgyy.xmu.edu.cn/idcallback"), "nav login page"},
		Action{chromedp.WaitVisible(`//*[@id="casLoginForm"]/p[4]/button`), "wait login page"},
		Action{chromedp.SendKeys(`//*[@id="username"]`, viper.GetString(ConfigKeyUsername)), "enter username"},
		Action{chromedp.SendKeys(`//*[@id="password"]`, viper.GetString(ConfigKeyPassword)), "enter password"},
		Action{chromedp.Click(`//*[@id="casLoginForm"]/p[4]/button`), "click login"},
		Action{chromedp.WaitVisible(`//*[@id="navbar"]/div/div[1]/a[1]/img`), "wait login"},
		Action{chromedp.Navigate("https://cgyy.xmu.edu.cn/my_reservations/slot"), "goto my_reservations page"},
		Action{chromedp.SetAttributes(`//*[@id="block-system-main"]/div/div/div/table/tbody/tr[1]/td[3]/span`, map[string]string{"id": IdDate}), "set id date"},
		Action{chromedp.SetAttributes(`//*[@id="block-system-main"]/div/div/div/table/tbody/tr[1]/td[4]`, map[string]string{"id": IdPeriod}), "set id period"},
		Action{chromedp.Evaluate(fmt.Sprintf(`document.getElementById("%s").innerHTML="%s"`, IdDate, date), nil), "change date"},
		Action{chromedp.Evaluate(fmt.Sprintf(`document.getElementById("%s").innerHTML="%s"`, IdPeriod, period), nil), "change period"},
		Action{chromedp.SetAttributes(`//*[@id="block-system-main"]/div/div/div/table/tbody`, map[string]string{"style": "font-size:12px;"}), "change font"},
		Action{chromedp.Emulate(device.IPhone13ProMax), "emulate iphone13"},
		Action{chromedp.FullScreenshot(&buf, 100), "screenshot"},
	).Run(); err != nil {
		return "", err
	}
	log.Info("Wooooo! Page hacked! Saving...")
	if err := os.WriteFile(abs, buf, 0777); err != nil {
		return "", err
	}
	log.Info("Hacked page SAVED!")
	return abs, nil
}
