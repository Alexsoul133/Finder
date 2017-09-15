package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// srcPath := ""
	srcName := "11"
	dstPath := filepath.Join("\\rackstation3", "YANDEX", "Ton")
	dstPath = "\\" + dstPath
	dstName := ""
	vf := ""
	time := ""

	// var commandArray = []string{
	// 	srcName=      ""
	// 	dstName=      "coid-9999-q1.mp4"
	// 	____vChannel:= "[0:v]"
	// 	_vfPrefilter:= ""
	// 	______vfCrop:= ""
	// 	_____vfScale:= "auto"
	// 	_____aFilter:= "[0:]" //tostereo,2left,2right
	// 	________time= ""
	// 	_______shift:= "" //shift[,from,to] //hh:mm:ss:ff or frames count
	// 	//   template: `${globalTemplate}`
	// }

	globalTemplate := strings.NewReplacer(
		"###srcName###", srcName,
		"###filter_complex###", vf,
		"###time###", time,
		"###dstPath###", dstPath,
		"###dstName###", dstName)
	command := globalTemplate.Replace("fflite -i \"###srcName###\" -map_metadata -1 -map_chapters -1 -filter_complex \"###filter_complex###\" -map [v] -vcodec libx264 -preset medium -crf 17 -pix_fmt yuv420p -g 0 -map [a] -acodec aac -ab 256k -metadata:s:a:0 language=rus -metadata:s:a:0 \"handler=Russian 2.0\" -disposition:1 default ###time### -n \"###dstPath######dstName###\"")

	cmd, err := exec.Command(command).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cmd))
}
