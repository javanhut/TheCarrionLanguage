# Carrion Language v0.1.8 Release Notes

**Release Date**: July 20, 2025  
**Version**: 0.1.8  
**Previous Version**: 0.1.7

## Overview

Carrion Language v0.1.8 introduces a comprehensive tooling ecosystem that transforms the development experience with three major new tools: **Sindri Testing Framework**, **Mimir Documentation System**, and **Bifrost Package Manager**. This release focuses on developer productivity, testing capabilities, and documentation accessibility while maintaining full backward compatibility.

## 🧪 Sindri Testing Framework

### What is Sindri?
Sindri is Carrion's built-in testing framework, named after the legendary dwarven forge. It provides a simple yet powerful way to write and run tests for your Carrion code.

### Key Features
- **Automatic Test Discovery**: Finds test functions using the "appraise" naming convention
- **Colored Terminal Output**: Green for passed tests, red for failures
- **Flexible Assertions**: Support for both boolean and value comparison assertions
- **Multiple Output Modes**: Summary and detailed reporting modes
- **Directory Testing**: Run all tests in a directory or specific files

### Basic Usage
```bash
# Install Sindri (included with Carrion installation)
make install

# Run tests
sindri appraise examples/sindri_demo.crl     # Single file
sindri appraise ./tests                      # Directory
sindri appraise -d                           # Detailed output
```

### Writing Tests
```carrion
spell appraise_arithmetic():
    check(2 + 2 == 4)
    check(10 - 3, 7)  # Value comparison format

spell test_appraise_strings():
    greeting = "Hello, " + "World!"
    check(greeting == "Hello, World!")
```

## 📚 Mimir Documentation System

### What is Mimir?
Mimir is Carrion's interactive documentation and help system, named after Odin's wise raven. It provides comprehensive documentation access both interactively and through command-line queries.

### Key Features
- **Interactive Documentation Browser**: Menu-driven exploration of language features
- **Command-Line Lookup**: Quick help for specific functions
- **Comprehensive Coverage**: Built-in functions, standard library, language features
- **Search Functionality**: Find functions by name or purpose
- **REPL Integration**: Seamless integration with the Carrion REPL

### Usage
```bash
# Interactive mode
mimir

# Quick function lookup
mimir scry print
mimir scry Array
mimir scry os

# List all functions
mimir list

# Browse by categories
mimir categories
```

### Interactive Experience
```
$ mimir
    MIMIR - The All-Seeing Helper
   ═══════════════════════════════════════
    Interactive Carrion Documentation
   ═══════════════════════════════════════

What knowledge do you seek?

  1.  Built-in Functions    - Core language functions
  2.  Standard Library      - Munin modules  
  3.  Language Features     - Syntax, control flow, OOP
  4.  Examples & Demos      - Working code examples
  5.  Search Functions      - Find specific functions
  6.  Tips & Tricks         - REPL shortcuts
```

## 📦 Bifrost Package Manager

### What is Bifrost?
Bifrost is Carrion's package manager, integrated as a Git submodule. It provides package management capabilities for the growing Carrion ecosystem.

### Integration
- **Git Submodule**: Integrated directly into the Carrion repository
- **Build System Updates**: Enhanced Makefile and installation scripts
- **Documentation Updates**: Updated ecosystem documentation

## 🔧 Enhanced Development Experience

### Build System Improvements
```bash
# New Makefile targets
make sindri          # Build testing framework
make mimir           # Build documentation system
make install         # Install all tools including new ones
make all            # Build everything including new tools
```

### Installation & Setup
- **Updated Installation Scripts**: `install/install.sh` now includes Sindri
- **Enhanced Setup**: `setup.sh` updated for complete development environment
- **Uninstall Support**: `install/uninstall.sh` removes all tools cleanly

### Example Files & Demos
- **Comprehensive Test Examples**: Multiple example files demonstrating testing patterns
- **Documentation Examples**: Working code examples integrated with Mimir
- **Visual Enhancements**: ASCII art and improved terminal experience

## 🏗️ Technical Implementation

