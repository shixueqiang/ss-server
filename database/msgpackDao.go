package database

import (
	"fmt"
	"io/ioutil"
	"os"
	"ss-server/models"

	"github.com/vmihailenco/msgpack"
)

func Marshal(p *models.Package) error {
	b, err := msgpack.Marshal(p)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(os.Getenv("GOBIN")+"/ss-server/datas", b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func QueryAll() (*models.Package, error) {
	b, err := ioutil.ReadFile(os.Getenv("GOBIN") + "/ss-server/datas")
	if err != nil {
		return new(models.Package), err
	}
	var item models.Package
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		return new(models.Package), err
	}
	return &item, nil
}

func InsertProfileToMsgpack(profile *models.Profile) error {
	p, err := QueryAll()
	if err != nil {
		fmt.Println(err)
	}
	p.Profiles = append(p.Profiles, *profile)
	err = Marshal(p)
	if err != nil {
		return err
	}
	return nil
}

func RemoveProfileFromMsgpack(originUrl string) error {
	p, err := QueryAll()
	if err != nil {
		return err
	}
	var index = -1
	for i := 0; i < len(p.Profiles); i++ {
		profile := p.Profiles[i]
		if originUrl == profile.OriginUrl {
			index = i
			break
		}
	}
	fmt.Printf("originUrl:%s index:%d\n", originUrl, index)
	if index >= 0 {
		p.Profiles = append(p.Profiles[:index], p.Profiles[index+1:]...)
	}
	err = Marshal(p)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProfileFromMsgpack(model *models.Profile) error {
	p, err := QueryAll()
	if err != nil {
		return err
	}
	var index = -1
	for i := 0; i < len(p.Profiles); i++ {
		profile := p.Profiles[i]
		if model.OriginUrl == profile.OriginUrl {
			index = i
			break
		}
	}
	if index >= 0 {
		p.Profiles = append(p.Profiles[:index], p.Profiles[index+1:]...)
	}
	p.Profiles = append(p.Profiles, *model)
	err = Marshal(p)
	if err != nil {
		return err
	}
	return nil
}
