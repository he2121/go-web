package hehe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRouter("GET", "/", nil)
	r.addRouter("GET", "/hello/:name", nil)
	r.addRouter("GET", "/hello/b/c", nil)
	r.addRouter("GET", "/hi/:name", nil)
	r.addRouter("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRouter("GET", "/hello/he2121")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "he2121" {
		t.Fatal("name should be equal to 'he2121'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

	n, params := r.getRouter("GET", "/assets/he2121/test/1.avi")
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/assets/*filepath" {
		t.Fatal("should match /assets/*filepath")
	}
	if params["filepath"] != "he2121/test/1.avi" {
		t.Fatal("filepath should be equal to 'he2121/test/1.avi'")
	}
	fmt.Printf("matched path: %s, params['filepath']: %s\n", n.pattern, ps["filepath"])
}

func TestGroup(t *testing.T) {
	engine := New()
	v1 := engine.Group("/v1")
	assert.True(t, v1.prefix == "/v1")
	assert.True(t, v1.engine == engine)
	v1.GET("/hello", func(c *Context) {
	})
	v1.GET("/hello/:name", func(c *Context) {
	})
	assert.True(t, engine.router.handlers["GET-/v1/hello"] != nil)
	assert.True(t, engine.router.handlers["GET-/v1/hello/:name"] != nil)

	v1 = engine.Group("/v2")
	v1.GET("/test", func(c *Context) {
	})
	assert.True(t, engine.router.handlers["GET-/v2/test"] != nil)
}
