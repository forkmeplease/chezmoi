# `gitHubLatestRelease` *owner-repo*

`gitHubLatestRelease` calls the GitHub API to retrieve the latest release about
the given *owner-repo*, returning structured data as defined by the
[GitHub Go API bindings][bindings].

Calls to `gitHubLatestRelease` are cached so calling `gitHubLatestRelease` with
the same *owner-repo* will only result in one call to the GitHub API.

!!! example

    ```
    {{ (gitHubLatestRelease "docker/compose").TagName }}
    ```

!!! hint

    Some fields in the returned object have type `*string`. Use the
    [`toString` template function][toString] to convert these to strings.

[bindings]: https://pkg.go.dev/github.com/google/go-github/v61/github#RepositoryRelease
[toString]: ../functions/toString.md
