package server

type Summary struct {
	Modules    []Module
	Files      []File
	NumFiles   int
	NumModules int
}

type Module struct {
	Name string
}

type File struct {
	Name string
}

func (s *Server) GetSummary() Summary {
	su := Summary{}

	su.Modules = s.getAllModules()
	su.Files = s.getAllFiles()
	su.NumFiles = len(su.Files)
	su.NumModules = len(su.Modules)

	return su
}

func (s *Server) getAllModules() []Module {
	allModules := s.dt.GetAllModules()

	modules := make([]Module, 0)

	for _, module := range allModules {
		modules = append(modules, Module{Name: module.Name})
	}

	return modules
}

func (s *Server) getAllFiles() []File {
	allFiles := s.dt.GetAllFiles()

	files := make([]File, 0)

	for _, file := range allFiles {
		files = append(files, File{Name: file.Name})
	}

	return files
}
