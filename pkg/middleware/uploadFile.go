package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// add profile image
func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		file, _, err := r.FormFile("image")

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error Retrieving the File")
			return
		}
		defer file.Close()
		const MAX_UPLOAD_SIZE = 10 << 20 // 10MB
		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size in 1mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

		data := tempFile.Name()
		filename := data[8:] // split uploads/

		// add filename to ctx
		ctx := context.WithValue(r.Context(), "dataFile", filename)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// add thumbnail housy
func UploadThumbnail(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		file, _, err := r.FormFile("thumbnail")

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error Retrieving the File")
			return
		}
		defer file.Close()
		const MAX_UPLOAD_SIZE = 10 << 20 // 10MB
		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size in 1mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

		data := tempFile.Name()
		filename := data[8:] // split uploads/

		// add filename to ctx
		ctx := context.WithValue(r.Context(), "dataFile", filename)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// multiple image
func UploadMultipleFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := r.ParseMultipartForm(200000) //grab the multipart form
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		formdata := r.MultipartForm

		//get the *fileheader
		files := formdata.File["image"]

		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}

			out, err := os.Create("/uploads" + files[i].Filename)

			defer out.Close()
			if err != nil {
				fmt.Fprintln(w, "Unable to create the file for writing. Check your write access privilege")
				return
			}

			_, err = io.Copy(out, file)

			if err != nil {
				fmt.Fprintln(w, err)
				return
			}

			fmt.Fprintf(w, "Files uploaded successfully")
			fmt.Fprintf(w, files[i].Filename)
		}
	})
}

// func UploadMultipleFiles(next http.HandlerFunc) http.HandlerFunc {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Parse the multipart form
//         err := r.ParseMultipartForm(200000)
//         if err != nil {
//             fmt.Fprintln(w, err)
//             return
//         }

//         formdata := r.MultipartForm

//         // Get the fileheaders
//         files := formdata.File["image"]

//         // Loop through the fileheaders and save each file
//         for i, _ := range files {
//             file, err := files[i].Open()
//             defer file.Close()
//             if err != nil {
//                 fmt.Fprintln(w, err)
//                 return
//             }

//             // Create a unique filename using the UUID package
//             filename := uuid.New().String() + filepath.Ext(files[i].Filename)
//             out, err := os.Create("/uploads/" + filename)

//             defer out.Close()
//             if err != nil {
//                 fmt.Fprintln(w, "Unable to create the file for writing. Check your write access privilege")
//                 return
//             }

//             _, err = io.Copy(out, file)

//             if err != nil {
//                 fmt.Fprintln(w, err)
//                 return
//             }
//         }

//         fmt.Fprintln(w, "Files uploaded successfully")
//     })
// }
