package main

import (
	"log"

	"github.com/bharathkkb/tfpkg-samples/sample-project-net/generated/network"
	"github.com/bharathkkb/tfpkg-samples/sample-project-net/generated/project_factory"
	"github.com/bharathkkb/tfpkg/pkg/tfgen"
)

func main() {
	// create a variable for project name
	projectNameTFVar := tfgen.NewVariable("project_name")

	// required attribs for project factory module
	reqProjParams := &project_factory.RequiredAttrib{
		Name:           projectNameTFVar.GetRef(), // reference output
		BillingAccount: "123",
		OrgId:          "456",
	}
	// optional attribs for project factory module
	optionalProjParams := &project_factory.OptionalAttrib{
		FolderId:        "789",
		RandomProjectId: true,
	}
	// instantiate project factory module
	projectModule := project_factory.New("net-project", reqProjParams, optionalProjParams)

	// required attribs for network module
	reqNetParams := &network.RequiredAttrib{
		ProjectId:   projectModule.GetProjectIdRef(), // reference output of project module to get project id
		NetworkName: "bar",
		Subnets: []map[string]string{
			{
				"subnet_name":   "subnet-01",
				"subnet_ip":     "10.10.10.0/24",
				"subnet_region": "us-west1",
			}, {
				"subnet_name":   "subnet-02",
				"subnet_ip":     "10.10.20.0/24",
				"subnet_region": "us-east1",
			}},
	}
	// instantiate network module
	netModule := network.New("prod-net", reqNetParams, &network.OptionalAttrib{})

	// map which file to write blocks to
	blocks := make(map[string][]tfgen.HCLBlock)
	blocks["main.tf"] = []tfgen.HCLBlock{projectModule, netModule}

	// add variable and output blocks
	tfVars := []*tfgen.Variable{projectNameTFVar}
	tfOPs := []*tfgen.Output{
		tfgen.NewOutput("project_id", projectModule.GetProjectIdRef()),
		tfgen.NewOutput("network_self_link", netModule.GetNetworkSelfLinkRef()),
	}

	// create root module
	root := tfgen.NewRootModule(
		tfgen.RootModuleWithBlocks(blocks),
		tfgen.RootModuleWithVars(tfVars),
		tfgen.RootModuleWithOutputs(tfOPs),
	)

	// write root module to disk
	err := root.Write("./test")
	if err != nil {
		log.Fatal(err)
	}
}
