/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

var testAccResourcePolicyNATRuleName = "nsxt_policy_nat_rule.test"
var testAccResourcePolicyNATRuleSourceNet = "14.1.1.3"
var testAccResourcePolicyNATRuleDestNet = "15.1.1.3"
var testAccResourcePolicyNATRuleTransNet = "16.1.1.3"

func TestAccResourceNsxtPolicyNATRule_minimalT0(t *testing.T) {
	name := getAccTestResourceName()
	action := model.PolicyNatRule_ACTION_REFLEXIVE

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccNsxtPolicyNATRuleCheckDestroy(state, name)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNsxtPolicyNATRuleTier0MinimalCreateTemplate(name, testAccResourcePolicyNATRuleSourceNet, testAccResourcePolicyNATRuleTransNet),
				Check: resource.ComposeTestCheckFunc(
					testAccNsxtPolicyNATRuleExists(testAccResourcePolicyNATRuleName),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "display_name", name),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "source_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "source_networks.0", testAccResourcePolicyNATRuleSourceNet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.0", testAccResourcePolicyNATRuleTransNet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "tag.#", "0"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "action", action),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "logging", "false"),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "path"),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "revision"),
				),
			},
		},
	})
}

func TestAccResourceNsxtPolicyNATRule_basicT1(t *testing.T) {
	name := getAccTestResourceName()
	updateName := getAccTestResourceName()
	snet := "22.1.1.2"
	dnet := "33.1.1.2"
	tnet := "44.1.1.2"
	action := model.PolicyNatRule_ACTION_DNAT

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccNsxtPolicyNATRuleCheckDestroy(state, name)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNsxtPolicyNATRuleTier1CreateTemplate(name, action, testAccResourcePolicyNATRuleSourceNet, testAccResourcePolicyNATRuleDestNet, testAccResourcePolicyNATRuleTransNet),
				Check: resource.ComposeTestCheckFunc(
					testAccNsxtPolicyNATRuleExists(testAccResourcePolicyNATRuleName),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "display_name", name),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "description", "Acceptance Test"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "destination_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "source_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "destination_networks.0", testAccResourcePolicyNATRuleDestNet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "source_networks.0", testAccResourcePolicyNATRuleSourceNet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.0", testAccResourcePolicyNATRuleTransNet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "tag.#", "2"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "action", action),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "logging", "false"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "firewall_match", model.PolicyNatRule_FIREWALL_MATCH_MATCH_EXTERNAL_ADDRESS),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "path"),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "revision"),
				),
			},
			{
				Config: testAccNsxtPolicyNATRuleTier1CreateTemplate(updateName, action, snet, dnet, tnet),
				Check: resource.ComposeTestCheckFunc(
					testAccNsxtPolicyNATRuleExists(testAccResourcePolicyNATRuleName),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "display_name", updateName),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "description", "Acceptance Test"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "destination_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "source_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "destination_networks.0", dnet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "source_networks.0", snet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.0", tnet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "tag.#", "2"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "action", action),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "logging", "false"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "firewall_match", model.PolicyNatRule_FIREWALL_MATCH_MATCH_EXTERNAL_ADDRESS),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "path"),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "revision"),
				),
			},
			{
				Config: testAccNsxtPolicyNATRuleTier1UpdateMultipleSourceNetworksTemplate(name, action, testAccResourcePolicyNATRuleSourceNet, snet, dnet, tnet),
				Check: resource.ComposeTestCheckFunc(
					testAccNsxtPolicyNATRuleExists(testAccResourcePolicyNATRuleName),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "display_name", name),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "description", "Acceptance Test"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "destination_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "source_networks.#", "2"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "destination_networks.0", dnet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.0", tnet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "tag.#", "2"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "action", action),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "logging", "false"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "firewall_match", model.PolicyNatRule_FIREWALL_MATCH_MATCH_EXTERNAL_ADDRESS),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "path"),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "revision"),
				),
			},
		},
	})
}

