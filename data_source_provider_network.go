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

// Download the latest ccp-client-library to your Git source directory
// go get -u github.com/CiscoSE/ccp-client-library

import (
	"errors"

	"github.com/CiscoSE/ccp-client-library/ccp"
	"github.com/hashicorp/terraform/helper/schema"
)

func datasourceProviderNetwork() *schema.Resource {
	return &schema.Resource{
		Read: datasourceProviderNetworkRead,

		Schema: map[string]*schema.Schema{

			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"gateway_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pools": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"nameservers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"total_ips": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"free_ips": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func datasourceProviderNetworkRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	networkProviderSubnet, err := client.GetNetworkProviderSubnetByName(d.Get("name").(string))

	if err != nil {
		return errors.New("UNABLE TO RETRIEVE DETAILS FOR NETWORK PROVIDER: " + d.Get("name").(string))
	}

	uuid := *networkProviderSubnet.UUID

	d.SetId(uuid)

	return setProviderNetworkDataSourceData(d, networkProviderSubnet)

}

func setProviderNetworkDataSourceData(d *schema.ResourceData, u *ccp.NetworkProviderSubnet) error {

	if err := d.Set("uuid", u.UUID); err != nil {
		return errors.New("CANNOT SET NETWORK PROVIDER UUID")
	}
	if err := d.Set("name", u.Name); err != nil {
		return errors.New("CANNOT SET NETWORK PROVIDER NAME")
	}
	if err := d.Set("ip_version", u.IPVersion); err != nil {
		return errors.New("CANNOT SET NETWORK PROVIDER IP VERSION")
	}
	if err := d.Set("gateway_ip", u.GatewayIP); err != nil {
		return errors.New("CANNOT SET NETWORK PROVIDER GATEWAY IP")
	}
	if err := d.Set("cidr", u.CIDR); err != nil {
		return errors.New("CANNOT SET NETWORK PROVIDER CIDR")
	}
	if err := d.Set("pools", u.Pools); err != nil {
		return errors.New("CANNOT SET NETWORK PROVIDER POOLS")
	}
	if err := d.Set("network", u.Network); err != nil {
		return errors.New("CANNOT SET NETWORK PROVIDER NETWORK")
	}
	if err := d.Set("nameservers", u.Nameservers); err != nil {
		return errors.New("CANNOT SET NETWORK PROVIDER NAMESERVERS")
	}
	if err := d.Set("total_ips", u.TotalIPs); err != nil {
		return errors.New("CANNOT SET NETWORK PROVIDER TOTAL IPS")
	}
	if err := d.Set("free_ips", u.FreeIPs); err != nil {
		return errors.New("CANNOT SET NETWORK PROVIDER FREE IPS")
	}

	return nil
}
