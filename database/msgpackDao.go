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
	err = ioutil.WriteFile(os.Getenv("GOBIN")+"/datas", b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func QueryAll() (*models.Package, error) {
	b, err := ioutil.ReadFile(os.Getenv("GOBIN") + "/datas")
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

func InsertBrook(brook *models.Brook) error {
	p, err := QueryAll()
	if err != nil {
		fmt.Println(err)
	}
	p.Brooks = append(p.Brooks, *brook)
	err = Marshal(p)
	if err != nil {
		return err
	}
	return nil
}

func RemoveBrook(originUrl string) error {
	p, err := QueryAll()
	if err != nil {
		return err
	}
	var index = -1
	for i := 0; i < len(p.Brooks); i++ {
		brook := p.Brooks[i]
		if originUrl == brook.OriginUrl {
			index = i
			break
		}
	}
	if index >= 0 {
		p.Brooks = append(p.Brooks[:index], p.Brooks[index+1:]...)
	}
	err = Marshal(p)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBrook(model *models.Brook) error {
	p, err := QueryAll()
	if err != nil {
		return err
	}
	for i := 0; i < len(p.Brooks); i++ {
		brook := p.Brooks[i]
		if brook.OriginUrl == model.OriginUrl {
			brook = *model
		}
	}
	err = Marshal(p)
	if err != nil {
		return err
	}
	return nil
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
	for i := 0; i < len(p.Profiles); i++ {
		profile := p.Profiles[i]
		if profile.OriginUrl == model.OriginUrl {
			profile = *model
		}
	}
	err = Marshal(p)
	if err != nil {
		return err
	}
	return nil
}