func TestAccResourceNsxtPolicyNATRule_basicT0(t *testing.T) {
	name := getAccTestResourceName()
	updateName := getAccTestResourceName()
	snet := "22.1.1.2"
	tnet := "44.1.1.2"
	action := model.PolicyNatRule_ACTION_REFLEXIVE

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccNsxtPolicyNATRuleCheckDestroy(state, updateName)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNsxtPolicyNATRuleTier0CreateTemplate(name, action, testAccResourcePolicyNATRuleSourceNet, testAccResourcePolicyNATRuleTransNet),
				Check: resource.ComposeTestCheckFunc(
					testAccNsxtPolicyNATRuleExists(testAccResourcePolicyNATRuleName),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "display_name", name),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "description", "Acceptance Test"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "source_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "source_networks.0", testAccResourcePolicyNATRuleSourceNet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.0", testAccResourcePolicyNATRuleTransNet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "tag.#", "2"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "action", action),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "logging", "false"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "firewall_match", model.PolicyNatRule_FIREWALL_MATCH_MATCH_EXTERNAL_ADDRESS),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "scope.#", "2"),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "path"),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "revision"),
				),
			},
			{
				Config: testAccNsxtPolicyNATRuleTier0CreateTemplate(updateName, action, snet, tnet),
				Check: resource.ComposeTestCheckFunc(
					testAccNsxtPolicyNATRuleExists(testAccResourcePolicyNATRuleName),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "display_name", updateName),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "description", "Acceptance Test"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "source_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.#", "1"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "source_networks.0", snet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "translated_networks.0", tnet),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "tag.#", "2"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "action", action),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "logging", "false"),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "firewall_match", model.PolicyNatRule_FIREWALL_MATCH_MATCH_EXTERNAL_ADDRESS),
					resource.TestCheckResourceAttr(testAccResourcePolicyNATRuleName, "scope.#", "2"),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "path"),
					resource.TestCheckResourceAttrSet(testAccResourcePolicyNATRuleName, "revision"),
				),
			},
		},
	})
}

func TestAccResourceNsxtPolicyNATRule_basicT1Import(t *testing.T) {
	name := getAccTestResourceName()
	action := model.PolicyNatRule_ACTION_DNAT

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccNsxtPolicyNATRuleCheckDestroy(state, name)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNsxtPolicyNATRuleTier1CreateTemplate(name, action, testAccResourcePolicyNATRuleSourceNet, testAccResourcePolicyNATRuleDestNet, testAccResourcePolicyNATRuleTransNet),
			},
			{
				ResourceName:      testAccResourcePolicyNATRuleName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccNSXPolicyNATRuleImporterGetID,
			},
		},
	})
}

func testAccNSXPolicyNATRuleImporterGetID(s *terraform.State) (string, error) {
	rs, ok := s.RootModule().Resources[testAccResourcePolicyNATRuleName]
	if !ok {
		return "", fmt.Errorf("NSX Policy NAT Rule resource %s not found in resources", testAccResourcePolicyNATRuleName)
	}
	resourceID := rs.Primary.ID
	if resourceID == "" {
		return "", fmt.Errorf("NSX Policy NAT Rule resource ID not set in resources ")
	}
	gwPath := rs.Primary.Attributes["gateway_path"]
	if gwPath == "" {
		return "", fmt.Errorf("NSX Policy NAT Rule Gateway Policy Path not set in resources ")
	}
	_, gwID := parseGatewayPolicyPath(gwPath)
	return fmt.Sprintf("%s/%s", gwID, resourceID), nil
}

func testAccNsxtPolicyNATRuleExists(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		connector := getPolicyConnector(testAccProvider.Meta().(nsxtClients))

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Policy NAT Rule resource %s not found in resources", resourceName)
		}

		resourceID := rs.Primary.ID
		if resourceID == "" {
			return fmt.Errorf("Policy NAT Rule resource ID not set in resources")
		}

		gwPath := rs.Primary.Attributes["gateway_path"]
		isT0, gwID := parseGatewayPolicyPath(gwPath)
		_, err := getNsxtPolicyNATRuleByID(connector, gwID, isT0, resourceID, testAccIsGlobalManager())
		if err != nil {
			return fmt.Errorf("Error while retrieving policy NAT Rule ID %s. Error: %v", resourceID, err)
		}

		return nil
	}
}

func testAccNsxtPolicyNATRuleCheckDestroy(state *terraform.State, displayName string) error {
	connector := getPolicyConnector(testAccProvider.Meta().(nsxtClients))
	for _, rs := range state.RootModule().Resources {

		if rs.Type != "nsxt_policy_nat_rule" {
			continue
		}

		resourceID := rs.Primary.Attributes["id"]
		gwPath := rs.Primary.Attributes["gateway_path"]
		isT0, gwID := parseGatewayPolicyPath(gwPath)
		_, err := getNsxtPolicyNATRuleByID(connector, gwID, isT0, resourceID, testAccIsGlobalManager())
		if err == nil {
			return fmt.Errorf("Policy NAT Rule %s still exists", displayName)
		}
	}
	return nil
}

func testAccNsxtPolicyNATRuleTier0MinimalCreateTemplate(name, sourceNet, translatedNet string) string {
	return testAccNsxtPolicyEdgeClusterReadTemplate(getEdgeClusterName()) +
		testAccNsxtPolicyTier0WithEdgeClusterTemplate("test", false) + fmt.Sprintf(`
resource "nsxt_policy_nat_rule" "test" {
  display_name         = "%s"
  gateway_path         = nsxt_policy_tier0_gateway.test.path
  action               = "%s"
  source_networks      = ["%s"]
  translated_networks  = ["%s"]
}
`, name, model.PolicyNatRule_ACTION_REFLEXIVE, sourceNet, translatedNet)
}

