package main

import (
	cstorage "github.com/containers/storage"
	"github.com/containers/storage/pkg/reexec"
	"github.com/cri-o/cri-o/pkg/config"
	"k8s.io/klog/v2"
)

func main() {
	if reexec.Init() {
		return
	}

	configIface, err := config.DefaultConfig()
	klog.Infof("Configuration is: %v, %v, %v, %v", configIface.Root, configIface.RunRoot, configIface.Storage, configIface.StorageOptions)
	if err != nil {
		klog.Fatal(err)
	}
	if configIface == nil {
		klog.Fatal("provided config is nil")
	}

	klog.Info("Getting Store")
	store, err := configIface.GetStore()
	if err != nil {
		klog.Fatal(err)
	}
	config := configIface.GetData()

	if config == nil {
		klog.Fatal("cannot create container server: interface is nil")
	}

	klog.Info("Running store.Check")
	checkOptions := cstorage.CheckEverything()
	report, err := store.Check(checkOptions)
	if err != nil {
		klog.Fatal(err)
	}

	klog.Info("Running store.Repair")
	options := cstorage.RepairOptions{
		RemoveContainers: true,
	}
	if errs := store.Repair(report, &options); len(errs) > 0 {
		klog.Fatal(err)
	}
}
