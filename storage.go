package main

import (
	"github.com/zonesan/clog"
	"golang.org/x/oauth2"
)

var labStore = make(map[string]*oauth2.Token)
var hubStore = make(map[string]*oauth2.Token)

type Storage interface {
	LoadTokenGitlab(user string) (*oauth2.Token, error)
	SaveTokenGitlab(user string, tok *oauth2.Token) error
	LoadTokenGithub(user string) (*oauth2.Token, error)
	SaveTokenGithub(user string, tok *oauth2.Token) error
}

type RedisStore struct {
}

func NewRedisStorage() *RedisStore {
	return &RedisStore{}
}

func (rs *RedisStore) LoadTokenGitlab(user string) (*oauth2.Token, error) {
	clog.Debug("loading user:", user)
	return labStore[user], nil
}

func (rs *RedisStore) SaveTokenGitlab(user string, tok *oauth2.Token) error {
	clog.Debugf("%v: %#v", user, tok)
	labStore[user] = tok
	return nil
}

func (rs *RedisStore) LoadTokenGithub(user string) (*oauth2.Token, error) {
	clog.Debug("called.")
	return nil, nil
}

func (rs *RedisStore) SaveTokenGithub(user string, tok *oauth2.Token) error {
	clog.Debug("called.")
	return nil
}
