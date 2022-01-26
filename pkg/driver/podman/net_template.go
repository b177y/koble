package podman

import "github.com/b177y/koble/pkg/driver"

type tmplNet struct {
	Net  driver.Network
	Opts driver.NetConfig
}

var INTERNAL_NET = `
{
   "args": {
      "podman_labels": {
         "koble": "true",
         "koble:name": "{{ .Net.Name }}",
         "koble:namespace": "{{ .Net.Namespace }}"
      }
   },
   "cniVersion": "0.4.0",
   "name": "{{ .Net.Id }}",
   "plugins": [
      {
         "type": "bridge",
         "bridge": "cni-podman1",
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

var EXTERNAL_NET = `
{
   "args": {
      "podman_labels": {
         "koble": "true",
         "koble:name": "{{ .Net.Name }}",
         "koble:namespace": "{{ .Net.Namespace }}"
      }
   },
   "cniVersion": "0.4.0",
   "name": "{{ .Net.Id }}",
   "plugins": [
      {
         "type": "bridge",
         "bridge": "cni-podman1",
         "isGateway": true,
         "ipMasq": true,
         "hairpinMode": true,
         "ipam": {
            "type": "host-local",
            "routes": [
               {
                  "dst": "0.0.0.0/0"
               }
            ],
            "ranges": [
               [
                  {
                     "subnet": "{{ .Opts.Subnet }}",
                     "gateway": "{{ .Opts.Gateway }}"
                  }
               ]
            ]
         }
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
