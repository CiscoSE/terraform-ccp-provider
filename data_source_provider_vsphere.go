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

func datasourceProviderVsphere() *schema.Resource {
	return &schema.Resource{
		Read: datasourceProviderVsphereRead,

		Schema: map[string]*schema.Schema{

			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"insecure_skip_verify": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceProviderVsphereRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	providerClientConfig, err := client.GetInfraProviderByName(d.Get("name").(string))

	if err != nil {
		return errors.New("UNABLE TO RETRIEVE DETAILS FOR VSPHERE PROVIDER: " + d.Get("name").(string))
	}

	uuid := *providerClientConfig.UUID

	d.SetId(uuid)

	return setProviderClientConfigDataSourceData(d, providerClientConfig)

}

func setProviderClientConfigDataSourceData(d *schema.ResourceData, u *ccp.ProviderClientConfig) error {

	if err := d.Set("uuid", u.UUID); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER  UUID")
	}
	if err := d.Set("name", u.Name); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER  NAME")
	}
	if err := d.Set("address", u.Address); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER  ADDRESS")
	}
	if err := d.Set("port", u.Port); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER  PORT")
	}
	if err := d.Set("username", u.Username); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER  USERNAME")
	}
	if err := d.Set("password", u.Password); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER PASSWORD")
	}
	if err := d.Set("insecure_skip_verify", u.InsecureSkipVerify); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER INSECURE SKIP VERIFY")
	}
	if err := d.Set("type", u.Type); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER TYPE")
	}
	if err := d.Set("description", u.Description); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER DESCRIPTION")
	}

	return nil
}
