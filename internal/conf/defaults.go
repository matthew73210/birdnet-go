// contains default values for settings
package conf

import (
	"time"

	"github.com/spf13/viper"
)

// Sets default values for the configuration.
func setDefaultConfig() {
	viper.SetDefault("debug", false)

	viper.SetDefault("main.name", "BirdNET-Go")
	viper.SetDefault("main.timeas24h", true)
	viper.SetDefault("main.log.enabled", true)
	viper.SetDefault("main.log.path", "birdnet.log")
	viper.SetDefault("main.log.rotation", RotationDaily)
	viper.SetDefault("main.log.maxsize", 1048576)
	viper.SetDefault("main.log.rotationday", time.Sunday)

	viper.SetDefault("birdnet.sensitivity", 1.0)
	viper.SetDefault("birdnet.threshold", 0.8)
	viper.SetDefault("birdnet.overlap", 0.0)
	viper.SetDefault("birdnet.threads", 0)
	viper.SetDefault("birdnet.locale", "en")
	viper.SetDefault("birdnet.latitude", 0.000)
	viper.SetDefault("birdnet.longitude", 0.000)
	viper.SetDefault("birdnet.locationfilterthreshold", 0.01)

	viper.SetDefault("realtime.interval", 15)
	viper.SetDefault("realtime.processingtime", false)

	viper.SetDefault("realtime.audioexport.enabled", true)
	viper.SetDefault("realtime.audioexport.path", "clips/")
	viper.SetDefault("realtime.audioexport.type", "wav")

	viper.SetDefault("realtime.log.enabled", false)
	viper.SetDefault("realtime.log.path", "birdnet.txt")

	viper.SetDefault("realtime.birdweather.enabled", false)
	viper.SetDefault("realtime.birdweather.debug", false)
	viper.SetDefault("realtime.birdweather.id", "00000")
	viper.SetDefault("realtime.birdweather.threshold", 0.8)
	viper.SetDefault("realtime.birdweather.locationaccuracy", 500)

	viper.SetDefault("realtime.rtsp.url", "")
	viper.SetDefault("realtime.rtsp.transport", "tcp")

	viper.SetDefault("realtime.mqtt.enabled", false)
	viper.SetDefault("realtime.mqtt.broker", "tcp://localhost:1883")
	viper.SetDefault("realtime.mqtt.topic", "birdnet")
	viper.SetDefault("realtime.mqtt.username", "birdnet")
	viper.SetDefault("realtime.mqtt.password", "secret")

	viper.SetDefault("realtime.privacyfilter.enabled", true)
	viper.SetDefault("realtime.dogbarkfilter.enabled", false)

	viper.SetDefault("realtime.telemetry.enabled", false)
	viper.SetDefault("realtime.telemetry.listen", "0.0.0.0:8090")

	viper.SetDefault("realtime.retention.enabled", false)
	viper.SetDefault("realtime.retention.minEvictionHours", 0)
	viper.SetDefault("realtime.retention.minClipsPerSpecies", 0)

	viper.SetDefault("webserver.enabled", true)
	viper.SetDefault("webserver.port", "8080")
	viper.SetDefault("webserver.autotls", false)

	viper.SetDefault("webserver.log.enabled", true)
	viper.SetDefault("webserver.log.path", "webui.log")
	viper.SetDefault("webserver.log.rotation", RotationDaily)
	viper.SetDefault("webserver.log.maxsize", 1048576)
	viper.SetDefault("webserver.log.rotationday", time.Sunday)

	viper.SetDefault("output.file.enabled", true)
	viper.SetDefault("output.file.path", "output/")
	viper.SetDefault("output.file.type", "table")

	viper.SetDefault("output.sqlite.enabled", true)
	viper.SetDefault("output.sqlite.path", "birdnet.db")

	viper.SetDefault("output.mysql.enabled", false)
	viper.SetDefault("output.mysql.username", "birdnet")
	viper.SetDefault("output.mysql.password", "secret")
	viper.SetDefault("output.mysql.database", "birdnet")
	viper.SetDefault("output.mysql.host", "localhost")
	viper.SetDefault("output.mysql.port", 3306)
}
