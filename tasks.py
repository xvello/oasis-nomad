from invoke import task


@task
def deps(ctx, update=False):
    ctx.run("go get -u github.com/golang/dep/cmd/dep")
    ctx.run("go get -u github.com/golang/lint/golint")
    ctx.run("go get -u github.com/gordonklaus/ineffassign")
    ctx.run("go get -u github.com/client9/misspell/cmd/misspell")

    if update:
        ctx.run("dep ensure -update")
    else:
        ctx.run("dep ensure -vendor-only")


@task
def lint(ctx):
    ctx.run("go vet ./cmd/... ./pkg/...")
    ctx.run("golint ./cmd/... ./pkg/...")
    ctx.run("misspell ./cmd/ ./pkg/")
    ctx.run("ineffassign ./cmd/ ./pkg/")


@task
def test(ctx):
    ctx.run("go test ./cmd/... ./pkg/...")


@task
def build(ctx):
    ctx.run("go build -o oasis ./cmd/oasis")
