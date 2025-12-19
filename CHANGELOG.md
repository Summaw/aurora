# Changelog

All notable changes to Aurora will be documented in this file.

## [1.0.0] - 2025-12-19

### Added
- Initial release
- ASCII art banner system with gradient support
- Structured logging with tree-style field display
- Multiple log levels: Trace, Debug, Info, Success, Warn, Error, Fatal, Panic
- Field chaining API (Str, Int, Bool, Dur, Time, Any, Err)
- 22 built-in gradient presets
- Custom gradient support (RGB and hex colors)
- UI components:
  - Tables with customizable borders
  - Progress bars with gradients
  - Animated spinners
  - Dividers/separators
  - Key-value displays
  - Bordered boxes
- JSON output mode for production
- Caller information (file:line)
- Context support for request tracing
- HTTP middleware helpers
- Three ASCII art fonts: block, slant, minimal
- Multiple border styles: rounded, sharp, double, heavy, ascii
- Zero external dependencies
