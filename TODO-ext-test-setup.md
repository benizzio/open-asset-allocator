# TODO: Configure `ext-test` as required CI check

Temporary setup steps for the new external integration test CI job.
Delete this file after completing the steps.

## 1. Create the `ext-integration` environment

- Go to **Settings > Environments > New environment**
- Name: `ext-integration`
- Add **Required reviewers** (this gates external tests behind manual approval)

## 2. Add `ext-test` as a required status check

- Go to **Settings > Branches > Branch protection rules** (edit the rule for `main`)
- Under **Require status checks to pass before merging**, add `ext-test`

> Note: The `ext-test` check may not appear in the dropdown until the workflow has run
> at least once on a PR. Push a PR first if needed, then come back and add it.
