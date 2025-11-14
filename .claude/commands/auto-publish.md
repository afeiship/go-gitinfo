# Auto Publish Command

This command automatically publishes the package to npm registry.

## Command

Execute the following command in the project root directory:

```bash
proxy npm run release -- --ci --npm.skipChecks
```

## Instructions

When executing auto-publish:
1. Ensure you are in the project root directory
2. Execute the command: `proxy npm run release -- --ci --npm.skipChecks`
3. The `--ci` flag indicates this is running in CI environment
4. The `--npm.skipChecks` flag skips npm registry checks
5. Wait for the command to complete and verify the output

## Notes

- This command will automatically:
  - Build the project
  - Run tests (if configured)
  - Bump version (if configured)
  - Publish to npm registry
  - Create git tags (if configured)

- Make sure all changes are committed before running this command
- Ensure you have proper npm authentication configured
- The `proxy` prefix is used to route through a proxy if needed

