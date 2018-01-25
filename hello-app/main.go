/**
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START all]
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

func main() {
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	server := http.NewServeMux()
	server.HandleFunc("/", hello)
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server))
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "I am a lady bug.\n")
	fmt.Fprintf(w, "This is mew and mewtwo.\n")
	fmt.Fprintf(w, "Version: 1.1.1\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
	    m := Image{256, 64}
        pic.ShowImage(m)
}

type Image struct{
    width int
    height int
}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
    img_func := func(x, y int) uint8 {
        //return uint8(x*y)
        //return uint8((x+y) / 2)
        return uint8(x^y)
    }
    v := img_func(x, y)
    return color.RGBA{v, v, 255, 255}
}
// [END all]
