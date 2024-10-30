// =================================================================================
//
//		gim - https://www.foxhollow.cc/projects/gim/
//
//	 go-import-media, aka gim, is a tool for automatically importing media
//	 from removable disks into a predefined folder structure automatically.
//
//		Copyright (c) 2024 Steve Cross <flip@foxhollow.cc>
//
//	This file was originally part of the localsend-go project created by
//	MeowRain. It was adapted for use by Steve Cross in the go-import-media
//	project on 2024-10-30.
//
//	    Copyright (c) 2024 MeowRain
//	    localsend-go - https://github.com/meowrain/localsend-go
//	    License: MIT (for complete text, see LICENSE file in localsend folder)
//
// =================================================================================
package handler

import (
	"encoding/json"
	"fmt"
	"gim/localsend/model"
	"net/http"
)

func RegisterHandler(config model.ConfigModel, message model.BroadcastMessage, w http.ResponseWriter, r *http.Request) {
	message.Announce = false

	res, err := json.Marshal(message)
	if err != nil {
		fmt.Println("json convert failed:", err)
		http.Error(w, "json convert failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)

	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		fmt.Println("Error writing file:", err)
		return
	}
}
