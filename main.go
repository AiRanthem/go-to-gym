package main

import (
	log "github.com/sirupsen/logrus"
	"go-to-gym/gym"
)

func main() {
	ctx, cancel := gym.GetContext()
	defer cancel()
	if err := gym.GoToGym(ctx); err != nil {
		log.WithError(err).Error("You cannot go to gym")
	} else {
		log.Info("YOU CAN GO TO GYM NOW !!!")
	}
}
