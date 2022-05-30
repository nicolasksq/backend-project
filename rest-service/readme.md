# Back-End Project

Please to read about our API go through postman collection, which has the documentation.

https://documenter.getpostman.com/view/16593738/Uz5CKcj4#97ab8da8-c9cf-4d00-abd0-e43b740c8256

### Test coverage 87.3%
There are a few things to improve, mostly related with naming, and testing.
Also we should have our owns error, like

```go 
type NotFound struct {
  Message string
  Err     err
}

func (nf *NotFound) Error() string {
  return fmt.Sprintf("Not found. err: %s - message: %s", nf.Message, nf.Err.Error())
}
``` 

