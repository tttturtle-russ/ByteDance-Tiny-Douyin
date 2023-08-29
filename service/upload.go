package service

func (svc *Service) UploadVideo(fileName string, title string, authorID int64) (err error) {
	return svc.d.UploadVideo(fileName, title, authorID)
}
