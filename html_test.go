package html

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTML(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html>
<html itemscope itemtype="https://schema.org/Article">
<head>
<title>DSN</title>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,minimum-scale=1,initial-scale=1">
<meta name="author" value="wmentor">
<meta name="keywords" content="wmentor,lemma,nlp,golang">
<meta name="description" content=" We often work with configuration strings like pq database connection…" />
<!-- Schema.org markup for Google+ -->
<meta itemprop="name" content="DSN">
<meta itemprop="description" content=" We often work with configuration strings like pq database connection…">
<meta itemprop="image" content="https://wmentor.ru/pics/go/5e67edc8.jpg">
<!-- Twitter Card data -->
<meta name="twitter:card" content="summary">
<meta name="twitter:site" content="@wmentorru">
<meta name="twitter:title" content="DSN">
<meta name="twitter:description" content=" We often work with configuration strings like pq database connection…">
<meta name="twitter:creator" content="@wmentorru">
<meta name="twitter:image" content="https://wmentor.ru/pics/go/5e67edc8.jpg">
<!-- Open Graph -->
<meta property="og:title" content="DSN" />
<meta property="og:type" content="article" />
<meta property="og:url" content="https://wmentor.ru/dsn" />
<meta property="og:image" content="https://wmentor.ru/pics/go/5e67edc8.jpg" />
<meta property="og:description" content=" We often work with configuration strings like pq database connection…" />
<meta property="og:site_name" content="WMentor.ru" />
<link rel="shortcut icon" href="/favicon.ico?v=1583781460" ; type="image/x-icon"/>
<link rel="icon" href="/favicon.ico?v=1583781460" ; type="image/x-icon"/>

<link rel="stylesheet" href="/css/styles.css?v=1583781460">
<link rel="stylesheet" href="/css/styles/a11y-light.css?v=1583781460">

<script src="https://ajax.googleapis.com/ajax/libs/jquery/2.2.4/jquery.min.js"></script>
<script src="/js/highlight.pack.js?v=1583781460"></script>
<script src="/js/rpc.js?v=1583781460"></script>
<script src="/js/share.js?v=1583781460"></script>
<body>
<div class="menu">
<div class="area">
<table width="100%" cellspacing="0" cellpadding="0" class="mobile_menu">
  <tr>
    <td style="WIDTH: 40px;"><img src="/images/menu.png?v=1583781460" style="" onclick="$('.menu_widget').css('display','block'); return false
;"/></td>
    <td style="COLOR:#FFF;TEXT-ALIGN:CENTER;" onclick="location.href='/';">WMENTOR.RU</td>
    <td style="WIDTH: 40px;"><img src="/images/search.png?v=1583781460" style="" onclick="$('.search_widget').css('display','block'); return f
alse;"/></td>
  </tr>
</table>
<img src="/images/logo.png?v=1583781460" height="48px" class="logo"/>
<div class="search"><table width="100%"><tr><form action="/" method="GET"><td><input type="text" id="search" name="search" placeholder="Search
..."></td><td><img src="/images/search.png?v=1583781460" height="48px"/></td></form></tr></table></div>
<table class="desktop_menu">
  <tr style="VERTICAL-ALIGN:MIDDLE;">
    <td><a href="/" class="name">WMentor.ru</a></td>
</table>
</div>
</div>
<div class="menu_widget">
<div class="close" onclick="$('.menu_widget').css('display','none'); return false;">X</div>
<a href="/">Main page</a><br/>
<a href="/golang">Golang</a><br/>
<a href="/hello">About</a><br/>
</div>
<div class="search_widget">
<div class="close" onclick="$('.search_widget').css('display','none'); return false;">X</div>
<form method="GET" action="/" id="mobile_search_form">
<table width="100%">
<tr>
  <td style="TEXT-ALIGN: CENTER;COLOR: #FFF;FONT-SIZE: 24px;">SEARCH</td>
</tr>
<tr>
<td style="TEXT-ALIGN:CENTER;"><input type="text" name="search" id="m_search" style="WIDTH:90%;"></td>
</tr>
<tr>
<td style="TEXT-ALIGN: CENTER;"><img src="/images/search.png" width="48px;" height="48px;" onclick="$('#mobile_search_form').submit(); return
false;" /></td>
</tr>
</table>
</form>
</div>
<div class="container">

  <div class="r-column">
    <br/>
    <div class="r-column-widget">
<div class="r-column-widget-head">top tags</div>
<div class="r-column-widget-body">
  <a href="/?tag=4">golang</a>
  <a href="/?tag=2">nlp</a>
  <a href="/?tag=6">go</a>
  <a href="/?tag=16">programming</a>
  <a href="/?tag=7">text mining</a>
  <a href="/?tag=14">lemma</a>
  <a href="/?tag=18">it</a>
  <a href="/?tag=10">lemmatization</a>
  <a href="/?tag=12">text processing</a>
  <a href="/?tag=3">tokenizer</a>
  <a href="/?tag=15">wmentor</a>
  <a href="/?tag=24">configuration</a>
  <a href="/?tag=26">data source name</a>
  <a href="/?tag=5">development</a>
</div>
</div>
<br/>
    <div class="r-column-widget">
<div class="r-column-widget-head">archive</div>
<div class="r-column-widget-body">
<a href="/?period=2020">2020</a>
<a href="/?period=2019">2019</a>
</div>
</div>
<br/>
    <div class="r-column-widget">
<div class="r-column-widget-head">archive</div>
<div class="r-column-widget-body">
<a href="/?period=2020">2020</a>
<a href="/?period=2019">2019</a>
</div>
</div>
<br/>
    <div class="r-column-widget">
<div class="r-column-widget-head">follow me</div>
<div class="r-column-widget-body">
<a href="https://twitter.com/wmentorru" target="_blank"><img class="social-icon-small social-tw" src="/images/social/tw.png?v=1583781460"></a>
<a href="https://t.me/wmentor" target="_blank"><img class="social-icon-small social-tg" src="/images/social/tg.png?v=1583781460"></a>
<a href="/rss.xml" title="RSS" target="_blank"><img class="social-icon-small social-rss" src="/images/social/rss.png?v=1583781460"></a>
<a href="https://github.com/wmentor" title="GitHub" target="_blank"><img class="social-icon-small social-github" src="/images/social/github.pn
g?v=1583781460"></a>
</div>
</div>
<br/>
  </div><div class="l-column">
<div>
<h1>DSN</h1>
<p>We often work with configuration strings like pq database connection:</p>

<pre><code class="plaintext">user=mylogin password=mypass database=mydb host=127.0.0.1 port=5432
</code></pre>

<figure>
<img src="/pics/go/5e67edc8.jpg"/>
</figure>

<p>The simplest way to parse them in Golang is regular expressions:</p>

<pre><code class="go">package main

import (
  "fmt"
  "regexp"
)

var rex = regexp.MustCompile("(\\w+)=(\\w+)")

func main() {
  conn := "user=mylogin password=mypass database=mydb host=127.0.0.1 port=5432 sslmode=true"

  data := rex.FindAllStringSubmatch(conn, -1)

  res := make(map[string]string)
  for _, kv := range data {
    k := kv[1]
    v := kv[2]
    res[k] = v
  }

  fmt.Println(res)
}
</code></pre>
<p>But if the value can contain a space character, then you need to add escaping support in a regular expression. And so on, each add[169/579]
ction makes processing harder. But there is an easier way (wmentor/dsn).</p>

<p>Install package:</p>

<pre><code class="plaintext">go get github.com/wmentor/dsn
</code></pre>

<p>Usage:</p>

<pre><code class="go">package main

import (
  "fmt"

  "github.com/wmentor/dsn"
)
func main() {                                                                                                                        [151/579]

  str := "user=mylogin passwd=mypass database=mydb port=5432 sslmode=true"

  ds, err := dsn.New(str)
  if err != nil {
    panic("invalid string")
  }

  // print user=mylogin
  fmt.Printf( "user=%s\n", ds.GetString("user","unknown") )

  // print passwd=mypass
  fmt.Printf( "passwd=%s\n", ds.GetString("passwd","nopass") )

  // host is not exists, print host=127.0.0.1
  fmt.Printf( "host=%s\n", ds.GetString("host","127.0.0.1") )

  // get int value and print port=5432
  fmt.Printf( "port=%d\n", ds.GetInt("port", 4321) )

  // print sslmode=true
  fmt.Printf( "sslmode=%t\n", ds.GetBool("sslmode", false) )

  // print keepalive=false
  fmt.Printf( "keepalive=%t\n", ds.GetBool("keepalive", false) )
}
</code></pre>

<p><b>dns.New</b> returns object <b>dsn.DSN</b> or an error. All get methods (<i>GetString</i>, <i>GetBool</i>, <i>GetInt</i>, <i>GetInt64</i$
, <i>GetFloat</i>) take 2 arguments - key name and default value. The default value is used when the key is missing or contains a invalid val$
e.</p>

<p>Moreover, dsn support escape some characters in key name and value (\s,\t,\r,\n,\=,\\,\",\').</p>

<script language="javascript">
<!--

var uid = "fewrretwertwertwert";

SharingOpts = {"url":"https://wmentor.ru/dsn?v=1","title":"DSN","text":" We often work with configuration strings like pq database connection…
","image":"https://wmentor.ru/pics/go/5e67edc8.jpg"};

-->
</script>

</body>
</html>
`))

	}))
	defer ts.Close()

	parser := New()
	err := parser.Get(ts.URL, nil)
	if err != nil {
		t.Fatal("HTML.Get failed")
	}

	txt := string(parser.Text())

	waitText := `WMENTOR.RU
WMentor.ru
X
Main page
Golang
About
X
top tags
golang nlp go programming text mining lemma it lemmatization text processing tokenizer wmentor configuration data source name development
archive
2020 2019
archive
2020 2019
follow me
DSN We often work with configuration strings like pq database connection:
user=mylogin password=mypass database=mydb host=127.0.0.1 port=5432 The simplest way to parse them in Golang is regular expressions:
package main

import (
  "fmt"
  "regexp"
)

var rex = regexp.MustCompile("(\\w+)=(\\w+)")

func main() {
  conn := "user=mylogin password=mypass database=mydb host=127.0.0.1 port=5432 sslmode=true"

  data := rex.FindAllStringSubmatch(conn, -1)

  res := make(map[string]string)
  for _, kv := range data {
    k := kv[1]
    v := kv[2]
    res[k] = v
  }

  fmt.Println(res)
} But if the value can contain a space character, then you need to add escaping support in a regular expression. And so on, each add[169/579]
ction makes processing harder. But there is an easier way (wmentor/dsn).
Install package:
go get github.com/wmentor/dsn Usage:
package main

import (
  "fmt"

  "github.com/wmentor/dsn"
)
func main() {                                                                                                                        [151/579]

  str := "user=mylogin passwd=mypass database=mydb port=5432 sslmode=true"

  ds, err := dsn.New(str)
  if err != nil {
    panic("invalid string")
  }

  // print user=mylogin
  fmt.Printf( "user=%s\n", ds.GetString("user","unknown") )

  // print passwd=mypass
  fmt.Printf( "passwd=%s\n", ds.GetString("passwd","nopass") )

  // host is not exists, print host=127.0.0.1
  fmt.Printf( "host=%s\n", ds.GetString("host","127.0.0.1") )

  // get int value and print port=5432
  fmt.Printf( "port=%d\n", ds.GetInt("port", 4321) )

  // print sslmode=true
  fmt.Printf( "sslmode=%t\n", ds.GetBool("sslmode", false) )

  // print keepalive=false
  fmt.Printf( "keepalive=%t\n", ds.GetBool("keepalive", false) )
} dns.New returns object dsn.DSN or an error. All get methods ( GetString , GetBool , GetInt , GetInt64 GetFloat ) take 2 arguments - key name and default value. The default value is used when the key is missing or contains a invalid val$
e.
Moreover, dsn support escape some characters in key name and value (\s,\t,\r,\n,\=,\\,\",\').
`

	if waitText != txt {
		t.Fatal("Invalid result")
	}

	totalUrls := 0
	parser.EachLink(func(str string) {
		totalUrls++
	})

	if totalUrls != 23 {
		t.Fatal("EachLink failed")
	}

	totalImg := 0
	parser.EachImage(func(str string) {
		totalImg++
	})

	if totalImg != 9 {
		t.Fatal("EachImg failed")
	}

	totalIframes := 0

	parser.EachIframe(func(str string) {
		totalIframes++
	})

	if totalIframes != 0 {
		t.Fatal("EachIframe failed")
	}
}
