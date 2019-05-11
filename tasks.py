from invoke import task


@task
def deps(ctx, update=False):
    ctx.run("go get -u github.com/golang/dep/cmd/dep")
    ctx.run("go get -u golang.org/x/lint/golint")
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
    """ Build the oasis binary """
    ctx.run("go build -o oasis ./cmd/oasis")

@task
def docker_runner(ctx):
    """ Build and push the CI docker image"""
    ctx.run("docker build --no-cache -t xvello/oasis-circleci-runner .circleci")
    ctx.run("docker push xvello/oasis-circleci-runner")
