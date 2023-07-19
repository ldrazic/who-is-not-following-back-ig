package instagram

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
)

type Service struct {
	igRepo Repository
}

func NewInstagramService(igRepo Repository) *Service {
	return &Service{igRepo}
}

func (s *Service) CompareFollowingThatDoesntFollowBack() {
	following, err := s.igRepo.GetFollowing()
	if err != nil {
		logrus.Error("error getting following", err)
		log.Fatal(err)
	}
	followers, err := s.igRepo.GetFollowers()
	if err != nil {
		logrus.Error("error getting followers", err)
		log.Fatal(err)
	}
	result := FindUsersNotInFollowers(following, followers)
	fmt.Print("Instagram following that doesnt follows back")
	for _, user := range result {
		fmt.Println(user.Username)
	}
}

func FindUsersNotInFollowers(following, followers []User) []User {
	// Create a map to store the users from the second array
	userMap := make(map[string]bool)
	for _, user := range followers {
		userMap[user.Username] = true
	}

	// Filter users from the first array that are not in the second array
	var result []User
	for _, user := range following {
		if !userMap[user.Username] {
			result = append(result, user)
		}
	}
	return result
}
