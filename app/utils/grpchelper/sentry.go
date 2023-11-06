package grpchelper

import "strings"

type ManageSentry struct {
}

func (*ManageSentry) ManageChildName(s string) string {
	// github.com/evenyosua18/auth2/app/infrastructure/server/grpc/service/accesstoken.(*ServiceAccessToken).PasswordGrant : infrastructure
	temp := strings.Split(s, "/")

	if len(temp) > 3 {
		return temp[4]
	}

	return s
}

func (*ManageSentry) ManageChildOperation(s string) string {
	// github.com/evenyosua18/auth2/app/infrastructure/server/grpc/service/accesstoken.(*ServiceAccessToken).PasswordGrant : PasswordGrant
	temp := strings.Split(s, ".")

	if len(temp) != 0 {
		return temp[len(temp)-1]
	}

	return s
}

func (*ManageSentry) ManageParentName(s string) string {
	// github.com/evenyosua18/auth2/app/infrastructure/server/grpc/service/accesstoken.(*ServiceAccessToken).PasswordGrant : infrastructure
	temp := strings.Split(s, "/")

	if len(temp) > 3 {
		return temp[4]
	}

	return s
}

func (*ManageSentry) ManageParentOperation(s string) string {
	// github.com/evenyosua18/auth2/app/infrastructure/server/grpc/middleware.OauthClientValidation.func1 : OauthClientValidation
	temp := strings.Split(s, ".")

	if len(temp) > 1 {
		return temp[len(temp)-2]
	}

	return s
}
