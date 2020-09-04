/*Copyright (c) 2019 Cisco and/or its affiliates.

This software is licensed to you under the terms of the Cisco Sample
Code License, Version 1.0 (the "License"). You may obtain a copy of the
License at

https://developer.cisco.com/docs/licenses

All use of the material herein must be in accordance with the terms of
the License. All rights not expressly granted by the License are
reserved. Unless required by applicable law or agreed to separately in
writing, software distributed under the License is distributed on an "AS
IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
or implied.*/

package main

import (
	"errors"

	"github.com/ccp-client-library/ccp"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceACIProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceACIProfileCreate,
		Read:   resourceACIProfileRead,
		Update: resourceACIProfileUpdate,
		Delete: resourceACIProfileDelete,

		Schema: map[string]*schema.Schema{

			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"apic_hosts": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"apic_username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"apic_password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"aci_vmm_domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"aci_infra_vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"vrf_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"l3_outside_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"l3_outside_network_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"aaep_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"nameservers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"control_plane_contract_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"node_vlan_start": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"node_vlan_end": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"pod_subnet_start": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_subnet_start": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"multicast_range": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"aci_tenant": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceACIProfileCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	nameservers := []string{}
	for _, server := range d.Get("nameservers").([]interface{}) {
		nameservers = append(nameservers, server.(string))
	}

	newACIProfile := ccp.ACIProfile{

		Name:                     ccp.String(d.Get("name").(string)),
		APICHosts:                ccp.String(d.Get("apic_hosts").(string)),
		APICUsername:             ccp.String(d.Get("apic_username").(string)),
		APICPassword:             ccp.String(d.Get("apic_password").(string)),
		ACIVMMDomainName:         ccp.String(d.Get("aci_vmm_domain_name").(string)),
		ACIInfraVLANID:           ccp.Int(d.Get("aci_infra_vlan_id").(int)),
		VRFName:                  ccp.String(d.Get("vrf_name").(string)),
		L3OutsidePolicyName:      ccp.String(d.Get("l3_outside_policy_name").(string)),
		L3OutsideNetworkName:     ccp.String(d.Get("l3_outside_network_name").(string)),
		AAEPName:                 ccp.String(d.Get("aaep_name").(string)),
		Nameservers:              &nameservers,
		ControlPlaneContractName: ccp.String(d.Get("control_plane_contract_name").(string)),
		NodeVLANStart:            ccp.Int(d.Get("node_vlan_start").(int)),
		NodeVLANEnd:              ccp.Int(d.Get("node_vlan_end").(int)),
		PodSubnetStart:           ccp.String(d.Get("pod_subnet_start").(string)),
		ServiceSubnetStart:       ccp.String(d.Get("service_subnet_start").(string)),
		MulticastRange:           ccp.String(d.Get("multicast_range").(string)),
		ACITenant:                ccp.String(d.Get("aci_tenant").(string)),
	}

	aciProfile, err := client.AddACIProfile(&newACIProfile)

	if err != nil {
		return errors.New(err.Error())
	}

	uuid := *aciProfile.UUID

	d.SetId(uuid)

	return setACIProfileResourceData(d, aciProfile)
}

func resourceACIProfileRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	aciProfile, err := client.GetACIProfileByName(d.Get("name").(string))

	if err != nil {
		return errors.New("UNABLE TO RETRIEVE DETAILS FOR ACI PROFILE: " + d.Get("name").(string))
	}

	return setACIProfileResourceData(d, aciProfile)

}

