package main

import (
	"fmt"
	"github.com/tyomarius05/go-academy/40-go-tutorialPostgres/config"
	"github.com/tyomarius05/go-academy/40-go-tutorialPostgres/src/modules/profile/model"
	"github.com/tyomarius05/go-academy/40-go-tutorialPostgres/src/modules/profile/repository"
)

func main() {
	fmt.Println("=====Ini merupakan contoh Koneksi ke DB Postgres di Golang=====")
	fmt.Println()

	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}

	/*
		profileRepositoryPostgres := repository.NewProfileRepositoryPostgres(db)
		myProfile := model.NewProfile()

		myProfile.ID = "2"
		myProfile.FirstName = "Ancilla"
		myProfile.LastName = "Intan"
		myProfile.Email = "ancillaintan@gmail.com"
		myProfile.Password = "98765"
		// err = saveProfile(myProfile, profileRepositoryPostgres)

		// myProfile := model.NewProfile()
		// myProfile.ID = "1"
		// myProfile.FirstName = "Tyo"
		// myProfile.LastName = "Marius"
		// myProfile.Email = "tyomarius05@gmail.com"
		// myProfile.Password = "12345"
		// err = updateProfile(myProfile, profileRepositoryPostgres)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// err = deleteProfile("1", profileRepositoryPostgres)

		myProfile2, err := getProfile("2", profileRepositoryPostgres)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(myProfile2)


		fmt.Println("===============")
		myProfiles, err := getProfiles(profileRepositoryPostgres)

		if err != nil {
			fmt.Println(err)
		}

		for _, v := range myProfiles {
			fmt.Println(v)
		}
	*/
}

func saveProfile(p *model.Profile, repo repository.ProfileRepository) error {
	err := repo.Save(p)

	if err != nil {
		return err
	}

	return nil
}

func updateProfile(p *model.Profile, repo repository.ProfileRepository) error {
	err := repo.Update(p.ID, p)

	if err != nil {
		return err
	}

	return nil
}

func deleteProfile(id string, repo repository.ProfileRepository) error {
	err := repo.Delete(id)

	if err != nil {
		return err
	}
	return nil
}

func getProfile(id string, repo repository.ProfileRepository) (*model.Profile, error) {
	profile, err := repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func getProfiles(repo repository.ProfileRepository) (model.Profiles, error) {
	profiles, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return profiles, nil
}
