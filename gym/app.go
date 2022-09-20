package gym

import (
	"context"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetContext() (context.Context, context.CancelFunc) {
	if viper.GetBool(ConfigKeyDebug) {
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
		)
		return chromedp.NewExecAllocator(context.Background(), opts...)
	} else {
		log.Info("debug mode")
		return chromedp.NewContext(context.Background())
	}
}

func GoToGym(ctx context.Context) error {

	log.Info("Start hacking...")
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://ids.xmu.edu.cn/authserver/login?service=http://cgyy.xmu.edu.cn/idcallback"),
		chromedp.Navigate("https://cgyy.xmu.edu.cn/my_reservations"),
	)

	if err != nil {
		return err
	}

	return nil
}
