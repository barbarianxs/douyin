package api

import (
       "bytes"
       "fmt"
       "github.com/disintegration/imaging"
       ffmpeg "github.com/u2takey/ffmpeg-go"
       "log"
       "os"
       "strings"
       "time"
       "strconv"
)



func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {

       buf := bytes.NewBuffer(nil)
       err = ffmpeg.Input(videoPath).
             Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
             Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
             WithOutput(buf, os.Stdout).
             Run()

       if err != nil {
             log.Fatal("生成缩略图失败1：", err)
             return "", err

       }

       img, err := imaging.Decode(buf)
       if err != nil {
             log.Fatal("生成缩略图失败2：", err)
             return "", err

       }
       var name = strconv.Itoa((int)(time.Now().UnixMilli()))
       err = imaging.Save(img, snapshotPath+"/"+name+".jpeg")
       log.Println("........................"+name+".jpeg"+"............................")
       if err != nil {
             log.Fatal("生成缩略图失败3：", err)
             return "", err

       }
       names := strings.Split(name, "\\")
      //  names := strings.Split(snapshotPath, "\\")
       snapshotName = names[len(names)-1] + ".jpeg"
       return
}

// func main() {
//        _, err := GetSnapshot("./test.mp4", "test", 1)
//        if err != nil {
//              return
//        }
// }
