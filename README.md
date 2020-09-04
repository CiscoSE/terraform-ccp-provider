# Cisco Container Platform Provider Plugin for Terraform

This is a provider plugin for Terraform which allows Terraform to interact with various Cisco Container Platform (CCP) resources. 

It is currently a __Proof of Concept__ and has been developed and tested against Cisco Container Platform 6.1 with Terraform version v0.13.0

Table of Contents
=================

  * [CCP Terraform Provider Plugin](#ccp-terraform-provider-plugin)
      * [Quick Start Calico](#quick-start-calico)
      * [Quick Start ACI CNI](#quick-start-aci-cni)
      * [Building and Installation](#building-and-installation)
      * [Guidelines and Limitations](#guidelines-and-limitations)
      * [License](#license)
 
## Quick Start Calico

The following example config can be found in the main.tf.dummydata-calico file. Remove the `.dummydata-calico` extension and replace the example config with your specific environment details.

```golang

/*
    These are the credentials used to login to CCP
*/

variable "username" {
    type = "string"
    default="my_ccp_admin_account"
}

variable "password" {
    type = "string"
    default="my_ccp_password"
}

variable "base_url" {
    type = "string"
    default="https://my_ccp_url:ccp_port"
}

provider "ccp" {
    username = "${var.username}"
    password = "${var.password}"
    base_url = "${var.base_url}"
}

/*
    This will create a new local user within CCP
*/
resource "ccp_user" "user" {
    firstname       = "Terrafom"
    lastname        = "Plugin"
    password        = "myPassword"
    username        = "builtByTerraform"
    role            = "Administrator"
}


/*
    This will create a new cluster within CCP with Calico as the CNI
*/

resource "ccp_cluster" "cluster" {
  provider_client_config_uuid = "1abc2-1abc2-1abc2-1abc2" //Calico provider
  name                        = "builtbyterraform"
  kubernetes_version          = "1.16.3"
  loadbalancer_ip_num         = 3
  type                        = "vsphere"
  ip_allocation_method = "ccpnet"
  subnet_uuid            = "d7a6f267-8545-4875-85c5-bf5b7f46b4f0" 
  infra {
      datacenter    = "vcenter-datacenter-name"
      cluster       = "vcenter-cluster-name"
      datastore     = "vcenter-datastore-name"
      resource_pool = " "
      networks = ["vcenter-network-name"] 

  }
  master_node_pool {
         name = "master-group"
         size = 1
         gpus=[]
         vcpus    = 2
         memory   = 16384
         template = "ccp-tenant-image-1.16.3-ubuntu18-6.1.1"
         ssh_user = "admin"
         ssh_key = "ssh-ed25519 AAAAC3fsdhSDFSDFbildsfDFSSDFbsdfFSDFSD"
         kubernetes_version = "1.16.3"

   }
  worker_node_pools     {
         name = "node-group"
         size = 4
         gpus=[]
         vcpus    = 2
         memory   = 16384
         template = "ccp-tenant-image-1.16.3-ubuntu18-6.1.1"
         ssh_user = "admin"
         ssh_key = "ssh-ed25519 AAAAC3fsdhSDFSDFbildsfDFSSDFbsdfFSDFSD"
         kubernetes_version = "1.16.3"
   }
   network_plugin {
      name ="calico"
      details {
        pod_cidr = "192.168.0.0/16"
      }
   }
}

```

## Quick Start ACI CNI

The following example config can be found in the main.tf.dummydata-aci file. Remove the `.dummydata-aci` extension and replace the example config with your specific environment details.

```golang

/*
    These are the credentials used to login to CCP
*/

variable "username" {
    type = "string"
    default="my_ccp_admin_account"
}

variable "password" {
    type = "string"
    default="my_ccp_password"
}

variable "base_url" {
    type = "string"
    default="https://my_ccp_url:ccp_port"
}

provider "ccp" {
    username = "${var.username}"
    password = "${var.password}"
    base_url = "${var.base_url}"
}

/*
    This will create a new local user within CCP
*/
resource "ccp_user" "user" {
    firstname       = "Terrafom"
    lastname        = "Plugin"
    password        = "myPassword"
    username        = "builtByTerraform"
    role            = "Administrator"
}

/*
    This will create a new ACI Profile to use with the new cluster.
*/

resource "ccp_aci_profile" "aci_profile" {
  
  name="builtbyterraform"
	apic_hosts= "10.1.1.1"
	apic_username= "admin"
	apic_password= "password"
	aci_vmm_domain_name= "DM_VMM"
	aci_infra_vlan_id= 4093
	vrf_name= "default"
	l3_outside_policy_name= "L3OUT_common"
	l3_outside_network_name= "eEPG_common"
	aaep_name= "AEP_ALL"
	nameservers= ["8.8.8.8"]
	control_plane_contract_name= "ANY-ANY"
	node_vlan_start= 3300
	node_vlan_end= 3400
	pod_subnet_start= "100.65.0.1/16"
	service_subnet_start= "100.100.0.1/16"
	multicast_range= "225.32.0.0/16"
	aci_tenant= "common"

}

/*
    This will create a new cluster within CCP using the ACI CNI
*/

resource "ccp_cluster" "cluster" {
  provider_client_config_uuid = "8b27074e-9ed8-4934-88ec-34gf43dgf" // ACI CNI provider
  name                        = "builtbyterraform"
  kubernetes_version          = "1.16.3"
  loadbalancer_ip_num         = 1 
  type                        = "vsphere"
  ip_allocation_method = "ccpnet"
  infra {
      datacenter    = "vcenter-datacenter-name"
      cluster       = "vcenter-cluster-name"
      datastore     = "vcenter-datastore-name"
      resource_pool = " "
      networks = [" "]
  }
  master_node_pool {
         name = "master-group"
         size = 1
         gpus=[]
         vcpus    = 2
         memory   = 16384
         template = "ccp-tenant-image-1.16.3-ubuntu18-6.1.1"
         ssh_user = "admin"
         ssh_key = "ssh-ed25519 AAAAC3fsdhSDFSDFbildsfDFSSDFbsdfFSDFSD"
         kubernetes_version = "1.16.3"

   }
  worker_node_pools     {
         name = "node-group"
         size = 4
         gpus=[]
         vcpus    = 2
         memory   = 16384
         template = "ccp-tenant-image-1.16.3-ubuntu18-6.1.1"
         ssh_user = "admin"
         ssh_key = "ssh-ed25519 AAAAC3fsdhSDFSDFbildsfDFSSDFbsdfFSDFSD"
         kubernetes_version = "1.16.3"
   }
   network_plugin {
      name="contiv-aci"
      details {
      }
   }

    routable_cidr = "10.140.2.0/24" 
    aci_profile_uuid = ccp_aci_profile.aci_profile.uuid

    depends_on = [ccp_aci_profile.aci_profile]
}

```

## Building and Installation

1. Clone provider repo to local machine.

`git clone https://github.com/conmurphy/terraform-provider-ccp.git`

2. From within the newly cloned directory, build the binary

`go build -o terraform-provider-ccp_v0.1.0`

3. Copy binary to local Terraform plugin directory.

As per the following document, third-party plugins should usually be installed in the user plugins directory, which is located at `~/.terraform.d/plugins`. 

`~/.terraform.d/plugins/<OS>_<ARCH>` or `%APPDATA%\terraform.d\plugins\<OS>_<ARCH>`	The user plugins directory, with explicit OS and architecture.

https://www.terraform.io/docs/extend/how-terraform-works.html#discovery

`cp terraform-provider-ccp_v0.1.0 ~/.terraform.d/plugins/cisco.com/ccp/ccp/0.1.0/darwin_amd64`

4. Initialise Terraform

` terraform init`

5. Ready to start planning and applying.

## Guidelines and Limitations


* Scaling: 
  * loadbalancer_ip_num can be increased or decreased
  * worker_node_pool.size can be increased or decreased
* Supports single worker node pool
* Has not been tested with GPUs
* Has not been tested with resource pools
* `networks` is required for the ACI CNI config however it can be left with whitespace as per the example config
* When increasing worker size the response will return straight away. Behind the scenes CCP will be adding a new worker node. This is the current behaviour of the API. Therefore the TFState file won’t contain the new worker node details. Once the node is ready you can run the `terraform refresh` command to refresh the state. If you run the refresh command before the node has completed you should see the phase as `creating`.

## License

This project is licensed to you under the terms of the [Cisco Sample Code License](./LICENSE).