func resourceACIProfileUpdate(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	nameservers := []string{}
	for _, server := range d.Get("nameservers").([]interface{}) {
		nameservers = append(nameservers, server.(string))
	}

	newACIProfile := ccp.ACIProfile{

		Name:                     ccp.String(d.Get("name").(string)),
		APICHosts:                ccp.String(d.Get("apic_hosts").(string)),
		APICUsername:             ccp.String(d.Get("apic_username").(string)),
		APICPassword:             ccp.String(d.Get("apic_password").(string)),
		ACIVMMDomainName:         ccp.String(d.Get("aci_vmm_domain_name").(string)),
		ACIInfraVLANID:           ccp.Int(d.Get("aci_infra_vlan_id").(int)),
		VRFName:                  ccp.String(d.Get("vrf_name").(string)),
		L3OutsidePolicyName:      ccp.String(d.Get("l3_outside_policy_name").(string)),
		L3OutsideNetworkName:     ccp.String(d.Get("l3_outside_network_name").(string)),
		AAEPName:                 ccp.String(d.Get("aaep_name").(string)),
		Nameservers:              &nameservers,
		ControlPlaneContractName: ccp.String(d.Get("control_plane_contract_name").(string)),
		NodeVLANStart:            ccp.Int(d.Get("node_vlan_start").(int)),
		NodeVLANEnd:              ccp.Int(d.Get("node_vlan_end").(int)),
		PodSubnetStart:           ccp.String(d.Get("pod_subnet_start").(string)),
		ServiceSubnetStart:       ccp.String(d.Get("service_subnet_start").(string)),
		MulticastRange:           ccp.String(d.Get("multicast_range").(string)),
		ACITenant:                ccp.String(d.Get("aci_tenant").(string)),
	}

	profile, err := client.PatchACIProfile(&newACIProfile, d.Get("uuid").(string))

	if err != nil {
		return errors.New(err.Error())
	}

	profile, err = client.GetACIProfileByName(d.Get("name").(string))

	if err != nil {
		return errors.New("UNABLE TO RETRIEVE DETAILS FOR ACI PROFILE: " + d.Get("name").(string))
	}

	return setACIProfileResourceData(d, profile)

}

func resourceACIProfileDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	err := client.DeleteACIProfile(d.Get("uuid").(string))

	if err != nil {
		return errors.New(err.Error())
	}

	d.SetId("")
	return nil
}

func setACIProfileResourceData(d *schema.ResourceData, u *ccp.ACIProfile) error {

	if err := d.Set("uuid", u.UUID); err != nil {
		return errors.New("CANNOT SET UUID")
	}
	if err := d.Set("name", u.Name); err != nil {
		return errors.New("CANNOT SET NAME")
	}
	if err := d.Set("apic_hosts", u.APICHosts); err != nil {
		return errors.New("CANNOT SET APIC HOSTS")
	}
	if err := d.Set("apic_username", u.APICUsername); err != nil {
		return errors.New("CANNOT SET APIC USERNAME")
	}
	if err := d.Set("apic_password", u.APICPassword); err != nil {
		return errors.New("CANNOT SET APIC PASSWORD")
	}
	if err := d.Set("aci_vmm_domain_name", u.ACIVMMDomainName); err != nil {
		return errors.New("CANNOT SET ACI VMM DOMAIN NAME")
	}
	if err := d.Set("aci_infra_vlan_id", u.ACIInfraVLANID); err != nil {
		return errors.New("CANNOT SET ACI INFRA VLAN ID")
	}
	if err := d.Set("vrf_name", u.VRFName); err != nil {
		return errors.New("CANNOT SET VRF NAME")
	}
	if err := d.Set("l3_outside_policy_name", u.L3OutsidePolicyName); err != nil {
		return errors.New("CANNOT SET L3 OUTSIDE POLICY NAME")
	}
	if err := d.Set("l3_outside_network_name", u.L3OutsideNetworkName); err != nil {
		return errors.New("CANNOT SET L3 OUTSIDE NETWORK NAME")
	}
	if err := d.Set("aaep_name", u.AAEPName); err != nil {
		return errors.New("CANNOT SET AAEP NAME")
	}
	if err := d.Set("nameservers", u.Nameservers); err != nil {
		return errors.New("CANNOT SET NAMESERVERS")
	}
	if err := d.Set("control_plane_contract_name", u.ControlPlaneContractName); err != nil {
		return errors.New("CANNOT SET CONTROL PLANE CONTRACT NAME")
	}
	if err := d.Set("node_vlan_start", u.NodeVLANStart); err != nil {
		return errors.New("CANNOT SET NODE VLAN START")
	}
	if err := d.Set("node_vlan_end", u.NodeVLANEnd); err != nil {
		return errors.New("CANNOT SET NODE VLAN END")
	}
	if err := d.Set("pod_subnet_start", u.PodSubnetStart); err != nil {
		return errors.New("CANNOT SET POD SUBNET START")
	}
	if err := d.Set("service_subnet_start", u.ServiceSubnetStart); err != nil {
		return errors.New("CANNOT SET SERVICE SUBNET START")
	}
	if err := d.Set("multicast_range", u.MulticastRange); err != nil {
		return errors.New("CANNOT SET MULTICAST RANGE")
	}
	if err := d.Set("aci_tenant", u.ACITenant); err != nil {
		return errors.New("CANNOT SET ACI TENANT")
	}

	return nil
}
