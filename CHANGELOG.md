# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- Support for style extensions on article pages
- Support for online editing

### Changed

- Refactored both frontend and backend

### Fixed

- Fixed backend bug

## [0.3.0]

### Added

- hange the stored Markdown data from Goldmark-rendered results to raw Markdown text in order to support future editing operations.
- Due to database changes, this is not compatible with the previous version.

## [0.2.2] - 2024-10-12

### Added

- Support for mkdocs material style admonition
  - Support for `!!!` & `???`
- Switched article page rendering to client-side

### Fixed

- Fixed an admonition bug: pointer overflow issue when consecutive admonition blocks are used

## [0.2.0] - 2024-09-20

### Added

- Support for branch switching, allowing a specific branch to be specified when creating a repository, making it easier to preview the display before merging into the mainline
