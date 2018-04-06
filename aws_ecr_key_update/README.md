# AWS ECR Key updater

## Prerequisite
Prepare your AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_DEFAULT_REGION, AWS_EMAIL_ADDRESS and put it into your environment
```
$ export AWS_ACCESS_KEY_ID=[Your AWS Access Key ID]
$ export AWS_SECRET_ACCESS_KEY=[Your AWS Secret Access Key]
$ export AWS_DEFAULT_REGION=[Your AWS Default Region (Not AZ)] # Default Region = ap-southeast-1
$ export AWS_EMAIL_ADDRESS=[Your Registered Email Address from AWS]
```

## Alternative
You can give out your access information by flags, there is a help
```
Usage of /tmp/go-build311400682/command-line-arguments/_obj/exe/run:
  -AWS_ACCESS_KEY_ID string
    	Your AWS Access Key ID
  -AWS_DEFAULT_REGION string
    	AWS Default Region: (Default: ap-southeast-1)
  -AWS_EMAIL_ADDRESS string
    	Registered email address for AWS
  -AWS_SECRET_ACCESS_KEY string
    	Your AWS Secret Access Key
```
## Run it
```
$ git clone https://github.com/vinnson/kubernetes-tools.git
$ ./kubernetes-tools/aws_ecr_key_update/bin/aws_ecr_key_update
```

## Build it on your own
```
$ git clone https://github.com/vinnson/kubernetes-tools.git
$ cd kubernetes-tools/aws_ecr_key_update
$ go build
$ ./aws_ecr_key_update
```

This is it, you may set it to your crontab and make sure it runs within every 6 hours.

## Crontab Sample
This is a sample of crontab command to run update in every hour.
```
0 * * * * /your_path_of_compiled_ecskey_update/ecskey_update >> /dev/null
```
