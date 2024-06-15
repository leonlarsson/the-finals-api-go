package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func HandleLeaderboard(w http.ResponseWriter, r *http.Request) {
	leaderboardId := r.PathValue("leaderboardId")
	platformId := r.PathValue("platformId")
	params := r.URL.Query().Encode()

	leaderboardIdIsValid, validLeaderboardIds := isLeaderboardIdValid(leaderboardId)
	platformIsValid, validPlatformIds := isPlatformIdValid(platformId)

	if !leaderboardIdIsValid {
		http.Error(w, fmt.Sprintf("Provided leaderboardId %s is not valid. Must be one of %s", leaderboardId, strings.Join(validLeaderboardIds, ", ")), http.StatusBadRequest)
		return
	}

	if platformId != "" && !platformIsValid {
		http.Error(w, fmt.Sprintf("Provided platformId %s is not valid. Must be one of %s", platformId, strings.Join(validPlatformIds, ", ")), http.StatusBadRequest)
		return
	}

	url := buildLeaderboardUrl(leaderboardId, platformId, params)

	res, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	json := string(body)
	w.Write([]byte(json))
}

func buildLeaderboardUrl(leaderboardId string, platformId string, params string) string {
	url := "https://api.the-finals-leaderboard.com/v1/leaderboard/" + leaderboardId

	if platformId != "" {
		url += "/" + platformId
	}

	if params != "" {
		url += "?" + params
	}

	return url
}

func isLeaderboardIdValid(leaderboardId string) (bool, []string) {
	validLeaderboardIds := []string{"cb1", "cb2", "ob", "s1", "s2", "s3", "s3worldtour"}
	for _, id := range validLeaderboardIds {
		if leaderboardId == id {
			return true, validLeaderboardIds
		}
	}
	return false, validLeaderboardIds
}

func isPlatformIdValid(platformId string) (bool, []string) {
	validPlatformIds := []string{"crossplay", "steam", "psn", "xbox"}
	for _, id := range validPlatformIds {
		if platformId == id {
			return true, validPlatformIds
		}
	}
	return false, validPlatformIds
}
