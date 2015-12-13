package main

import (
  "fmt"
  "os"
  "strings"
  "path"
  "io/ioutil"
  "encoding/json"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
)

const (
  DefaultUser = "ubuntu"
  StorageDest = "$HOME/.awssh/machines"
)

func fileExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

func fail(err error) {
  if err != nil {
    panic(err)
  }
}

func showHelpAndExit() {
  fmt.Println("Lol cya")
  os.Exit(1)
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

type Machine struct {
  Name string `json:"name"`
  User string `json:"user"`
  Address string `json:"host"`
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

func createStorageIfNotExists() {
  expandedDest := os.ExpandEnv(StorageDest)
  storageExists, err := fileExists(expandedDest)
  fail(err)

  if !storageExists {
    err = os.MkdirAll(expandedDest, 0777)
    fail(err)
  }
}

func main() {
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
          Name: name,
          User: user,
        }
        machine.Save()
      }
    }
  }

}
