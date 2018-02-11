# AWS ECS Key updater

## Prerequisite
Prepare your AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_DEFAULT_REGION, AWS_EMAIL_ADDRESS and put it into your environment
```
$ export AWS_ACCESS_KEY_ID=[Your AWS Access Key ID]
$ export AWS_SECRET_ACCESS_KEY=[Your AWS Secret Access Key]
$ export AWS_DEFAULT_REGION=[Your AWS Default Region (Not AZ)]
$ export AWS_EMAIL_ADDRESS=[Your Registered Email Address from AWS]
```

## Build and run it
```
$ git clone https://github.com/vinnson/kubernetes-tools.git
$ cd kubernetes-tools/ecskey_update
$ go build
$ ./ecskey_update
```

This is it, you may set it to your crontab and make sure it runs within every 6 hours.
