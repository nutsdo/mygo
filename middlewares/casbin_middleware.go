package middlewares

import (
	"github.com/casbin/casbin"
	"os"
	"github.com/kataras/iris"
)

type CmConfig struct {
	casbinModelPath string
	casbinPolicyPath string
}

type CasbinMiddleware struct {
	Ctx iris.Context
}

func relativePath() string{
	base,err := os.Getwd()
	if err != nil {
		return err.Error()
	}
	return base + "/middlewares"
}

func (cm *CasbinMiddleware) configPath() *CmConfig{
	cc := new(CmConfig)
	cc.casbinModelPath = relativePath() + "/casbinmodel.conf"
	cc.casbinPolicyPath = relativePath() + "/casbinpolicy.csv"
	return cc
}

func (cm *CasbinMiddleware) NewCasbin() *casbin.Enforcer {

	enforcer := casbin.NewEnforcer(cm.configPath().casbinModelPath,cm.configPath().casbinPolicyPath)

	return enforcer
}