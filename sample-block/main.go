package main

import (
	"log"

	"github.com/bharathkkb/tfpkg/pkg/tfgen"
)

func main() {
	// custom module
	mod := tfgen.NewModule("sample", tfgen.ModuleWithSource("terraform-google-modules/cloud-storage/google//modules/simple_bucket"))
	mod.AddAttribute("name", "sample-bucket")
	mod.AddAttribute("project_id", "sample-bucket-project")
	mod.AddAttribute("location", "us-east1")

	// map which file to write blocks to
	blocks := make(map[string][]tfgen.HCLBlock)
	blocks["main.tf"] = []tfgen.HCLBlock{mod}

	// create root module
	root := tfgen.NewRootModule(
		tfgen.RootModuleWithBlocks(blocks),
	)

	// write root module to disk
	err := root.Write("./test")
	if err != nil {
		log.Fatal(err)
	}
}
