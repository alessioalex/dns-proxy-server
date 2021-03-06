package controller

import (
	log "github.com/mageddo/go-logging"
	"net/http"
	"encoding/json"
	"github.com/mageddo/dns-proxy-server/events/local"
	"golang.org/x/net/context"
	"github.com/mageddo/dns-proxy-server/utils"
)

func init(){

	Get("/env/active", func(ctx context.Context, res http.ResponseWriter, req *http.Request, url string){
		res.Header().Add("Content-Type", "application/json")
		utils.GetJsonEncoder(res).Encode(local.EnvVo{Name: local.GetConfiguration(ctx).ActiveEnv})
	})

	Put("/env/active", func(ctx context.Context, res http.ResponseWriter, req *http.Request, url string){

		logger := log.NewLog(ctx)
		logger.Infof("m=/env/active/, status=begin")
		res.Header().Add("Content-Type", "application/json")

		var envVo local.EnvVo
		json.NewDecoder(req.Body).Decode(&envVo)
		logger.Infof("m=/env/active/, status=parsed-host, env=%+v", envVo)
		err := local.GetConfiguration(ctx).SetActiveEnv(ctx, envVo)
		if err != nil {
			logger.Infof("m=/env/, status=error, action=create-env, err=%+v", err)
			BadRequest(res, err.Error())
			return
		}
		logger.Infof("m=/env/active/, status=success, action=active-env")

	})

	Get("/env/", func(ctx context.Context, res http.ResponseWriter, req *http.Request, url string){
		res.Header().Add("Content-Type", "application/json")
		utils.GetJsonEncoder(res).Encode(local.GetConfiguration(ctx).Envs)
	})

	Post("/env/", func(ctx context.Context, res http.ResponseWriter, req *http.Request, url string){
		logger := log.NewLog(ctx)
		res.Header().Add("Content-Type", "application/json")
		logger.Infof("m=/env/, status=begin, action=create-env")
		var envVo local.EnvVo
		json.NewDecoder(req.Body).Decode(&envVo)
		logger.Infof("m=/env/, status=parsed-host, env=%+v", envVo)
		err := local.GetConfiguration(ctx).AddEnv(ctx, envVo)
		if err != nil {
			logger.Infof("m=/env/, status=error, action=create-env, err=%+v", err)
			BadRequest(res, err.Error())
			return
		}
		logger.Infof("m=/env/, status=success, action=create-env")
	})

	Delete("/env/", func(ctx context.Context, res http.ResponseWriter, req *http.Request, url string){
		logger := log.NewLog(ctx)
		res.Header().Add("Content-Type", "application/json")
		logger.Infof("m=/env/, status=begin, action=delete-env")
		var env local.EnvVo
		json.NewDecoder(req.Body).Decode(&env)
		logger.Infof("m=/env/, status=parsed-host, action=delete-env, env=%+v", env)
		err := local.GetConfiguration(ctx).RemoveEnvByName(ctx, env.Name)
		if err != nil {
			logger.Infof("m=/env/, status=error, action=delete-env, err=%+v", err)
			BadRequest(res, err.Error())
			return
		}
		logger.Infof("m=/env/, status=success, action=delete-env")
	})
}
