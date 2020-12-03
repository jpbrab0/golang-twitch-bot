package main

import (
	"fmt"
	"os"
    "time"
    "log"
    "strings"
    "io/ioutil"
    "math/rand"
    "github.com/joho/godotenv"
	"github.com/gempir/go-twitch-irc/v2"
    "github.com/christopher-dG/go-obs-websocket"
)
func getDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    panic(err)
  }

  return os.Getenv(key)
}

func main(){
    rand.Seed(time.Now().UTC().UnixNano())
    c := obsws.Client{Host: "localhost", Port: 4444}

    if err := c.Connect(); err != nil {
        log.Fatal(err)
	}
	defer c.Disconnect()

    tokenBot := getDotEnvVariable("TOKEN_BOT")
    client := twitch.NewClient("brab0bot", tokenBot) 
    client.OnConnect(func() {
        client.Say("jpbrab0","O bot ta on!")
    })
    client.OnUserNoticeMessage(func(message twitch.UserNoticeMessage){
        fmt.Println(message)
    })
    client.OnPrivateMessage(func(message twitch.PrivateMessage){
        fmt.Println(message.User.Name+":"+ message.Message)
        if strings.HasPrefix(message.Message, "!ban"){
            nicknameBan := strings.Split(message.Message, "!ban")[1]
            client.Say("jpbrab0",message.User.Name +  " baniu" + nicknameBan)
            numRandom := rand.Intn(1000 - 100) + 100
            if(numRandom > 100 && numRandom < 499 ){
                client.Say("jpbrab0", "/timeout "+message.User.Name+" 15")
                client.Say("jpbrab0", message.User.Name+" foi banido(a)")
            } else {
                client.Say("jpbrab0", "/timeout "+nicknameBan+" 15")
                client.Say("jpbrab0", nicknameBan+" foi banido(a)")
            }
        } else if strings.HasPrefix(message.Message, "!discord"){
            nickname := strings.Split(message.Message, "!discord")[1]
            if(nickname == ""){
                client.Say("jpbrab0", "Discord -> https://discord.io/caverna")
            } else {
                client.Say("jpbrab0", nickname + "," + " Discord -> https://discord.io/caverna")
            }
        } else if strings.HasPrefix(message.Message, "!os"){
            client.Say("jpbrab0", "Ubuntu 20.10 Shell: Zsh Theme: WhiteSur-dark")
        } else if strings.HasPrefix(message.Message, "!hoje"){
            data, err := ioutil.ReadFile("hoje.txt")
            if err != nil {
                fmt.Println("File reading error", err)
                return
            }
            client.Say("jpbrab0", string(data))
        } else if strings.HasPrefix(message.Message, "!sethj"){
            if message.User.Name == "jpbrab0" {
                hoje := []byte(strings.Split(message.Message, "!sethj")[1])
                err := ioutil.WriteFile("hoje.txt", hoje, 0777)

                if err != nil {
                    fmt.Println(err)
                }

                client.Say("jpbrab0", "O arquivo hoje.txt foi atualizado.")
            }
        }
    })
    client.Join("jpbrab0")
    client.Connect()
}   