### Core Integration
- **Enhanced Evaluator**: Added `check()` builtin function for test assertions
- **Improved Error Handling**: Better error reporting for test execution
- **Module System Updates**: Enhanced function registry and error reporting

### File Structure
```
cmd/
  sindri/           # Testing framework
    main.go
    go.mod
  mimir/            # Documentation system  
    main.go
    *.go
  assets/           # Shared assets (ASCII art, etc.)
bifrost/            # Package manager (submodule)
docs/
  Sindri.md         # Testing framework documentation
  Mimir.md          # Documentation system reference
examples/
  sindri_demo.crl   # Test examples
  sindri_test_*.crl # More test patterns
```

## 📖 Documentation Updates

### New Documentation
- **Sindri.md**: Complete testing framework guide with examples
- **Mimir.md**: Documentation system reference and usage
- **Updated Module Documentation**: Enhanced for new tooling ecosystem

### Enhanced Existing Docs
- **README Updates**: Tool ecosystem overview
- **Language Documentation**: Updated to reference new tools
- **Installation Guides**: Updated for new tool installation

## 🔄 Migration & Compatibility

### Backward Compatibility
- **Zero Breaking Changes**: All existing Carrion code continues to work unchanged
- **Optional Tooling**: New tools enhance but don't replace existing workflows
- **Existing APIs**: All functions and modules remain the same

### New Commands
```bash
# Testing workflow
sindri appraise mytest.crl

# Documentation workflow  
mimir scry <function>
mimir list

# Enhanced REPL experience
carrion  # REPL now suggests using Mimir for help
```

### Recommended Workflow
1. **Development**: Write Carrion code as usual
2. **Testing**: Use Sindri to create and run tests
3. **Documentation**: Use Mimir for quick function reference
4. **Package Management**: Use Bifrost for external packages (future)

## 🚀 Getting Started with v0.1.8

### For Existing Users
1. **Update Installation**: Run `make install` to get new tools
2. **Try Sindri**: Create a test file with `appraise_` functions
3. **Explore Mimir**: Run `mimir` to explore documentation
4. **Continue Development**: All existing code works unchanged

### For New Users
1. **Install Carrion**: Follow standard installation process
2. **Learn with Mimir**: Use `mimir` to explore language features
3. **Test with Sindri**: Write tests using `check()` assertions
4. **Develop**: Use REPL and tools for interactive development

## 🎯 What's Next

### Future Enhancements
- **Enhanced Sindri Features**: Code coverage, benchmarking, parallel testing
- **Expanded Mimir Capabilities**: Interactive tutorials, fuzzy search
- **Bifrost Development**: Package publishing, dependency management
- **IDE Integration**: Language server, syntax highlighting, debugging

### Community & Ecosystem
- **Testing Best Practices**: Community-driven testing patterns
- **Documentation Contributions**: Community documentation via Mimir
- **Package Ecosystem**: Growing library of Carrion packages via Bifrost

## 🛠️ Technical Notes

### Performance
- **Fast Test Execution**: Minimal overhead for test running
- **Responsive Documentation**: Quick interactive navigation
- **Efficient Build Process**: Parallel compilation of tools

### Quality Assurance
- **Comprehensive Testing**: All new tools thoroughly tested
- **Documentation Coverage**: Complete documentation for all features
- **Error Handling**: Robust error reporting and recovery

## 📋 Summary

Carrion Language v0.1.8 represents a major milestone in the language's development, introducing a complete tooling ecosystem that enhances every aspect of the development experience:

- 🧪 **Sindri**: Comprehensive testing framework for reliable code
- 📚 **Mimir**: Interactive documentation for learning and reference  
- 📦 **Bifrost**: Package manager for ecosystem growth
- 🔧 **Enhanced Tooling**: Improved build system and installation
- 📖 **Better Documentation**: Comprehensive guides and references

This release maintains 100% backward compatibility while providing powerful new tools that make Carrion development more productive, reliable, and enjoyable.

**Download Carrion v0.1.8 today and experience the enhanced development ecosystem!**

---

**Contributors**: Carrion Language Team  
**Documentation**: Complete guides available in `docs/`  
**Support**: GitHub Issues, Community Forums  
**License**: Same as Carrion Language