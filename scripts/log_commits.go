package main

import (
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "time"
)

func main() {
  cmd := exec.Command("git", "log", "-n", "3", "--pretty=format:%h - %an, %ar : %s")
  out, err := cmd.Output()
  if err != nil {
    fmt.Printf("Error ejecutandogit log %v\n, err")
    os.Exit(1)
  }

 // creacion de la carpeta log
  logDir := filepath.Join("..", "log")
  if _, err := os.Stat(logDir); os.IsNotExist(err) {
    //Definiendo permisos de escritura
    err = os.Mkdir (logDir, 0755)
    if err != nil {
      fmt.Printf("Error creando directorio %s: %v\n", logDir,err);
      os.Exit(1)
    }
  }

  //Generar nombre del archivo
  currentTime := time.Now().Format("2006-01-02_15-04-05") // la mascara de YYYY-MM-DD HH-II-SS
  logFile := filepath.Join(logDir, fmt.Sprintf("commits_%s.txt",currentTime))

  //Escrivimos el archivo
  contigut := fmt.Sprintf("Se han escrito los ultimos 3 commits del repositorio:\n%s", string(out))
  os.WriteFile(logFile, []byte(contigut), 0644)
  if err != nil {
    fmt.Printf("se ha producido un error creando en %s %V\n", logFile,err)
    os.Exit(1)
  }

  fmt.Printf("Se ha creado el archivo del log en %s\n", logFile)
}
                           
  
