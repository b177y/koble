package podman

var NET = `
{
   "args": {
      "podman_labels": {
         "koble": "true",
         "koble:lab": "{{ .Lab }}",
         "koble:name": "{{ .Name }}",
         "koble:namespace": "{{ .Namespace }}"
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
