/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/rollexxx23/spotify-cli/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("$HOME")
	viper.SetConfigName(".spotify-cli")
	viper.SetConfigType("json")
	_ = viper.SafeWriteConfig()
	_ = viper.ReadInConfig()
	cmd.Execute()
}
