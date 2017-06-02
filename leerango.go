func ReadRangeFromSet (parentBucketName string,detailBucketName string) ([][]byte,error) {
  var buf [][]byte
  var elem []byte

  err:=Db.View(func(tx *bolt.Tx) error {
    // Bucket padre
    p := tx.Bucket([]byte(parentBucketName))
    // Bucket hijo
    d := tx.Bucket([]byte(detailBucketName))
      // Iterar resultados
      c := p.Cursor()

      for k, v := c.First(); k != nil; k, v = c.Next() {
        elem = d.Get(v)
        buf= append(buf,elem)
      }
      return nil
  })
  return buf,err
}

