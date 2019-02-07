package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString(`app.env`))
	fmt.Println(viper.GetString("app.env"))
	fmt.Println(viper.GetBool("app.debug"))
}


func main()  {
	
}