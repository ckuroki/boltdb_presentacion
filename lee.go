
func ReadUnique (bucketName string, key []byte) ([]byte,error) {
  var buf[]byte

  // View abre una transacción de lectura 

  err:= Db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(bucketName))
      // Trae solo una clave
      buf = b.Get(key)
      return nil
  })
  return buf,err
}

