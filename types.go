package main

type Adapter struct {
	Label       string
	Name        string
	Description string
}

var adapters = []Adapter{
	Adapter{
		Label:       "executable",
		Name:        "Periodic executable",
		Description: "Runs an executable on a schedule",
	},
	Adapter{
		Label:       "directory",
		Name:        "Directory watcher",
		Description: "Watches a directory and uploads newly appearing files",
	},
}
