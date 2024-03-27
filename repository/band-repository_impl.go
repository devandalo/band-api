package repository

import "band_api/model"

type BandRepository interface {
	Save(band model.Band)
	Update(band model.Band)
	Delete(bandId int)
	FindById(bandId int) (band model.Band, err error)
	FindAll() []model.Band
}