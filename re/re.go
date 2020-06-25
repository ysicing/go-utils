// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package re

import (
	"regexp"
	"strings"
)

const (
	mail           = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	configMapsName = "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$"
	envsName       = "^[-._a-zA-Z][-._a-zA-Z0-9]*$"
	phone          = "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	domain         = "[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?"
	git            = "^(http(s)?|git|ssh).*(\\.git)(/)?$"
)

//VerifyEmailFormat email verify
func VerifyEmailFormat(code string) bool {
	reg := regexp.MustCompile(mail)
	return reg.MatchString(code)
}

// VerifyK8sConfigMapsName 校验k8s configmaps name
func VerifyK8sConfigMapsName(code string) bool {
	return regexp.MustCompile(configMapsName).MatchString(code)
}

// VerifyK8sEnvs 校验k8s envs name
func VerifyK8sEnvs(code string) bool {
	return regexp.MustCompile(envsName).MatchString(code)
}

//VerifyMobileFormat mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	reg := regexp.MustCompile(phone)
	return reg.MatchString(mobileNum)
}

//VerifyDomainFormat domain verify
func VerifyDomainFormat(domain string) bool {
	reg := regexp.MustCompile(domain)
	return reg.MatchString(domain)
}

//VerifyGitFormat git addr verify
func VerifyGitFormat(gitaddr string) bool {
	reg := regexp.MustCompile(git)
	return reg.MatchString(gitaddr)
}

// ParseGitAddr 解析git repo
func ParseGitAddr(gitAddr string) (string, string) {
	gitAddr = strings.Replace(gitAddr, ".git", "", -1)
	addr := strings.Split(gitAddr, ":")
	names := strings.Split(addr[len(addr)-1], "/")
	owner := names[len(names)-2]
	repo := names[len(names)-1]
	return owner, repo
}

// ParseGitAddrv2 获取项目仓库组织和项目仓库名
func ParseGitAddrv2(gitaddr string) (string, string) {
	if strings.HasPrefix(gitaddr, "git") {
		s0 := strings.Split(gitaddr, ":")[1]
		s1 := strings.Split(s0, ".")[0]
		s2 := strings.Split(s1, "/")
		return s2[0], s2[1]
	}
	s0 := strings.Split(gitaddr, "/")
	s1 := strings.Split(s0[len(s0)-1], ".")[0]
	s2 := s0[len(s0)-2]
	return s2, s1
}
