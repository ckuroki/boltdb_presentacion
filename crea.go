  func CrearActualizar (bucketName string, key []byte, val []byte) error {
    // Abre una transaccion de escritura
    tx, err := Db.Begin(true)
    if err != nil {
      return err
    }
    defer tx.Rollback() 

    // Trae el bucket, si no existe, lo crea
    b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

    // Guarda en el bucket.
    err = b.Put(key, val)
    if  err != nil {
      return err
    }

    // Cierra la transaccion
    if err:= tx.Commit(); err != nil {
      return  err 
    }
  return nil
  }
