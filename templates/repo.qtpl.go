// This file is automatically generated by qtc from "repo.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line repo.qtpl:1
package templates

//line repo.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line repo.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line repo.qtpl:1
func (r Repository) StreamPage(qw422016 *qt422016.Writer) {
	//line repo.qtpl:1
	qw422016.N().S(`
<!DOCTYPE HTML>
<html>

<head>
  <title>`)
	//line repo.qtpl:6
	qw422016.E().S(r.Repository.ID.Owner)
	//line repo.qtpl:6
	qw422016.N().S(`/`)
	//line repo.qtpl:6
	qw422016.E().S(r.Repository.ID.Name)
	//line repo.qtpl:6
	qw422016.N().S(` - GradleUpdate</title>
  `)
	//line repo.qtpl:7
	StreamStylesheets(qw422016)
	//line repo.qtpl:7
	qw422016.N().S(`
</head>
<body>

`)
	//line repo.qtpl:11
	StreamHeader(qw422016)
	//line repo.qtpl:11
	qw422016.N().S(`

<div class="container">
  <section class="text-center">
    <div class="jumbotron">
      <img src="`)
	//line repo.qtpl:16
	qw422016.E().S(r.Repository.AvatarURL)
	//line repo.qtpl:16
	qw422016.N().S(`" alt="avatar" width="160" height="160" class="img-circle"/>
      <h2>`)
	//line repo.qtpl:17
	qw422016.E().S(r.Repository.Description)
	//line repo.qtpl:17
	qw422016.N().S(`</h2>
      <p>`)
	//line repo.qtpl:18
	qw422016.E().S(r.Repository.ID.Owner)
	//line repo.qtpl:18
	qw422016.N().S(`/`)
	//line repo.qtpl:18
	qw422016.E().S(r.Repository.ID.Name)
	//line repo.qtpl:18
	qw422016.N().S(`</p>
    </div>

    <p><img src="`)
	//line repo.qtpl:21
	qw422016.E().S(r.BadgeURL)
	//line repo.qtpl:21
	qw422016.N().S(`" alt="badge"/></p>
    `)
	//line repo.qtpl:22
	if r.UpToDate {
		//line repo.qtpl:22
		qw422016.N().S(`
    <p>Gradle is up-to-date.</p>
    `)
		//line repo.qtpl:24
	} else {
		//line repo.qtpl:24
		qw422016.N().S(`
    <form method="POST" action="`)
		//line repo.qtpl:25
		qw422016.E().S(r.SendPullRequestURL)
		//line repo.qtpl:25
		qw422016.N().S(`">
      <p>
        <button type="submit" class="btn btn-link">
          Send a Pull Request for Gradle `)
		//line repo.qtpl:28
		qw422016.E().V(r.LatestVersion)
		//line repo.qtpl:28
		qw422016.N().S(`
        </button>
      </p>
    </form>
    `)
		//line repo.qtpl:32
	}
	//line repo.qtpl:32
	qw422016.N().S(`
    <form>
      <p>
        <label class="text-uppercase" for="badge-markdown">Markdown:</label>
        <input type="text" id="badge-markdown" class="form-control input-text-monospace"
               value="[![Gradle Status](`)
	//line repo.qtpl:37
	qw422016.E().S(r.BaseURL)
	//line repo.qtpl:37
	qw422016.E().S(r.BadgeURL)
	//line repo.qtpl:37
	qw422016.N().S(`)](`)
	//line repo.qtpl:37
	qw422016.E().S(r.BaseURL)
	//line repo.qtpl:37
	qw422016.E().S(r.ThisURL)
	//line repo.qtpl:37
	qw422016.N().S(`)"/>
      </p>
    </form>
  </section>
</div>

`)
	//line repo.qtpl:43
	StreamFooter(qw422016)
	//line repo.qtpl:43
	qw422016.N().S(`

</body>
</html>
`)
//line repo.qtpl:47
}

//line repo.qtpl:47
func (r Repository) WritePage(qq422016 qtio422016.Writer) {
	//line repo.qtpl:47
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line repo.qtpl:47
	r.StreamPage(qw422016)
	//line repo.qtpl:47
	qt422016.ReleaseWriter(qw422016)
//line repo.qtpl:47
}

//line repo.qtpl:47
func (r Repository) Page() string {
	//line repo.qtpl:47
	qb422016 := qt422016.AcquireByteBuffer()
	//line repo.qtpl:47
	r.WritePage(qb422016)
	//line repo.qtpl:47
	qs422016 := string(qb422016.B)
	//line repo.qtpl:47
	qt422016.ReleaseByteBuffer(qb422016)
	//line repo.qtpl:47
	return qs422016
//line repo.qtpl:47
}
