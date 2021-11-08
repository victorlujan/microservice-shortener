package shortener

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T)  {
	repo := NewMockRepository()
	service := NewRedirectService(repo)
	_, err := service.Find("abc")
	assert.ErrorIs(t ,err, ErrRedirectNotFound)	
}

func TestStoreAndFind(t *testing.T)  {
	repo := NewMockRepository()
	service := NewRedirectService(repo)
	redirect := Redirect{
		URL: "http://www.google.com",
	}
	err := service.Store(&redirect)
	assert.Nil(t,err)
	res,err:= service.Find(redirect.Code)
	assert.Equal(t,redirect.URL,res.URL)
}

func TestInvalid(t *testing.T)  {

	//assert.Error.Is(t, err, ErrRedirectInvalid)
	//assert.Error.Is(t, err, ErrRedirectNotFound)
}