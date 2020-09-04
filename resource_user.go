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

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"firstname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lastname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"disable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {

	/*username := d.Get("username").(string)
	d.SetId(username)

	client := m.(*ccp.Client)

	newUser := ccp.User{

		FirstName: ccp.String(d.Get("firstname").(string)),
		LastName:  ccp.String(d.Get("lastname").(string)),
		Password:  ccp.String(d.Get("password").(string)),
		Username:  ccp.String(d.Get("username").(string)),
		Disable:   ccp.Bool(d.Get("disable").(bool)),
		Role:      ccp.String(d.Get("role").(string)),
	}

	user, err := client.AddUser(&newUser)

	if err != nil {
		return errors.New(err.Error())
	}

	return setUserResourceData(d, user)*/

	return nil
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {

	/*
		client := m.(*ccp.Client)
		user, err := client.GetUser(d.Get("username").(string))

		if err != nil {
			return errors.New("UNABLE TO RETRIEVE DETAILS FOR USER: " + d.Get("username").(string))
		}

		return setUserResourceData(d, user)
	*/

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {

	/*client := m.(*ccp.Client)

	newUser := ccp.User{
		Username:  ccp.String(d.Get("username").(string)),
		FirstName: ccp.String(d.Get("firstname").(string)),
		LastName:  ccp.String(d.Get("lastname").(string)),
		Password:  ccp.String(d.Get("password").(string)),
		Disable:   ccp.Bool(d.Get("disable").(bool)),
		Role:      ccp.String(d.Get("role").(string)),
	}

	user, err := client.PatchUser(&newUser)

	if err != nil {
		return errors.New(err.Error())
	}

	user, err = client.GetUser(d.Get("username").(string))

	if err != nil {
		return errors.New("UNABLE TO RETRIEVE DETAILS FOR USER: " + d.Get("username").(string))
	}

	return setUserResourceData(d, user)
	*/
	return nil
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {

	/*client := m.(*ccp.Client)

	err := client.DeleteUser(d.Get("username").(string))

	if err != nil {
		return errors.New(err.Error())
	}

	d.SetId("")
	return nil
	*/
	return nil
}

/*func setUserResourceData(d *schema.ResourceData, u *ccp.User) error {

	if err := d.Set("firstname", u.FirstName); err != nil {
		return errors.New("CANNOT SET FIRST NAME")
	}
	if err := d.Set("lastname", u.LastName); err != nil {
		return errors.New("CANNOT SET LAST NAME")
	}
	if err := d.Set("username", u.Username); err != nil {
		return errors.New("CANNOT SET USERNAME")
	}
	if err := d.Set("password", u.Password); err != nil {
		return errors.New("CANNOT SET PASSWORD")
	}
	if err := d.Set("disable", u.Disable); err != nil {
		return errors.New("CANNOT SET DISABLE FIELD")
	}
	if err := d.Set("role", u.Role); err != nil {
		return errors.New("CANNOT SET ROLE")
	}
	return nil
}*/
