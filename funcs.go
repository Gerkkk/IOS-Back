package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) fetchPeople(c *gin.Context) {
	Users := []Person{Person{
		Id:           0,
		Name:         "Anton",
		Tag:          "antony",
		Status:       "Eating food",
		Follows:      []int{1, 2},
		FollowersNum: 10,
		Avatar:       nil,
	}, Person{
		Id:           1,
		Name:         "Bob",
		Tag:          "biba",
		Status:       "Living life",
		Follows:      []int{0, 2},
		FollowersNum: 100,
		Avatar:       nil,
	},
		Person{
			Id:           2,
			Name:         "Chris",
			Tag:          "boba",
			Status:       "Making love",
			Follows:      []int{0, 1},
			FollowersNum: 1000,
			Avatar:       nil,
		}}

	c.JSON(http.StatusOK, map[string]interface{}{
		"people": Users,
	})
}
