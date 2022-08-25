// This file is generated by github.com/bharathkkb/tfpkg. Do not edit.
package network

import "github.com/bharathkkb/tfpkg/pkg/tfgen"

// Inputs marked as required by the module
type RequiredAttrib struct {
	Subnets     []map[string]string
	ProjectId   string
	NetworkName string
}

// Inputs with a default value in the module
type OptionalAttrib struct {
	DeleteDefaultInternetGatewayRoutes bool
	Description                        string
	AutoCreateSubnetworks              bool
	SharedVpcHost                      bool
	SecondaryRanges                    map[string][]string
	Routes                             []map[string]string
	FirewallRules                      string
	Mtu                                int
	RoutingMode                        string
}
type Mod struct {
	tfgen.Module
	Required *RequiredAttrib
	Optional *OptionalAttrib
}

func New(name string, required *RequiredAttrib, optional *OptionalAttrib) Mod {
	mod := tfgen.NewModule(name, tfgen.ModuleWithSource("terraform-google-modules/network/google"))
	n := Mod{
		Module:   *mod,
		Optional: optional,
		Required: required,
	}
	n.AddAttribute("subnets", required.Subnets)
	n.AddAttribute("project_id", required.ProjectId)
	n.AddAttribute("network_name", required.NetworkName)
	n.AddAttribute("description", optional.Description)
	n.AddAttribute("firewall_rules", optional.FirewallRules)
	n.AddAttribute("mtu", optional.Mtu)
	n.AddAttribute("routing_mode", optional.RoutingMode)
	n.AddAttribute("delete_default_internet_gateway_routes", optional.DeleteDefaultInternetGatewayRoutes)
	n.AddAttribute("auto_create_subnetworks", optional.AutoCreateSubnetworks)
	n.AddAttribute("shared_vpc_host", optional.SharedVpcHost)
	n.AddAttribute("secondary_ranges", optional.SecondaryRanges)
	n.AddAttribute("routes", optional.Routes)
	return n
}

// Getters for output references from the module
func (m Mod) GetNetworkRef() string {
	return m.GetRef("network")
}

func (m Mod) GetSubnetsRef() string {
	return m.GetRef("subnets")
}

func (m Mod) GetNetworkNameRef() string {
	return m.GetRef("network_name")
}

func (m Mod) GetSubnetsIpsRef() string {
	return m.GetRef("subnets_ips")
}

func (m Mod) GetSubnetsPrivateAccessRef() string {
	return m.GetRef("subnets_private_access")
}

func (m Mod) GetSubnetsSecondaryRangesRef() string {
	return m.GetRef("subnets_secondary_ranges")
}

func (m Mod) GetProjectIdRef() string {
	return m.GetRef("project_id")
}

func (m Mod) GetSubnetsIdsRef() string {
	return m.GetRef("subnets_ids")
}

func (m Mod) GetSubnetsRegionsRef() string {
	return m.GetRef("subnets_regions")
}

func (m Mod) GetRouteNamesRef() string {
	return m.GetRef("route_names")
}

func (m Mod) GetNetworkIdRef() string {
	return m.GetRef("network_id")
}

func (m Mod) GetNetworkSelfLinkRef() string {
	return m.GetRef("network_self_link")
}

func (m Mod) GetSubnetsNamesRef() string {
	return m.GetRef("subnets_names")
}

func (m Mod) GetSubnetsSelfLinksRef() string {
	return m.GetRef("subnets_self_links")
}

func (m Mod) GetSubnetsFlowLogsRef() string {
	return m.GetRef("subnets_flow_logs")
}

// Setters for setting references as inputs
func (m *Mod) SetSecondaryRangesAsRef(ref string) {
	m.AddAttribute("secondary_ranges", ref)
}

func (m *Mod) SetDeleteDefaultInternetGatewayRoutesAsRef(ref string) {
	m.AddAttribute("delete_default_internet_gateway_routes", ref)
}

func (m *Mod) SetDescriptionAsRef(ref string) {
	m.AddAttribute("description", ref)
}

func (m *Mod) SetAutoCreateSubnetworksAsRef(ref string) {
	m.AddAttribute("auto_create_subnetworks", ref)
}

func (m *Mod) SetSharedVpcHostAsRef(ref string) {
	m.AddAttribute("shared_vpc_host", ref)
}

func (m *Mod) SetSubnetsAsRef(ref string) {
	m.AddAttribute("subnets", ref)
}

func (m *Mod) SetRoutingModeAsRef(ref string) {
	m.AddAttribute("routing_mode", ref)
}

func (m *Mod) SetRoutesAsRef(ref string) {
	m.AddAttribute("routes", ref)
}

func (m *Mod) SetFirewallRulesAsRef(ref string) {
	m.AddAttribute("firewall_rules", ref)
}

func (m *Mod) SetMtuAsRef(ref string) {
	m.AddAttribute("mtu", ref)
}

func (m *Mod) SetProjectIdAsRef(ref string) {
	m.AddAttribute("project_id", ref)
}

func (m *Mod) SetNetworkNameAsRef(ref string) {
	m.AddAttribute("network_name", ref)
}
