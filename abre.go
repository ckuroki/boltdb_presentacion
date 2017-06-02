import (
  "log"
  "github.com/boltdb/bolt"
)

func main() {
  // Abre la base de datos en el directorio actual
  // Crea la base si el archivo no existe
  db, err := bolt.Open("blog.db", 0600, nil)

  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  //...

}
