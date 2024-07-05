package service

import (
	"gin-rest-api/model"
)

type Tag struct {
	Id         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return model.ExistTagByName(t.Name)
}

func (t *Tag) ExistById() (bool, error) {
	return model.ExistTagById(t.Id)
}

func (t *Tag) Add() error {
	return model.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return model.EditTag(t.Id, data)
}

func (t *Tag) Delete() error {
	return model.DeleteTag(t.Id)
}

func (t *Tag) Count() (int64, error) {
	return model.GetTagTotal(t.getMaps())
}

func (t *Tag) GetAll() ([]model.Tag, error) {
	var (
		tags []model.Tag
	)

	tags, err := model.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}

func (t *Tag) GetTag() (model.Tag, error) {
	tag, err := model.GetTagById(t.Id)
	if err != nil {
		return tag, err
	}

	return tag, nil
}
