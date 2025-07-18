package service

import (
	"errors"
	"hs-web-service/internal/crawler"
)

type Service interface {
	GetPlayerByAccountID() error
}

type Implementation struct {
	craw crawler.Crawler
}

func New(craw crawler.Crawler) (*Implementation, error) {
	if craw == nil {
		return nil, errors.New("craw cannot be nil")
	}
	return &Implementation{craw: craw}, nil
}
