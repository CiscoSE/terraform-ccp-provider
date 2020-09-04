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
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CCP_USERNAME", nil),
				Description: "Username used to access Cisco Container Platform",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CCP_PASSWORD", nil),
				Description: "Password used to access Cisco Container Platform",
			},
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CCP_URL", nil),
				Description: "URL to the Cisco Container Platform",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ccp_user":        resourceUser(),
			"ccp_cluster":     resourceCluster(),
			"ccp_aci_profile": resourceACIProfile(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Base_url: d.Get("base_url").(string),
	}

	return config.Client(), nil
}
