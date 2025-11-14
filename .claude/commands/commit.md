# Git Commit Message Generator

Generate a git commit message following the Conventional Commits specification.

## Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

## Commit Types

- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `build`: Changes that affect the build system or external dependencies
- `ci`: Changes to CI configuration files and scripts
- `chore`: Other changes that don't modify src or test files
- `revert`: Reverts a previous commit

## Rules

1. **Subject line**:
   - Use imperative, present tense: "change" not "changed" nor "changes"
   - Don't capitalize the first letter
   - No period (.) at the end
   - Maximum 50 characters (soft limit)
   - Be concise and descriptive

2. **Body** (optional but recommended for complex changes):
   - Use imperative, present tense
   - Explain the "what" and "why" vs. "how"
   - Wrap at 72 characters
   - Can include bullet points for multiple changes
   - Reference issues/PRs if applicable

3. **Footer** (optional):
   - Breaking changes: `BREAKING CHANGE: <description>`
   - Issue references: `Closes #123`, `Fixes #456`

4. **Scope** (optional):
   - The scope should be the name of the package affected
   - Examples: `parser`, `utils`, `api`, `config`

## Examples

### Simple fix
```
fix: correct typo in README
```

### Feature with scope
```
feat(parser): add support for custom delimiters
```

### Complex change with body
```
refactor(url-parser): simplify regex patterns

Break down complex regex into semantic constants:
- domainSuffix: domain suffix pattern
- hostnamePattern: hostname pattern
- ownerPattern: username/organization pattern
- repoPattern: repository name pattern
- gitSuffix: optional .git suffix

This improves code readability and maintainability while
preserving existing functionality.
```

### Breaking change
```
feat(api): change authentication method

BREAKING CHANGE: API now requires Bearer token instead of API key.
Update your client code to use Authorization header.
```

## Instructions

When generating a commit message:
1. Analyze the git diff to understand what changed
2. Determine the appropriate type and scope
3. Write a clear, concise subject line
4. Add a detailed body if the change is complex or needs explanation
5. Include breaking changes or issue references in the footer if applicable
6. Ensure the message follows all formatting rules above
7. **After generating the commit message, automatically execute `git push` command to push the changes to the remote repository**
