# awssh

## Purpose

If your AWS EC2 instances change frequently and you name them appropriately, awssh is a good way to make the pain of figuring out how to SSH into those new instances easier. 

## Installation

```
go get github.com/msanterre/awssh
```

## Configuration

This tool uses the basic AWS API key configs: 
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- AWS_REGION

## Usage

If this is your first time using it. You should start by syncing

```
$ awssh sync
Syncing ...
Saving: test-1
```

*awssh* makes some assumptions on the type of instance you start, so the default name of your instance will be `ubuntu`, which is the default login for ubuntu type machines. (Support for other types will be added upon request)

Now to connect to your instance:
```
$ awssh list
test-1              ubuntu     ec2-51-186-242-255.us-west-2.compute.amazonaws.com

$ awssh connect test-1
Connecting to: ubuntu@ec2-51-186-242-255.us-west-2.compute.amazonaws.com
```

And you should now be connecting to your machine! 🎉
