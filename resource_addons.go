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

func resourceAddons() *schema.Resource {
	return &schema.Resource{
		Create: resourceAddonsCreate,
		Read:   resourceAddonsRead,
		Update: resourceAddonsUpdate,
		Delete: resourceAddonsDelete,

		Schema: map[string]*schema.Schema{

			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"kubernetes_dashboard": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"monitoring": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"logging": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"istio": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"harbor": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"kubeflow": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"hx_csi": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"addon_details": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"helm_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceAddonsCreate(d *schema.ResourceData, m interface{}) error {

	//Check no conflicts exist - CCP does not allow Kubeflow or Harbor to be enabled when Istio is enabled

	if d.Get("istio").(bool) && (d.Get("harbor").(bool) || d.Get("kubeflow").(bool)) {
		return errors.New("Kubeflow and Harbor cannot be enabled when Istio is enabled")
	}

	client := m.(*ccp.Client)

	if d.Get("kubernetes_dashboard").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "kubernetes-dashboard")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "kubernetes-dashboard")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	if d.Get("monitoring").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "monitoring")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "monitoring")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	if d.Get("logging").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "logging")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "logging")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	if d.Get("kubeflow").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "kubeflow")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "kubeflow")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	if d.Get("istio").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "istio")
		if err != nil {
			return errors.New(err.Error())
		}

	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "istio")

		if err != nil {
			return errors.New(err.Error())
		}

	}

	if d.Get("harbor").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "harbor")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "harbor")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	if d.Get("hx_csi").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "hx-csi")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "hx-csi")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	clusterAddons, err := client.GetClusterInstalledAddons(d.Get("uuid").(string))

	if err != nil {
		return errors.New(err.Error())
	}

	d.SetId(d.Get("uuid").(string))

	return setAddonsResourceData(d, clusterAddons, d.Get("uuid").(string))

}

func resourceAddonsRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	addons, err := client.GetClusterInstalledAddons(d.Get("uuid").(string))

	if err != nil {
		return errors.New("UNABLE TO RETRIEVE DETAILS FOR CLUSTER ADDONS: " + d.Get("uuid").(string))
	}

	return setAddonsResourceData(d, addons, d.Get("uuid").(string))

}

func resourceAddonsUpdate(d *schema.ResourceData, m interface{}) error {

	//Check no conflicts exist - CCP does not allow Kubeflow or Harbor to be enabled when Istio is enabled

	if d.Get("istio").(bool) && (d.Get("harbor").(bool) || d.Get("kubeflow").(bool)) {
		return errors.New("Kubeflow and Harbor cannot be enabled when Istio is enabled")
	}

	client := m.(*ccp.Client)

	if d.Get("kubernetes_dashboard").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "kubernetes-dashboard")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "kubernetes-dashboard")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	if d.Get("monitoring").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "monitoring")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "monitoring")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	if d.Get("logging").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "logging")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "logging")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	if d.Get("kubeflow").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "kubeflow")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "kubeflow")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	if d.Get("istio").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "istio")
		if err != nil {
			return errors.New(err.Error())
		}

	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "istio")

		if err != nil {
			return errors.New(err.Error())
		}

	}

	if d.Get("harbor").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "harbor")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "harbor")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	if d.Get("hx_csi").(bool) {
		err := client.InstallAddon(d.Get("uuid").(string), "hx-csi")

		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		err := client.DeleteAddon(d.Get("uuid").(string), "hx-csi")

		if err != nil {
			return errors.New(err.Error())
		}
	}

	clusterAddons, err := client.GetClusterInstalledAddons(d.Get("uuid").(string))

	if err != nil {
		return errors.New(err.Error())
	}

	return setAddonsResourceData(d, clusterAddons, d.Get("uuid").(string))

}

func resourceAddonsDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	err := client.DeleteAddon(d.Get("uuid").(string), "kubernetes-dashboard")

	if err != nil {
		return errors.New("Can't delete addon - Kubernetes Dashboard. " + err.Error())
	}

	err = client.DeleteAddon(d.Get("uuid").(string), "monitoring")

	if err != nil {
		return errors.New("Can't delete addon - Monitoring. " + err.Error())
	}

	err = client.DeleteAddon(d.Get("uuid").(string), "logging")

	if err != nil {
		return errors.New("Can't delete addon - Logging. " + err.Error())
	}

	err = client.DeleteAddon(d.Get("uuid").(string), "kubeflow")

	if err != nil {
		return errors.New("Can't delete addon - Kubeflow. " + err.Error())
	}

	err = client.DeleteAddon(d.Get("uuid").(string), "istio")

	if err != nil {
		return errors.New("Can't delete addon - ISTIO. " + err.Error())
	}

	err = client.DeleteAddon(d.Get("uuid").(string), "harbor")

	if err != nil {
		return errors.New("Can't delete addon - Harbor. " + err.Error())
	}

	err = client.DeleteAddon(d.Get("uuid").(string), "hx-csi")

	if err != nil {
		return errors.New("Can't delete addon - HX-CSI. " + err.Error())
	}

	d.SetId("")

	return nil
}

func setAddonsResourceData(d *schema.ResourceData, u *ccp.ClusterInstalledAddons, uuid string) error {

	if err := d.Set("uuid", uuid); err != nil {
		return errors.New("CANNOT SET CLUSTER ADDON UUID")
	}

	// The boolean values toggle the installation or removal of the addons. Still wanted to include the
	// status of each installed addon for reference which is the function of this next piece of code

	addonDetailsOut := make([]interface{}, 0, 0)
	addonDetailsIn := make(map[string]interface{})

	results := &u.Results

	for _, addons := range *results {

		addonDetailsIn = make(map[string]interface{})

		addonDetailsIn["helm_status"] = addons.AddonStatus.HelmStatus
		addonDetailsIn["status"] = addons.AddonStatus.Status
		addonDetailsIn["name"] = addons.Name
		addonDetailsIn["namespace"] = addons.Namespace
		addonDetailsIn["description"] = addons.Description
		addonDetailsIn["display_name"] = addons.DisplayName

		addonDetailsOut = append(addonDetailsOut, addonDetailsIn)

	}

	if err := d.Set("addon_details", addonDetailsOut); err != nil {
		return errors.New("CANNOT SET ADDON DETAILS")
	}

	return nil
}
