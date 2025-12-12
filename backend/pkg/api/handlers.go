package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"terje-configurator/pkg/db"
	"terje-configurator/pkg/parser"
)

type TreeNode struct {
	Name     string      `json:"name"`
	Path     string      `json:"path"`
	Type     string      `json:"type"` // "file" or "folder"
	Children []*TreeNode `json:"children,omitempty"`
}

var RootPath string

func SetRootPath(path string) {
	RootPath = path
}

func EnableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func HandleTree(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	tree, err := buildTree(RootPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tree)
}

func buildTree(root string) ([]*TreeNode, error) {
	var nodes []*TreeNode

	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		name := entry.Name()
		fullPath := filepath.Join(root, name)

		relPath, _ := filepath.Rel(RootPath, fullPath)
		relPath = filepath.ToSlash(relPath)

		node := &TreeNode{
			Name: name,
			Path: relPath,
		}

		if entry.IsDir() {
			node.Type = "folder"
			children, err := buildTree(fullPath)
			if err == nil {
				node.Children = children
			}
			nodes = append(nodes, node)
		} else {
			ext := strings.ToLower(filepath.Ext(name))
			if ext == ".cfg" || ext == ".xml" {
				node.Type = "file"
				nodes = append(nodes, node)
			}
		}
	}
	return nodes, nil
}

func HandleFile(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method == "GET" {
		getFile(w, r)
	} else if r.Method == "POST" {
		saveFile(w, r)
	}
}

func getFile(w http.ResponseWriter, r *http.Request) {
	relPath := r.URL.Query().Get("path")
	fullPath := filepath.Join(RootPath, relPath)

	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	ext := strings.ToLower(filepath.Ext(fullPath))
	var result interface{}

	if ext == ".cfg" {
		result, err = parser.ParseCFG(string(content))
	} else if ext == ".xml" {
		result, err = parser.ParseXML(string(content))
	} else {
		http.Error(w, "Unsupported file type", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

type FileUpdate struct {
	LineIndex int    `json:"lineIndex"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

type SaveRequest struct {
	Path    string       `json:"path"`
	Updates []FileUpdate `json:"updates"`
}

func saveFile(w http.ResponseWriter, r *http.Request) {
	var req SaveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fullPath := filepath.Join(RootPath, req.Path)
	contentBytes, err := ioutil.ReadFile(fullPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	content := string(contentBytes)

	// Save snapshot
	db.SaveSnapshot(req.Path, content)

	ext := strings.ToLower(filepath.Ext(fullPath))
	var newContent string

	if ext == ".cfg" {
		file, _ := parser.ParseCFG(content)
		// Apply updates
		for _, update := range req.Updates {
			if update.LineIndex >= 0 && update.LineIndex < len(file.Lines) {
				line := file.Lines[update.LineIndex]
				if line.Type == parser.LineTypeConfig {
					line.Value = update.Value
				}
			}
		}
		newContent = file.String()
	} else if ext == ".xml" {
		file, _ := parser.ParseXML(content)
		// Apply updates
		for _, update := range req.Updates {
			if update.LineIndex >= 0 && update.LineIndex < len(file.Lines) {
				line := file.Lines[update.LineIndex]
				for i, seg := range line.Segments {
					if seg.IsAttribute && seg.AttrName == update.Key {
						line.Segments[i].AttrValue = update.Value
					}
				}
			}
		}
		newContent = file.String()
	}

	err = ioutil.WriteFile(fullPath, []byte(newContent), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleHistory(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	path := r.URL.Query().Get("path")
	history, err := db.GetHistory(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(history)
}

type RestoreRequest struct {
	ID uint `json:"id"`
}

func HandleRestore(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	var req RestoreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	snapshot, err := db.GetSnapshot(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Current content to save as new snapshot before restore?
	// Logic: When restoring, we overwrite current. Should we save the "bad" current state?
	// Yes, to allow undoing the restore.
	fullPath := filepath.Join(RootPath, snapshot.FilePath)
	currentBytes, _ := ioutil.ReadFile(fullPath)
	if len(currentBytes) > 0 {
		db.SaveSnapshot(snapshot.FilePath, string(currentBytes))
	}

	err = ioutil.WriteFile(fullPath, []byte(snapshot.Content), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
