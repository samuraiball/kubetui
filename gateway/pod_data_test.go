package gateway_test

var (
	podRowData = `{
    "apiVersion": "v1",
    "items": [
        {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "creationTimestamp": "2020-02-15T07:56:13Z",
                "generateName": "sample-deployment-7bb5fc6bc6-",
                "labels": {
                    "app": "sample-app",
                    "pod-template-hash": "7bb5fc6bc6"
                },
                "name": "sample-deployment-7bb5fc6bc6-9btmd",
                "namespace": "default",
                "ownerReferences": [
                    {
                        "apiVersion": "apps/v1",
                        "blockOwnerDeletion": true,
                        "controller": true,
                        "kind": "ReplicaSet",
                        "name": "sample-deployment-7bb5fc6bc6",
                        "uid": "f5f9e8fd-354a-4bbd-9cbb-556fb8d35f34"
                    }
                ],
                "resourceVersion": "518424",
                "selfLink": "/api/v1/namespaces/default/pods/sample-deployment-7bb5fc6bc6-9btmd",
                "uid": "9166f767-8e57-4488-853e-93aee3b5f15b"
            },
            "spec": {
                "containers": [
                    {
                        "image": "nginx:1.12",
                        "imagePullPolicy": "IfNotPresent",
                        "name": "nginx-container",
                        "ports": [
                            {
                                "containerPort": 80,
                                "protocol": "TCP"
                            }
                        ],
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "default-token-qzfsf",
                                "readOnly": true
                            }
                        ]
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "enableServiceLinks": true,
                "nodeName": "minikube",
                "priority": 0,
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "default",
                "serviceAccountName": "default",
                "terminationGracePeriodSeconds": 30,
                "tolerations": [
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/not-ready",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/unreachable",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    }
                ],
                "volumes": [
                    {
                        "name": "default-token-qzfsf",
                        "secret": {
                            "defaultMode": 420,
                            "secretName": "default-token-qzfsf"
                        }
                    }
                ]
            },
            "status": {
                "conditions": [
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2020-02-15T07:56:23Z",
                        "status": "True",
                        "type": "Initialized"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2020-02-15T07:56:25Z",
                        "status": "True",
                        "type": "Ready"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2020-02-15T07:56:25Z",
                        "status": "True",
                        "type": "ContainersReady"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2020-02-15T07:56:23Z",
                        "status": "True",
                        "type": "PodScheduled"
                    }
                ],
                "containerStatuses": [
                    {
                        "containerID": "docker://566326b9d78aa7ba9f71b30525b410d4fea7e08e3ed2de3c0791ec29e64de851",
                        "image": "nginx:1.12",
                        "imageID": "docker-pullable://nginx@sha256:72daaf46f11cc753c4eab981cbf869919bd1fee3d2170a2adeac12400f494728",
                        "lastState": {},
                        "name": "nginx-container",
                        "ready": true,
                        "restartCount": 0,
                        "state": {
                            "running": {
                                "startedAt": "2020-02-15T07:56:25Z"
                            }
                        }
                    }
                ],
                "hostIP": "10.27.40.254",
                "phase": "Running",
                "podIP": "172.17.0.17",
                "qosClass": "BestEffort",
                "startTime": "2020-02-15T07:56:23Z"
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "creationTimestamp": "2020-02-15T07:56:13Z",
                "generateName": "sample-deployment-7bb5fc6bc6-",
                "labels": {
                    "app": "sample-app",
                    "pod-template-hash": "7bb5fc6bc6"
                },
                "name": "sample-deployment-7bb5fc6bc6-hhd2l",
                "namespace": "default",
                "ownerReferences": [
                    {
                        "apiVersion": "apps/v1",
                        "blockOwnerDeletion": true,
                        "controller": true,
                        "kind": "ReplicaSet",
                        "name": "sample-deployment-7bb5fc6bc6",
                        "uid": "f5f9e8fd-354a-4bbd-9cbb-556fb8d35f34"
                    }
                ],
                "resourceVersion": "518430",
                "selfLink": "/api/v1/namespaces/default/pods/sample-deployment-7bb5fc6bc6-hhd2l",
                "uid": "583eeb19-7ded-4313-a80a-d84f9632e422"
            },
            "spec": {
                "containers": [
                    {
                        "image": "nginx:1.12",
                        "imagePullPolicy": "IfNotPresent",
                        "name": "nginx-container",
                        "ports": [
                            {
                                "containerPort": 80,
                                "protocol": "TCP"
                            }
                        ],
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "default-token-qzfsf",
                                "readOnly": true
                            }
                        ]
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "enableServiceLinks": true,
                "nodeName": "minikube",
                "priority": 0,
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "default",
                "serviceAccountName": "default",
                "terminationGracePeriodSeconds": 30,
                "tolerations": [
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/not-ready",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/unreachable",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    }
                ],
                "volumes": [
                    {
                        "name": "default-token-qzfsf",
                        "secret": {
                            "defaultMode": 420,
                            "secretName": "default-token-qzfsf"
                        }
                    }
                ]
            },
            "status": {
                "conditions": [
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2020-02-15T07:56:23Z",
                        "status": "True",
                        "type": "Initialized"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2020-02-15T07:56:26Z",
                        "status": "True",
                        "type": "Ready"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2020-02-15T07:56:26Z",
                        "status": "True",
                        "type": "ContainersReady"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2020-02-15T07:56:23Z",
                        "status": "True",
                        "type": "PodScheduled"
                    }
                ],
                "containerStatuses": [
                    {
                        "containerID": "docker://33a748e19652ac76f116a97a03fc470b957e107a63a836b67069a870b9caa698",
                        "image": "nginx:1.12",
                        "imageID": "docker-pullable://nginx@sha256:72daaf46f11cc753c4eab981cbf869919bd1fee3d2170a2adeac12400f494728",
                        "lastState": {},
                        "name": "nginx-container",
                        "ready": true,
                        "restartCount": 0,
                        "state": {
                            "running": {
                                "startedAt": "2020-02-15T07:56:25Z"
                            }
                        }
                    }
                ],
                "hostIP": "10.27.40.254",
                "phase": "Running",
                "podIP": "172.17.0.19",
                "qosClass": "BestEffort",
                "startTime": "2020-02-15T07:56:23Z"
            }
        }
    ],
    "kind": "List",
    "metadata": {
        "resourceVersion": "",
        "selfLink": ""
    }
}
`
)
