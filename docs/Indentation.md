# Indentation Handling in Carrion

Carrion uses indentation‐sensitive syntax similar to Python, so both the lexer and parser cooperate to track block boundaries. This note documents the mechanics so it is easier to reason about future changes.

## Indentation Style

Carrion requires consistent indentation throughout each file:

- **Spaces**: Use spaces (recommended: 4 spaces per level)
- **Tabs**: Use tabs (displayed as 4 spaces)

**You cannot mix tabs and spaces in the same file.** The first indented line determines the style for the entire file. Mixing styles produces an error:

```
Error: inconsistent indentation: expected spaces, got tabs
```

Or if you mix within the same line:

```
Error: mixed tabs and spaces in indentation
```

### Why This Restriction?

Mixed indentation causes unpredictable behavior because tabs display differently across editors. Enforcing consistency ensures code looks and behaves the same everywhere.

## Lexer responsibilities

The lexer (`src/lexer/lexer.go`) measures the number of leading spaces or tabs on each physical line. When the indent level increases or decreases it emits synthetic `INDENT`/`DEDENT` tokens before the next real token (blank lines are ignored). These tokens carry the column where the first non‑space character appears; the parser uses that column number to understand how far the dedent walked back.

The lexer also tracks the indentation style (`indentStyle` field) starting from the first indented line. If a subsequent line uses a different style, or mixes tabs and spaces, the lexer emits an `INDENT_ERROR` token with a descriptive error message.

## Parser bookkeeping

The parser maintains two stacks:

- `contextStack`: existing structure that tracks logical constructs (`"grim"`, `"if"`, etc.) for semantic checks.
- `indentStack`: newly added stack that records the numeric indentation depth currently in effect. It is initialised to `[0]` and updated automatically in `nextToken()` whenever an `INDENT` or `DEDENT` token is seen.

`getIndentLevel()` simply reads the top of `indentStack`, ensuring any part of the parser can query the current physical indentation.

## Block parsing logic

`parseBlockStatement()` now captures the block’s indent depth when it starts (`blockIndent := getIndentLevel()`). As it consumes statements, any intervening `DEDENT` tokens are inspected:

- If the dedent only returns to the block’s indent (or increases/decreases within nested sub‑blocks), parsing continues.
- If the dedent drops *below* the block’s indent, the block is complete and the loop stops, letting the caller handle the lower indentation.

This prevents premature termination when a nested structure temporarily dedents (e.g., returning to the class body after finishing a `spell`). Without this guard the `grim BTree` definition in `src/munin/datastructures.crl` lost the `find()` and `iter()` spells because the parser bailed out at the dedent that followed `print_tree()`.

## Practical testing guidance

When touching indentation logic, run at least:

```bash
GOCACHE=$(pwd)/.gocache go run ./src test_simple_btree.crl
GOCACHE=$(pwd)/.gocache go run ./src test_all_datastructures.crl
```

Those files exercise deeply nested grimoires and ensure methods declared late in the block remain attached to their owning class. Add new regression tests near similar structures if you uncover another corner case.