func testAccNsxtPolicyNATRuleTier1CreateTemplate(name string, action string, sourceNet string, destNet string, translatedNet string) string {
	return testAccNsxtPolicyEdgeClusterReadTemplate(getEdgeClusterName()) +
		testAccNsxtPolicyTier1WithEdgeClusterTemplate("test", false) + fmt.Sprintf(`
resource "nsxt_policy_nat_rule" "test" {
  display_name         = "%s"
  description          = "Acceptance Test"
  gateway_path         = nsxt_policy_tier1_gateway.test.path
  action               = "%s"
  source_networks      = ["%s"]
  destination_networks = ["%s"]
  translated_networks  = ["%s"]
  logging              = false
  firewall_match       = "%s"

  tag {
    scope = "scope1"
    tag   = "tag1"
  }

  tag {
    scope = "scope2"
    tag   = "tag2"
  }
}
`, name, action, sourceNet, destNet, translatedNet, model.PolicyNatRule_FIREWALL_MATCH_MATCH_EXTERNAL_ADDRESS)
}

func testAccNsxtPolicyNATRuleTier1UpdateMultipleSourceNetworksTemplate(name string, action string, sourceNet1 string, sourceNet2 string, destNet string, translatedNet string) string {
	return testAccNsxtPolicyEdgeClusterReadTemplate(getEdgeClusterName()) +
		testAccNsxtPolicyTier1WithEdgeClusterTemplate("test", false) + fmt.Sprintf(`
resource "nsxt_policy_nat_rule" "test" {
  display_name         = "%s"
  description          = "Acceptance Test"
  gateway_path         = nsxt_policy_tier1_gateway.test.path
  action               = "%s"
  source_networks      = ["%s", "%s"]
  destination_networks = ["%s"]
  translated_networks  = ["%s"]
  logging              = false
  firewall_match       = "%s"

  tag {
    scope = "scope1"
    tag   = "tag1"
  }

  tag {
    scope = "scope2"
    tag   = "tag2"
  }
}
`, name, action, sourceNet1, sourceNet2, destNet, translatedNet, model.PolicyNatRule_FIREWALL_MATCH_MATCH_EXTERNAL_ADDRESS)
}

func testAccNsxtPolicyNATRuleTier0CreateTemplate(name string, action string, sourceNet string, translatedNet string) string {

	var transportZone string
	interfaceSite := ""
	if testAccIsGlobalManager() {
		transportZone = `
data "nsxt_policy_transport_zone" "test" {
  transport_type = "VLAN_BACKED"
  is_default     = true
  site_path      = data.nsxt_policy_site.test.path
}`
		interfaceSite = "site_path = data.nsxt_policy_site.test.path"
	} else {
		transportZone = fmt.Sprintf(`
data "nsxt_policy_transport_zone" "test" {
  display_name = "%s"
}`, getVlanTransportZoneName())
	}
	return testAccNsxtPolicyEdgeClusterReadTemplate(getEdgeClusterName()) +
		transportZone + testAccNsxtPolicyTier0WithEdgeClusterTemplate("test", true) + fmt.Sprintf(`
resource "nsxt_policy_vlan_segment" "test" {
  count               = 2
  transport_zone_path = data.nsxt_policy_transport_zone.test.path
  display_name        = "interface-test-${count.index}"
  vlan_ids            = [10 + count.index]
}

resource "nsxt_policy_tier0_gateway_interface" "test" {
  count        = 2
  display_name = "gwinterface-test-${count.index}"
  type         = "SERVICE"
  gateway_path = nsxt_policy_tier0_gateway.test.path
  segment_path = nsxt_policy_vlan_segment.test[count.index].path
  subnets      = ["1.1.${count.index}.2/24"]
  %s
}

resource "nsxt_policy_nat_rule" "test" {
  display_name         = "%s"
  description          = "Acceptance Test"
  gateway_path         = nsxt_policy_tier0_gateway.test.path
  action               = "%s"
  source_networks      = ["%s"]
  translated_networks  = ["%s"]
  logging              = false
  firewall_match       = "%s"
  scope                = [nsxt_policy_tier0_gateway_interface.test[1].path, nsxt_policy_tier0_gateway_interface.test[0].path]

  tag {
    scope = "scope1"
    tag   = "tag1"
  }

  tag {
    scope = "scope2"
    tag   = "tag2"
  }
}

`, interfaceSite, name, action, sourceNet, translatedNet, model.PolicyNatRule_FIREWALL_MATCH_MATCH_EXTERNAL_ADDRESS)
}
