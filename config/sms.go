package config

import (
	"GoHub/pkg/config"
)

func init() {
	config.Add("sms", func() map[string]interface{} {
		aliyunConfig := map[string]interface{}{
			"access_key_id":     config.Env("SMS_ALIYUN_ACCESS_ID"),
			"access_key_secret": config.Env("SMS_ALIYUN_ACCESS_SECRET"),
			"sign_name":         config.Env("SMS_ALIYUN_SIGN_NAME"),
			"template_code":     config.Env("SMS_ALIYUN_TEMPLATE_CODE"),
		}

		return map[string]interface{}{
			"aliyun": aliyunConfig,
		}
	})
}
