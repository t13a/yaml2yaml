# Yaml2yaml

A pretty YAML formatter using [go-yaml](https://github.com/go-yaml/yaml).

## Getting started

You can build from the source, but the fastest way is to run [the Docker container](https://hub.docker.com/r/t13a/yaml2yaml). The following is an example of formatting YAML containing a multi-line string. You can see that the quotes and the escape sequences has been removed and the string is comfortably formatted.

```sh
$ echo 'say: "hello\\nworld!"' | docker run --rm -i t13a/yaml2yaml
say: |-
  hello
  world!
```

Of course, it can be also used as JSON to YAML converter.

```sh
$ echo '{"say": "hello\\nworld!"}' | docker run --rm -i t13a/yaml2yaml
say: |-
  hello
  world!
```

In addition, you can output in JSON format by specifying `-json-output` flag.

```sh
$ echo 'say: "hello\\nworld!"' | docker run --rm -i t13a/yaml2yaml -json-output
{
  "say": "hello\nworld!"
}
```

What was `2yaml`? Let's not care about such a small thing.
