package filehandler

type FileHandler struct {
	path string
}

func CreateNewFileHander(filepath string) *FileHandler {
	var newFh FileHandler

	return &newFh
}

func (f *FileHandler) RunHandler() {

}
