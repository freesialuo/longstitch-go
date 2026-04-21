# Implementation References

## GUI

- Fyne package quickstart and app lifecycle:
  - https://github.com/fyne-io/docs.fyne.io/blob/main/content/docs/api/v2/fyne/pkg.md
- File dialogs (`ShowFileOpen`, `ShowFileSave`):
  - https://github.com/fyne-io/docs.fyne.io/blob/main/content/docs/api/v2/dialog/pkg.md
- Extension filters (`NewExtensionFileFilter`):
  - https://github.com/fyne-io/docs.fyne.io/blob/main/content/docs/api/v2/storage/filefilter.md

## Image stitching

- `image.Decode` and decoder registration side effects:
  - https://go.dev/pkg/image
- `image.DecodeConfig` usage for size checks:
  - https://go.dev/pkg/image
- `draw.Draw` composition model:
  - https://go.dev/pkg/image/draw

## CI/CD

- setup-go action usage and go-version-file:
  - https://github.com/actions/setup-go
- upload-artifact v4 usage and artifact naming in matrix:
  - https://github.com/actions/upload-artifact/blob/v4/README.md
- storing artifacts in workflows:
  - https://docs.github.com/en/actions/using-workflows/storing-workflow-data-as-artifacts
- release creation and assets upload:
  - https://github.com/softprops/action-gh-release
- GITHUB_TOKEN least-privilege guidance:
  - https://docs.github.com/en/actions/using-jobs/assigning-permissions-to-jobs
