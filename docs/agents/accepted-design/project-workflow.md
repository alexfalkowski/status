# Accepted design for the project workflow

These entries are part of the [accepted-design index](../accepted-design.md).

## Docker publication

Status relies on its existing release workflow and ships Docker images
frequently. A repository-local check that starts the container and probes its
health before image publication is optional.

Raise the missing check as a reliability or project gap only when concrete
evidence shows current release breakage or the repository owner decides to
require that check locally.

[CircleCI release workflow](../../../.circleci/config.yml)
