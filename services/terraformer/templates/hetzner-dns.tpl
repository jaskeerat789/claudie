provider "google" {
    credentials = "${file("{{.Provider.Credentials}}")}"
    project = "{{.Project}}"
    alias = "dns"
}

data "google_dns_managed_zone" "hetzner-zone" {
  name = "{{.DNSZone}}"
}

{{- $clusterName := .ClusterName }}
{{- $clusterHash := .ClusterHash }}
{{- $hostnameHash := .HostnameHash }}
{{- range $nodepool := .NodePools}}

resource "google_dns_record_set" "{{$nodepool.Name}}-{{$clusterName}}" {
  provider = google.dns
  name = "{{ $hostnameHash }}.${data.google_dns_managed_zone.hetzner-zone.dns_name}"
  type = "A"
  ttl  = 300

  managed_zone = data.google_dns_managed_zone.hetzner-zone.name

  rrdatas = [
        for node in hcloud_server.{{$nodepool.Name}} :node.ipv4_address
    ]
}

output "{{$clusterName}}-{{$clusterHash}}" {
  value = { {{$clusterName}}-{{$clusterHash}} = google_dns_record_set.{{$nodepool.Name}}-{{$clusterName}}.name }
}
{{- end}}