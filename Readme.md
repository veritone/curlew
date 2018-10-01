# Curlew

![Curlew](curlew.jpg)

## Introduction

Curlew (a bird with a long, narrow beak) will get a single object from an S3 bucket.

## Contents

- [Download](#download)
- [Usage](#usage)
- [Permissions](#permissions)
- [Dev](#dev)

## Download

Download source:

```bash
go get git@github.com:veritone/curlew.git
```

Download a release:

```bash
# OS X
curl -L -o /usr/bin/curlew https://github.com/veritone/curlew/releases/download/0.1.0/curlew-0.1.0-darwin-amd64

# Linux
curl -L -o /usr/bin/curlew https://github.com/veritone/curlew/releases/download/0.1.0/curlew-0.1.0-linux-amd64

# Update permissions in all cases
chmod a+x /usr/bin/curlew
```

## Usage

Curlew prints object to stdout & errors to stderr; you need to redirect stdout to file:

```bash
dep ensure
go install

curlew s3://s3-bucket-here/path/to/object.txt > object.txt
```

## Permissions

Make sure the caller has a policy to allow access to the object

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::s3-bucket-here/path/to/object.txt"
    }
  ]
}
```

## Dev

When you need to change versions, update the `version.txt` then push to Github so releases can be made

```bash
git tag "$(cat version.txt)"
git push upstream "$(cat version.txt)"
```
