terraform {
  required_providers {
    {{ if .Hetzner }}
    hcloud = {
      source = "hetznercloud/hcloud"
      version = "1.35.1"
    }
    {{ end }}
    {{ if .Gcp }}
    google = {
      source = "hashicorp/google"
      version = "4.31.0"
    }
    {{ end }}
    {{if .Aws }}
    aws = {
      source = "hashicorp/aws"
      version = "4.31.0"
    }
    {{ end }}
    {{if .Oci }}
    oci = {
      source = "oracle/oci"
      version = "4.94.0"
    }
    {{ end }}
    {{if .Azure }}
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.26.0"
    }
    {{ end }}
    {{ if .Cloudflare }}
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 3.32.0"
    }
    {{ end }}
    {{ if .HetznerDNS }}
    hetznerdns = {
      source = "timohirt/hetznerdns"
      version = "2.2.0"
    }
    {{ end }}
  }
}