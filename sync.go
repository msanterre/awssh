package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
  "path"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var cmdSync = &Command{
	Usage: "sync",
	Short: "syncs your ec2 machines with awssh",
	Run:   runSync,
}

func runSync(cmd *Command, args []string) {
  createStorageIfNotExists()

	svc := ec2.New(session.New(), &aws.Config{})
	validateRegion(svc)
	validateCredentials(svc)

	fmt.Println(*svc.Config.Region)

	resp, err := svc.DescribeInstances(nil)
	fail(err)

	var machine *Machine

	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {

			// The AWS sdk will fail if the instance isn't running
			if *instance.State.Name == "running" {
				address := instanceAddress(instance)
				name := instanceName(instance)
				user := instanceUser(instance)

				fmt.Println(address, name, user)
				machine = &Machine{
					Address: address,
					Name:    name,
					User:    user,
				}
				machine.Save()
			}
		}
	}
}

func validateRegion(svc *ec2.EC2) {
	if len(*svc.Config.Region) == 0 {
		fmt.Println("[error] AWS_REGION not set")
		showHelpAndExit()
	}
}

func validateCredentials(svc *ec2.EC2) {
	credentials, err := (svc.Config.Credentials.Get())
	fail(err)

	if len(credentials.AccessKeyID) == 0 {
		fmt.Println("[error] AWS_ACCESS_KEY_ID not set")
		showHelpAndExit()
	}
	if len(credentials.SecretAccessKey) == 0 {
		fmt.Println("[error] AWS_SECRET_ACCESS_KEY not set")
		showHelpAndExit()
	}
}

func formatName(name string) string {
	lowercased := strings.ToLower(name)
	return strings.Replace(lowercased, " ", "-", -1)
}

func instanceAddress(instance *ec2.Instance) string {
	address := *instance.PublicDnsName
	if len(address) == 0 {
		address = *instance.PublicIpAddress
	}
	return address
}

func instanceName(instance *ec2.Instance) string {
	for _, value := range instance.Tags {
		if *value.Key == "Name" {
			return formatName(*value.Value)
		}
	}
	return "n/a"
}

func instanceUser(instance *ec2.Instance) string {
	return DefaultUser
}

func (machine *Machine) Save() {
	fmt.Println("Saving: ", machine.Name)
	filePath := path.Join(os.ExpandEnv(StorageDest), machine.Name)
	machineJson, err := json.Marshal(machine)
	if err == nil {
		err = ioutil.WriteFile(filePath, machineJson, 0777)
		if err == nil {
			return
		}
	}
	fmt.Println("Could not save: ", machine.Name)
}

`
