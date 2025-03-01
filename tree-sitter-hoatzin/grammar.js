/**
 * @file A Rust-inspired language for learning programming language development.
 * @author John Hiatt <john.william.hiatt@gmail.com>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "hoatzin",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
