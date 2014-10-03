package godecode
import "testing"
func TestDecode(t *testing.T){
  var gd GoDecode
  gd.init("data")
  m := make(map[string]string)
  m["hello world"] = "hello world"
  m["南无阿弥陀佛"] = "Nan Wu A Mi Tuo Fo"
  m["Κνωσός"] = "Knosos"
  m["あみだにょらい"] = "amidaniyorai"
  for k, v := range m {
    if gd.decode(k) != v {
      t.Errorf("[ %s ] Error: [ %s ] != [ %s ]", k, v, gd.decode(k))
    }
  }
}
func TestInitals(t *testing.T){
  var gd GoDecode
  gd.init("data")
  m := make(map[string]string)
  m["Hello world."] = "Hw"
  m["南无阿弥陀佛"] = "NWAMTF"
  m["Κνωσός"] = "K"
  m["あみだにょらい"] = "a"
  m["小小姑娘\n清早起床\n\r提着花篮\t上市场。"] = "XXGN\nQZQC\n\rTZHL\tSSC"
  for k, v := range m {
    if gd.initials(k) != v {
      t.Errorf("[ %s ] Error: [ %s ] != [ %s ]", k, v, gd.decode(k))
    }
  }
}
