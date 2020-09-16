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

func resourceProviderVsphere() *schema.Resource {
	return &schema.Resource{
		Create: resourceProviderVsphereCreate,
		Read:   resourceProviderVsphereRead,
		Update: resourceProviderVsphereUpdate,
		Delete: resourceProviderVsphereDelete,

		Schema: map[string]*schema.Schema{

			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"insecure_skip_verify": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceProviderVsphereCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	newProviderClientConfig := ccp.ProviderClientConfig{

		Name:               ccp.String(d.Get("name").(string)),
		Description:        ccp.String(d.Get("description").(string)),
		Type:               ccp.String(d.Get("type").(string)),
		Address:            ccp.String(d.Get("address").(string)),
		Username:           ccp.String(d.Get("username").(string)),
		Password:           ccp.String(d.Get("password").(string)),
		Port:               ccp.Int64(d.Get("port").(int64)),
		InsecureSkipVerify: ccp.Bool(d.Get("insecure_skip_verify").(bool)),
	}

	providerClientConfig, err := client.AddVsphereProviderClientConfig(&newProviderClientConfig)

	if err != nil {
		return errors.New(err.Error())
	}

	uuid := *providerClientConfig.UUID

	d.SetId(uuid)

	return setProviderClientConfigResourceData(d, providerClientConfig)
}

func resourceProviderVsphereRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	providerClientConfig, err := client.GetInfraProviderByName(d.Get("name").(string))

	if err != nil {
		return errors.New("UNABLE TO RETRIEVE DETAILS FOR VSPHERE PROVIDER: " + d.Get("name").(string))
	}

	return setProviderClientConfigResourceData(d, providerClientConfig)

}

func resourceProviderVsphereUpdate(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	nameservers := []string{}
	for _, server := range d.Get("nameservers").([]interface{}) {
		nameservers = append(nameservers, server.(string))
	}

	newProviderClientConfig := ccp.ProviderClientConfig{

		Name:               ccp.String(d.Get("name").(string)),
		Description:        ccp.String(d.Get("description").(string)),
		Type:               ccp.String(d.Get("type").(string)),
		Address:            ccp.String(d.Get("address").(string)),
		Username:           ccp.String(d.Get("username").(string)),
		Password:           ccp.String(d.Get("password").(string)),
		Port:               ccp.Int64(d.Get("port").(int64)),
		InsecureSkipVerify: ccp.Bool(d.Get("insecure_skip_verify").(bool)),
	}

	provider, err := client.PatchProviderClientConfig(&newProviderClientConfig, d.Get("uuid").(string))

	if err != nil {
		return errors.New(err.Error())
	}

	provider, err = client.GetInfraProviderByName(d.Get("name").(string))

	if err != nil {
		return errors.New("UNABLE TO RETRIEVE DETAILS FOR VSPHERE PROVIDER: " + d.Get("name").(string))
	}

	return setProviderClientConfigResourceData(d, provider)

}

func resourceProviderVsphereDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	err := client.DeleteProviderClientConfig(d.Get("uuid").(string))

	if err != nil {
		return errors.New(err.Error())
	}

	d.SetId("")
	return nil
}

func setProviderClientConfigResourceData(d *schema.ResourceData, u *ccp.ProviderClientConfig) error {

	if err := d.Set("uuid", u.UUID); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER UUID")
	}
	if err := d.Set("name", u.Name); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER NAME")
	}
	if err := d.Set("description", u.Description); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER DESCRIPTION")
	}
	if err := d.Set("type", u.Type); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER TYPE")
	}
	if err := d.Set("address", u.Address); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER ADDRESS")
	}
	if err := d.Set("username", u.Username); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER USERNAME")
	}
	if err := d.Set("password", u.Password); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER PASSWORD")
	}
	if err := d.Set("port", u.Port); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER  PORT")
	}
	if err := d.Set("insecure_skip_verify", u.InsecureSkipVerify); err != nil {
		return errors.New("CANNOT SET VSPHERE PROVIDER INSECURE SKIP VERIFY CONDITION")
	}

	return nil
}
