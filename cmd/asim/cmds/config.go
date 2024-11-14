package cmds

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFilePath string
)

var config struct {
	Address struct {
		Search  string `mapstructure:"search"`
		Country string `mapstructure:"country"`
		City    string `mapstructure:"city"`
		State   string `mapstructure:"state"`
	} `mapstructure:"address"`
}

func init() {
	cobra.OnInitialize(initConfig)

	flags := rootCmd.Flags()

	// Special flag for config
	flags.StringVarP(&configFilePath, "config", "c", "", "A file path to load as additional configuration.")

	flags.String("address.country", "US", "The country of the address to look around.")
	flags.StringP("address.state", "s", "SC", "The state of the address to look around.")
	flags.StringP("address.city", "C", "", "The city of the address to look around.")
	flags.StringP("address.search", "a", "", "The address to search for (ex: 123 Some Rd).")

	err := viper.BindPFlags(flags)

	if err != nil {
		panic(err)
	}
}

func initConfig() {
	if configFilePath != "" {
		viper.SetConfigFile(configFilePath)
	}

	viper.SetEnvPrefix("ASIM")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	// Ignore errors here because we don't necessarily need a config file
	_ = viper.ReadInConfig()

	err := viper.Unmarshal(&config)

	if err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}
}
