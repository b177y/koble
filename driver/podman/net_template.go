package podman

var NET = `
{
   "args": {
      "podman_labels": {
         "netkit": "true",
         "netkit:lab": "{{ .Lab }}",
         "netkit:name": "{{ .Name }}",
         "netkit:namespace": "{{ .Namespace }}"
      }
   },
   "cniVersion": "0.4.0",
   "name": "{{ .Fullname }}",
   "plugins": [
      {
         "type": "bridge",
         "bridge": "cni-podman2",
         "isGateway": false,
         "hairpinMode": false,
         "ipam": {}
      },
      {
         "type": "portmap",
         "capabilities": {
            "portMappings": true
         }
      },
      {
         "type": "firewall",
         "backend": ""
      },
      {
         "type": "tuning"
      }
   ]
}
`
