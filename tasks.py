from invoke import task


@task
def deps(ctx, update=False):
    """ Install binary deps and checkout vendors """
    ctx.run("go get -u github.com/golang/dep/cmd/dep")
    ctx.run("go get -u golang.org/x/lint/golint")
    ctx.run("go get -u github.com/gordonklaus/ineffassign")
    ctx.run("go get -u github.com/client9/misspell/cmd/misspell")
    ctx.run("go get -u github.com/docker/distribution/cmd/registry")

    if update:
        ctx.run("dep ensure -update")
    else:
        ctx.run("dep ensure -vendor-only")


@task
def lint(ctx):
    """ Run the unit tests """
    ctx.run("go vet ./cmd/... ./pkg/... ./tests/...")
    ctx.run("golint ./cmd/... ./pkg/... ./tests/...")
    ctx.run("misspell ./cmd/ ./pkg/ ./tests/")
    ctx.run("ineffassign ./cmd/ ./pkg/ ./tests/")


@task
def test(ctx):
    """ Run the unit tests """
    ctx.run("go test ./cmd/... ./pkg/...")


@task
def itest(ctx, dockerize=False):
    """ Run the integration tests """
    if dockerize:
        ctx.run("docker run --rm -v $PWD:/go/src/github.com/xvello/oasis-nomad xvello/oasis-circleci-runner inv itest")
    else:
        build(ctx)
        ctx.run("go test ./tests/... -v")


@task
def build(ctx):
    """ Build the oasis binary """
    ctx.run("go build -o oasis ./cmd/oasis")


@task
def docker_runner(ctx):
    """ Build and push the CI docker image"""
    ctx.run("docker build --no-cache -t xvello/oasis-circleci-runner .circleci")
    ctx.run("docker push xvello/oasis-circleci-runner")
