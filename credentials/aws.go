package requestcredentials

import (
	"fmt"
	"net"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

// Hard-code credentials are not recommended.
// Hard-coding credentials in your application can make it difficult
// to manage and rotate those credentials. Use this method for small
// personal scripts or testing purposes only. Do not submit code with
//credentials to source control.
const (
	AWS_ACCESS_KEY_ID     string = ""
	AWS_SECRET_ACCESS_KEY string = ""
	AWS_ACCOUNT           string = ""
)

func Generate(role string) (*sts.AssumeRoleOutput, error) {
	creds := credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, "")
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}
	cfgs := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)
	svc := sts.New(session.New(), cfgs)

	// WE USE THIS TO ATTACH USERS IP TO CLOUDWATCH LOGS.
	tempRoleSessionName := GetUserIP()
	Params := &sts.AssumeRoleInput{
		RoleArn:         aws.String(fmt.Sprintf("arn:aws:iam::%s:role/%s", AWS_ACCOUNT, role)),
		RoleSessionName: aws.String(tempRoleSessionName),
		DurationSeconds: aws.Int64(900),
	}

	resp, err := svc.AssumeRole(Params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetUserIP() string {
	addrs, _ := net.InterfaceAddrs()
	var ip string

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = fmt.Sprint(ipnet.IP.String())
				break
			}
		}
	}
	return ip
}
