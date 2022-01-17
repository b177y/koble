package podman

var NET = `
{
   "args": {
      "podman_labels": {
         "koble": "true",
         "koble:name": "{{ .Name }}",
         "koble:namespace": "{{ .Namespace }}"
      }
   },
   "cniVersion": "0.4.0",
   "name": "{{ .Id }}",
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
