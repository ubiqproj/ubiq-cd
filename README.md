# ubiq-cd

Uniq CD is a declarative continuous delivery tool to built GitOps and PULL-type pipelines on just one server.

## Usage

```bash
make build
```

```bash
# server
./ubiqcd
```

```bash
# create pipeline
./ubiqctl apply -f app.toml
./ubiqctl get applications
```
