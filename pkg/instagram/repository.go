package instagram

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/ldrazic/who-is-not-following-back-ig/internal/config"
	"github.com/ldrazic/who-is-not-following-back-ig/pkg/shared"
	"github.com/sirupsen/logrus"
)

const IG_URL_BASE = "https://www.instagram.com/api/v1"

type Repository struct {
	httpClient shared.HTTPClient
}

func NewInstagramRepository(httpClient shared.HTTPClient) *Repository {
	return &Repository{httpClient}
}

type Storage interface {
	GetFollowers() ([]User, error)
}

func (r *Repository) GetFollowing() ([]User, error) {
	var following []User
	countAmount := 192
	lastId := countAmount
	for {
		response, err := r.getFollowingFromAPI(countAmount, lastId)
		if err != nil {
			logrus.Error("Error in response getting following", err)
			return nil, err
		}
		if len(response) == 0 {
			if countAmount > 3 {
				countAmount = countAmount / 2
			} else {
				break
			}
		}
		following = append(following, response...)
		lastId = len(following) + countAmount
	}
	return following, nil
}
func (r *Repository) GetFollowers() ([]User, error) {
	var followers []User
	nextMaxId := ""
	countAmount := 192
	for {
		changeMaxId := true
		response, responseNextMaxId, err := r.getFollowersFromAPI(nextMaxId, countAmount)
		if err != nil {
			logrus.Error("Error in response getting following", err)
			return nil, err
		}
		if responseNextMaxId == "" {
			if countAmount > 3 {
				countAmount = countAmount / 2
				changeMaxId = false
			} else {
				break
			}
		} else {
			followers = append(followers, response...)
		}
		if changeMaxId {
			nextMaxId = responseNextMaxId
		}
	}
	return followers, nil
}
func (r *Repository) getFollowersFromAPI(nextMaxId string, countAmount int) ([]User, string, error) {
	url := ""
	if nextMaxId != "" {
		url = fmt.Sprintf("%s/friendships/%s/followers/?count=%d&max_id=%s", IG_URL_BASE, config.Config.InstagramUserID, countAmount, nextMaxId)
	} else {
		url = fmt.Sprintf("%s/friendships/%s/followers/?count=%d&search_surface=follow_list_page", IG_URL_BASE, config.Config.InstagramUserID, countAmount)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logrus.Error("Cannot get followers", err)
		return nil, "", err
	}
	cookie, err := base64.StdEncoding.DecodeString(config.Config.Base64InstagramCookie)
	if err != nil {
		logrus.Error("Cannot get followers", err)
		return nil, "", err
	}
	req.Header.Set("Cookie", string(cookie))
	req.Header.Set("X-Ig-App-Id", config.Config.InstagramAppID)
	var response FollowingResponse
	res, err := r.httpClient.Do(req, &response)
	if res.StatusCode >= 200 && res.StatusCode <= 300 && err == nil {
		return response.Users, response.NextMaxID, nil
	}
	if err != nil {
		return nil, "", err
	}
	return nil, "", errors.New("cannot get response error getting following")
}
func (r *Repository) getFollowingFromAPI(countAmount, lastId int) ([]User, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/friendships/%s/following/?count=%d&max_id=%d", IG_URL_BASE, config.Config.InstagramUserID, countAmount, lastId), nil)
	if err != nil {
		logrus.Error("Cannot build request following body", err)
		return nil, err
	}
	cookie, err := base64.StdEncoding.DecodeString(config.Config.Base64InstagramCookie)
	if err != nil {
		logrus.Error("Cannot get followers", err)
		return nil, err
	}
	req.Header.Set("Cookie", string(cookie))
	req.Header.Set("X-Ig-App-Id", config.Config.InstagramAppID)
	var response FollowingResponse
	res, err := r.httpClient.Do(req, &response)
	if res.StatusCode >= 200 && res.StatusCode <= 300 && err == nil {
		return response.Users, nil
	}
	if err != nil {
		return nil, err
	}
	return nil, errors.New("cannot get response error getting following")
}
