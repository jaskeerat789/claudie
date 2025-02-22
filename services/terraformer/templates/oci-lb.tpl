{{- $clusterName := .ClusterName}}
{{- $clusterHash := .ClusterHash}}
{{$index :=  0}}

variable "default_compartment_id" {
  type = string
  default = "{{(index .NodePools 0).Provider.OciCompartmentOcid}}"
}
{{ range $i, $region := .Regions }}
provider "oci" {
  tenancy_ocid = "{{(index $.NodePools 0).Provider.OciTenancyOcid}}"
  user_ocid = "{{(index $.NodePools 0).Provider.OciUserOcid}}"
  fingerprint = "{{(index $.NodePools 0).Provider.OciFingerprint}}"
  private_key_path = "{{(index $.NodePools 0).Provider.SpecName}}" 
  region = "{{ $region }}"
  alias = "lb-nodepool-{{ $region }}"
}

resource "oci_core_vcn" "claudie_vcn_{{ $region }}" {
  provider = oci.lb-nodepool-{{ $region }}
  compartment_id = var.default_compartment_id
  display_name = "{{ $clusterName }}-{{ $clusterHash }}-vcn"
  cidr_blocks = ["10.0.0.0/16"]
}

resource "oci_core_internet_gateway" "claudie_gateway-{{ $region }}" {
  provider = oci.lb-nodepool-{{ $region }}
  compartment_id = var.default_compartment_id
  display_name   = "{{ $clusterName }}-{{ $clusterHash }}-gateway"
  vcn_id         = oci_core_vcn.claudie_vcn_{{ $region }}.id
  enabled = true
}  

resource "oci_core_default_security_list" "claudie_security_rules-{{ $region }}" {
  provider = oci.lb-nodepool-{{ $region }}
  manage_default_resource_id = oci_core_vcn.claudie_vcn_{{ $region }}.default_security_list_id
  display_name   = "{{ $clusterName }}-{{ $clusterHash }}_security_rules"

  egress_security_rules {  
    destination = "0.0.0.0/0"
    protocol    = "all"
    description = "Allow all egress"
  }

  ingress_security_rules {
    protocol = "1"
    source   = "0.0.0.0/0"
    description = "Allow all ICMP"
  }

  ingress_security_rules {
    protocol    = "6"
    source      = "0.0.0.0/0"
    tcp_options {
      min = "22"
      max = "22"
    }
    description = "Allow SSH connections"
  }

  {{range $role := index $.Metadata "roles"}}
  ingress_security_rules {
    protocol    = "{{ protocolToOCIProtocolNumber $role.Protocol}}"
    source      = "0.0.0.0/0"
    tcp_options {
      max = "{{ $role.Port }}"
      min = "{{ $role.Port }}"
    }
    description = "LoadBalancer port defined in the manifest"
  }
  {{end}}

  ingress_security_rules {
    protocol    = "17"
    source      = "0.0.0.0/0"
    udp_options {
      max = "51820"
      min = "51820"
    }
    description = "Allow Wireguard VPN port"
  }
}

resource "oci_core_default_route_table" "claudie_routes-{{ $region }}" {
  provider = oci.lb-nodepool-{{ $region }}
  manage_default_resource_id = oci_core_vcn.claudie_vcn_{{ $region }}.default_route_table_id

  route_rules {
    destination       = "0.0.0.0/0"
    network_entity_id = oci_core_internet_gateway.claudie_gateway-{{ $region }}.id
    destination_type  = "CIDR_BLOCK"
  }
}
{{- end }}

{{ range $i, $nodepool := .NodePools }}
resource "oci_core_subnet" "{{ $nodepool.Name }}-subnet" {
  provider = oci.lb-nodepool-{{ $nodepool.Region }}
  vcn_id = oci_core_vcn.claudie_vcn_{{ $nodepool.Region }}.id
  cidr_block = "{{getCIDR "10.0.0.0/24" 2 $i}}"
  compartment_id = var.default_compartment_id
  display_name = "{{ $clusterName }}-{{ $clusterHash }}-subnet"
  security_list_ids = [oci_core_vcn.claudie_vcn_{{ $nodepool.Region }}.default_security_list_id]
  route_table_id = oci_core_vcn.claudie_vcn_{{ $nodepool.Region }}.default_route_table_id
  dhcp_options_id   = oci_core_vcn.claudie_vcn_{{ $nodepool.Region }}.default_dhcp_options_id
  availability_domain = "{{ $nodepool.Zone }}"
}

resource "oci_core_instance" "{{ $nodepool.Name }}" {
  provider = oci.lb-nodepool-{{ $nodepool.Region }}
  compartment_id = var.default_compartment_id
  count = {{ $nodepool.Count }}
  availability_domain = "{{ $nodepool.Zone }}"
  shape = "{{ $nodepool.ServerType }}"
  display_name = "{{ $clusterName }}-{{ $clusterHash }}-{{ $nodepool.Name }}-${count.index + 1}"

  metadata = {
      ssh_authorized_keys = file("./public.pem")
      user_data = base64encode(<<EOF
      #cloud-config
      runcmd:
        # Allow Claudie to ssh as root
        - sed -n 's/^.*ssh-rsa/ssh-rsa/p' /root/.ssh/authorized_keys > /root/.ssh/temp
        - cat /root/.ssh/temp > /root/.ssh/authorized_keys
        - rm /root/.ssh/temp
        - echo 'PermitRootLogin without-password' >> /etc/ssh/sshd_config && echo 'PubkeyAuthentication yes' >> /etc/ssh/sshd_config && echo "PubkeyAcceptedKeyTypes=+ssh-rsa" >> sshd_config && service sshd restart
        # Disable iptables
        # Accept all traffic to avoid ssh lockdown via iptables firewall rules
        - iptables -P INPUT ACCEPT
        - iptables -P FORWARD ACCEPT
        - iptables -P OUTPUT ACCEPT
        # Flush and cleanup
        - iptables -F
        - iptables -X
        - iptables -Z
        # Make changes persistent
        - netfilter-persistent save
      EOF
      )
  }
  
  source_details {
    source_id   = "{{ $nodepool.Image }}"
    source_type = "image"
    boot_volume_size_in_gbs = "{{ $nodepool.DiskSize }}"
  }
  create_vnic_details {
    assign_public_ip = true
    subnet_id = oci_core_subnet.{{ $nodepool.Name }}-subnet.id
  }
}

output "{{$nodepool.Name}}" {
  value = {
    for node in oci_core_instance.{{$nodepool.Name}}:
    node.display_name => node.public_ip
  }
}
{{end}}
