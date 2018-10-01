# Curlew

![Curlew](curlew.jpg)

## Introduction

Curlew (a bird with a long, narrow beak) will get a single object from an S3 bucket.

## Contents

- [Usage](#usage)
- [Permissions](#permissions)
- [Dev](#dev)

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
