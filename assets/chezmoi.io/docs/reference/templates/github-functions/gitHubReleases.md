# `gitHubReleases` *owner-repo*

`gitHubReleases` calls the GitHub API to retrieve the first page of releases for
the given *owner-repo*, returning structured data as defined by the
[GitHub Go API bindings][github-go].

Calls to `gitHubReleases` are cached so calling `gitHubReleases` with the same
*owner-repo* will only result in one call to the GitHub API.

!!! example

    ```
    {{ (index (gitHubReleases "docker/compose") 0).TagName }}
    ```

!!! note

    The maximum number of items returned by `gitHubReleases` is determined by
    default page size for the GitHub API.

!!! warning

    The values returned by `gitHubReleases` are not directly queryable via the
    [`jq`][jq] function and must instead be converted through JSON:

    ```
    {{ gitHubReleases "docker/compose" | toJson | fromJson | jq ".[0].tag_name" }}
    ```

!!! hint

    Some fields in the returned object have type `*string`. Use the
    [`toString` template function][toString] to convert these to strings.

[github-go]: https://pkg.go.dev/github.com/google/go-github/v61/github#RepositoryRelease
[jq]: /reference/templates/functions/jq.md
[toString]: ../functions/toString.md
