package services

func (service Service) GetAllDetectedSmoker() (interface{}, error) {
	data := service.Model.GetAllData()
	return data, nil
}