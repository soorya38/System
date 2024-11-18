package handlers

import (
	"fmt"
	"log"
	"net/http"
	"profile/database"
	"profile/models"
)

func RegisterHandlers() {
	http.HandleFunc("/create-profile", handleCreateProfile)
	http.HandleFunc("/delete-profile/", handleDeleteProfile)
}

func StartServer(PORT int) error {
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%v", PORT), nil); err != nil {
		return fmt.Errorf("unable to start server")
	}

	return nil
}

func handleCreateProfile(w http.ResponseWriter, r *http.Request) {
	// connect with mongodb
	// create user name password in postgres
	// student id: incrementing number
	// get id from postgres
	// create profile with default values

	// this data is from front-end
	// only for test purposes
	user := models.Profile{
		ID:           1,
		Name:         "test",
		UserName:     "test_n",
		Password:     "test",
		ProfileImage: "https://test.com/profile.jpg",
		Email:        "test@mail.com",
		Phone:        "+1234567890",
		Bio:          "test bio",
		Role:         "test_user",
	}

	if err := database.CreateProfile(user); err != nil {
		log.Printf("%v", err)
	}
}

func handleDeleteProfile(w http.ResponseWriter, r *http.Request) {

}
