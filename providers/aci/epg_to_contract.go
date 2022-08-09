package aci

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const EPGToContractClass = "fvRsProv"

var missTargetEPGToContractResource []string

type EPGToContractGenerator struct {
	ACIService
}

func (a *EPGToContractGenerator) InitResources() error {
	const epgToContractResourceName = "aci_epg_to_contract"

	if clientImpl == nil {
		_, err := a.createClient()
		if err != nil {
			return err
		}
	}

	client := clientImpl

	baseURL := "/api/node/class"
	dnURL := fmt.Sprintf("%s/%s.json", baseURL, EPGToContractClass)

	EPGToContractCont, err := client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	EPGToContractCount, err := strconv.Atoi(stripQuotes(EPGToContractCont.S("totalCount").String()))
	provCount := EPGToContractCount
	if err != nil {
		return err
	}

	for i := 0; i < EPGToContractCount; i++ {
		EPGToContractDN := stripQuotes(EPGToContractCont.S("imdata").Index(i).S(EPGToContractClass, "attributes", "dn").String())

		if checkTarget(epgToContractResourceName, EPGToContractCont.S("imdata").Index(i).S(EPGToContractClass, "attributes")) {
			if filterChildrenDn(EPGToContractDN, client.parentResource) != "" {
				resource := terraformutils.NewResource(
					EPGToContractDN,
					resourceNamefromDn(EPGToContractClass, (EPGToContractDN), i),
					epgToContractResourceName,
					"aci",
					map[string]string{
						"contract_type": "provider",
					},
					[]string{
						"annotation",
						"match_t",
						"prio",
						"description",
					},
					map[string]interface{}{},
				)
				resource.SlowQueryRequired = true
				a.Resources = append(a.Resources, resource)
			}
		} else {
			missTargetEPGToContractResource = append(missTargetEPGToContractResource, fmt.Sprintf("%s:%s", epgToContractResourceName, EPGToContractDN))
		}
	}
	// Consumer
	dnURL = fmt.Sprintf("%s/%s.json", baseURL, "fvRsCons")

	EPGToContractCont, err = client.GetViaURL(dnURL)

	if err != nil {
		return err
	}

	EPGToContractCount, err = strconv.Atoi(stripQuotes(EPGToContractCont.S("totalCount").String()))

	if err != nil {
		return err
	}

	for i := 0; i < EPGToContractCount; i++ {
		EPGToContractDN := stripQuotes(EPGToContractCont.S("imdata").Index(i).S("fvRsCons", "attributes", "dn").String())

		if checkTarget(epgToContractResourceName, EPGToContractCont.S("imdata").Index(i).S("fvRsCons", "attributes")) {
			if filterChildrenDn(EPGToContractDN, client.parentResource) != "" {
				resource := terraformutils.NewResource(
					EPGToContractDN,
					resourceNamefromDn("fvRsCons", GetMOName(EPGToContractDN), i+provCount),
					epgToContractResourceName,
					"aci",
					map[string]string{
						"contract_type": "consumer",
					},
					[]string{
						"annotation",
						"match_t",
						"prio",
						"description",
					},
					map[string]interface{}{},
				)
				resource.SlowQueryRequired = true
				a.Resources = append(a.Resources, resource)
			}
		} else {
			missTargetEPGToContractResource = append(missTargetEPGToContractResource, fmt.Sprintf("%s:%s", epgToContractResourceName, EPGToContractDN))
		}
	}
	printMissTarget(missTargetEPGToContractResource)
	return nil
}
