// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// RBDPersistentVolumeSource Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
//
// swagger:model RBDPersistentVolumeSource
type RBDPersistentVolumeSource struct {

	// fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	FSType string `json:"fsType,omitempty"`

	// image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// Required: true
	Image *string `json:"image"`

	// keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Keyring string `json:"keyring,omitempty"`

	// monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// Required: true
	Monitors []string `json:"monitors"`

	// pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Pool string `json:"pool,omitempty"`

	// readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	ReadOnly bool `json:"readOnly,omitempty"`

	// secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	SecretRef *SecretReference `json:"secretRef,omitempty"`

	// user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	User string `json:"user,omitempty"`
}
