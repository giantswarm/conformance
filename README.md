# conformance

TODO

Build e2e binary to use subcommands

```bash
go test -v -c -o e2e
```

```bash
./e2e create-cluster --provider aws --cluster-name my-cluster --owner giantswarm --pod-cidr 10.0.0.0/24 --release 200.0.0 --credential my-credential --master-zone eu-central-1a,eu-central-1b,eu-central-1c --region eu-central-1
```

Outside of the cluster add flag `--kubeconfig` to pass Kubernetes configuration

Currently does nothing, except printing CR's
```bash
INFO[0000] Running e2e tests for provider 'aws'
INFO[0000] Dumping CR's:
INFO[0000] {
	"kind": "Cluster",
	"apiVersion": "cluster.x-k8s.io/v1alpha2",
	"metadata": {
		"name": "u7d9g",
		"namespace": "default",
		"creationTimestamp": null,
		"labels": {
			"cluster-operator.giantswarm.io/version": "",
			"giantswarm.io/cluster": "u7d9g",
			"giantswarm.io/organization": "giantswarm",
			"release.giantswarm.io/version": "200.0.0"
		}
	},
	"spec": {
		"infrastructureRef": {
			"kind": "AWSCluster",
			"namespace": "default",
			"name": "u7d9g",
			"apiVersion": "infrastructure.giantswarm.io/v1alpha2"
		}
	},
	"status": {
		"infrastructureReady": false,
		"controlPlaneInitialized": false
	}
}
INFO[0000] {
	"kind": "AWSCluster",
	"apiVersion": "infrastructure.giantswarm.io/v1alpha2",
	"metadata": {
		"name": "u7d9g",
		"namespace": "default",
		"creationTimestamp": null,
		"labels": {
			"aws-operator.giantswarm.io/version": "",
			"giantswarm.io/cluster": "u7d9g",
			"giantswarm.io/organization": "giantswarm",
			"release.giantswarm.io/version": "200.0.0"
		}
	},
	"spec": {
		"cluster": {
			"description": "my-cluster",
			"dns": {
				"domain": "eu-central-1a,eu-central-1b,eu-central-1c"
			},
			"kubeProxy": {},
			"oidc": {
				"claims": {}
			}
		},
		"provider": {
			"credentialSecret": {
				"name": "my-credential",
				"namespace": "giantswarm"
			},
			"master": {
				"availabilityZone": "",
				"instanceType": "m5.xlarge"
			},
			"pods": {
				"cidrBlock": "10.0.0.0/24",
				"externalSNAT": false
			},
			"region": "eu-central-1"
		}
	},
	"status": {
		"cluster": {},
		"provider": {
			"network": {}
		}
	}
}
INFO[0000] {
	"kind": "G8sControlPlane",
	"apiVersion": "infrastructure.giantswarm.io/v1alpha2",
	"metadata": {
		"name": "jjrc8",
		"namespace": "default",
		"creationTimestamp": null,
		"labels": {
			"cluster-operator.giantswarm.io/version": "",
			"giantswarm.io/cluster": "u7d9g",
			"giantswarm.io/control-plane": "jjrc8",
			"giantswarm.io/organization": "giantswarm",
			"release.giantswarm.io/version": "200.0.0"
		}
	},
	"spec": {
		"replicas": 1,
		"infrastructureRef": {
			"kind": "AWSControlPlane",
			"namespace": "default",
			"name": "jjrc8",
			"apiVersion": "infrastructure.giantswarm.io/v1alpha2"
		}
	},
	"status": {}
}
INFO[0000] {
	"kind": "AWSControlPlane",
	"apiVersion": "infrastructure.giantswarm.io/v1alpha2",
	"metadata": {
		"name": "jjrc8",
		"namespace": "default",
		"creationTimestamp": null,
		"labels": {
			"aws-operator.giantswarm.io/version": "",
			"giantswarm.io/cluster": "u7d9g",
			"giantswarm.io/control-plane": "jjrc8",
			"giantswarm.io/organization": "giantswarm",
			"release.giantswarm.io/version": "200.0.0"
		}
	},
	"spec": {
		"availabilityZones": [
			""
		],
		"instanceType": "m5.xlarge"
	}
}
PASS
ok  	github.com/giantswarm/conformance	0.458s
```
