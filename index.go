package main

import (
        "fmt"
        "html/template"
        "log"
        "net/http"
        qrcode "github.com/yougg/go-qrcode"

)

func generateQRCode(w http.ResponseWriter, r *http.Request) {
        // Obtener los datos del formulario
        nombre := r.FormValue("nombre")
        telefono := r.FormValue("telefono")
        telefonolaboral := r.FormValue("telefonolaboral")
        email := r.FormValue("email")
        apellido := r.FormValue("apellido")
        direccion := r.FormValue("direccion")
        empresa := r.FormValue("empresa")
        titulo := r.FormValue("titulo")
        cargo := r.FormValue("cargo")
        web := r.FormValue("web")
        instagram := r.FormValue("instagram")
        facebook := r.FormValue("facebook")
        linkedin := r.FormValue("linkedin")
        website := r.FormValue("website")
        vcard := ""

        if website != "" {
                vcard += website
        }else{
                 fmt.Println("Inside vcard")
                // Construir el contenido de la vCard
                vcard += "BEGIN:VCARD\n" +
                        "VERSION:3.0\n" +
                        "N:" + nombre + ";" + apellido + ";;;\n" +
                        "TEL;TYPE=CELL:" + telefono + "\n";

                if web != "" {
                        vcard += "URL:" + web + "\n"
                }



                if telefonolaboral != "" {
                        vcard += "TEL;TYPE=WORK:" + telefonolaboral + "\n"
                }
                if cargo != "" {
                        vcard += "ROLE:" + cargo + ";\n"
                }

                if empresa != "" {
                        vcard += "ORG:" + empresa + ";\n"
                }

                if email != "" {
                        vcard += "EMAIL:" + email + ";\n"
                }

                if direccion != "" {
                        vcard += "ADR:;;" + direccion + ";;;\n"
                }


                if titulo != "" {
                        vcard += "TITLE:" + titulo + ";\n"
                }


                if instagram != "" {
                        vcard += "X-SOCIALPROFILE;type=instagram:" + instagram + "\n"
                }

                if facebook != "" {
                        vcard += "X-SOCIALPROFILE;type=facebook:" + facebook + "\n"
                }

                if linkedin != "" {
                        vcard += "X-SOCIALPROFILE;type=linkedin:" + linkedin + "\n"
                }

                vcard += "END:VCARD"
                fmt.Println(vcard)
        }
        
         fmt.Println(vcard)
        var png []byte
        png, err := qrcode.Encode(vcard, qrcode.Medium, 256, 256, 1)
        
        // Obtener la imagen QR en bytes


        // Establecer las cabeceras de respuesta HTTP
        w.Header().Set("Content-Type", "image/png")

        // Escribir los bytes de la imagen QR en el ResponseWriter
        _, err = w.Write(png)
        if err != nil {
                log.Fatal(err)
        }
        
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
                // Mostrar el formulario HTML
                tmpl := template.Must(template.ParseFiles("index.html"))
                err := tmpl.Execute(w, nil)
                if err != nil {
                        log.Fatal(err)
                }
        } else if r.Method == "POST" {
                // Procesar la solicitud POST
                generateQRCode(w, r)
        }
}

func qrHandler(w http.ResponseWriter, r *http.Request) {

                generateQRCode(w, r)
        
}

func main() {
        http.HandleFunc("/", indexHandler)
        http.HandleFunc("/qr", qrHandler)
        err := http.ListenAndServe(":8002", nil)
        if err != nil {
                fmt.Println(err)
        }
}

