# AWS SAMPLE IAM Role/Policy Requirements

## App
```
  Group: "arn:aws:iam::AWS_ACCOUNT:group/GROUP_NAME",
  User: 'sample_user",
  PolicyARN: 'arn:aws:iam::AWS_ACCOUNT:policy/POLICY_NAME
  Policy: 
    {
      "Version": "2012-10-17",
      "Statement": [
          {
              "Effect": "Allow",
              "Action": "sts:AssumeRole",
              "Resource": "arn:aws:iam::AWS_ACCOUNT:role/ROLE_NAME"
          }
      ]
    }
```

## Sample Role Assumed Role Admin Access. This is not Best Practice and Should NEVER be used. This is just an example. 
```
  RoleARN: "arn:aws:iam::AWS_ACCOUNT:role/ROLE_NAME",
  Policy:
    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Action": "*",
                "Resource": "*"
            }
        ]
    }
```