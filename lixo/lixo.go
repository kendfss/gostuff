package main

import (
    "fmt"
    "os"
    "path"
    "path/filepath"
    "image"
    "image/jpeg"
    "log"
    "testing"
    "time"
    
    
    // "github.com/buchanae/ink"
    // . "github.com/buchanae/ink/color"
    // . "github.com/buchanae/ink/dd"
    // "github.com/buchanae/ink/gfx"
)

func dummy(args...interface{}) {
    if len(args) < 0 {
        tau := testing.T{}
        
        fmt.Println(
            // os.Executable(), 
            filepath.Join("one", "two"),
            path.Join("one", "two"),
            tau,
            jpeg.DefaultQuality,
            image.YCbCrSubsampleRatio444,
            time.Second,
            
        )
        log.Fatal("dog the bollock")
    }
}

func Assert(err error, fstr string) {
    if err != nil {
        fstr += "\n\t%v"
        str := fmt.Sprintf(fstr, err) + "\n"
        panic(str)
    }
}

func Check(err error, fstr string) {
    if err != nil {
        fstr += "\n\t%v"
        str := fmt.Sprintf(fstr, err) + "\n"
        log.Fatalf(str)
    }
}
func LoadImageFace(filePath string) image.Image {
    f, err := os.Open(filePath)
    defer f.Close()
    Assert(err, "")
    img, ext, err := image.Decode(f)
    fmt.Println(ext)
    // img, err := jpeg.Decode(f)
    if err != nil {
        panic(fmt.Sprintf("Couldn't decode image file:\n\t\"%s\"\n%s", filePath, err))
    }
    return img
}

// func Ink(doc gfx.Doc) {
//     t := Triangle{
//         XY{0.2, 0.2},
//         XY{0.8, 0.2},
//         XY{0.5, 0.8},
//     }
//     s := gfx.Fill{Shape: t}.Shader()
//     s.Set("a_color", []RGBA{
//         Red, Green, Blue,
//     })
//     s.Draw(doc)
// }

func main() {
    // exe, err := os.Executable()
    // Check(err, "")
    // config, err := os.UserConfigDir()
    // Check(err, "")
    // cache, err := os.UserCacheDir()
    // Check(err, "")
    // fmt.Println(
    // )
    // dir, err := os.UserHomeDir()
    // Check(err, "")
    // im := LoadImageFace("C:/Users/Kenneth/Pictures/repatriation pics/MusicPics/AlbumArt_{B5020207-474E-4720-0A2D-43031AF9F600}_Large.jpg")
    // fmt.Println(im.Bounds())
    // fmt.Println(im.ColorModel())
    // fmt.Println(string(os.PathSeparator))
    
    // switch today := time.Now(); {
    //     case today.Day() < 5:
    //         fmt.Println("Clean your house.")
    //     case today.Day() <= 10:
    //         fmt.Println("Buy some wine.")
    //     case today.Day() > 15:
    //         fmt.Println("Visit a doctor.")
    //     case today.Day() > 17:
    //         fmt.Println("lol.")
    //     case today.Day() == 25:
    //         fmt.Println("Buy some food.")
    //     default:
    //         fmt.Println("No information available for that day.")
    // }
    
    dict := map[string]int{}
    fmt.Println(dict["shit"])
    
    
}
